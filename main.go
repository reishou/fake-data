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
	Name     string `faker:"name"`
	Email    string `faker:"email"`
	Password string `faker:"password"`
	Birthday string `faker:"date"`
	Phone    string `faker:"e_164_phone_number"`
	Address  string `faker:"mac_address"`
}

func (u User) toString() []string {
	return []string{
		u.Name,
		u.Email,
		u.Password,
		u.Birthday,
		u.Phone,
		u.Address,
	}
}

func logTime() {
	now := time.Now()
	fmt.Printf("%d\n", now.Unix())
}

func touchCsv(name string) (*csv.Writer, *os.File) {
	file, err := os.Create("csv/" + name + ".csv")

	if err != nil {
		log.Fatalf("Fail creating file: %s", err)
	}

	return csv.NewWriter(file), file
}

func writeCsv(writer *csv.Writer, file *os.File, n int64) {
	var i int64
	for i < n {
		i += 1
		user := User{}
		_ = faker.FakeData(&user)
		_ = writer.Write(user.toString())
	}

	writer.Flush()
	_ = file.Close()
}

func main() {
	args := os.Args

	for index, arg :=range args{
		if index > 0  {
			if file, err := os.Stat("template/" + arg + ".yml"); err == nil {
				fmt.Print(file)
			}
		}
	}
	//file, _ = os.ReadFile("template/")
	//writer, file := touchCsv("data")
	//logTime()
	//n := math.Pow(10, 7)
	//writeCsv(writer, file, int64(n))
	//logTime()
}
