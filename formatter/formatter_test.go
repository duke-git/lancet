package formatter

import (
	"github.com/duke-git/lancet/utils"
	"testing"
)

func TestComma(t *testing.T) {
	comma(t, "", "","")
	comma(t, "aa", "","")
	comma(t, "123", "","123")
	comma(t, "12345", "","12,345")
	comma(t, 12345, "","12,345")
	comma(t, 12345, "$","$12,345")
	comma(t, 12345, "¥","¥12,345")
	comma(t, 12345.6789, "","12,345.6789")
}

func comma(t *testing.T, test interface{}, symbol string, expected interface{})  {
	res:= Comma(test, symbol)
	if res != expected {
		utils.LogFailedTestInfo(t, "Comma", test, expected, res)
		t.FailNow()
	}
}