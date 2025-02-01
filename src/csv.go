package console_apps

import (
	"encoding/csv"
	"log"
	"os"
)

type CSV struct {
	Filename string
	Data     [][]string
}

func (c CSV) Write() {
	file, err := os.Create(c.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.WriteAll(c.Data)
}

func (c CSV) Destroy() {
	os.Remove(c.Filename)
}
