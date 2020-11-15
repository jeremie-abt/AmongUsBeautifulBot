package main

import (
	"io/ioutil"
	"os"
)

func ReadJson(jsonPath string) interface{} {
	jsonByteFile, err := os.Open(jsonPath)

	if err != nil {
		return nil
	}

	_, err = ioutil.ReadAll(jsonByteFile)
	if err != nil {
		return nil
	}
	// TODO : Parser / trouver a quelle structure ca va etc ...
	// en vraie c un peu galere
	return nil
}
