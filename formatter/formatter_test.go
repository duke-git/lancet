package formatter

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestComma(t *testing.T) {
	comma(t, "", "", "")
	comma(t, "aa", "", "")
	comma(t, "aa.a", "", "")
	comma(t, []int{1}, "", "")
	comma(t, "123", "", "123")
	comma(t, "12345", "", "12,345")
	comma(t, 12345, "", "12,345")
	comma(t, 12345, "$", "$12,345")
	comma(t, 12345, "¥", "¥12,345")
	comma(t, 12345.6789, "", "12,345.6789")
}

func comma(t *testing.T, test interface{}, symbol string, expected interface{}) {
	res := Comma(test, symbol)
	if res != expected {
		internal.LogFailedTestInfo(t, "Comma", test, expected, res)
		t.FailNow()
	}
}
