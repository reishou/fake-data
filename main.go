package main

import (
	"encoding/csv"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"log"
	"math"
	"os"
	"time"
)

type User struct {
	name     string
	email    string
	password string
	birthday time.Time
	phone    string
	address  string
}

func (u User) toString() []string {
	return []string{
		u.name,
		u.email,
		u.password,
		u.birthday.Format("Y-m-d"),
		u.phone,
		u.address,
	}
}

func logTime() {
	now := time.Now()
	fmt.Printf("%d\n", now.Unix())
}

func touchCsv(name string) (*csv.Writer, *os.File) {
	file, err := os.Create("csv/%s.csv")

	if err != nil {
		log.Fatalf("Fail creating file: %s", err)
	}

	return csv.NewWriter(file), file
}

func writeCsv(writer *csv.Writer, file *os.File, n int64) {
	var i int64
	for i < n {
		i += 1
		var user User
		_ = faker.FakeData(&user)
		_ = writer.Write(user.toString())
	}

	writer.Flush()
	_ = file.Close()
}

func main() {
	writer, file := touchCsv("data")
	logTime()
	n := math.Pow(10, 7)
	writeCsv(writer, file, int64(n))
	logTime()
}
