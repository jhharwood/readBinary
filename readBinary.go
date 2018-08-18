package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

//this type represnts a record with three fields
type Payload struct {
	Time      float64
	Latrad    float64
	Lonrad    float64
	Alt       float64
	Ewspeed   float64
	Nsspeed   float64
	Vertspeed float64
	Roll      float64
	Pitch     float64
	Heading   float64
	Wander    float64
	Ewacc     float64
	Nsacc     float64
	Vertacc   float64
	Xacc      float64
	Yacc      float64
	Zacc      float64
}

func main() {

	readFile()

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile() {

	file, err := os.Open("sbet_RTX_NAD83_180808_1742_a.out")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	m := Payload{}
	newfile, err := os.Create("sbet.txt")
	check(err)
	defer newfile.Close()
	w := bufio.NewWriter(newfile)
	for true {
		data := readNextBytes(file, 136)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &m)
		if err != nil {
			log.Fatal("binary.Read failed", err)
		} else {
			if err == io.EOF {
				break
			}
			//fmt.Println(m)
			_, err = fmt.Fprintf(w, "%v\n", m)
		}
	}
	w.Flush()

}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}
