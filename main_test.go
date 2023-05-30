package main

import ("testing")

type PasswordTest struct {
	arg1 string; expected int
}

var pwTest = []PasswordTest{
	{"aA1", 3},
	{"1445D1cd", 0},
	{"!ecse34dF", 0},
	{"aaa124kebndlfhb", 1},
	{"hello,test1444", 1},
	{"a", 5},
	{"WERWEDSAE314", 1},
	{"ABCDEFGHIJKLMNOPQRSTUvwxyz1", 8},
	{"1111111", 2},
}

func TestStrong(t *testing.T){

    for _, test := range pwTest{
		if output := strongPassword(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
