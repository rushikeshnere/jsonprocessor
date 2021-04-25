package jsonprocessor

import (
	"encoding/json"
)

func MergeJSONDataInArray(jsonData1, jsonData2 []byte) ([]byte, error) {
	var dataMapOne map[string]interface{}
	var dataMapTwo map[string]interface{}
	var meragedJSONDataArray [2]map[string]interface{}

    err := json.Unmarshal(jsonData1, &dataMapOne)
    if err != nil {
		return nil, err
    }

	e := json.Unmarshal(jsonData2, &dataMapTwo)
    if e != nil {
		return nil, e
    }

	meragedJSONDataArray[0] = dataMapOne
	meragedJSONDataArray[1] = dataMapTwo

	mergedJsonData, errorInMarshalling := json.Marshal(meragedJSONDataArray)
	if(errorInMarshalling != nil) {
		return nil, errorInMarshalling
	}

	return mergedJsonData, nil
}

func MergeJSONDataInExistingJSONArray(jsonData1, jsonData2 []byte) ([]byte, error) {
	var dataMapOne []map[string]interface{}
	var dataMapTwo map[string]interface{}

    err := json.Unmarshal(jsonData1, &dataMapOne)
    if err != nil {
		return nil, err
    }

	var meragedJSONDataArray []map[string]interface{}

	e := json.Unmarshal(jsonData2, &dataMapTwo)
    if e != nil {
		return nil, e
    }
	
	for _, value := range dataMapOne {
		meragedJSONDataArray = append(meragedJSONDataArray, value)
	}

	meragedJSONDataArray = append(meragedJSONDataArray, dataMapTwo)

	mergedJsonData, errorInMarshalling := json.Marshal(meragedJSONDataArray)
	if(errorInMarshalling != nil) {
		return nil, errorInMarshalling
	}

	return mergedJsonData, nil
}