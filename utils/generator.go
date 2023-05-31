package utils

import (
	"math/rand"
)

func GeneratePages(
	processesQuantity,
	maxPageQuantity,
	minPageNum,
	maxPageNum,
	localityMaximumFrequency,
	localityMaximumHistoryLength,
	localityMaximumLength int) []Page {

	pages := make([][]Page, processesQuantity)
	pagesQuantity := 0
	for i := 0; i < processesQuantity; i++ {
		currentProcessPageQuantity := rand.Intn(maxPageQuantity)
		pagesQuantity += currentProcessPageQuantity
		pages[i] = GenerateProcessPages(i,
			currentProcessPageQuantity,
			minPageNum+maxPageNum*i,
			maxPageNum,
			localityMaximumFrequency,
			localityMaximumHistoryLength,
			localityMaximumLength)
	}

	return MergePages(pages, pagesQuantity)

}

func MergePages(processesPages [][]Page, pagesQuantity int) []Page {
	pages := make([]Page, pagesQuantity)
	idx := 0
	for idx < pagesQuantity {
		currentProcess := rand.Intn(len(processesPages))

		// .pop()
		if len(processesPages[currentProcess]) == 0 {
			continue
		}

		pages[idx] = processesPages[currentProcess][len(processesPages[currentProcess])-1]
		processesPages[currentProcess] = processesPages[currentProcess][:len(processesPages[currentProcess])-1]
		idx++
	}
	return pages
}

// GeneratePages generate pages of length pagesQuantity with random numbers from 0 to maxPageNum
func GenerateProcessPages(
	processNumber,
	pagesQuantity,
	minPageNum,
	maxPageNum,
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
				minPageNum + rand.Intn(maxPageNum+1),
				processNumber,
			}
		}
	}
	return pages
}
