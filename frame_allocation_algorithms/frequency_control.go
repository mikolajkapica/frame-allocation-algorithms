package frame_allocation_algorithms

import (
	"frame-allocation-algorithms/page_replacement_algorithms"
	"frame-allocation-algorithms/utils"
	"math/rand"
)

func FrequencyControlAllocation(ram utils.RAM, processes []utils.Process, trashingCheckInterval, dt int, l, u, h float32) (int, int) {
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
			//if len(ram.Frames[currentProcess.Id]) == 0 && len(currentProcess.Pages) > 0 {
			//	if ram.FramesAvailable > 0 {
			//		ram.Frames[currentProcess.Id] = make([]utils.Frame, 1)
			//	}
			//}
			index = rand.Intn(len(processes))
			currentProcess = &processes[index]
		}
		//|| len(ram.Frames[currentProcess.Id]) == 0

		if len(ram.Frames[currentProcess.Id]) == 0 {
			if ram.FramesAvailable > 0 {
				ram.Frames[currentProcess.Id] = make([]utils.Frame, 1)
				ram.FramesAvailable--
			} else {
				// find which process has the most frames
				maxFrames := 0
				maxFramesProcess := &utils.Process{}
				for _, process := range processes {
					if len(ram.Frames[process.Id]) > maxFrames {
						maxFrames = len(ram.Frames[process.Id])
						maxFramesProcess = &process
					}
				}
				// remove one frame from the process with the most frames
				if len(ram.Frames[maxFramesProcess.Id]) > 0 {
					ram.Frames[maxFramesProcess.Id] = ram.Frames[maxFramesProcess.Id][:len(ram.Frames[maxFramesProcess.Id])-1]
					ram.FramesAvailable++

					// add one frame to the current process
					ram.Frames[currentProcess.Id] = make([]utils.Frame, 1)
					ram.FramesAvailable--
				}

			}
		}

		currentPage := currentProcess.Pages[0]
		isPageFault := page_replacement_algorithms.LRU(&ram, currentPage, currentProcess, trashingCheckInterval, &trashingCounter)
		pagesQuantity--

		if isPageFault {
			pageFaultsCounter++
		}

		if len(currentProcess.Pages) == 0 {
			ram.FramesAvailable += len(ram.Frames[currentProcess.Id])
			ram.Frames[currentProcess.Id] = []utils.Frame{}
		}

		currentProcessHistoryLength := len(currentProcess.History)
		if currentProcessHistoryLength%dt == 0 && currentProcessHistoryLength > dt {
			ppf := float32(utils.SumOfTrue(currentProcess.HistoryOfPageFaults[currentProcessHistoryLength-dt:])) / float32(dt)
			if ppf <= l {
				// we decrease the number of frames by one of this process
				if len(ram.Frames[currentProcess.Id]) > 0 {
					ram.Frames[currentProcess.Id] = ram.Frames[currentProcess.Id][:len(ram.Frames[currentProcess.Id])-1]
				}
			} else if ppf >= u && ppf < h {
				// we increase the number of frames by one of this process
				if ram.FramesAvailable > 0 {
					ram.Frames[currentProcess.Id] = append(ram.Frames[currentProcess.Id], utils.Frame{})
					ram.FramesAvailable--
				}
			} else if ppf >= h {
				// we remove all frames of this process
				ram.FramesAvailable += len(ram.Frames[currentProcess.Id])
				ram.Frames[currentProcess.Id] = []utils.Frame{}
				p := 0
				for ram.FramesAvailable > 0 {
					ram.Frames[p%len(processes)] = append(ram.Frames[p%len(processes)], utils.Frame{})
					ram.FramesAvailable--
					p++
				}
			}
		}

		//utils.DisplayRAM(ram, isPageFault, *currentProcess, currentPage)
	}

	return pageFaultsCounter, trashingCounter
}
