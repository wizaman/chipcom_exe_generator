package exe6

import (
	"fmt"

	"chipcom/lib/util"
)

const TargetName = "exe6"
const OutputRoot = "ロックマンエグゼ6"

func Generate() error {
	fmt.Println(">>> " + TargetName)

	context, err := createContext()
	if err != nil {
		return err
	}

	return process(context)
}

func loadJson[T any](name string) (*T, error) {
	path := util.CreateInputJsonPath(TargetName, name)
	return util.LoadJson[T](path)
}

func writeCsv[T any](name string, data *T) error {
	path := util.CreateOutputCsvPath(OutputRoot, name)
	return util.WriteCsv(path, data)
}
