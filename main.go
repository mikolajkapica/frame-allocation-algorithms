package main

import (
	"fmt"
	"frame-allocation-algorithms/frame_allocation_algorithms"
	"frame-allocation-algorithms/utils"
	"math/rand"
)

func main() {
	const framesQuantity int = 25
	const processesQuantity int = 5
	const pagesQuantityBase int = 40
	const pagesQuantityDivergence int = 10
	const minPageNum int = 1
	const maxPageNum int = 30
	const localityMaximumFrequency = 75
	const localityMaximumHistoryLength = 20
	const localityMaximumLength = 15

	const trashingCheckInterval int = 10
	const trashingCheckLast int = 20
	const trashingMax int = 4

	ram := utils.RAM{
		Frames:          make([][]utils.Frame, processesQuantity),
		FramesQuantity:  framesQuantity,
		FramesAvailable: framesQuantity,
	}

	processes1 := generateProcessesWithPages(processesQuantity, pagesQuantityBase, pagesQuantityDivergence, minPageNum, maxPageNum, localityMaximumFrequency, localityMaximumHistoryLength, localityMaximumLength)
	processes2 := make([]utils.Process, len(processes1))
	processes3 := make([]utils.Process, len(processes1))
	processes4 := make([]utils.Process, len(processes1))
	copy(processes2, processes1)
	copy(processes3, processes1)
	copy(processes4, processes1)

	fmt.Println("Equal allocation")
	pageFaults, trashing := frame_allocation_algorithms.EqualAllocation(ram, processes1, trashingCheckInterval, trashingCheckLast, trashingMax)
	fmt.Println("Page faults: ", pageFaults, " Trashing: ", trashing)
	fmt.Println()

	fmt.Println("Proportional allocation")
	pageFaults, trashing = frame_allocation_algorithms.ProportionalAllocation(ram, processes2, trashingCheckInterval, trashingCheckLast, trashingMax)
	fmt.Println("Page faults: ", pageFaults, " Trashing: ", trashing)
	fmt.Println()

	fmt.Println("Frequency control allocation")
	pageFaults, trashing = frame_allocation_algorithms.FrequencyControlAllocation(ram, processes3, trashingCheckInterval, trashingCheckLast, trashingMax, 15, 0.3, 0.5, 0.9)
	fmt.Println("Page faults: ", pageFaults, " Trashing: ", trashing)
	fmt.Println()

	fmt.Println("Zone allocation")
	pageFaults, trashing = frame_allocation_algorithms.ZonalAllocation(ram, processes4, trashingCheckInterval, trashingCheckLast, trashingMax, 15)
	fmt.Println("Page faults: ", pageFaults, " Trashing: ", trashing)
	fmt.Println()
}

func generateProcessesWithPages(processesQuantity int, pagesQuantityBase int, pagesQuantityDivergence int, minPageNum int, maxPageNum int, localityMaximumFrequency int, localityMaximumHistoryLength int, localityMaximumLength int) []utils.Process {
	//rand.Seed(100)
	processes := make([]utils.Process, processesQuantity)

	for i, process := range processes {
		process.Id = i
		process.Pages = utils.GenerateProcessPages(process, pagesQuantityBase+rand.Intn(pagesQuantityDivergence+1), minPageNum+i*maxPageNum, maxPageNum, localityMaximumFrequency, localityMaximumHistoryLength, localityMaximumLength)
		processes[i] = process
	}

	return processes
}
