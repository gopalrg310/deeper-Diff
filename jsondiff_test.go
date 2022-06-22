package main

import (
	"testing"
	"strings"
)

const(
  content1=`[
	{
		color: "red",
		value: "#f00"
	},
	{
		color: "green",
		value: "#0f0"
	},
	{
		color: "blue",
		value: "#00f"
	},
	{
		color: "cyan",
		value: "#0ff"
	},
	{
		color: "magenta",
		value: "#f0f"
	},
	{
		color: "yellow",
		value: "#ff0"
	},
	{
		color: "black",
		value: "#000"
	}
]`
  content2=`[
	{
		color: "red",
		value: "#f00"
	},
	{
		color: "green",
		value: "#0f0"
	},
	{
		color: "blue",
		value: "#00f"
	},
	{
		color: "violet",
		value: "#7f0"
	},
	{
		color: "Indigo",
		value: "#4b0"
	},
	{
		color: "yellow",
		value: "#ff0"
	},
	{
		color: "black",
		value: "#000"
	}
]`
)

func Test_Jsondiff(t *testing.T){
  	got, old, newval, err := GetDiffJSONValue("bool",content1, content2, false)
	want :=true
	if got != want {
		t.Errorf("got %q, wanted %q and found error %q", got, want, err)
	}
	if !strings.Contains(fmt.Sprint(old),"true"){
		t.Errorf("got %q, wanted %q and found error %q", got, want, err)
	}
	got, old, newval, err = GetDiffJSONValue("string",content1, content2, false)
	if got != want {
		t.Errorf("got %q, wanted %q and found error %q", got, want, err)
	}
	if !strings.Contains(fmt.Sprint(old),"cyan") || !strings.Contains(fmt.Sprint(newval),"violet") {
		t.Errorf("got %q, wanted %q and found error %q", got, want, err)
	}
}
