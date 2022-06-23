package structs

// structs from old layout.go

type Pos struct {
	Col int
	Row int
}

type Pair [2]Pos
type Finger int

type Layout struct {
	Name         string
	Keys         [][]string
	Keymap       map[string]Pos
	Fingermatrix map[Pos]Finger
	Fingermap    map[Finger][]Pos
	Total        float64
}

type FreqPair struct {
	Ngram string  `json:"ngram"`
	Count float64 `json:"count"`
}

type TrigramValues struct {
	RightInwardRolls  int
	RightOutwardRolls int
	LeftInwardRolls   int
	LeftOutwardRolls  int
	Alternates        int
	Onehands          int
	Redirects         int
	Total             int
}
