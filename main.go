package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/cheggaaa/pb/v3"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math"
	"os"
)

type Schema struct {
	Schema []string
}

func touchCsv(name string) (*csv.Writer, *os.File) {
	file, err := os.Create(fmt.Sprintf("csv/%s.csv", name))

	if err != nil {
		log.Fatalf("Fail creating file: %s", err)
	}

	return csv.NewWriter(file), file
}

func writeCsv(writer *csv.Writer, file *os.File, count int64, schema Schema) {
	var i int64

	bar := pb.Full.Start(int(count))

	for i < count {
		i += 1
		_ = writer.Write(makeData(schema))
		bar.Increment()
	}

	bar.Finish()
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

type FakerFunctionString func() string
type FakerFunctionFloat64 func() float64
type FakerFunctionInt64 func() int64

var mapperString = map[string]FakerFunctionString{
	faker.CreditCardNumber:      faker.CCNumber,
	faker.CreditCardType:        faker.CCType,
	faker.EmailTag:              faker.Email,
	faker.DomainNameTag:         faker.DomainName,
	faker.IPV4Tag:               faker.IPv4,
	faker.IPV6Tag:               faker.IPv6,
	faker.PASSWORD:              faker.Password,
	faker.JWT:                   faker.Jwt,
	faker.PhoneNumber:           faker.Phonenumber,
	faker.MacAddressTag:         faker.MacAddress,
	faker.URLTag:                faker.URL,
	faker.UserNameTag:           faker.Username,
	faker.TollFreeNumber:        faker.TollFreePhoneNumber,
	faker.E164PhoneNumberTag:    faker.E164PhoneNumber,
	faker.TitleMaleTag:          faker.TitleMale,
	faker.TitleFemaleTag:        faker.TitleFemale,
	faker.FirstNameTag:          faker.FirstName,
	faker.FirstNameMaleTag:      faker.FirstNameMale,
	faker.FirstNameFemaleTag:    faker.FirstNameFemale,
	faker.LastNameTag:           faker.LastName,
	faker.NAME:                  faker.Name,
	faker.GENDER:                faker.Gender,
	faker.DATE:                  faker.Date,
	faker.TIME:                  faker.TimeString,
	faker.MonthNameTag:          faker.MonthName,
	faker.YEAR:                  faker.YearString,
	faker.DayOfWeekTag:          faker.DayOfWeek,
	faker.DayOfMonthTag:         faker.DayOfMonth,
	faker.TIMESTAMP:             faker.Timestamp,
	faker.CENTURY:               faker.Century,
	faker.TIMEZONE:              faker.Timezone,
	faker.TimePeriodTag:         faker.Timeperiod,
	faker.WORD:                  faker.Word,
	faker.SENTENCE:              faker.Sentence,
	faker.PARAGRAPH:             faker.Paragraph,
	faker.CurrencyTag:           faker.Currency,
	faker.AmountWithCurrencyTag: faker.AmountWithCurrency,
	faker.HyphenatedID:          faker.UUIDHyphenated,
	faker.ID:                    faker.UUIDDigit,
}

var mapperFloat64 = map[string]FakerFunctionFloat64{
	faker.LATITUDE:  faker.Latitude,
	faker.LONGITUDE: faker.Longitude,
}

var mapperInt64 = map[string]FakerFunctionInt64{
	faker.UnixTimeTag: faker.UnixTime,
}

func callFuncByName(name string) string {
	if _, exist := mapperString[name]; exist {
		return mapperString[name]()
	}
	if _, exist := mapperFloat64[name]; exist {
		return fmt.Sprintf("%f", mapperFloat64[name]())
	}
	if _, exist := mapperInt64[name]; exist {
		return fmt.Sprintf("%d", mapperInt64[name]())
	}
	return ""
}

func exportCsv(name string, count int64) {
	schema := getSchema(name)
	writer, file := touchCsv(name)
	writeCsv(writer, file, count, schema)
}

func main() {
	countPtr := flag.Int64("count", int64(math.Pow(10, 3)), "An int")
	tempPtr := flag.String("template", "default", "A template name")
	flag.Parse()
	exportCsv(*tempPtr, *countPtr)
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
