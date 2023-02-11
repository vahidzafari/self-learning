package main

import (
	"encoding/xml"
	"fmt"
)

// The idea behind XML and Go is the same as with JSON and Go. You put tags in Go
// structures in order to specify the XML tags and you can still serialize and deserialize
// XML records using xml.Unmarshal() and xml.Marshal(), which are found in the
// encoding/xml package.

type Employee struct {
	XMLName   xml.Name `xml:"employee"`
	ID        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Height    float32  `xml:"height,omitempty"`
	Comment   string   `xml:",comment"`
	Address
}

type Address struct {
	City, Country string
}

func main() {
	r := Employee{ID: 7, FirstName: "Mihalis", LastName: "Tsoukalos"}
	r.Comment = "Technical Writer + DevOps"
	r.Address = Address{"SomeWhere 12", "12312, Greece"}
	output, err := xml.MarshalIndent(&r, "	", "	")
	if err != nil {
		fmt.Println("Error:", err)
	}
	output = []byte(xml.Header + string(output))
	fmt.Printf("%s\n", output)
}
