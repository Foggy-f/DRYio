package DRYio

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadString(S string) string {

	file, err := os.Open(S)
	S = "" //очищаем S от данных, так как она используется далее
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
		S = S + (string(data[:n]))
	}
	return S
}

func ReadStrSlise(S string) []string {

	words := strings.Fields(ReadString(S)) //преобразуем строку ,разделенную пробелами, в слайс

	return words
}

func WriteF(S string, data []byte) int {

	file, err := os.OpenFile(S, os.O_APPEND|os.O_WRONLY, 0644)
	S = "" //очищаем S от данных, так как она используется далее
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	//data := []byte(InText)
	N, err := file.Write(data)
	if err != nil {
		fmt.Println(file, err)
	}
	//склеиваем в одну строку
	//S = "ok"

	return N
}
