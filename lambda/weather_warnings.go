package main

// ある地域に出ている気象警報・注意報一つ
type AreaWeatherWarning struct {
	// 気象警報・注意報コード
	code string
	// 気象警報・注意報の状態
	status string
}

// ある地域に出ている気象警報・注意報の配列
type AreaWeatherWarnings []AreaWeatherWarning

// 地域コードに対応する地域に出ている気象警報・注意報を取得する
func FetchAreaWeatherWarnings(areaCode string) (AreaWeatherWarnings, error) {
	return nil, nil
}