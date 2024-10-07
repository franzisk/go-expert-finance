package util

import "time"

// 2006 faz referência a um formato de ano em datas
// 01 faz referência a um formato de mês em datas
// 02 faz referência a um formato de dia em datas
// 15 faz referência a um formato de horas 0-24 em datas
// 04 faz referência a um formato de minutos em datas
// 05 faz referência a um formato de segundos em datas
const layout = "2006-01-02T15:04:05"

// StringToDateTime é usada para converter uma string de data/hora para Time
func StringToDateTime(value string) time.Time {
	dateTime, _ := time.Parse(layout, value)
	return dateTime
}
