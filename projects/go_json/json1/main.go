package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main()  {
	 // Open our jsonFile
	 jsonFile, err := os.Open("browserhistory.json")
	 // if we os.Open returns an error then handle it
	 if err != nil {
		 fmt.Println(err)
	 }
 
	 fmt.Println("Successfully Opened browserhistory.json")
	 // defer the closing of our jsonFile so that we can parse it later on
	 defer jsonFile.Close()
 
	 // read our opened xmlFile as a byte array.
	 byteValue, _ := ioutil.ReadAll(jsonFile)
 
	 browse, err := UnmarshalBrowserHistories(byteValue)
	 if	err != nil {
		fmt.Println(err)
	 }
	 // we iterate through every user within our browserhistory array and
	 // print out the browserhistory title and ids
	 // as just an example
	 idMap := make(map[int64]string)
	 for i := 0; i < len(browse); i++ {
		idMap[browse[i].ID] = browse[i].Title 
	 }

	 for k, v := range idMap{
		fmt.Println("key: ", k, "value: ", v)
	 }
}

type BrowserHistories []BrowserHistory

func UnmarshalBrowserHistories(data []byte) (BrowserHistories, error) {
	var r BrowserHistories
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BrowserHistories) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type BrowserHistory struct {
	ID                 int64  `json:"id"`
	Title              string `json:"title"`
	URL                string `json:"url"`
	LastVisitTimeLocal string `json:"lastVisitTimeLocal"`
	LastVisitTimeUTC   string `json:"lastVisitTimeUTC"`
	TypedCount         int64  `json:"typedCount"`
	VisitCount         int64  `json:"visitCount"`
}