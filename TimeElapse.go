package DRYio

import (
	"fmt"
	"time"
)

func cookies(k int32, a []int32) int32

func TimeElapse(somefunc func(A []int32) int32) {

	//проверка на время
	start := time.Now()
	res := somefunc()
	elapsed := time.Since(start)

	// проверка на результат
	//	if r != res {
	//		panic(fmt.Sprintf("%s - expected: %d result: %d\n", path, r, res))
	//	} else {
	fmt.Printf(" - ok in %s result:%v\n", elapsed, res)
	//	}

}
