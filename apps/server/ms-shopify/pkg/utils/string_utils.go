package utils

import (
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"time"
)

func RandomWordsCode(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.NewSource(time.Now().UnixNano()) // Seed the random number generator

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func ParseIDToUUID(id string) uuid.UUID {
	parsedUuid, _ := uuid.Parse(id)
	return parsedUuid
}

func ParseIdToInt(id string) int {
	parsedId, _ := strconv.Atoi(id)
	return parsedId
}
