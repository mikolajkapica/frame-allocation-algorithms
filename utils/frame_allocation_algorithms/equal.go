package frame_allocation_algorithms

import (
	"fmt"
	"frame-allocation-algorithms/utils"
)

func EqualAllocation(pages []utils.Page, processesQuantity, framesQuantity, trashingCheckInterval int) (int, float32) {
	frames := GetEqualFrames(framesQuantity, processesQuantity)
	processes := make([]utils.Process, processesQuantity)

	pageFaults := 0
	trashingCheckedCounter := 0
	trashing := 0

	for i := 0; i < len(pages); i++ {
		utils.Lru(pages, i, &frames, &pageFaults, &processes, &trashing, &trashingCheckedCounter, trashingCheckInterval)
	}
	trashingPercentage := float32(trashing) / float32(trashingCheckedCounter)
	fmt.Println(len(pages))
	return pageFaults, trashingPercentage
}

func GetEqualFrames(framesQuantity int, processesQuantity int) [][]utils.Page {
	framesForEachProcess := framesQuantity / processesQuantity
	frames := make([][]utils.Page, processesQuantity)
	for i := 0; i < processesQuantity; i++ {
		frames[i] = make([]utils.Page, framesForEachProcess)
	}
	return frames
}
