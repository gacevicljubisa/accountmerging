package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	am "github.com/gacevicljubisa/accountmerging/merger"
)

func main() {
	file := "accounts.json"
	accounts, err := getAccounts(file)
	if err != nil {
		fmt.Printf("failed to get '%s': %s\n", file, err.Error())
		return
	}

	accountMerger := am.AccountMerger{}
	persons := accountMerger.Merge(accounts)

	err = write(os.Stdout, persons)
	if err != nil {
		fmt.Printf("failed to write 'persons': %s\n", err.Error())
		return
	}
}

func getAccounts(file string) (accounts []am.Account, err error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("open accounts.json: %s", err.Error())
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("read opened jsonFile as a byte slice: %s", err.Error())
	}

	err = json.Unmarshal(byteValue, &accounts)
	if err != nil {
		return nil, fmt.Errorf("unmarshal jsonFile's contents into 'accounts': %s", err.Error())
	}
	return
}

func write(writer io.Writer, persons []am.Person) error {
	b, err := json.MarshalIndent(persons, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal 'persons' to json: %s", err.Error())
	}
	n, err := writer.Write(b)
	if err != nil {
		return fmt.Errorf("write json failed at byte %v: %s", n, err.Error())
	}
	return nil
}
