package utils

import (
	"net/url"
	"strings"
	"unicode"
)

func NormalizeString(input string) string {
	// Step 1: Trim leading and trailing spaces
	normalized := strings.TrimSpace(input)

	// Step 2: Convert to lowercase
	normalized = strings.ToLower(normalized)

	// Step 3: Replace multiple spaces with a single space
	normalized = strings.Join(strings.Fields(normalized), " ")

	// Step 4: Make url unescaped
	normalized, _ = url.PathUnescape(normalized)

	// Step 5: Remove non-alphanumeric characters
	normalized = removeNonAlphanumeric(normalized)

	return normalized
}

// Helper function to remove non-alphanumeric characters
func removeNonAlphanumeric(input string) string {
	var result []rune
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			result = append(result, r)
		}
	}
	return string(result)
}
