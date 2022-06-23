package structs

// results of keyboard analyze
// base on old PrintAnalysis func
// cuted: Name, Keys, Keymap -> only stats
type Analysis struct {
	// for overall rolls just get sum
	// Left rolls
	LeftInwardRolls  float64 `json:"leftInwardRolls"`
	LeftOutwardRolls float64 `json:"leftOutwardRolls"`

	// right rolls
	RightInwardRolls  float64 `json:"rightInwardRolls"`
	RightOutwardRolls float64 `json:"rightOutwardRolls"`

	Alternates float64 `json:"alternates"`
	Onehands   float64 `json:"onehands"`
	Redirects  float64 `json:"redirects"`

	WeightedFingerSpeed    []float64 `json:"weightedFingerSpeed"`
	UnweightedFingerSpeed  []float64 `json:"unweightedFingerSpeed"`
	WeightedHighestSpeed   float64   `json:"weightedHighestSpeed"`
	UnweightedHighestSpeed float64   `json:"unweightedHighestSpeed"`

	LeftIndexUsage  float64 `json:"leftIndexUsage"`
	RightIndexUsage float64 `json:"rightIndexUsage"`

	Sfbs  float64 `json:"sfbs"`
	Dsfbs float64 `json:"dsfbs"`
	Lsbs  float64 `json:"lsbs"`

	TopSfbs      []FreqPair `json:"topSfbs"`
	WorstBigrams []FreqPair `json:"worstBigrams"`

	Score float64 `json:"score"`
}
