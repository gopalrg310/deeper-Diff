# deeper-Diff

This will be used to find the json diff from the code itself.

# To Find the Json diff in boolean
Try GetjsonDiffInBool with two comparable json

# To Find the Json diff in Value
Try GetjsonDiffInValue with two comparable json


When there is a comparison between two different json which can be even million line this fuction will give you the exact difference

Access it by `go get github.com/gopalrg310/json-diff`

To Run test cases
`go test -v`


And do 
`go build`

filename1 and filename2 is to take the input file from the user to do comparison

`./json-diff -filename1=json1.json -filename2=json2.json`

To build Docker image `docker build -t json-diff`
