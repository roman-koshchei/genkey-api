package structs

// results of keyboard analyze
// base on old PrintAnalysis func
// cuted: Name, Keys, Keymap -> only stats
type Analysis struct {
	// for overall rolls just get sum
	// Left rolls
	LeftInwardRolls  float32 `json:"leftInwardRolls"`
	LeftOutwardRolls float32 `json:"leftOutwardRolls"`

	// right rolls
	RightInwardRolls  float32 `json:"rightInwardRolls"`
	RightOutwardRolls float32 `json:"rightOutwardRolls"`

	Alternates float32 `json:"alternates"`
	Onehands   float32 `json:"onehands"`
	Redirects  float32 `json:"redirects"`

	WeightedFingerSpeed    float32 `json:"weightedFingerSpeed"`
	UnweightedFingerSpeed  float32 `json:"unweightedFingerSpeed"`
	WeightedHighestSpeed   float32 `json:"weightedHighestSpeed"`
	UnweightedHighestSpeed float32 `json:"unweightedHighestSpeed"`

	LeftIndexUsage  float32 `json:"leftIndexUsage"`
	RightIndexUsage float32 `json:"rightIndexUsage"`

	Sfbs  float32 `json:"sfbs"`
	Dsfbs float32 `json:"dsfbs"`
	Lsbs  float32 `json:"lsbs"`

	TopSfbs      []FreqPair `json:"topSfbs"`
	WorstBigrams []FreqPair `json:"worstBigrams"`

	Score float32 `json:"score"`
}
