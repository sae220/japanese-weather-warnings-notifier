package main

import (
	"regexp"
)

// 全国地方公共団体コード
type AreaCode string

// 全国地方公共団体コードが有効か確認する
func (areaCode AreaCode) IsValid() bool {
	// 6文字の数字
	if match, _ := regexp.MatchString(`^\d{6}$`, string(areaCode)); !match {
		return false
	}
	// 検査数字の検査
	return areaCode.hasValidCheckDigit()
}

// 全国地方公共団体コードの検査数字を検査する
//
// 第１桁から第５桁までの数字に、それぞれ６．５．４．３．２を乗じて算出した積の和を求め、その和を１１で除し、商と剰余（以下「余り数字」という。）を求めて、１１と余り数字との差の下１桁の数字を検査数字とする（全国地方公共団体コード仕様より）
func (areaCode AreaCode) hasValidCheckDigit() bool {
	checkDigit := int(areaCode[5] - '0')
	// 気象庁で用いられている検査数字がないものの場合は良しとする
	if checkDigit == 0 {
		return true
	}

	var calculatedDigit int
	for i := 0; i < 5; i++ {
		calculatedDigit += (6 - i) * int(areaCode[i]-'0')
	}
	calculatedDigit = (11 - calculatedDigit%11) % 10

	return calculatedDigit == checkDigit
}

// 地方公共団体が位置している都道府県のコードを取得する
func (areaCode AreaCode) PrefectureAreaCodeWithoutCheckDigit() AreaCode {
	var prefectureAreaCode string
	for i := 0; i < 5; i++ {
		if i < 2 {
			prefectureAreaCode += string(areaCode[i])
		} else {
			prefectureAreaCode += "0"
		}
	}
	return AreaCode(prefectureAreaCode)
}
