package DRYio

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadF(S string) string {

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

func ReadSlise(S string) []string {

	words := strings.Fields(ReadF(S)) //преобразуем строку ,разделенную пробелами, в слайс

	return words
}
