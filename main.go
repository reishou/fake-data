package main

import (
	"encoding/csv"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math"
	"os"
	"time"
)

type Schema struct {
	Schema []string
}

func logTime() {
	now := time.Now()
	fmt.Println(now.Unix())
}

func touchCsv(name string) (*csv.Writer, *os.File) {
	file, err := os.Create(fmt.Sprintf("csv/%s.csv", name))

	if err != nil {
		log.Fatalf("Fail creating file: %s", err)
	}

	return csv.NewWriter(file), file
}

func writeCsv(writer *csv.Writer, file *os.File, n int64, schema Schema) {
	var i int64
	for i < n {
		i += 1
		_ = writer.Write(makeData(schema))
	}

	writer.Flush()
	_ = file.Close()
}

func makeData(schema Schema) []string {
	var data []string

	for _, v := range schema.Schema {
		value := callFuncByName(v)
		data = append(data, value)
	}

	return data
}

func callFuncByName(name string) string {
	if name == "lat" {
		return fmt.Sprintf("%f", faker.Latitude())
	} else if name == "long"{
		return fmt.Sprintf("%f", faker.Longitude())
	} else if name == "cc_number"{
		return faker.CCNumber()
	} else if name == "cc_type"{
		return faker.CCType()
	}
	return faker.Sentence()
}

func exportCsv(name string, schema Schema) {
	writer, file := touchCsv(name)
	logTime()
	n := 2 * math.Pow(10, 6)
	writeCsv(writer, file, int64(n), schema)
	logTime()
}

func main() {
	implementArgs(func(arg string) {
		schema := getSchema(arg)
		exportCsv(arg, schema)
	})
}

func getSchema(arg string) Schema {
	source, err := ioutil.ReadFile("template/" + arg + ".yml")
	if err != nil {
		panic(err)
	}
	var schema Schema
	_ = yaml.Unmarshal(source, &schema)
	return schema
}

func implementArgs(f func(arg string)) {
	args := os.Args
	for index, arg := range args {
		if index > 0 {
			f(arg)
		}
	}
}
