package exe6

import (
	"errors"
	"fmt"
	"strings"

	"chipcom/lib"
)

type Context struct {
	BattleChipList *[]*BattleChip
	TraderList     *[]*Trader
}

func createContext() (*Context, error) {
	context := &Context{}

	battleChipList, err := loadJson[[]*BattleChip]("BattleChip")
	if err != nil {
		return nil, err
	}
	context.BattleChipList = battleChipList

	traderList, err := loadJson[[]*Trader]("Trader")
	if err != nil {
		return nil, err
	}
	context.TraderList = traderList

	return context, nil
}

func process(context *Context) error {
	if err := processBattleChip(context); err != nil {
		return err
	}

	if err := processTrader(context); err != nil {
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

func processTrader(context *Context) error {
	normal := []*CsvChipTrader{}
	spSky := []*CsvChipTrader{}
	spAkihara := []*CsvChipTrader{}
	spGreen := []*CsvChipTrader{}
	bug := []*CsvChipTrader{}
	crossover10 := []*CsvCrossoverTrader{}
	crossover20 := []*CsvCrossoverTrader{}
	crossover30 := []*CsvCrossoverTrader{}
	crossover50 := []*CsvCrossoverTrader{}

	for _, trader := range *context.TraderList {
		invalid := false

		// 種別ごとの分類
		switch trader.Type {
		case TraderType.Normal:
			normal = *trader.toCsvChipTrader()
		case TraderType.Special:
			switch trader.Variant {
			case TraderVariant.SkyTown:
				spSky = *trader.toCsvChipTrader()
			case TraderVariant.Akiharacho:
				spAkihara = *trader.toCsvChipTrader()
			case TraderVariant.GreenTown:
				spGreen = *trader.toCsvChipTrader()
			default:
				invalid = true
			}
		case TraderType.Bug:
			bug = *trader.toCsvChipTrader()
		case TraderType.Crossover:
			switch trader.Variant {
			case TraderVariant.Crossover10:
				crossover10 = *trader.toCsvCrossoverTrader()
			case TraderVariant.Crossover20:
				crossover20 = *trader.toCsvCrossoverTrader()
			case TraderVariant.Crossover30:
				crossover30 = *trader.toCsvCrossoverTrader()
			case TraderVariant.Crossover50:
				crossover50 = *trader.toCsvCrossoverTrader()
			default:
				invalid = true
			}
		default:
			invalid = true
		}

		if invalid {
			return errors.New(fmt.Sprintf("Invalid Trader (Type: {%v}, Variant{%v})", trader.Type, trader.Variant))
		}
	}

	if err := writeCsv("チップトレーダー/チップトレーダー(アスタランド)", &normal); err != nil {
		return err
	}
	if err := writeCsv("チップトレーダー/チップトレーダーSP(スカイタウン)", &spSky); err != nil {
		return err
	}
	if err := writeCsv("チップトレーダー/チップトレーダーSP(秋原町)", &spAkihara); err != nil {
		return err
	}
	if err := writeCsv("チップトレーダー/チップトレーダーSP(グリーンタウン)", &spGreen); err != nil {
		return err
	}
	if err := writeCsv("チップトレーダー/バグピーストレーダー", &bug); err != nil {
		return err
	}
	if err := writeCsv("チップトレーダー/シンボクトレーダー(10)", &crossover10); err != nil {
		return err
	}
	if err := writeCsv("チップトレーダー/シンボクトレーダー(20)", &crossover20); err != nil {
		return err
	}
	if err := writeCsv("チップトレーダー/シンボクトレーダー(30)", &crossover30); err != nil {
		return err
	}
	if err := writeCsv("チップトレーダー/シンボクトレーダー(50)", &crossover50); err != nil {
		return err
	}

	return nil
}
