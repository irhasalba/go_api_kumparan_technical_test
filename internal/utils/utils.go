package utils

import (
	"regexp"
	"strings"
	"unicode"
)

func Slugify(title string) string {
	slug := strings.ToLower(title)

	slug = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			return r
		}
		return ' '
	}, slug)

	re := regexp.MustCompile(`\s+`)
	slug = re.ReplaceAllString(slug, "-")

	slug = strings.Trim(slug, "-")

	return slug
}

func PageToOffset(page int) int {
	var offset int
	if page > 1 {
		offset = page - 1*10
	} else {
		offset = 0 * 10
	}
	return offset
}
