package frame_allocation_algorithms

import (
	"frame-allocation-algorithms/page_replacement_algorithms"
	"frame-allocation-algorithms/utils"
	"math/rand"
)

func ZonalAllocation(ram utils.RAM, processes []utils.Process, trashingCheckInterval, trashingCheckLast, trashingMax, dt int) (int, int) {
	c := dt / 2
	pagesQuantity := 0
	for _, process := range processes {
		pagesQuantity += len(process.Pages)
	}

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

		// WSS
		if len(currentProcess.History) > c {
			framesToBe := findHowManyUnique(currentProcess.History[len(currentProcess.History)-c : len(currentProcess.History)])
			if framesToBe > len(ram.Frames[index]) {
				for ram.FramesAvailable > 0 {
					ram.Frames[index] = append(ram.Frames[index], utils.Frame{})
					ram.FramesAvailable--
				}
			} else if framesToBe < len(ram.Frames[index]) {
				ram.FramesAvailable += len(ram.Frames[index]) - framesToBe
				ram.Frames[index] = ram.Frames[index][:framesToBe]
			}
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

func findHowManyUnique(pages []utils.Page) int {
	unique := make([]utils.Page, 0)
	for _, page := range pages {
		if !utils.Contains(unique, page) {
			unique = append(unique, page)
		}
	}
	return len(unique)
}
