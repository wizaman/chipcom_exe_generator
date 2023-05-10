package exe6

type GameVersionDef struct {
	Gregar string
	Falzar string
	Gate   string
}

var GameVersion = GameVersionDef{
	"G",
	"F",
	"X",
}

type AttributeDef struct {
	Normal string
	Fire   string
	Water  string
	Elec   string
	Wood   string
	Sword  string
	Break  string
	Cursor string
	Wind   string
	Object string
	Number string
}

var Attribute = AttributeDef{
	"Normal",
	"Fire",
	"Water",
	"Elec",
	"Wood",
	"Sword",
	"Break",
	"Cursor",
	"Wind",
	"Object",
	"Number",
}

type BattleChipClassDef struct {
	Standard string
	Mega     string
	Giga     string
	Secret   string
}

var BattleChipClass = BattleChipClassDef{
	"Standard",
	"Mega",
	"Giga",
	"Secret",
}

type AdditionalEffectDef struct {
	KnockBack     string
	Invisible     string
	Blow          string
	Sword         string
	Break         string
	AntiInvisible string
	AntiGround    string
	Paralysis     string
	Bubble        string
	Confusion     string
	Bind          string
	Blind         string
	Bug           string
	PanelNormal   string
	PanelCrack    string
	PanelHole     string
	PanelGreen    string
	PanelIce      string
	PanelPoison   string
	PanelHoly     string
	PanelComing   string
	PanelGoing    string
}

var AdditionalEffect = AdditionalEffectDef{
	"KnockBack",
	"Invisible",
	"Blow",
	"Sword",
	"Break",
	"AntiInvisible",
	"AntiGround",
	"Paralysis",
	"Bubble",
	"Confusion",
	"Bind",
	"Blind",
	"Bug",
	"PanelNormal",
	"PanelCrack",
	"PanelHole",
	"PanelGreen",
	"PanelIce",
	"PanelPoison",
	"PanelHoly",
	"PanelComing",
	"PanelGoing",
}

type TraderTypeDef struct {
	Normal    string
	Special   string
	Bug       string
	Crossover string
}

var TraderType = TraderTypeDef{
	"Normal",
	"Special",
	"Bug",
	"Crossover",
}

type TraderVariantDef struct {
	AsterLand   string
	SkyTown     string
	Akiharacho  string
	GreenTown   string
	Crossover10 string
	Crossover20 string
	Crossover30 string
	Crossover50 string
}

var TraderVariant = TraderVariantDef{
	"アスタランド",
	"スカイタウン",
	"秋原町",
	"グリーンタウン",
	"10",
	"20",
	"30",
	"50",
}

type BattleChip struct {
	ID         string
	SortId     int
	Class      string
	No         int
	Version    string
	Name       string
	Attribute  string
	Power      int
	HP         int
	Lifetime   int
	Event      bool
	EffectList []string
	Capacity   int
	Limit      int
	Rarity     int
	CodeList   []string
}

type TraderBattleChip struct {
	ID       string
	Name     string
	Rarity   int
	CodeList []string
}

type Trader struct {
	ID             string
	Type           string
	Variant        string
	HasAsterisk    bool
	RarityCount    *map[string]int
	BattleChipList *[]TraderBattleChip
}
