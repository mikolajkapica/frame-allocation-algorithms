package utils

func Contains(array []Page, element Page) bool {
	for _, value := range array {
		if value == element {
			return true
		}
	}
	return false
}
