package util

import "testing"

func TestStringToDateTime(t *testing.T) {
	var result = StringToDateTime("2024-09-25T11:45:21")

	if result.Year() != 2024 {
		t.Errorf("Ano inválido")
	}
}
