package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// add items_offset parameter
func AddOffset(url string, offset int) string {
	if strings.Contains(url, "items_offset") {
		return url
	}
	if strings.Contains(url, "?") {
		return url + "&items_offset=" + strconv.Itoa(offset)
	}
	return url + "?items_offset=" + strconv.Itoa(offset)
}

// remove duplicates
func Unique(input []string) []string {
	keys := make(map[string]bool)
	var result []string
	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return result
}

// random sleep time
func RandomDelay() time.Duration {
	return time.Duration(rand.Intn(4)+2) * time.Second
}

// random message
var messages = []string{
	"Just a moment, gathering the data!",
	"Almost there, please hang on!",
	"Scraping in progress, stay with us!",
	"Loading your results, one more second...",
}
func GetRandomMessage() string {
	return messages[rand.Intn(len(messages))]
}
