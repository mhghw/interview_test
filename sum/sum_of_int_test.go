package sum_test

import (
	"int_test/sum"
	"testing"
)

type TestCase struct {
	str       string
	actualSum int64
}

var TestCases = []TestCase{
	{"iK9cvyY0kaqwpWf8", 17},
	{"J5aclE8PURdMcirE", 13},
	{"fBCsJDlmDlTD32Sm", 5},
	{"LJh5Eo0md7u45Uxg", 21},
	{"XABVsVTiXXW60Jzk", 6},
	{"6u3HrEB7IjUD7rJK", 23},
	{"rOpr9e1CMEnaw0u0", 10},
}

func TestSumIntegers(t *testing.T) {
	t.Parallel()
	for _, testCase := range TestCases {
		if s := sum.SumIntegers(testCase.str); s != testCase.actualSum {
			t.Errorf("For string: %v expected %v but got %v", testCase.str, testCase.actualSum, s)
		}
	}
}
