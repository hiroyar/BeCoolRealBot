package helpers

import (
	"log"
	"strconv"
)

func FromStringToInt64(text string) int64 {
	str, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal("Ошибка перевода string в int64")
	}

	return int64(str)
}

func FromStringToInt(text string) int {
	str, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal("Ошибка перевода string в int")
	}
	return str
}

func FromStringToBool(text string) bool {
	boolean, err := strconv.ParseBool(text)
	if err != nil {
		log.Fatal("Ошибка перевода string в bool")
	}

	return boolean
}
