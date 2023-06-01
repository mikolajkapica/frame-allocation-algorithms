package frame_allocation_algorithms

import (
	"frame-allocation-algorithms/page_replacement_algorithms"
	"frame-allocation-algorithms/utils"
	"math/rand"
)

func EqualAllocation(ram utils.RAM, processes []utils.Process, trashingCheckInterval, trashingCheckLast, trashingMax int) (int, int) {
	pagesQuantity := 0
	for _, process := range processes {
		pagesQuantity += len(process.Pages)
	}

	// get equal frames for each process in ram
	numberOfFramesForEachProcess := ram.FramesQuantity / len(processes)
	for i := 0; i < len(processes); i++ {
		ram.Frames[i] = make([]utils.Frame, numberOfFramesForEachProcess)
	}

	for pagesQuantity > 0 {
		// check if its trashing
		if pagesQuantity%trashingCheckInterval == 0 {
			for i := 0; i < len(processes); i++ {
				if trashingCheckLast < len(processes[i].HistoryOfPageFaults) {
					pageFaultsLately := utils.SumOfTrue(processes[i].HistoryOfPageFaults[len(processes[i].HistoryOfPageFaults)-trashingCheckLast:])
					if pageFaultsLately > trashingMax {
						processes[i].TrashedTimes++
					}
				}
			}
		}

		// get random not frozen process
		index := rand.Intn(len(processes))
		currentProcess := &processes[index]
		for currentProcess.IsFrozen {
			index += 1
			index = index % len(processes)
			currentProcess = &processes[index]
		}

		currentPage := currentProcess.Pages[0]
		currentProcess.AddPageToHistory(currentPage)

		isPageFault := page_replacement_algorithms.LRU(&ram, currentPage, currentProcess)

		if isPageFault {
			currentProcess.PageFaults++
		}

		currentProcess.RemovePage()
		if len(currentProcess.Pages) == 0 {
			currentProcess.IsFrozen = true
		}

		pagesQuantity--
		//utils.DisplayRAM(ram, isPageFault, *currentProcess, currentPage)
	}

	pageFaultsCounter := 0
	trashingCounter := 0
	for _, process := range processes {
		pageFaultsCounter += process.PageFaults
		trashingCounter += process.TrashedTimes
	}

	return pageFaultsCounter, trashingCounter
}
