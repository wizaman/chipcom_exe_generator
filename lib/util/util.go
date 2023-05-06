package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocarina/gocsv"

	"chipcom/lib"
)

func CreateInputJsonPath(paths ...string) string {
	paths = append([]string{lib.InputJsonRoot}, paths...)
	return strings.Join(paths, "/") + ".json"
}

func CreateOutputCsvPath(paths ...string) string {
	paths = append([]string{lib.OutputCsvRoot}, paths...)
	return strings.Join(paths, "/") + ".csv"
}

func GetAbsPath(path string) (string, error) {
	return filepath.Abs(filepath.Clean(path))
}

func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func LoadJson[T any](path string) (*T, error) {
	path, err := GetAbsPath(path)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Loading: %v\n", path)
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	var value T
	decoder := json.NewDecoder(fp)
	if err := decoder.Decode(&value); err != nil {
		return nil, err
	}

	return &value, nil
}

func WriteCsv[T any](path string, data *T) error {
	path, err := GetAbsPath(path)
	if err != nil {
		return nil
	}

	// 親ディレクトリがなければ作る
	parent := filepath.Dir(path)
	if !IsDir(parent) {
		err := os.MkdirAll(parent, os.ModePerm)
		if err != err {
			return err
		}
	}

	fmt.Printf("Writing: %v\n", path)
	fp, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fp.Close()

	return gocsv.MarshalFile(data, fp)
}
