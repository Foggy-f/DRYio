package DRYio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// read byte slise
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
func ReadStringFromSTD(s string) (string, error) {

	file, err := os.Open(s)
	if err != nil {
		return "", err
	}
	defer file.Close()

	r := bufio.NewScanner(file)
	// handle first encountered error while reading
	if err = r.Err(); err != nil {
		err = fmt.Errorf("error while reading file %s: %s", s, err)
		return "", err
	}
	st := ""
	for r.Scan() {
		st = r.Text()
	}

	return st, err
}

func WaitInput() {

	r := bufio.NewScanner(os.Stdin)
	// handle first encountered error while reading
	if err := r.Err(); err != nil {
		return
	}
	fmt.Println("Нажмите любую клаишую")
	for r.Scan() {
		if r.Text() != "" {
			break
		}
	}
}
