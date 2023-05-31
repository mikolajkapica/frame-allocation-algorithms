package frame_allocation_algorithms

import (
	"frame-allocation-algorithms/utils"
)

func ProportionalAllocation(pages []utils.Page, processesQuantity, framesQuantity, trashingCheckInterval int) (int, float32) {
	frames := GetProportionalFrames(pages, processesQuantity, framesQuantity)
	processes := make([]utils.Process, processesQuantity)

	pageFaults := 0
	trashingCheckedCounter := 0
	trashing := 0

	for i := 0; i < len(pages); i++ {
		utils.Lru(pages, i, &frames, &pageFaults, &processes, &trashing, &trashingCheckedCounter, trashingCheckInterval)
	}

	trashingPercentage := float32(trashing) / float32(trashingCheckedCounter)
	return pageFaults, trashingPercentage
}

func GetProportionalFrames(pages []utils.Page, processesQuantity int, framesQuantity int) [][]utils.Page {
	// equation: frames[i] = (pages of i-th process / all pages) * framesQuantity
	pagesOfEachProcess := make([]int, processesQuantity)
	for i := 0; i < len(pages); i++ {
		pagesOfEachProcess[pages[i].ProcessNumber]++
	}

	sum := 0
	for i := 0; i < len(pagesOfEachProcess); i++ {
		sum += pagesOfEachProcess[i]
	}

	framesUsed := 0
	frames := make([][]utils.Page, processesQuantity)
	for i := 0; i < processesQuantity; i++ {
		framesForProcess := int((float32(pagesOfEachProcess[i]) / float32(sum)) * float32(framesQuantity))
		frames[i] = make([]utils.Page, framesForProcess)
		framesUsed += framesForProcess
	}

	// if there is a remainder and there is a frame that is not used then add it to the last process
	if framesUsed < framesQuantity {
		for i := 0; i < len(frames); i++ {
			if len(frames[i]) == 0 {
				frames[i] = make([]utils.Page, 1)
				framesUsed++
				if framesUsed == framesQuantity {
					break
				}
			}
		}
	}

	// if theres a remainder add it to the process with the most frames
	if framesUsed < framesQuantity {
		max := 0
		maxFramesIndex := 0
		for i := 0; i < len(frames); i++ {
			if len(frames[i]) > max {
				max = len(frames[i])
				maxFramesIndex = i
			}
		}
		for framesUsed < framesQuantity {
			frames[maxFramesIndex] = append(frames[maxFramesIndex], utils.Page{})
			framesUsed++
		}
	}
	return frames
}
