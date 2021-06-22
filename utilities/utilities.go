package utilities

import (
	"fmt"
	"strconv"
	"strings"
)

// JoinInt, join integer with delim
func JoinInt(input []int, delim string) (string, error) {
	if len(input) < 1 {
		return "", fmt.Errorf("array empty")
	}
	output := make([]string, len(input))
	for i, v := range input {
		output[i] = strconv.Itoa(v)
	}

	return strings.Join(output, delim), nil
}

// GenerateKey, join strings for key generation purposes, will automatically trim any excess spaces and make it to lower
func GenerateKey(data ...string) string {
	for idx := range data {
		data[idx] = strings.TrimSpace(data[idx])
		data[idx] = strings.ToLower(data[idx])
	}
	return strings.Join(data, "|")
}
