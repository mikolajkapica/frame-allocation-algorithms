package page_replacement_algorithms

import "frame-allocation-algorithms/utils"

func LRU(ram *utils.RAM, currentPage utils.Page, process *utils.Process, trashingCheckInterval int, trashingCounter *int) bool {
	currentProcessFrames := ram.Frames[process.Id]

	isPageFault := false
	isNewFrame := false
	if !utils.FramesContainPage(currentProcessFrames, currentPage) {
		process.PageFaults++
		if len(currentProcessFrames) > 0 {
			frameToReplace := NotUsedForLongestTimeInPast(process.History, currentProcessFrames)
			//if ram.Frames[process.Id][frameToReplace].Page.Id == 0 {
			//	isNewFrame = true
			//}
			ram.Frames[process.Id][frameToReplace] = utils.Frame{Page: currentPage}
		}
		if !isNewFrame {
			isPageFault = true
		}
	}

	process.AddPageToHistory(currentPage)
	process.RemovePage()
	process.HistoryOfPageFaults = append(process.HistoryOfPageFaults, isPageFault)

	// trashing
	if len(process.HistoryOfPageFaults)%trashingCheckInterval == 0 {
		lastPageFaults := process.HistoryOfPageFaults[len(process.HistoryOfPageFaults)-trashingCheckInterval:]
		trashingOccured := true

		for i := range lastPageFaults {
			// if there was no page fault at some point then there is no trashing
			if !lastPageFaults[i] {
				trashingOccured = false
				break
			}
		}
		if trashingOccured {
			*trashingCounter++
		}
	}

	return isPageFault
}

func NotUsedForLongestTimeInPast(history []utils.Page, frames []utils.Frame) int {
	checkedFrames := make([]utils.Page, 0)
	// add frames that are used in past
	for i := len(history) - 1; i >= 0; i-- {
		// if there is a frame that is used in future and is not checked yet then add it to checkedFrames
		if utils.FramesContainPage(frames, history[i]) && !utils.Contains(checkedFrames, history[i]) {
			checkedFrames = append(checkedFrames, history[i])
			// if every frame except the last one is checked then return the last one
			if len(checkedFrames) == len(frames)-1 {
				for j := 0; j < len(frames); j++ {
					if !utils.Contains(checkedFrames, frames[j].Page) {
						return j
					}
				}
			}
		}
	}
	for i := 0; i < len(frames); i++ {
		if !utils.Contains(checkedFrames, frames[i].Page) {
			return i
		}
	}
	return 0
}
