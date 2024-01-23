package helpers

import "regexp"

func IsImage(filename string) bool {
	re := regexp.MustCompile(`\.(jpg|jpeg|png)$`)
	return re.MatchString(filename)
}