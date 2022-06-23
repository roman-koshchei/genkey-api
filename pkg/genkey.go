package genkey

// mian file for genkey

import . "github.com/roman-koshchei/genkey-api/structs"

var Data TextData
var isConfigLoaded bool = false

// Analysis
func Analyze(keys string, fingers string) Analysis {
	if !isConfigLoaded {
		loadConfig()
	}

	//layout := loadLayout(keys, fingers);

	return Analysis{} // tmp
}

func loadConfig() {

	Data = loadData() // from text.go
	loadWeights()     // from config.go

	isConfigLoaded = true
	// old loaded layouts
	// we should return layout given to us
	// Layouts = make(map[string]Layout)
	// LoadLayoutDir()

	// was
	// checkLayoutProvided(args)
	// PrintAnalysis(getLayout(args[1])) // args[1] = name of layout
}
