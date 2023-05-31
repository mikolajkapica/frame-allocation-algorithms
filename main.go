package main

import (
	"fmt"
	"frame-allocation-algorithms/utils"
	"frame-allocation-algorithms/utils/frame_allocation_algorithms"
	//"frame-allocation-algorithms/utils/page_replacement_algorithms"
)

func main() {
	const framesQuantity int = 30
	const processesQuantity int = 5
	const maxPageQuantity int = 100
	const minPageNum int = 0
	const maxPageNum int = 100
	const localityMaximumFrequency = 100
	const localityMaximumHistoryLength = 20
	const localityMaximumLength = 70
	const trashingCheckInterval int = 5
	const simulationLoops int = 100

	pages := utils.GeneratePages(processesQuantity, maxPageQuantity, minPageNum, maxPageNum, localityMaximumFrequency, localityMaximumHistoryLength, localityMaximumLength)

	equalFaults, equalTrashing := frame_allocation_algorithms.EqualAllocation(pages, processesQuantity, framesQuantity, trashingCheckInterval)
	fmt.Println("Equal faults: ", equalFaults, " trashing: ", equalTrashing)
	proportionalFaults, proportionalTrashing := frame_allocation_algorithms.ProportionalAllocation(pages, processesQuantity, framesQuantity, trashingCheckInterval)
	fmt.Println("Proportional faults: ", proportionalFaults, " trashing: ", proportionalTrashing)
	pageFaultFrequencyControlFaults, controlTrashing := frame_allocation_algorithms.PageFaultFrequencyControl(pages, processesQuantity, framesQuantity, trashingCheckInterval, 5, 0.1, 0.2, 0.3)
	fmt.Println("Page fault frequency control faults: ", pageFaultFrequencyControlFaults, " trashing: ", controlTrashing)
}
