package main

import (
	"fmt"

	score_printing2 "github.com/choi-yh/example-golang/TIL/design_pattern/observer_pattern/score_printing"
)

func main() {
	scoreRecord := score_printing2.ScoreRecord{}
	dataSheetView := score_printing2.DataSheetView{}
	dataSheetView.DataSheetView(scoreRecord, 3)

	scoreRecord.SetDataSheetView(3, dataSheetView)

	for i := 1; i <= 5; i++ {
		score := i * 10
		fmt.Printf("Adding %d \n", score)

		scoreRecord.AddScore(score)
	}
}
