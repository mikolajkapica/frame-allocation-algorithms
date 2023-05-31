package utils

import (
	"math/rand"
)

func GenerateProcessPages(
	process Process,
	pagesQuantity,
	minimumPageNumber,
	maximumPageNumber,
	localityMaximumFrequency,
	localityMaximumHistoryLength,
	localityMaximumLength int) []Page {

	pages := make([]Page, pagesQuantity)

	localityFrequency := rand.Intn(localityMaximumFrequency) + 1
	localityHistoryLength := rand.Intn(localityMaximumHistoryLength) + 1
	localityLength := rand.Intn(localityMaximumLength) + 1

	for i := 0; i < pagesQuantity; i++ {
		// after localityMaximumFrequency of pages
		if i%localityFrequency == 0 && i >= localityHistoryLength {
			// we take random number (0, localityMaximumHistoryLength) of last pages
			history := pages[i-localityHistoryLength : i]
			enteredTime := i
			// for random number (0, localityMaximumLength) times
			for i := i; i < enteredTime+localityLength; i++ {
				if len(pages) == i {
					break
				}
				// generate random number out of these pages from history
				pages[i] = history[rand.Intn(len(history))]
			}
		} else {
			// else generate random number out of all pages
			pages[i] = Page{
				minimumPageNumber + rand.Intn(maximumPageNumber+1),
				process.Id,
			}
		}
	}
	return pages
}
