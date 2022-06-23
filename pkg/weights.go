package genkey

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hjson/hjson-go"
)

// load weights
func loadWeights() {
	b, err := os.ReadFile("config/weights.hjson") //ioutil.ReadFile("config/weights.hjson")
	if err != nil {
		fmt.Printf("There was an issue reading the weights file.\nPlease make sure there is a 'weights.hjson' in this directory.")
		panic(err)
	}

	var dat map[string]interface{}

	err = hjson.Unmarshal(b, &dat)
	if err != nil {
		panic(err)
	}

	j, _ := json.Marshal(dat)
	json.Unmarshal(j, &Weight)
}
