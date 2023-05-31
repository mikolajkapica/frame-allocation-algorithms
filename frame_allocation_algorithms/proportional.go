package frame_allocation_algorithms

import (
	"frame-allocation-algorithms/page_replacement_algorithms"
	"frame-allocation-algorithms/utils"
	"math/rand"
)

func ProportionalAllocation(ram utils.RAM, processes []utils.Process, trashingCheckInterval int) (int, int) {
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

		//utils.DisplayRAM(ram, isPageFault, *currentProcess, currentPage)
	}

	return pageFaultsCounter, trashingCounter
}

func ProportionalFrames(ram *utils.RAM, processes []utils.Process, pagesQuantity int) {
	for i := 0; i < len(processes); i++ {
		// equation: frames[i] = (pages of i-th process / all pages) * framesQuantity
		framesToAllocate := int((float32(len(processes[i].Pages)) / float32(pagesQuantity)) * float32(ram.FramesQuantity))
		if framesToAllocate == 0 {
			framesToAllocate = 1
		}
		ram.Frames[i] = make([]utils.Frame, framesToAllocate)
		ram.FramesAvailable -= framesToAllocate
	}

	p := 0
	for ram.FramesAvailable > 0 {
		p = p % len(processes)
		ram.Frames[p] = append(ram.Frames[p], utils.Frame{})
		ram.FramesAvailable--
		p++
	}

	p = 0
	for ram.FramesAvailable < 0 {
		p = p % len(processes)
		if len(ram.Frames[p]) > 1 {
			ram.Frames[p] = ram.Frames[p][:len(ram.Frames[p])-1]
			ram.FramesAvailable++
		}
		p++
	}
}
