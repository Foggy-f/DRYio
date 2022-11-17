package DRYio

import (
	"fmt"
	//"io"
	"os"
	//"strings"
)

func WriteF(S string, data []byte) int {

	file, err := os.OpenFile(S, os.O_APPEND|os.O_WRONLY, 0644)
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

	return N
}

func WriteNew(S string) string {

	file, err := os.Create(S)

	if err != nil {
		fmt.Println(file, err)
	}
	defer file.Close()

	return "ok"
}
