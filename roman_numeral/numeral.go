package roman_numeral

import "strings"

type romanNumeral struct {
	Value int
	Roman string
}

var allRomanNumerals = []romanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Roman)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	arabic := 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Roman) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Roman)
		}
	}

	return arabic
}
