package main

import (
	"encoding/json"
	"fmt"
	"io"
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

	err = writeJson(os.Stdout, persons)
	if err != nil {
		fmt.Printf("failed to write 'persons': %s\n", err.Error())
	}
}

func getAccounts(file string) (accounts []am.Account, err error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("open file %s: %s", file, err.Error())
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&accounts)
	if err != nil {
		return nil, fmt.Errorf("decode jsonFile's contents into 'accounts': %s", err.Error())
	}
	return
}

func writeJson(writer io.Writer, persons []am.Person) error {
	enc := json.NewEncoder(writer)
	enc.SetIndent("", "  ")
	if err := enc.Encode(persons); err != nil {
		return fmt.Errorf("json encode: %s", err.Error())
	}
	return nil
}
