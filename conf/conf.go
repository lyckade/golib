package conf

import (
	"encoding/json"
	"io/ioutil"
)

//LoadConf loads the configuration from a json file into a
//struct
func LoadConf(confFile string, confStruc interface{}) {
	var err error
	var b []byte

	b, err = ioutil.ReadFile(confFile)
	errorLog(err)
	err = json.Unmarshal(b, &confStruc)
	errorLog(err)

}

func errorLog(err error) {
	if err != nil {
		panic(err)
	}
}
