package score_printing

// ScoreRecord 점수 저장
type ScoreRecord struct {
	Scores []int
}

func (r *ScoreRecord) SetDataSheetView(score int, dataSheetView DataSheetView) {
	r.AddScore(score)
	dataSheetView.Update()
}

func (r *ScoreRecord) AddScore(score int) {
	r.Scores = append(r.Scores, score)
}

func (r *ScoreRecord) GetScoreRecord() []int {
	return r.Scores
}
