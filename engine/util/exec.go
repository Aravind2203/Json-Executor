package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Resp struct {
	Message string
}
type Block struct {
	Id         int
	Stype      string
	Method     string
	Url        string
	Identifier string
	Next       int
}
type MainBlock struct {
	Blocks []Block
}

type Binding map[string]interface{}
type Mapping map[int]Block

func MainExecute(data []byte) []byte {
	mapping := make(Mapping)
	binding := make(Binding)
	fmt.Println("It's here")
	var m MainBlock
	err := json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(m.Blocks[0])
	}
	for i := 0; i < len(m.Blocks); i++ {
		mapping[m.Blocks[i].Id] = m.Blocks[i]
	}
	j := 1
	for j != -1 {
		if mapping[j].Stype == "REQUEST" {
			url := mapping[j].Url
			r, err := http.Get(url)
			if err != nil {
				break
			}
			body, _ := ioutil.ReadAll(r.Body)
			binding[mapping[j].Identifier] = string(body)
		} else if mapping[j].Stype == "RETURN" {
			resp, _ := json.Marshal(
				binding[mapping[j].Identifier],
			)
			return resp
		}
		j = mapping[j].Next
	}
	resp, _ := json.Marshal(
		Resp{
			Message: "Welcome to the execution engine",
		},
	)
	return resp

}
