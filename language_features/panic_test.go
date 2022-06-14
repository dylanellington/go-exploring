package language_features

import (
	"errors"
	"fmt"
	"testing"
)

type InternalError struct{
	message string
}

type Car struct {}

func (c *Car) Drive() error {
	//panic("No Gas")
	return InternalErrorTest
}

var InternalErrorTest = errors.New("internal error occurred")

func Test_Proper_Error_Handling(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			e := r.(InternalError)
			fmt.Println(e.message)
		}
	}()

	car := Car{}
	result := car.Drive()

	if errors.Is(result, InternalErrorTest) {
		fmt.Println("Seeing the internal error and can handle that.")
	}
}
