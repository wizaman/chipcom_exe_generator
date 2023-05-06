package exe6

import (
	"chipcom/lib/util"
)

type Context struct {
	BattleChipList *[]BattleChip
}

func loadJson[T any](name string) (*T, error) {
	path := util.CreateInputJsonPath(TargetName, name)
	return util.LoadJson[T](path)
}

func loadContext() (*Context, error) {
	context := &Context{}

	battleChipList, err := loadJson[[]BattleChip]("BattleChip")
	if err != nil {
		return nil, err
	}
	context.BattleChipList = battleChipList

	return context, nil
}

func process(context *Context) error {
	if err := processBattleChip(context); err != nil {
		return err
	}

	return nil
}

func processBattleChip(context *Context) error {
	return nil
}
