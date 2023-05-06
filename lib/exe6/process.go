package exe6

import (
	"strings"

	"chipcom/lib"
)

type Context struct {
	BattleChipList *[]*BattleChip
}

func createContext() (*Context, error) {
	context := &Context{}

	battleChipList, err := loadJson[[]*BattleChip]("BattleChip")
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
	standards := []*CsvBattleChip{}
	megas := []*CsvBattleChip{}
	gigas := []*CsvBattleChip{}
	secrets := []*CsvBattleChip{}

	codeMap := map[string][]string{}
	codes := []*CsvBattleChipCode{}

	for _, chip := range *context.BattleChipList {
		// クラス別分類
		switch chip.Class {
		case BattleChipClass.Standard:
			standards = append(standards, chip.toCsvBattleChip())
		case BattleChipClass.Mega:
			megas = append(megas, chip.toCsvBattleChip())
		case BattleChipClass.Giga:
			gigas = append(gigas, chip.toCsvBattleChip())
		case BattleChipClass.Secret:
			secrets = append(secrets, chip.toCsvBattleChip())
		}

		// チップコード別分類
		for _, code := range chip.CodeList {
			names, ok := codeMap[code]
			if ok {
				names = append(names, chip.Name)
				codeMap[code] = names
			} else {
				names = []string{chip.Name}
				codeMap[code] = names
			}
		}
	}

	for _, code := range lib.ChipCodeList {
		names, ok := codeMap[code]
		if ok {
			codes = append(codes, &CsvBattleChipCode{code, strings.Join(names, "\n"), len(names)})
		} else {
			codes = append(codes, &CsvBattleChipCode{code, "", 0})
		}
	}

	if err := writeCsv("バトルチップ/スタンダードチップ", &standards); err != nil {
		return err
	}
	if err := writeCsv("バトルチップ/メガクラスチップ", &megas); err != nil {
		return err
	}
	if err := writeCsv("バトルチップ/ギガクラスチップ", &gigas); err != nil {
		return err
	}
	if err := writeCsv("バトルチップ/シークレットチップ", &secrets); err != nil {
		return err
	}
	if err := writeCsv("バトルチップ/チップコード別一覧", &codes); err != nil {
		return err
	}

	return nil
}
