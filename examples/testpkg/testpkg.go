package testpkg

import (
	"fmt"
	"math/rand"
)

func AlwaysSucceed() error {
	return nil
}

func AlwaysFailing() error {
	return fmt.Errorf("hope is the key to life")
}

func MaySucceed() error {
	if rand.Intn(3) != 0 {
		return fmt.Errorf("oops, no luck, try again")
	}
	return nil
}
