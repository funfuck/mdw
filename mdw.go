package mdw

import (
	"fmt"

	"github.com/funfuck/pii"
)

func NewMDW(f func(interface{}) (string, error)) func() error {
	return func() error {

		str, err := f(nil)
		fmt.Printf("str: %v\n", str)
		fmt.Printf("err: %v\n", err)

		s, e := pii.Any(nil)
		fmt.Printf("s: %v\n", s)
		fmt.Printf("e: %v\n", e)

		return err
	}
}
