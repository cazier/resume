package marshall

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	var json_string string = `{
		"basics": {
		  "name": "John Doe",
		  "label": "Programmer",
		  "image": "",
		  "email": "john@gmail.com",
		  "phone": "(912) 555-4321",
		  "url": "https://johndoe.com",
		  "summary": "A summary of John Doe…",
		  "location": {
			"address": "2712 Broadway St",
			"postalCode": "CA 94115",
			"city": "San Francisco",
			"countryCode": "US",
			"region": "California"
		  },
		  "profiles": [{
			"network": "Twitter",
			"username": "john",
			"url": "https://twitter.com/john"
		  },
		  {
			"network": "Reddit",
			"username": "john",
			"url": "https://reddit.com/john"
		  }]
		}
	}`
	var input []byte = []byte(json_string)
	fmt.Println(len(input))

	var resume Resume = LoadJsonFile("../../resume.json") //(input)
	// var resume Resume = LoadJsonString(input)

	// t.Log(resume)
	t.Log(resume)

}
