package helpers

import "strconv"

func FromStringToInt64(text string) int64 {
	chatId, _ := strconv.Atoi(text)
	return int64(chatId)
}
