package floatsfile

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func LoadBinary(path string, floats64Num int) []float64 {

	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	r := bytes.NewReader(f)
	floatArr := make([]float64, floats64Num)

	err = binary.Read(r, binary.LittleEndian, &floatArr)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	return floatArr
}

func LoadIntBinary(path string, floats64Num int) []int64 {

	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	r := bytes.NewReader(f)
	intArr := make([]int64, floats64Num)

	err = binary.Read(r, binary.LittleEndian, &intArr)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	return intArr
}
