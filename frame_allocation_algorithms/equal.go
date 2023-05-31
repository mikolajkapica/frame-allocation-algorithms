package frame_allocation_algorithms

import (
	"frame-allocation-algorithms/page_replacement_algorithms"
	"frame-allocation-algorithms/utils"
	"math/rand"
)

func EqualAllocation(ram utils.RAM, processes []utils.Process, trashingCheckInterval int) (int, int) {
	// make frames for each process in ram
	framesForEachProcess := ram.FramesQuantity / len(processes)
	for i := 0; i < len(processes); i++ {
		ram.Frames[i] = make([]utils.Frame, framesForEachProcess)
		ram.FramesAvailable -= framesForEachProcess
	}

	// if there are frames left, add them to to first processes
	p := 0
	for ram.FramesAvailable > 0 {
		ram.Frames[p] = append(ram.Frames[p], utils.Frame{})
		ram.FramesAvailable--
	}

	// calculate how much pages are in all processes
	pagesLeft := 0
	for _, process := range processes {
		pagesLeft += len(process.Pages)
	}

	trashingCounter := 0
	pageFaultsCounter := 0

	//run lru
	for pagesLeft > 0 {
		// get random process
		index := 0
		currentProcess := &utils.Process{Pages: make([]utils.Page, 0)}
		for len(currentProcess.Pages) == 0 {
			index = rand.Intn(len(processes))
			currentProcess = &processes[index]
		}

		currentPage := currentProcess.Pages[0]
		isPageFault := page_replacement_algorithms.LRU(&ram, currentPage, currentProcess, trashingCheckInterval, &trashingCounter)
		pagesLeft--

		if isPageFault {
			pageFaultsCounter++
		}

		//utils.DisplayRAM(ram, isPageFault, *currentProcess, currentPage)
	}

	return pageFaultsCounter, trashingCounter
}
