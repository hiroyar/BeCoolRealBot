package telegram

import "math/rand"

func getTimeForNotification() string {
	periods := []string{"13:00:00", "14:00:00", "15:00:00", "16:00:00", "17:00:00", "18:00:00"}
	randIndex := rand.Intn(len(periods))
	return periods[randIndex]
}
