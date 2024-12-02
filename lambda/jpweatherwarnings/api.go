package jpweatherwarnings

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// 気象警報・注意報を取得するAPIのレスポンスの形式
type (
	WeatherWarningApiResponse struct {
		AreaTypes AreaTypes `json:"areaTypes"`
	}

	AreaTypes []AreaType

	AreaType struct {
		Areas Areas `json:"areas"`
	}

	Areas []Area

	Area struct {
		Code     string   `json:"code"`
		Warnings Warnings `json:"warnings"`
	}

	Warnings []Warning

	Warning struct {
		Code   WeatherWarningCode `json:"code"`
		Status string             `json:"status"`
	}
)

// APIのレスポンスのAreaType
const (
	Prefecture = iota
	Cities
)

func (warnings Warnings) String() (text string) {
	for _, warning := range warnings {
		text += fmt.Sprintf("%s %s\n", WeatherWarningTypeByCode[warning.Code].Name(), warning.Status)
	}
	return
}

// 全国地方公共団体コードに対応する地域に出ている気象警報・注意報を取得する
func FetchAreaWeatherWarnings(areaCode AreaCode) (Warnings, error) {
	if !areaCode.IsValid() {
		return nil, errors.New("全国地方公共団体コードが無効です。https://www.soumu.go.jp/denshijiti/code.htmlを確認してください")
	}
	if areaCode.IsPrefectureAreaCode() {
		return nil, errors.New("都道府県ではなく市町村のコードを入力してください")
	}

	url := fmt.Sprintf("https://www.jma.go.jp/bosai/warning/data/warning/%s.json", areaCode.PrefectureCodeForAPI())
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTPリクエスト時にエラーが発生しました: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTPリクエストのレスポンスが%dでした", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// JSONにする
	var weatherWarnings WeatherWarningApiResponse
	err = json.Unmarshal(body, &weatherWarnings)
	if err != nil {
		return nil, fmt.Errorf("HTTPレスポンスをJSONに変換できませんでした: %s", err)
	}

	// 該当する地方公共団体コードのものだけ取り出す
	for _, area := range weatherWarnings.AreaTypes[Cities].Areas {
		if area.Code == areaCode.CodeForAPI() {
			return area.Warnings, nil
		}
	}
	return nil, errors.New("該当する地方公共団体コードのデータが見つかりませんでした")
}
