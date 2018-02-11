package main

import (
	"io"
	"os"
	"fmt"
	"encoding/csv"
	"strconv"
	"reflect"
)

func main() {

	csvFile, _ := os.Open("exam.csv")
	var powers Powers
	powers, _ = readCSV(csvFile)
	csvFile.Close()
	//fmt.Println(powers)
	fmt.Println(reflect.TypeOf(powers))
	csvFile2, _ := os.Create("toFile.csv")
	powers.ToCsv(csvFile2)
	csvFile2.Close()
}
type Power struct {
	TimeStamp string
	DeltaTime float64
	AvgCount int
}
func readCSV(b io.Reader) ([]Power, error) {
	var powers Powers
	r := csv.NewReader(b)
	fmt.Println(r.Read())
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			break//return nil, err
		}
		TimeStamp := record[0]
		DeltaTime,_ := strconv.ParseFloat(record[1],64)
		AvgCount,_ := strconv.Atoi(record[2])
		p := Power{TimeStamp, float64(DeltaTime), int(AvgCount)}
		powers = append(powers, p)
	}
	return powers, nil
}

type Powers []Power

func (powers *Powers) ToCsv(w io.Writer) {
	n := csv.NewWriter(w)
	n.Write([]string{"TimeStamp", "DeltaT", "AvgCt"})
	for _, power := range *powers {
		n.Write([]string{power.TimeStamp, strconv.FormatFloat(power.DeltaTime,'g',1,64), string(power.AvgCount)})
	}
	n.Flush()
}