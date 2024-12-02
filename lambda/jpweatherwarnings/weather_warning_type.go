package jpweatherwarnings

// 気象警報・注意報のコード
type WeatherWarningCode string

// 気象警報・注意報の種類
type WeatherWarningType struct {
	Name1 string
	Name2 string
	Elem  string
	Level int
}

// 気象警報・注意報の表示名
func (weatherWarningType WeatherWarningType) Name() string {
	return weatherWarningType.Name1 + weatherWarningType.Name2
}

// 気象警報・注意報の種類の一覧
var WeatherWarningTypeByCode = map[WeatherWarningCode]WeatherWarningType{
	"33":  {"大雨", "特別警報", "rain", 50},
	"03":  {"大雨", "警報", "rain", 30},
	"10":  {"大雨", "注意報", "rain", 20},
	"04":  {"洪水", "警報", "flood", 30},
	"18":  {"洪水", "注意報", "flood", 20},
	"35":  {"暴風", "特別警報", "wind", 40},
	"05":  {"暴風", "警報", "wind", 30},
	"15":  {"強風", "注意報", "wind", 20},
	"32":  {"暴風雪", "特別警報", "wind_snow", 40},
	"02":  {"暴風雪", "警報", "wind_snow", 30},
	"13":  {"風雪", "注意報", "wind_snow", 20},
	"36":  {"大雪", "特別警報", "snow", 40},
	"06":  {"大雪", "警報", "snow", 30},
	"12":  {"大雪", "注意報", "snow", 20},
	"37":  {"波浪", "特別警報", "wave", 40},
	"07":  {"波浪", "警報", "wave", 30},
	"16":  {"波浪", "注意報", "wave", 20},
	"38":  {"高潮", "特別警報", "tide", 40},
	"08":  {"高潮", "警報", "tide", 40},
	"19":  {"高潮", "注意報", "tide", 20},
	"19+": {"高潮", "注意報", "tide", 30},
	"14":  {"雷", "注意報", "thunder", 20},
	"17":  {"融雪", "注意報", "snow_melting", 20},
	"20":  {"濃霧", "注意報", "fog", 20},
	"21":  {"乾燥", "注意報", "dry", 20},
	"22":  {"なだれ", "注意報", "avalanche", 20},
	"23":  {"低温", "注意報", "cold", 20},
	"24":  {"霜", "注意報", "frost", 20},
	"25":  {"着氷", "注意報", "ice_accretion", 20},
	"26":  {"着雪", "注意報", "snow_accretion", 20},
}
