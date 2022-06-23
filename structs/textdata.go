package structs

type TextData struct {
	Letters      map[string]int     `json:"letters"`
	Bigrams      map[string]int     `json:"bigrams"`
	Trigrams     map[string]int     `json:"trigrams"`
	TopTrigrams  []FreqPair         `json:"toptrigrams"`
	Skipgrams    map[string]float64 `json:"skipgrams"`
	TotalBigrams int
	Total        int
}
