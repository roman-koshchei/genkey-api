package genkey

import (
	"encoding/json"
	"os"

	. "github.com/paragoda/genkey-api/structs"
)

func loadData() TextData {
	b, err := os.ReadFile("config/data.json") //ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	var localData TextData

	err = json.Unmarshal(b, &localData)

	if err != nil {
		panic(err)
	}

	return localData
}
