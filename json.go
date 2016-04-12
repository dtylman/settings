package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (v View) loadJSON(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &v)
}

func (v View) printJSON() error {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func (v View) saveJSON(fileName string) error {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, 0664)
}
