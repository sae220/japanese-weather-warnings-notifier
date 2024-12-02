package jpweatherwarnings_test

import (
	"testing"

	"lambda/jpweatherwarnings"
)

func TestAreaCodeIsValid(t *testing.T) {
	tests := map[string]struct {
		areaCodeText string
		result       bool
	}{
		"正常":     {"062014", true},
		"文字数違い":  {"0620144", false},
		"数字以外":   {"abcdef", false},
		"検査数字違い": {"062016", false},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got, expected := jpweatherwarnings.AreaCode(test.areaCodeText).IsValid(), test.result; got != expected {
				t.Errorf("AreaCode(%q).IsValid() returned %v; expected %v", test.areaCodeText, got, expected)
			}
		})
	}
}
