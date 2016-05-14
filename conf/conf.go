package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//LoadConf loads the configuration from a json file into a
//struct
func LoadConf(confFile string, confStruc interface{}, logger *log.Logger) {
	var err error
	var b []byte

	b, err = ioutil.ReadFile(confFile)
	errorLog(err, logger)
	err = json.Unmarshal(b, &confStruc)
	errorLog(err, logger)

}

func errorLog(err error, logger *log.Logger) {
	if err != nil {
		logger.Println(err)
	}
}
