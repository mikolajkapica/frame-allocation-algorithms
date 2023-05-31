package frame_allocation_algorithms

import (
	"fmt"
	"frame-allocation-algorithms/page_replacement_algorithms"
	"frame-allocation-algorithms/utils"
	"math/rand"
)

func ZoneAllocation(ram utils.RAM, processes []utils.Process, trashingCheckInterval, dt int) (int, int) {
	// make frames for each process in ram
	pagesQuantity := 0
	for _, process := range processes {
		pagesQuantity += len(process.Pages)
	}

	ProportionalFrames(&ram, processes, pagesQuantity)

	trashingCounter := 0
	pageFaultsCounter := 0

	//run lru
	for pagesQuantity > 0 {
		// get random process
		index := 0
		currentProcess := &utils.Process{Pages: make([]utils.Page, 0)}
		for len(currentProcess.Pages) == 0 {
			index = rand.Intn(len(processes))
			currentProcess = &processes[index]
		}

		currentPage := currentProcess.Pages[0]
		isPageFault := page_replacement_algorithms.LRU(&ram, currentPage, currentProcess, trashingCheckInterval, &trashingCounter)
		pagesQuantity--

		if isPageFault {
			pageFaultsCounter++
		}

		currentProcessHistoryLength := len(currentProcess.History)

		c := dt / 2
		framesToAdd := 0
		if currentProcessHistoryLength%c == 0 && currentProcessHistoryLength > dt {
			for i := 0; i < len(processes); i++ {
				framesToHave := findHowManyUnique(currentProcess.History[currentProcessHistoryLength-dt:])
				framesToAdd = framesToHave - len(ram.Frames[i])
				if framesToAdd > 0 {
					for j := 0; j < framesToAdd; j++ {
						if ram.FramesAvailable > 0 {
							ram.Frames[i] = append(ram.Frames[i], utils.Frame{})
							ram.FramesAvailable--
						}
					}
				} else if framesToAdd < 0 {
					for j := 0; j < -framesToAdd; j++ {
						if len(ram.Frames[i]) > 0 {
							ram.Frames[i] = ram.Frames[i][:len(ram.Frames[i])-1]
							ram.FramesAvailable++
						}
					}
				}
			}
			fmt.Println("Zone allocation for process", currentProcess.Id, "added frames:", framesToAdd)
		}

		utils.DisplayRAM(ram, isPageFault, *currentProcess, currentPage)
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
