package panic

import (
	"errors"
	"fmt"
	"testing"
)

func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("finally", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("error occurs"))
}
