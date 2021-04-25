# jsonprocessor
Aim of this library to provide different functionalities to process JSON objects. 

Currently it provides two functionalities :-
1. Merge two JSON objects into array of objects
2. Merge a JSON object into existing array of JSON objects

## How to use it
example 
```
package main

import (
	"fmt"
	"github.com/rushikeshnere/jsonprocessor"
)

func main() {	
	data1 := []byte(`{
        "Id" : 1,
        "Name": "John Doe",
        "Occupation": "gardener"
    }`)

	data2 := []byte(`{
        "Id" : 5,
        "Name": "May",
        "Occupation": "LPQO"
    }`)

	data3 := []byte(`[{
        "Id" : 1,
        "Name": "John Doe",
        "Occupation": "gardener"
    },
    {
        "Id" : 2,
        "Name": "Rushikesh",
        "Occupation": "COEP"
    },
    {
        "Id" : 4,
        "Name": "Chintu",
        "Occupation": "POEUNF"
    }]`)

    data, err := jsonprocessor.MergeJSONDataInArray(data1, data2)
    if(err != nil) {
        panic(err)
    }

	fmt.Println(string(data))
	fmt.Println("")

    data, err = jsonprocessor.MergeJSONDataInExistingJSONArray(data3, data2)

    if(err != nil) {
        panic(err)
    }
	fmt.Println(string(data))
}
```