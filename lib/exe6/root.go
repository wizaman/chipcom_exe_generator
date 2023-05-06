package exe6

import "fmt"

func Generate() error {
	fmt.Println(">>> " + TargetName)

	context, err := loadContext()
	if err != nil {
		return err
	}

	return process(context)
}
