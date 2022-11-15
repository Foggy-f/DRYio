package DRYio

import (
	"fmt"
	//"io"
	"os"
	//"strings"
)

func WriteFi(S, InText string) string {

	file, err := os.Open(S)
	S = "" //очищаем S от данных, так как она используется далее
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := []byte(InText)
	file.Write(data)
	//склеиваем в одну строку
	S = "ok"

	return S
}
