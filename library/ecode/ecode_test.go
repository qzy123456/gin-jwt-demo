package ecode

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestEqualError(t *testing.T) {
	fmt.Println(errors.Is(LoginConflict, testErr()))
	fmt.Println(LoginConflict)
	fmt.Println(LoginConflict.Message())
	fmt.Println(LoginConflict.Code())
	fmt.Println(LoginConflict.Message())
	fmt.Println(LoginConflict.Details())
}

func testErr() error {
	return LoginConflict
}