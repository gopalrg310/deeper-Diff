package main

import (
	"testing"
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

func Test_Jsondiff(t *testing.T){
  got, err := interpret_bf(helloworld, "")
	want := "Hello World!\n"
	if got != want {
		t.Errorf("got %q, wanted %q and found error %q", got, want, err)
	}
	got, err = interpret_bf(WorldLine, "")
	want = "WorldLine Ingenico"
	if got != want {
		t.Errorf("got %q, wanted %q and found error %q", got, want, err)
	}
}
