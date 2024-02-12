package cfg

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
)

const FileExtYaml = ".yaml"
const FileExtYml = ".yml"

var AllowedFileTypes = []string{FileExtYaml, FileExtYml}

func Read(cfg interface{}, path string) error {
	return readConfigFile(cfg, path)
}

func readConfigFile(cfg interface{}, path string) error {
	ext := filepath.Ext(path)
	if valid := validateFileExt(ext); !valid {
		return fmt.Errorf("unsupported file format: %s", ext)
	}

	return parseFile(cfg, path)
}

func parseFile(cfg interface{}, path string) error {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_SYNC, 0)
	if err != nil {
		return err
	}

	defer func() {
		if er := f.Close(); er != nil {
			err = er
		}
	}()

	return decodeYAML(f, cfg)
}

func decodeYAML(r io.Reader, cfg interface{}) error {
	return yaml.NewDecoder(r).Decode(cfg)
}

func validateFileExt(val string) bool {
	for _, item := range AllowedFileTypes {
		if item == val {
			return true
		}
	}

	return false
}
