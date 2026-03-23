package utils

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"strings"
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

func GenerateDisplayName() string {
	fruits := []string{
		"Apple", "Banana", "Cherry", "Date", "Elderberry", "Fig",
		"Grape", "Honeydew", "Kiwi", "Lemon", "Mango",
		"Nectarine", "Orange", "Papaya", "Quince", "Raspberry",
		"Strawberry", "Tangerine", "Watermelon",
	}

	combination := make([]string, 2)
	for i := 0; i < 2; i++ {
		combination[i] = fruits[rand.Intn(len(fruits))]
	}

	number := fmt.Sprintf("%d", rand.Intn(9999))
	combinationName := strings.Join(combination, "_")

	return combinationName + "" + number
}

func ParseDaysString(s string) (time.Duration, error) {
	s = strings.TrimSpace(s)
	parts := strings.Fields(s) // Split by whitespace

	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid format: expected 'x days', got '%s'", s)
	}

	if parts[1] != "days" && parts[1] != "day" {
		return 0, fmt.Errorf("invalid unit: expected 'days' or 'day', got '%s'", parts[1])
	}

	days, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number '%s': %w", parts[0], err)
	}

	return time.Duration(days * 24 * float64(time.Hour)), nil
}
