package utils

import (
	"log"
	"strings"
)

func FormatPrice(price string, decimals int) string {
	/*
		if decimal (dec) >= len(price) => add the missing dec to the front
		else add the dec in specificed position
	*/
	builder := strings.Builder{}
	runeArr := make([]rune, 0)
	log.Println("***:", price)

	if len(price) <= decimals {
		log.Println("less:")
		diff := decimals - len(price)
		builder.WriteString("0.")

		for diff > 0 {
			builder.WriteString("0")
			diff--
		}
		builder.WriteString(price)
	} else {
		log.Println("more:")
		diff := len(price) - decimals
		runeArr = []rune(price)
		left := runeArr[:diff]
		right := runeArr[diff:]

		builder.WriteString(string(left))
		builder.WriteString(".")
		builder.WriteString(string(right))
	}
	runeArr = []rune(builder.String())

	for end := len(runeArr) - 1; end >= 0; end-- {
		if runeArr[end] != '0' {
			runeArr = runeArr[:end+1]
			break
		}
	}
	log.Println(string(runeArr))
	return string(runeArr)
}
