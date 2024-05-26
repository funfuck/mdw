package mdw

import (
	"fmt"

	"github.com/funfuck/pii"
)

func NewMDW() func() error {
	return func() error {

		str, err := pii.Any(nil)
		fmt.Printf("str: %v\n", str)
		fmt.Printf("err: %v\n", err)

		return err
	}
}
