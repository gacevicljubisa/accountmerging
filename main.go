package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	am "github.com/gacevicljubisa/accountmerging/merger"
)

func main() {
	fmt.Println("hello")

	// Open our jsonFile
	jsonFile, err := os.Open("accounts.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Successfully opened accounts.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte slice.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// we initialize our Accounts slice
	var accounts []am.Account

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'accounts' which we defined above
	err = json.Unmarshal(byteValue, &accounts)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	accountMerger := am.AccountMerger{}
	persons := accountMerger.Merge(accounts)

	b, err := json.MarshalIndent(persons, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	os.Stdout.Write(b)
}
