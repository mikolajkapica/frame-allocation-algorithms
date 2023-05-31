package utils

import "fmt"

type Frame struct {
	//Number int
	Page Page
}

func (frame1 *Frame) Equals(frame2 Frame) bool {
	return frame1.Page.Equals(frame2.Page)
}

type Page struct {
	Id        int
	ProcessID int
}

func (page1 *Page) Equals(page2 Page) bool {
	return page1.Id == page2.Id && page1.ProcessID == page2.ProcessID
}

type Process struct {
	Id                  int
	Pages               []Page
	PageFaults          int
	History             []Page
	HistoryOfPageFaults []bool
}

func (process *Process) RemovePage() {
	process.Pages = process.Pages[1:]
}

func (process *Process) AddPageToHistory(page Page) {
	process.History = append(process.History, page)
}

func (process *Process) Equals(process2 Process) bool {
	if process.Id != process2.Id || process.PageFaults != process2.PageFaults {
		return false
	}

	if len(process.Pages) != len(process2.Pages) {
		return false
	}

	for i := range process.Pages {
		if !process.Pages[i].Equals(process2.Pages[i]) {
			return false
		}
	}

	return true
}

type RAM struct {
	Frames          [][]Frame
	FramesQuantity  int
	FramesAvailable int
}

func DisplayRAM(ram RAM, isPageFault bool, currentProcess Process, currentPage Page) {
	fmt.Printf("ProcessID: %2d PageID: %3d", currentProcess.Id, currentPage.Id)

	fmt.Print(" | ")
	for i := 0; i < len(ram.Frames); i++ {
		if i == currentProcess.Id {
			if isPageFault {
				fmt.Print("\033[31m")
			} else {
				// green
				fmt.Print("\033[32m")
			}
		}
		for j := 0; j < len(ram.Frames[i]); j++ {
			fmt.Printf("%-4d", ram.Frames[i][j].Page.Id)
		}
		fmt.Print("\033[0m")
		fmt.Print(" | ")
	}
	fmt.Println()
}
