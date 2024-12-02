package jpweatherwarnings_test

import (
	"testing"

	"lambda/jpweatherwarnings"
)

func TestAreaCode_IsValid(t *testing.T) {
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

func TestAreaCode_IsPrefectureAreaCode(t *testing.T) {
	tests := map[string]struct {
		areaCodeText string
		result       bool
	}{
		"正常":  {"060003", true},
		"市町村": {"062014", false},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got, expected := jpweatherwarnings.AreaCode(test.areaCodeText).IsPrefectureAreaCode(), test.result; got != expected {
				t.Errorf("AreaCode(%q).IsPrefectureAreaCode() returned %v; expected %v", test.areaCodeText, got, expected)
			}
		})
	}
}

func TestAreaCode_CodeForAPI(t *testing.T) {
	areaCode := jpweatherwarnings.AreaCode("062014")
	expected := "0620100"
	if got := areaCode.CodeForAPI(); got != expected {
		t.Errorf("AreaCode(%q).CodeForAPI() returned %v; expected %v", string(areaCode), got, expected)
	}
}

func TestAreaCode_PrefectureCodeForAPI(t *testing.T) {
	areaCode := jpweatherwarnings.AreaCode("062014")
	expected := "060000"
	if got := areaCode.PrefectureCodeForAPI(); got != expected {
		t.Errorf("AreaCode(%q).PrefectureCodeForAPI() returned %v; expected %v", string(areaCode), got, expected)
	}
}
