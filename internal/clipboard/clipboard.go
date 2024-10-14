package clipboard

import (
	"fmt"

	"golang.design/x/clipboard"
)

func Init() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r) //nolint:goerr113
		}
	}()

	err = clipboard.Init()
	return err
}

func WriteText(value string) <-chan struct{} {
	return clipboard.Write(clipboard.FmtText, []byte(value))
}
