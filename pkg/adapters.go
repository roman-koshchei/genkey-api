package genkey

// get key rows and finger rows
// from strings that contain all keys
func RowsFromTogether(keys string, fingers string) ([]string, []string) {
	var keyRows []string
	var fingerRows []string

	// topRowLength = 13, homeRowLength = 11 bottomRowLength = 10
	lengths := [3]int{13, 11, 10} //  length for each row

	start := 0 // start position of new row
	end := 0
	for _, length := range lengths {
		end += length
		keyRows = append(keyRows, keys[start:end])
		start += length
	}

	start = 0
	end = 0
	for _, length := range lengths {
		end += length
		fingerRows = append(fingerRows, fingers[start:end])
		start += length
	}

	return keyRows, fingerRows
}

// top, home, bot = rows
func RowsFormDivided(topKeys string, homeKeys string, botKeys string, topFingers string, homeFingers string, botFingers string) ([]string, []string) {
	return []string{topKeys, homeKeys, botKeys}, []string{topFingers, homeFingers, botFingers}
}
