package genkey

import . "github.com/roman-koshchei/genkey-api/structs"

// global variables for genkey working
// TODO: delete unnesessary

var StaggerFlag bool
var SlideFlag bool
var DynamicFlag bool
var ImproveFlag bool
var ImproveLayout Layout

var FingerNames = [8]string{"LP", "LR", "LM", "LI", "RI", "RM", "RR", "RP"}

var Layouts map[string]Layout
var GeneratedFingermap map[Finger][]Pos
var GeneratedFingermatrix map[Pos]Finger

var SwapPossibilities []Pos

var Analyzed int

var Weight WeightData

//  struct {
// 	FSpeed struct {
// 		SFB       float64
// 		DSFB      float64
// 		KeyTravel float64
// 		KPS       [8]float64
// 	}
// 	Dist struct {
// 		Lateral float64
// 	}
// 	Score struct {
// 		FSpeed       float64
// 		IndexBalance float64
// 		LSB          float64

// 		TrigramPrecision int
// 		LeftInwardRoll   float64
// 		LeftOutwardRoll  float64
// 		RightInwardRoll  float64
// 		RightOutwardRoll float64
// 		Alternate        float64
// 		Redirect         float64
// 		Onehand          float64
// 	}
// }
