package DRYio

import (
	"fmt"
	"io"
	"os"
	"strings"
)
//read byte slise
func ReadByteSlise(S string) []byte {

	file, err := os.Open(S)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 1)
	B := []byte{}
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		//склеиваем в одну строку
		B = append(B, data[:n]...)
	}
	return B
}
func ReadString(S string) string {
	var St string
	file, err := os.Open(S)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 5000)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		//склеиваем в одну строку
		St = St + (string(data[:n]))
	}
	return St
}

func ReadStrSlise(S string) []string {

	words := strings.Fields(ReadString(S)) //преобразуем строку ,разделенную пробелами, в слайс

	return words
}
