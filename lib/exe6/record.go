package exe6

type CsvBattleChip struct {
	No         string `csv:"No"`
	Name       string `csv:"名称"`
	Attribute  string `csv:"属性"`
	Power      int    `csv:"攻撃力"`
	HP         int    `csv:"-"`
	Lifetime   int    `csv:"-"`
	EffectList string `csv:"特性"`
	Capacity   int    `csv:"容量"`
	Limit      int    `csv:"枚数制限"`
	Rarity     string `csv:"レア度"`
	CodeList   string `csv:"コード"`
}

type CsvBattleChipCode struct {
	Code           string `csv:"チップコード"`
	BattleChipList string `csv:"バトルチップ"`
	Count          int    `csv:"種類数"`
}

type CsvChipTrader struct {
	BattleChipName string `csv:"バトルチップ"`
	CodeList       string `csv:"チップコード"`
	Rarity         string `csv:"レア度"`
}

type CsvCrossoverTrader struct {
	BattleChipName string `csv:"バトルチップ"`
	CodeList       string `csv:"チップコード"`
	Rarity         string `csv:"レア度"`
	Probability    string `csv:"確率"`
}
