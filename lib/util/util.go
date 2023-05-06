package util

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"chipcom/lib"
)

func LoadJson[T any](path string) (*T, error) {
	fmt.Printf("Loading: %v", path)
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

func CreateInputJsonPath(paths ...string) string {
	paths = append([]string{lib.InputJsonRoot}, paths...)
	return strings.Join(paths, "/") + ".json"
}

func CreateOutputCsvPath(paths ...string) string {
	paths = append([]string{lib.OutputCsvRoot}, paths...)
	return strings.Join(paths, "/") + ".csv"
}
