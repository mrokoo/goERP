package valueobj

type Price struct {
	Purchase float64 `json:"purchase"`
	Retail   float64 `json:"retail"`
	Grade1   float64 `json:"grade1"`
	Grade2   float64 `json:"grade2"`
	Grade3   float64 `json:"grade3"`
}
