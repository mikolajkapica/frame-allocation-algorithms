package frame_allocation_algorithms

import (
	"frame-allocation-algorithms/utils"
)

func PageFaultFrequencyControl(pages []utils.Page, processesQuantity, framesQuantity, trashingCheckInterval, dt int, l, u, h float32) (int, float32) {
	frames := GetProportionalFrames(pages, processesQuantity, framesQuantity)
	processes := make([]utils.Process, processesQuantity)

	framesAvailable := 0
	pageFaults := 0
	trashingCheckedCounter := 0
	trashing := 0

	for i := 0; i < len(pages); i++ {
		utils.Lru(pages, i, &frames, &pageFaults, &processes, &trashing, &trashingCheckedCounter, trashingCheckInterval)

		if i%dt == 0 {
			for j := 0; j < len(processes); j++ {
				currentProcess := processes[j]
				if len(currentProcess.LastPagesIsFault) < dt {
					continue
				}

				ppf := float32(SumOfTrue(currentProcess.LastPagesIsFault[len(currentProcess.LastPagesIsFault)-dt:])) / float32(dt)
				//fmt.Println(ppf)

				if ppf <= l {
					// we decrease the number of frames by one of this process
					if len(frames[j]) > 0 {
						frames[j] = frames[j][:len(frames[j])-1]
					}
				} else if ppf >= u && ppf < h {
					// we increase the number of frames by one of this process
					if framesAvailable > 0 {
						frames[j] = append(frames[j], utils.Page{})
						framesAvailable--
					}
				} else if ppf >= h {
					// we remove all frames of this process
					framesAvailable += len(frames[j])
					frames[j] = []utils.Page{}
				}

			}
		}
	}

	trashingPercentage := float32(trashing) / float32(trashingCheckedCounter)

	return pageFaults, trashingPercentage
}

func SumOfTrue(a []bool) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		if a[i] {
			sum++
		}
	}
	return sum
}
