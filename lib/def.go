package lib

import (
	"chipcom/lib/exe6"
)

const AllTargetName = "all"

func GetTargets() []string {
	targets := []string{
		exe6.TargetName,
	}
	return targets
}
