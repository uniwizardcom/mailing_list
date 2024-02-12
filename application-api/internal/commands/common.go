package commands

import (
	"application-api/mailbox"
	"application-api/postgresql"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
)

type Command interface {
	GetType() string
}

type CommandStructure struct {
	Type string
	Data []byte
}

type CommandHandlers struct {
	Postgresql *postgresql.PostgreSQL
	Mailbox    *mailbox.MailBox
	register   map[string]Command
}

func (ch *CommandHandlers) Init() {
	ch.register = make(map[string]Command)
}

func (ch *CommandHandlers) Add(command Command) {
	ch.register[command.GetType()] = command
}

func (ch *CommandHandlers) Execute(commandTypeExec string, data []byte) error {
	commandType, commandExec := divideToTypeSufix(commandTypeExec)

	fieldPostgresql := reflect.ValueOf(ch.register[commandType]).Elem().FieldByName("Postgresql")
	if fieldPostgresql.IsValid() && fieldPostgresql.CanSet() {
		fieldPostgresql.Set(reflect.ValueOf(ch.Postgresql))
	}

	fieldMailbox := reflect.ValueOf(ch.register[commandType]).Elem().FieldByName("Mailbox")
	if fieldMailbox.IsValid() && fieldMailbox.CanSet() {
		fieldMailbox.Set(reflect.ValueOf(ch.Mailbox))
	}

	funcName := cases.Title(language.English, cases.Compact).String(commandExec)
	meth := reflect.ValueOf(ch.register[commandType]).MethodByName(funcName)
	if !meth.IsValid() {
		return fmt.Errorf("based on [%s], function [%s] is not valid in [%s]", commandTypeExec, funcName, commandType)
	}

	inputs := make([]reflect.Value, 1)
	inputs[0] = reflect.ValueOf(data)
	meth.Call(inputs)

	return nil
}

func divideToTypeSufix(content string) (string, string) {
	contLen := len(content) - 1
	var delimiter byte = '\\'

	for contLen > 0 {
		if content[contLen] == delimiter {
			return content[0:contLen], content[contLen+1:]
		}
		contLen--
	}

	return content, ""
}
