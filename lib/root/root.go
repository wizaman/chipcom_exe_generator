package root

import (
	"fmt"
	"os"

	"chipcom/lib"
	"chipcom/lib/exe6"
)

func GetTargets() []string {
	targets := []string{
		exe6.TargetName,
	}
	return targets
}

func Execute() error {
	targets := parseTargets()

	for _, target := range targets {
		switch target {
		case exe6.TargetName:
			if err := exe6.Generate(); err != nil {
				return err
			}
		default:
			fmt.Fprintf(os.Stderr, "Invalid target: %s\n", target)
		}
	}

	return nil
}

func parseTargets() []string {
	targets := []string{}

	// とりあえずターゲット名だけ与えられる想定
	// 必要になるまでパーサは導入しない
	for i, target := range os.Args {
		// 実行プログラムは無視
		if i == 0 {
			continue
		}

		if target == lib.AllTargetName {
			return GetTargets()
		}

		targets = append(targets, target)
	}

	return targets
}
