package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var jsonText string

type Objetos struct {
	Y      int
	Height int
	X      int
	Width  int
}

var decoded []Objetos

func main() {
	fJSON, err := ioutil.ReadFile("./city1.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(fJSON, &decoded)

	objeto := make([][4]int, len(decoded))

	for i := 0; i < len(decoded); i++ {
		objeto[i][0] = decoded[i].X
		objeto[i][1] = decoded[i].Width
		objeto[i][2] = decoded[i].Y
		objeto[i][3] = decoded[i].Height
	}

	fmt.Println(objeto)
}