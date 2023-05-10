package exe6

import (
	"fmt"
	"strconv"
	"strings"
)

func toCsvAttribute(attribute string) string {
	switch attribute {
	case Attribute.Fire:
		return "炎"
	case Attribute.Water:
		return "水"
	case Attribute.Elec:
		return "電気"
	case Attribute.Wood:
		return "木"
	case Attribute.Sword:
		return "ソード"
	case Attribute.Break:
		return "ブレイク"
	case Attribute.Cursor:
		return "カーソル"
	case Attribute.Wind:
		return "風"
	case Attribute.Object:
		return "置き物"
	case Attribute.Number:
		return "数値付加"
	}

	return "無"
}

func toCsvEffect(effect string) string {
	// TODO: 変換テーブルを外部化したい
	switch effect {
	case AdditionalEffect.KnockBack:
		return "[の]"
	case AdditionalEffect.Invisible:
		return "[無]"
	case AdditionalEffect.Blow:
		return "[吹]"
	case AdditionalEffect.Break:
		return "[ブ]"
	case AdditionalEffect.Sword:
		return "[ソ]"
	case AdditionalEffect.AntiInvisible:
		return "[イ]"
	case AdditionalEffect.AntiGround:
		return "[地]"
	case AdditionalEffect.Paralysis:
		return "[マ]"
	case AdditionalEffect.Bubble:
		return "[泡]"
	case AdditionalEffect.Confusion:
		return "[混]"
	case AdditionalEffect.Bind:
		return "[移]"
	case AdditionalEffect.Blind:
		return "[盲]"
	case AdditionalEffect.Bug:
		return "[バ]"
	}
	return ""
}

func (p *BattleChip) getCsvNo() string {
	if p.No >= 900 {
		return "-"
	}
	no := strconv.Itoa(p.No)
	return no
}

func (p *BattleChip) getCsvName() string {
	no := p.Name
	switch p.Version {
	case GameVersion.Gregar:
		no += "\n(グレイガ版)"
	case GameVersion.Falzar:
		no += "\n(ファルザー版)"
	case GameVersion.Gate:
		no += "\n(ゲート限定)"
	}
	return no
}

func (p *BattleChip) getCsvAttribute() string {
	return toCsvAttribute(p.Attribute)
}

func (p *BattleChip) getCsvEffectList() string {
	effects := []string{}

	if p.Event {
		effects = append(effects, "[暗]")
	}

	for _, effect := range p.EffectList {
		s := toCsvEffect(effect)
		if s != "" {
			effects = append(effects, s)
		}
	}

	return strings.Join(effects, "\n")
}

func (p *BattleChip) getCsvRarity() string {
	// return strings.Repeat("☆", p.Rarity)
	if p.Rarity > 0 {
		return "☆" + strconv.Itoa(p.Rarity)
	}
	return "-"
}

func (p *BattleChip) getCsvCodeList() string {
	return strings.Join(p.CodeList, "")
}

func (p *BattleChip) toCsvBattleChip() *CsvBattleChip {
	return &CsvBattleChip{
		p.getCsvNo(),
		p.getCsvName(),
		p.getCsvAttribute(),
		p.Power,
		p.HP,
		p.Lifetime,
		p.getCsvEffectList(),
		p.Capacity,
		p.Limit,
		p.getCsvRarity(),
		p.getCsvCodeList(),
	}
}

func (p *TraderBattleChip) getCsvCodeList() string {
	return strings.Join(p.CodeList, "")
}

func (p *TraderBattleChip) toCsvChipTrader() *CsvChipTrader {
	return &CsvChipTrader{
		p.Name,
		p.getCsvCodeList(),
		p.Rarity,
	}
}

func (p *TraderBattleChip) toCsvCrossoverTrader(count int, total int) *CsvCrossoverTrader {
	return &CsvCrossoverTrader{
		p.Name,
		p.getCsvCodeList(),
		p.Rarity,
		fmt.Sprintf("%d/%d", count, total),
	}
}

func (p *Trader) toCsvChipTrader() *[]*CsvChipTrader {
	traderChips := []*CsvChipTrader{}

	for _, chip := range *p.BattleChipList {
		traderChips = append(traderChips, chip.toCsvChipTrader())
	}

	return &traderChips
}

func (p *Trader) toCsvCrossoverTrader() *[]*CsvCrossoverTrader {
	traderChips := []*CsvCrossoverTrader{}

	isVariant10 := p.Variant == "10"
	total := len(*p.BattleChipList)
	if isVariant10 {
		// 6種だが2種が倍率2倍なので総数を補正
		total += 2
	}

	for _, chip := range *p.BattleChipList {
		count := 1
		if isVariant10 {
			// 特定チップの排出率が2倍
			if chip.Name == "ジャンゴ" || chip.Name == "ハクシャクEX" {
				count = 2
			}
		}
		traderChips = append(traderChips, chip.toCsvCrossoverTrader(count, total))
	}

	return &traderChips
}
