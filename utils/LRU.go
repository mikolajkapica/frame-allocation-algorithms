package utils

func Lru(pages []Page, i int, frames *[][]Page, pageFaults *int, processes *[]Process, trashing, trashinCheckedCounter *int, trashingCheckInterval int) {
	currentPage := pages[i]
	if !Contains((*frames)[currentPage.ProcessNumber], currentPage) {
		*pageFaults++
		if len((*frames)[currentPage.ProcessNumber]) != 0 {
			currentFrame := NotUsedForLongestTimeInPast(pages, (*frames)[currentPage.ProcessNumber], i)
			(*frames)[currentPage.ProcessNumber][currentFrame] = currentPage
			(*processes)[currentPage.ProcessNumber].LastPagesIsFault = append((*processes)[currentPage.ProcessNumber].LastPagesIsFault, true)
		}
	} else {
		(*processes)[currentPage.ProcessNumber].LastPagesIsFault = append((*processes)[currentPage.ProcessNumber].LastPagesIsFault, false)
	}
	(*processes)[currentPage.ProcessNumber].ReferenceCounter++

	*trashing = CalculateTrashing((*processes)[currentPage.ProcessNumber], trashingCheckInterval, trashinCheckedCounter, *trashing)
}

func CalculateTrashing(process Process, trashingCheckInterval int, trashingCheckedCounter *int, trashing int) int {
	if process.ReferenceCounter%trashingCheckInterval == 0 && process.ReferenceCounter != 0 {
		isTrashing := AllElementsAreTrue(process, trashingCheckInterval)
		*trashingCheckedCounter++
		if isTrashing {
			trashing++
		}
	}
	return trashing
}

func AllElementsAreTrue(process Process, lastElements int) bool {
	lastArePageFaults := process.LastPagesIsFault[len(process.LastPagesIsFault)-lastElements:]
	allTrue := true
	for k := 0; k < len(lastArePageFaults); k++ {
		if !lastArePageFaults[k] {
			allTrue = false
			break
		}
	}
	return allTrue
}

func NotUsedForLongestTimeInPast(pages []Page, frames []Page, currentPageIndex int) int {
	checkedFrames := make([]Page, 0)
	// add frames that are used in past
	for i := currentPageIndex - 1; i >= 0; i-- {
		// if there is a frame that is used in future and is not checked yet then add it to checkedFrames
		if Contains(frames, pages[i]) && !Contains(checkedFrames, pages[i]) {
			checkedFrames = append(checkedFrames, pages[i])
			// if every frame except the last one is checked then return the last one
			if len(checkedFrames) == len(frames)-1 {
				for j := 0; j < len(frames); j++ {
					if !Contains(checkedFrames, frames[j]) {
						return j
					}
				}
			}
		}
	}
	for i := 0; i < len(frames); i++ {
		if !Contains(checkedFrames, frames[i]) {
			return i
		}
	}
	return 0
}

type Process struct {
	LastPagesIsFault []bool
	ReferenceCounter int
}
