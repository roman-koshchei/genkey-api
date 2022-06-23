package genkey

import (
	"math"
	"sort"
	"strconv"
	"strings"

	. "github.com/roman-koshchei/genkey-api/structs"
)

// make layout from for example:
// keys = qwertyuiop[]\asdfghjkl;'zxcvbnm,./
// fingers= 0123344567777012334456770123344567
func loadLayout(keys string, fingers string) Layout {
	var layout Layout

	// topRowLength = 13, homeRowLength = 11 bottomRowLength = 10

	lengths := [3]int{13, 11, 10} //  length for each row
	start := 0                    // start position of new row
	end := 0

	layout.Keys = make([][]string, 3)
	for row, length := range lengths {
		end += length
		// rune = code of char
		for _, rune := range keys[start:end] {
			key := strings.ToLower(string(rune))

			layout.Keys[row] = append(layout.Keys[row], key)
			layout.Total += float64(data.Letters[key])
		}
		start += length
	}

	start = 0
	end = 0

	layout.Fingermatrix = make(map[Pos]Finger, 3)
	layout.Fingermap = make(map[Finger][]Pos)
	for row, length := range lengths {
		end += length
		// rune = code of char
		for col, rune := range fingers[start:end] {
			fingerNum, err := strconv.Atoi(string(rune))

			if err != nil {
				// error
			}

			finger := Finger(fingerNum)
			layout.Fingermatrix[Pos{Col: col, Row: row}] = finger
			layout.Fingermap[finger] = append(layout.Fingermap[finger], Pos{Col: col, Row: row})
		}
		start += length
	}

	layout.Keymap = GenKeymap(layout.Keys) // need

	return layout
}

// +
func GenKeymap(keys [][]string) map[string]Pos {
	keymap := make(map[string]Pos)
	for y, row := range keys {
		for x, v := range row {
			keymap[v] = Pos{Col: x, Row: y}
		}
	}
	return keymap
}

// +
func FingerSpeed(l *Layout, weighted bool) []float64 {
	speeds := []float64{0, 0, 0, 0, 0, 0, 0, 0}
	sfbweight := Weight.FSpeed.SFB
	dsfbweight := Weight.FSpeed.DSFB
	for f, posits := range l.Fingermap {
		for i := 0; i < len(posits); i++ {
			for j := i; j < len(posits); j++ {
				p1 := &posits[i]
				p2 := &posits[j]
				k1 := &l.Keys[p1.Row][p1.Col]
				k2 := &l.Keys[p2.Row][p2.Col]

				sfb := float64(data.Bigrams[*k1+*k2])
				dsfb := data.Skipgrams[*k1+*k2]
				if i != j {
					sfb += float64(data.Bigrams[*k2+*k1])
					dsfb += data.Skipgrams[*k2+*k1]
				}

				dist := twoKeyDist(*p1, *p2, true) + (2 * Weight.FSpeed.KeyTravel)
				speeds[f] += ((sfbweight * sfb) + (dsfbweight * dsfb)) * dist
			}
		}
		if weighted {
			speeds[f] /= Weight.FSpeed.KPS[f]
		}
		speeds[f] = 800 * speeds[f] / l.Total
	}
	return speeds
}

// if add dynamic flag
func DynamicFingerSpeed(l *Layout, weighted bool) []float64 {
	speeds := []float64{0, 0, 0, 0, 0, 0, 0, 0}
	sfbweight := Weight.FSpeed.SFB
	dsfbweight := Weight.FSpeed.DSFB
	for f, posits := range l.Fingermap {
		for i := 0; i < len(posits); i++ {
			var highestsfb float64
			var highestdsfb float64
			var highestdist float64
			var highestspeed float64
			for j := 0; j < len(posits); j++ {
				p1 := &posits[i]
				p2 := &posits[j]
				k1 := &l.Keys[p1.Row][p1.Col]
				k2 := &l.Keys[p2.Row][p2.Col]

				sfb := float64(data.Bigrams[*k1+*k2])
				dsfb := data.Skipgrams[*k1+*k2]

				dist := twoKeyDist(*p1, *p2, true) + (2 * Weight.FSpeed.KeyTravel)
				speed := ((sfbweight * sfb) + (dsfbweight * dsfb)) * dist
				if sfb > highestsfb {
					highestsfb = sfb
					highestdsfb = dsfb
					highestdist = dist
					highestspeed = speed
				}
				speeds[f] += speed
			}
			newspeed := (dsfbweight * highestdsfb) * highestdist
			speeds[f] -= highestspeed
			speeds[f] += newspeed
		}
		if weighted {
			speeds[f] /= Weight.FSpeed.KPS[f]
		}
		speeds[f] = 800 * speeds[f] / l.Total
	}
	return speeds
}

// +
func SFBs(l Layout, skipgrams bool) float64 {
	var count float64
	for _, posits := range l.Fingermap {
		for i := 0; i < len(posits); i++ {
			for j := i; j < len(posits); j++ {
				if i == j {
					continue
				}
				p1 := &posits[i]
				p2 := &posits[j]
				k1 := &l.Keys[p1.Row][p1.Col]
				k2 := &l.Keys[p2.Row][p2.Col]
				if !skipgrams {
					count += float64(data.Bigrams[*k1+*k2] + data.Bigrams[*k2+*k1])
				} else {
					count += data.Skipgrams[*k1+*k2] + data.Skipgrams[*k2+*k1]
				}
			}
		}
	}
	return count
}

// if add dynamic flag
func DynamicSFBs(l Layout) float64 {
	var count float64
	for _, posits := range l.Fingermap {
		for i := 0; i < len(posits); i++ {
			var highest float64
			for j := 0; j < len(posits); j++ {
				if i == j {
					continue
				}
				p1 := &posits[i]
				p2 := &posits[j]
				k1 := &l.Keys[p1.Row][p1.Col]
				k2 := &l.Keys[p2.Row][p2.Col]
				sfb := float64(data.Bigrams[*k1+*k2])
				if sfb > highest {
					highest = sfb
				}
				count += sfb
			}
			count -= highest
		}
	}
	return count
}

// +
func SortFreqList(pairs []FreqPair) {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})
}

// +
func ListSFBs(l Layout, skipgrams bool) []FreqPair {
	var list []FreqPair
	for _, posits := range l.Fingermap {
		for i := 0; i < len(posits); i++ {
			// since this is output, reversed sfbs cannot
			// be shortcut, so we iterate through all
			// combinations without mirroring (j starts at
			// 0 instead of i)
			for j := 0; j < len(posits); j++ {
				if i == j {
					continue
				}
				p1 := &posits[i]
				p2 := &posits[j]
				k1 := &l.Keys[p1.Row][p1.Col]
				k2 := &l.Keys[p2.Row][p2.Col]
				var count float64
				ngram := *k1 + *k2
				if !skipgrams {
					count = float64(data.Bigrams[ngram])
				} else {
					count = data.Skipgrams[ngram]
				}
				list = append(list, FreqPair{Ngram: ngram, Count: count})
			}
		}
	}

	return list
}

// if add dynamic flag
func ListDynamic(l Layout) ([]FreqPair, []FreqPair) {
	sfbs := ListSFBs(l, false)
	SortFreqList(sfbs)
	var escaped []FreqPair
	var real []FreqPair
	highestfound := make(map[Pos]bool)
	for _, bg := range sfbs {
		prefix := l.Keymap[string(bg.Ngram[0])]
		if highestfound[prefix] {
			real = append(real, bg)
		} else {
			escaped = append(escaped, bg)
			highestfound[prefix] = true
		}
	}

	return escaped, real
}

// +
func ListWorstBigrams(l Layout) []FreqPair {
	var bigrams []FreqPair
	sfbweight := Weight.FSpeed.SFB
	dsfbweight := Weight.FSpeed.DSFB
	for f, posits := range l.Fingermap {
		for i := 0; i < len(posits); i++ {
			for j := i; j < len(posits); j++ {
				p1 := &posits[i]
				p2 := &posits[j]
				k1 := &l.Keys[p1.Row][p1.Col]
				k2 := &l.Keys[p2.Row][p2.Col]
				sfb := float64(data.Bigrams[*k1+*k2])
				dsfb := data.Skipgrams[*k1+*k2]
				if i != j {
					sfb += float64(data.Bigrams[*k2+*k1])
					dsfb += data.Skipgrams[*k2+*k1]
				}

				dist := twoKeyDist(*p1, *p2, true) + (2 * Weight.FSpeed.KeyTravel)
				cost := 100 * (((sfbweight * sfb) + (dsfbweight * dsfb)) * dist) / Weight.FSpeed.KPS[f]
				bigrams = append(bigrams, FreqPair{Ngram: *k1 + *k2, Count: cost})
			}
		}
	}
	return bigrams
}

// FastTrigrams approximates trigram counts with a given precision
// (precision=0 gives full data). It returns a count of {rolls,
// alternates, onehands, redirects, total}
// +
func FastTrigrams(l Layout, precision int) TrigramValues {
	var tgs TrigramValues

	if precision == 0 {
		precision = len(data.TopTrigrams)
	}

	for _, tg := range data.TopTrigrams[:precision] {
		km1, ok1 := l.Keymap[string(tg.Ngram[0])]
		km2, ok2 := l.Keymap[string(tg.Ngram[1])]
		km3, ok3 := l.Keymap[string(tg.Ngram[2])]

		if !ok1 || !ok2 || !ok3 {
			continue
		}

		f1 := l.Fingermatrix[km1]
		f2 := l.Fingermatrix[km2]
		f3 := l.Fingermatrix[km3]

		tgs.Total += int(tg.Count)

		if f1 != f2 && f2 != f3 {
			h1 := (f1 >= 4)
			h2 := (f2 >= 4)
			h3 := (f3 >= 4)

			if h1 == h2 && h2 == h3 {
				dir1 := f1 < f2
				dir2 := f2 < f3

				if dir1 == dir2 {
					tgs.Onehands += int(tg.Count)
				} else {
					tgs.Redirects += int(tg.Count)
				}
			} else if h1 != h2 && h2 != h3 {
				tgs.Alternates += int(tg.Count)
			} else {
				rollhand := h2
				rollfirst := (h1 == rollhand)
				var first Finger
				var second Finger
				if rollfirst {
					first = f1
					second = f2
				} else {
					first = f2
					second = f3
				}
				if rollhand == false { // left hand
					if first < second { // inward roll
						tgs.LeftInwardRolls += int(tg.Count)
						//println("Left Inward Roll: ", tg.Ngram)
					} else {
						tgs.LeftOutwardRolls += int(tg.Count)
						//println("Left Outward Roll: ", tg.Ngram)
					}
				} else if rollhand == true { // right hand
					if first > second { // inward roll
						tgs.RightInwardRolls += int(tg.Count)
						//println("Right Inward Roll: ", tg.Ngram)
					} else {
						tgs.RightOutwardRolls += int(tg.Count)
						//println("Right Outward Roll:", tg.Ngram)
					}
				}
			}
		}
	}

	return tgs
}

// +
func IndexUsage(l Layout) (float64, float64) {
	left := 0
	right := 0

	for _, pos := range l.Fingermap[3] {
		key := l.Keys[pos.Row][pos.Col]
		left += data.Letters[key]
	}
	for _, pos := range l.Fingermap[4] {
		key := l.Keys[pos.Row][pos.Col]
		right += data.Letters[key]
	}

	return (100 * float64(left) / l.Total), (100 * float64(right) / l.Total)
}

// +
func LSBs(l Layout) int {
	var count int

	// LI LM
	for _, p1 := range l.Fingermap[3] {
		for _, p2 := range l.Fingermap[2] {
			var dist float64
			if StaggerFlag {
				dist = math.Abs(staggeredX(p1.Col, p1.Row) - staggeredX(p2.Col, p2.Row))
			} else {
				dist = math.Abs(float64(p1.Col - p2.Col))
			}
			if dist >= 2 {
				k1 := l.Keys[p1.Row][p1.Col]
				k2 := l.Keys[p2.Row][p2.Col]
				count += data.Bigrams[k1+k2]
				count += data.Bigrams[k2+k1]
			}
		}
	}

	// RI RM
	for _, p1 := range l.Fingermap[4] {
		for _, p2 := range l.Fingermap[5] {
			var dist float64
			if StaggerFlag {
				dist = math.Abs(staggeredX(p1.Col, p1.Row) - staggeredX(p2.Col, p2.Row))
			} else {
				dist = math.Abs(float64(p1.Col - p2.Col))
			}
			if dist >= 2 {
				k1 := l.Keys[p1.Row][p1.Col]
				k2 := l.Keys[p2.Row][p2.Col]
				count += data.Bigrams[k1+k2]
				count += data.Bigrams[k2+k1]
			}
		}
	}

	// LP LR
	for _, p1 := range l.Fingermap[0] {
		for _, p2 := range l.Fingermap[1] {
			var dist float64
			if StaggerFlag {
				dist = math.Abs(staggeredX(p1.Col, p1.Row) - staggeredX(p2.Col, p2.Row))
			} else {
				dist = math.Abs(float64(p1.Col - p2.Col))
			}
			if dist >= 2 {
				k1 := l.Keys[p1.Row][p1.Col]
				k2 := l.Keys[p2.Row][p2.Col]
				count += data.Bigrams[k1+k2]
				count += data.Bigrams[k2+k1]
			}
		}
	}

	// RP RR
	for _, p1 := range l.Fingermap[7] {
		for _, p2 := range l.Fingermap[6] {
			var dist float64
			if StaggerFlag {
				dist = math.Abs(staggeredX(p1.Col, p1.Row) - staggeredX(p2.Col, p2.Row))
			} else {
				dist = math.Abs(float64(p1.Col - p2.Col))
			}
			if dist >= 2 {
				k1 := l.Keys[p1.Row][p1.Col]
				k2 := l.Keys[p2.Row][p2.Col]
				count += data.Bigrams[k1+k2]
				count += data.Bigrams[k2+k1]
			}
		}
	}
	return count
}

func ListLSBs(l Layout) []FreqPair {
	var list []FreqPair
	for _, p1 := range l.Fingermap[3] {
		for _, p2 := range l.Fingermap[2] {
			var dist float64
			if StaggerFlag {
				dist = math.Abs(staggeredX(p1.Col, p1.Row) - staggeredX(p2.Col, p2.Row))
			} else {
				dist = math.Abs(float64(p1.Col - p2.Col))
			}
			if dist >= 2 {
				k1 := l.Keys[p1.Row][p1.Col]
				k2 := l.Keys[p2.Row][p2.Col]
				list = append(list, FreqPair{Ngram: k1 + k2, Count: float64(data.Bigrams[k1+k2])})
				list = append(list, FreqPair{Ngram: k2 + k1, Count: float64(data.Bigrams[k2+k1])})
			}
		}
	}

	for _, p1 := range l.Fingermap[4] {
		for _, p2 := range l.Fingermap[5] {
			var dist float64
			if StaggerFlag {
				dist = math.Abs(staggeredX(p1.Col, p1.Row) - staggeredX(p2.Col, p2.Row))
			} else {
				dist = math.Abs(float64(p1.Col - p2.Col))
			}
			if dist >= 2 {
				k1 := l.Keys[p1.Row][p1.Col]
				k2 := l.Keys[p2.Row][p2.Col]
				list = append(list, FreqPair{Ngram: k1 + k2, Count: float64(data.Bigrams[k1+k2])})
				list = append(list, FreqPair{Ngram: k2 + k1, Count: float64(data.Bigrams[k2+k1])})
			}
		}
	}
	return list
}

// Count total score of layout
func Score(l Layout) float64 {
	var score float64
	s := &Weight.Score
	if s.FSpeed != 0 {
		var speeds []float64
		if !DynamicFlag {
			speeds = FingerSpeed(&l, true)
		} else {
			speeds = DynamicFingerSpeed(&l, true)
		}
		total := 0.0
		for _, s := range speeds {
			total += s
		}
		score += s.FSpeed * total
	}
	if s.LSB != 0 {
		score += s.LSB * 100 * float64(LSBs(l)) / l.Total
	}
	if s.TrigramPrecision != -1 {
		tri := FastTrigrams(l, s.TrigramPrecision)
		score += s.LeftInwardRoll * (100 - (100 * float64(tri.LeftInwardRolls) / float64(tri.Total)))
		score += s.RightInwardRoll * (100 - (100 * float64(tri.RightInwardRolls) / float64(tri.Total)))
		score += s.LeftOutwardRoll * (100 - (100 * float64(tri.LeftOutwardRolls) / float64(tri.Total)))
		score += s.RightOutwardRoll * (100 - (100 * float64(tri.RightOutwardRolls) / float64(tri.Total)))
		score += s.Alternate * (100 - (100 * float64(tri.Alternates) / float64(tri.Total)))
		score += s.Onehand * (100 - (100 * float64(tri.Onehands) / float64(tri.Total)))
		score += s.Redirect * (100 * float64(tri.Redirects) / float64(tri.Total))
	}

	if s.IndexBalance != 0 {
		left, right := IndexUsage(l)
		score += s.IndexBalance * math.Abs(right-left)
	}

	return score
}

// unnesasery now
func ColRow(pos int) (int, int) {
	var col int
	var row int
	if pos < 10 {
		col = pos
		row = 0
	} else if pos < 20 {
		col = pos - 10
		row = 1
	} else if pos < 30 {
		col = pos - 20
		row = 2
	}

	return col, row
}

// unnesasery now
func Similarity(a, b []string) int {
	var score int
	for i := 0; i < 30; i++ {
		weight := 1
		if i >= 10 && i <= 13 {
			weight = 2
		} else if i >= 16 && i <= 19 {
			weight = 2
		}
		if a[i] == b[i] {
			score += weight
		}
	}
	return score
}

// unnesasery now
func DuplicatesAndMissing(l Layout) ([]string, []string) {
	counts := make(map[string]int)
	// collect counts of each key
	for _, row := range l.Keys {
		for _, c := range row {
			counts[c] += 1
		}
	}
	// then check duplicates and missing
	duplicates := make([]string, 0)
	missing := make([]string, 0)
	for _, r := range []rune("abcdefghijklmnopqrstuvwxyz,./;'") {
		c := string(r)
		if counts[c] == 0 {
			missing = append(missing, c)
		} else if counts[c] > 1 {
			duplicates = append(duplicates, c)
		}
	}
	return duplicates, missing
}

// unnesasery now
func staggeredX(c, r int) float64 {
	var sx float64
	if r == 0 {
		sx = float64(c) - 0.25
	} else if r == 2 {
		sx = float64(c) + 0.5
	} else {
		sx = float64(c)
	}
	return sx
}

// unnesasery now
func twoKeyDist(a, b Pos, weighted bool) float64 {
	var ax float64
	var bx float64

	if StaggerFlag {
		ax = staggeredX(a.Col, a.Row)
		bx = staggeredX(b.Col, b.Row)
	} else {
		ax = float64(a.Col)
		bx = float64(b.Col)
	}

	x := ax - bx
	y := float64(a.Row - b.Row)

	var dist float64
	if weighted {
		dist = (Weight.Dist.Lateral * x * x) + (y * y)
	} else {
		dist = math.Sqrt((x * x) + (y * y))
	}
	return dist
}
