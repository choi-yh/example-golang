package score_printing

import "fmt"

// DataSheetView 목록 형태로 출력
type DataSheetView struct {
	ScoreRecord
	ViewCount int
}

func (v *DataSheetView) DataSheetView(scoreRecord ScoreRecord, viewCount int) {
	v.ScoreRecord = scoreRecord
	v.ViewCount = viewCount
}

func (v *DataSheetView) Update() {
	record := v.GetScoreRecord()
	v.displayScores(record, v.ViewCount)
}

func (v *DataSheetView) displayScores(record []int, viewCount int) {
	fmt.Printf("List of %d entries: ", viewCount)
	for _, r := range record {
		fmt.Printf("%d ", r)
	}
	fmt.Println()
}
