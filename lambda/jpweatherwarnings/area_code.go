// 全国地方公共団体コードのバリデーションと気象警報・注意報を取得するために必要なコードへの変換を行う
//
// 全国地方公共団体コードについては[総務省｜地方行政のデジタル化｜全国地方公共団体コード](https://www.soumu.go.jp/denshijiti/code.html)を参照

package jpweatherwarnings

import (
	"regexp"
)

// 全国地方公共団体コード
type AreaCode string

// 全国地方公共団体コードが有効か確認する
func (areaCode AreaCode) IsValid() bool {
	// 6桁の数字かつ検査数字が正しい
	return areaCode.isSixDigits() && areaCode.hasValidCheckDigit()
}

// 全国地方公共団体コードが6桁の数字か確認する
func (areaCode AreaCode) isSixDigits() bool {
	match, _ := regexp.MatchString(`^\d{6}$`, string(areaCode))
	return match
}

// 全国地方公共団体コードの検査数字を検査する
//
// 第１桁から第５桁までの数字に、それぞれ６．５．４．３．２を乗じて算出した積の和を求め、その和を１１で除し、商と剰余（以下「余り数字」という。）を求めて、１１と余り数字との差の下１桁の数字を検査数字とする（全国地方公共団体コード仕様より）
func (areaCode AreaCode) hasValidCheckDigit() bool {
	checkDigit := int(areaCode[5] - '0')

	var calculatedDigit int
	for i := 0; i < 5; i++ {
		calculatedDigit += (6 - i) * int(areaCode[i]-'0')
	}
	calculatedDigit = (11 - calculatedDigit%11) % 10

	return calculatedDigit == checkDigit
}

// 都道府県のコードか確認する
func (areaCode AreaCode) IsPrefectureAreaCode() bool {
	return string(areaCode[2:5]) == "000"
}

// 気象警報・注意報を取得するために必要なコードを取得する
func (areaCode AreaCode) CodeForAPI() string {
	return string(areaCode[:5]) + "00"
}

// 気象警報・注意報を取得するために必要な都道府県コードを取得する
func (areaCode AreaCode) PrefectureCodeForAPI() string {
	return string(areaCode[:2]) + "0000"
}
