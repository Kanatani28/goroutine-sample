package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("# main() Start")
	alphabetList := []string{"A", "B", "C", "D", "E"}
	for _, alphabet := range alphabetList {
		// goroutineとして実行
		go sub(alphabet)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("# main() End")
}

// goroutineで実行する関数
func sub(val string) {
	fmt.Println("## sub() " + val)
}
