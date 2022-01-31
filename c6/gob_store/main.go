package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	COntent string
	Author  string
}

func store(data interface{}, fileName string) {
	buffer := new(bytes.Buffer)
	encode := gob.NewEncoder(buffer)
	err := encode.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileName, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, fileName string) {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 1, COntent: "hello world", Author: "ssau sheong"}
	store(post, "post1")
	var postRead Post
	load(&postRead, "post1")
	fmt.Println(postRead)
}
