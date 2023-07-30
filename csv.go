package seed

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadTempLateFile(template string) [][]string {
	file, err := os.Open(template)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("read all error")
		panic(err)
	}
	return records
}

func WriteFile(records [][]string, output string) error {
	newFile, err := os.Create(output)
	if err != nil {
		fmt.Println("create file error")
		panic(err)
	}
	defer newFile.Close()
	w := csv.NewWriter(newFile)
	err = w.WriteAll(records)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}
