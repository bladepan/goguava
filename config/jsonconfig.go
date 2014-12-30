package config

import (
	"log"
	//"container/list"
	"encoding/json"
	"os"
)

type config struct {
	values map[string]interface{}
}

func NewConfig(filename string) (result *config, err error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Error in opening %s : %#v \n", filename, err)
		return
	}
	decoder := json.NewDecoder(file)
	value := make(map[string]interface{})
	err = decoder.Decode(&value)
	if err != nil {
		log.Printf("Error in unmarshal %s : %#v \n", filename, err)
		return
	}
	result = &config{}
	result.values = value
	return
}

func (c *config) GetString(key string) string {
	val := c.values[key]
	if val != nil {
		v, ok := val.(string)
		if !ok {
			log.Println("the " + key + " is not a string")
		}
		return v
	}
	return ""
}
