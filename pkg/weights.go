package genkey

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hjson/hjson-go"
)

// load weights
func loadWeights() {
	b, err := ioutil.ReadFile("../configs/weights.hjson")
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
