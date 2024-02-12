package data

import "time"

type DBTableCustomers struct {
	CustomersId        uint64    `json:"customers_id"`
	CustomersEmail     string    `json:"customers_email"`
	CustomersCreatedAt time.Time `json:"customers_createdat"`
}

type DBTableGroups struct {
	GroupsId        uint64    `json:"groups_id"`
	GroupsToken     string    `json:"groups_token"`
	GroupsCreatedAt time.Time `json:"groups_createdat"`
}

type DBTableMessages struct {
	MessagesId         uint64    `json:"messages_id"`
	MessagesCustomerId uint64    `json:"messages_customer_id"`
	MessagesGroupId    uint64    `json:"messages_group_id"`
	MessagesTitle      string    `json:"messages_title"`
	MessagesContent    string    `json:"messages_content"`
	MessagesCreatedAt  time.Time `json:"messages_createdat"`
}
