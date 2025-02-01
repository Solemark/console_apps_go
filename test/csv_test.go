package console_apps

import (
	"encoding/csv"
	"os"
	"reflect"
	console_apps "solemarc/go/console_apps/src"
	"testing"
)

func TestFileExists(t *testing.T) {
	filename := "test.csv"
	data := [][]string{
		{"Greeting", "Count to 5", "3rd Header"},
		{"Hello World!", "1, 2, 3, 4, 5", "3rd Col. Data"},
		{"How are you?", "5, 4, 3, 2, 1", "More random data"},
	}
	csv := console_apps.CSV{filename, data}
	csv.Write()
	_, err := os.Stat(filename)
	if err != nil {
		t.Fatalf("Couldn't find %s file\n%s", filename, err)
	}
	csv.Destroy()
}

func TestFileHasData(t *testing.T) {
	filename := "test.csv"
	data := [][]string{
		{"Greeting", "Count to 5", "3rd Header"},
		{"Hello World!", "1, 2, 3, 4, 5", "3rd Col. Data"},
		{"How are you?", "5, 4, 3, 2, 1", "More random data"},
	}
	c := console_apps.CSV{filename, data}
	c.Write()
	input, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	reader := csv.NewReader(input)
	result, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	} else {
		if !reflect.DeepEqual(data, result) {
			t.Log("Expected:\n", data)
			t.Log("result:\n", result)
			t.Fatalf("Expected array does not match result!")
		}
	}
	c.Destroy()
}
