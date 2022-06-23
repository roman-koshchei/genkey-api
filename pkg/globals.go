package genkey

import . "github.com/roman-koshchei/genkey-api/structs"

// global variables for genkey working
// TODO: delete unnesessary

var StaggerFlag bool // flag from args (not used now)
var SlideFlag bool   // flag from args (not used now)
var DynamicFlag bool // flag from args (not used now)

//var ImproveFlag bool
//var ImproveLayout Layout

var FingerNames = [8]string{"LP", "LR", "LM", "LI", "RI", "RM", "RR", "RP"}

//var GeneratedFingermap map[Finger][]Pos
//var GeneratedFingermatrix map[Pos]Finger

//var SwapPossibilities []Pos

var Weight WeightData
