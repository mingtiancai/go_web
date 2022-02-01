package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:"author`
}

func main() {
	post := Post{
		Id:      "1",
		Content: "Hello world",
		Author: Author{
			Id:   "2",
			Name: "Sau Sheng",
		},
	}
	// output, err := xml.Marshal(&post)
	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to xml ", err)
		return
	}
	err = ioutil.WriteFile("post.xml",
		[]byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file: ", err)
		return
	}
}
