package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

type DeeperDiff struct {
	exact       interface{}
	comparative interface{}
}

// This fucntion will be used to find the deeper diff which is similar like http://www.jsondiff.com/

/*input
types -  Whether it needs to be bool or string of exact difference
Exact - compare content of orginal
comparitive - compare content of comparitive
Changes - Use this as initial false that used internally
*/
func GetDiffJSONValue(types string, Exact interface{}, comparitive interface{}, Changes bool) (bool, interface{}, interface{}, error) {
	var data1, data2 interface{}
	exactduplicate := make(map[string]interface{})
	var err error
	switch reflect.TypeOf(comparitive).Kind().String() {
	case "map":
		if reflect.TypeOf(comparitive) == reflect.TypeOf(Exact) {
			result_map := make(map[string]interface{})
			old_map := make(map[string]interface{})
			for key, value := range comparitive.(map[string]interface{}) {
				if avail_value, ok := Exact.(map[string]interface{})[key]; !ok {
					Changes = true
					instant_result1, instant_result2 := GetDiffOutput(types, avail_value, value)
					result_map[key] = instant_result1
					old_map[key] = instant_result2
				} else if reflect.DeepEqual(avail_value, value) {
					exactduplicate[key] = true
				} else {
					exactduplicate[key] = true
					value_type := reflect.ValueOf(value).Kind().String()
					if reflect.TypeOf(avail_value) != reflect.TypeOf(value) {
						Changes = true
						instant_result1, instant_result2 := GetDiffOutput(types, avail_value, value)
						result_map[key] = instant_result1
						old_map[key] = instant_result2
					} else if value_type == "map" || value_type == "slice" {
						if Changes, data1, data2, err = GetDiffJSONValue(types, avail_value, value, Changes); data1 != nil {
							result_map[key] = data1
							old_map[key] = data2
						} else if err != nil {
							return false, nil, nil, err
						}
					} else {
						Changes = true
						instant_result1, instant_result2 := GetDiffOutput(types, avail_value, value)
						result_map[key] = instant_result1
						old_map[key] = instant_result2
					}
				}
			}
			for key, value := range Exact.(map[string]interface{}) {
				if _, ok := exactduplicate[key]; !ok {
					Changes = true
					instant_result1, instant_result2 := GetDiffOutput(types, value, nil)
					result_map[key] = instant_result1
					old_map[key] = instant_result2
				}
			}
			if len(result_map) > 0 {
				return Changes, result_map, old_map, nil
			}
			return Changes, nil, nil, nil
		} else {
			Changes = true
			instant_result1, instant_result2 := GetDiffOutput(types, Exact, comparitive)
			return Changes, instant_result1, instant_result2, nil
		}

	case "slice":
		var resultant_val, newslice_val interface{}
		if reflect.TypeOf(Exact) == reflect.TypeOf(comparitive) {
			result_slice := make([]interface{}, 0)
			old_slice := make([]interface{}, 0)
			if _, ok := comparitive.([]string); ok {
				for i, j := range comparitive.([]string) {
					if len(Exact.([]string)) <= i {
						if Changes, resultant_val, newslice_val, err = GetDiffJSONValue(types, nil, j, Changes); resultant_val != nil {
							result_slice = append(result_slice, resultant_val)
							old_slice = append(old_slice, newslice_val)
						}
					} else if Changes, resultant_val, newslice_val, err = GetDiffJSONValue(types, Exact.([]string)[i], j, Changes); resultant_val != nil {
						result_slice = append(result_slice, resultant_val)
						old_slice = append(old_slice, newslice_val)
					}
				}
				if len(Exact.([]string)) > len(comparitive.([]string)) {
					for _, j := range Exact.([]string)[len(comparitive.([]string)):] {
						instant_result1, instant_result2 := GetDiffOutput(types, j, nil)
						result_slice = append(result_slice, instant_result1)
						old_slice = append(old_slice, instant_result2)
					}
				}
				if len(result_slice) > 0 {
					Changes = true
					return Changes, result_slice, old_slice, nil
				}
				return Changes, nil, nil, nil
			} else {
				for i, j := range comparitive.([]interface{}) {
					if len(Exact.([]interface{})) <= i {
						if Changes, resultant_val, newslice_val, err = GetDiffJSONValue(types, nil, j, Changes); resultant_val != nil {
							result_slice = append(result_slice, resultant_val)
							old_slice = append(old_slice, newslice_val)
						}
					} else if Changes, resultant_val, newslice_val, err = GetDiffJSONValue(types, Exact.([]interface{})[i], j, Changes); resultant_val != nil {
						result_slice = append(result_slice, resultant_val)
						old_slice = append(old_slice, newslice_val)
					}
				}
				if len(Exact.([]interface{})) > len(comparitive.([]interface{})) {
					for _, j := range Exact.([]interface{})[len(comparitive.([]interface{})):] {
						instant_result1, instant_result2 := GetDiffOutput(types, j, nil)
						result_slice = append(result_slice, instant_result1)
						old_slice = append(old_slice, instant_result2)
					}
				}
				if len(result_slice) > 0 {
					Changes = true
					return Changes, result_slice, old_slice, nil
				}
				return Changes, nil, nil, nil
			}
		} else {
			Changes = true
			return Changes, comparitive, Exact, nil
		}

	default:
		if !reflect.DeepEqual(comparitive, Exact) {
			Changes = true
			instant_result1, instant_result2 := GetDiffOutput(types, Exact, comparitive)
			return Changes, instant_result1, instant_result2, nil
		} else {
			return Changes, nil, nil, nil
		}
	}
}
func (diff *DeeperDiff) GetDiffJSON(types string) (bool, interface{}, interface{}, error) {
	return GetDiffJSONValue(types, diff.exact, diff.comparative, false)
}

func GetDiffOutput(types string, old, new interface{}) (interface{}, interface{}) {
	switch types {
	case "bool":
		return true, nil
	default:
		return new, old
	}
}

//Function will prove the key whch got modified as tru or false
func (diff *DeeperDiff) GetjsonDiffInBool() (bool, interface{}, interface{}, error) {
	return diff.GetDiffJSON("bool")
}

//Function will exactly show that what value got changed
func (diff *DeeperDiff) GetjsonDiffInValue() (bool, interface{}, interface{}, error) {
	return diff.GetDiffJSON("value")
}

func main() {
	filename1 := "json1.json"
	filename2 := "json2.json"
	file1, err := ioutil.ReadFile(filename1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	file2, err := ioutil.ReadFile(filename2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	value1, value2 := make([]interface{}, 0), make([]interface{}, 0)
	json.Unmarshal(file1, &value1)
	json.Unmarshal(file2, &value2)
	deeperdata := new(DeeperDiff)
	deeperdata.exact = value1
	deeperdata.comparative = value2
	diff, old, new, err := deeperdata.GetjsonDiffInBool()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if diff {
		fmt.Println(ConvertJson(old))
		//[{"id":true,"type":true},{"id":true,"type":true}]
	}
	diff, old, new, err = deeperdata.GetjsonDiffInValue()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if diff {
		fmt.Println(ConvertJson(old))
		fmt.Println(ConvertJson(new))
		/*
			[{"id":"2489651046","type":"CreatedEvent"},{"id":"2489651052","type":"PushedEvent"}]
			[{"id":"2489651045","type":"CreateEvent"},{"id":"2489651051","type":"PushEvent"}]
		*/

	}
}
func ConvertJson(data interface{}) string {
	jsondata, _ := json.Marshal(&data)
	return string(jsondata)
}
