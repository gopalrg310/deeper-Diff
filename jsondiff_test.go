package main

import (
	"testing"
	"strings"
	"fmt"
	"encoding/json"
)

const(
  content1=`[
	{
		"color": "red",
		"value": "#f00"
	},
	{
		"color": "green",
		"value": "#0f0"
	},
	{
		"color": "blue",
		"value": "#00f"
	},
	{
		"color": "cyan",
		"value": "#0ff"
	},
	{
		"color": "magenta",
		"value": "#f0f"
	},
	{
		"color": "yellow",
		"value": "#ff0"
	},
	{
		"color": "black",
		"value": "#000"
	}
]`
  content2=`[
	{
		"color": "red",
		"value": "#f00"
	},
	{
		"color": "green",
		"value": "#0f0"
	},
	{
		"color": "blue",
		"value": "#00f"
	},
	{
		"color": "violet",
		"value": "#7f0"
	},
	{
		"color": "Indigo",
		"value": "#4b0"
	},
	{
		"color": "yellow",
		"value": "#ff0"
	},
	{
		"color": "black",
		"value": "#000"
	}
]`
)

func Test_Jsondiff(t *testing.T){
	val1:=make([]interface{},0)
	val2:=make([]interface{},0)

	json.Unmarshal([]byte(content1),&val1)
	json.Unmarshal([]byte(content2),&val2)
	
  	got, newval, old, err := GetjsonDiffInBool(val1, val2)
	want :=true
	/* the value of old will be 
	[map[color:true value:true] map[color:true value:true]]
	*/
	if got != want {
		t.Errorf("got %v, wanted %v and found error %v", got, want, err)
	}
	if !strings.Contains(fmt.Sprint(old),"true"){
		t.Errorf("got %v, wanted %v and found error %v", got, want, err)
	}
	got, newval, old, err = GetjsonDiffInValue(val1, val2)
	/*
	value of
	Old - [map[color:cyan value:#0ff] map[color:magenta value:#f0f]]
	New - [map[color:violet value:#7f0] map[color:Indigo value:#4b0]]
	*/
	if got != want {
		t.Errorf("got %v, wanted %v and found error %v", got, want, err)
	}
	if !strings.Contains(fmt.Sprint(old),"cyan") || !strings.Contains(fmt.Sprint(newval),"violet") {
		t.Errorf("got %v, wanted %v and found error %v", got, want, err)
	}
}
