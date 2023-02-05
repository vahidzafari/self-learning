package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// You have a utility that processes its command-line arguments; if everything goes
// as expected, then you get the supported types of command-line arguments and
// everything goes smoothly. But when something unexpected occurs the
// map[string]interface{} map is here to help.

// Using map[string]interface{} for storing these unknown JSON records (arbitrary data) is a
// good choice because map[string]interface{} is good at storing JSON records of an unknown type.

// Remember that the biggest advantage you get from using a map[string]interface{}
// map or any map that stores an interface{} value in general, is that you still have
// your data in its original state and data type. If you use map[string]string instead,
// or anything similar, then any data you have is going to be converted into a string,
// which means that you are going to lose information about the original data type and
// the structure of the data you are storing in the map.

var JSONrecord = `{
	"Flag": true,
	"Array": ["a","b","c"],
	"Entity": {
		"a1": "b1",
		"a2": "b2",
		"Value": -456,
		"Null": null
	},
	"Message": "Hello Go!"
	}`

func typeSwitch(m map[string]interface{}) {
	for k, v := range m {
		switch c := v.(type) {
		case string:
			fmt.Println("Is a string!", k, c)
		case float64:
			fmt.Println("Is a float64!", k, c)
		case bool:
			fmt.Println("Is a Boolean!", k, c)
		case map[string]interface{}:
			fmt.Println("Is a map!", k, c)
			typeSwitch(v.(map[string]interface{}))
		default:
			fmt.Printf("...Is %v: %T!\n", k, c)
		}
	}
	return
}

func exploreMap(m map[string]interface{}) {
	for k, v := range m {
		embMap, ok := v.(map[string]interface{})
		// If it is a map, explore deeper
		if ok {
			fmt.Printf("{\"%v\": \n", k)
			exploreMap(embMap)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v: %v\n", k, v)
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("*** Using default JSON record.")
	} else {
		JSONrecord = os.Args[1]
	}
	JSONMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(JSONrecord), &JSONMap)

	if err != nil {
		fmt.Println(err)
		return
	}

	exploreMap(JSONMap)
	fmt.Println("-----------------------")
	typeSwitch(JSONMap)
}
