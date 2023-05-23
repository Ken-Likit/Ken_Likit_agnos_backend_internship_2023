package main

import "testing"

type PasswordTest struct {
	arg1 string; expected int
}

var pwTest = []PasswordTest{
	PasswordTest{"aA1", 3},
	PasswordTest{"1445D1cd", 0},
	PasswordTest{"!ecse34dF", 0},
	PasswordTest{"aaa124kebndlfhb", 2},
	PasswordTest{"hello,test1444", 2},
	PasswordTest{"a", 7},
	PasswordTest{"WERWEDSAE314", 1},
	PasswordTest{"ABCDEFGHIJKLMNOPQRSTUvwxyz1", 8},
}

func strongTest(t *testing.T) {

	for _, test := range pwTest{
		if output := strongPassword(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
