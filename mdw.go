package mdw

import "fmt"

func NewMDW(f func(interface{}) (string, error)) func() error {
	return func() error {

		str, err := f(nil)
		fmt.Printf("str: %v\n", str)
		fmt.Printf("err: %v\n", err)

		return err
	}
}
