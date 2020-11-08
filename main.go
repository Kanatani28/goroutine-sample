package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("# main() Start")
	alphabetList := []string{"A", "B", "C", "D", "E"}
	// ① WaitGroupを宣言する
	wg := new(sync.WaitGroup)
	for _, alphabet := range alphabetList {
		// ② 終了待ちするgoroutineの数を設定していく（今回の場合最終的に5になる）
		wg.Add(1)
		// ③ goroutineとして実行
		go sub(alphabet, wg)
	}
	// ⑤ WaitGroupに設定された数だけDoneが実行されるまで待機
	wg.Wait()
	fmt.Println("# main() End")
}

// goroutineで実行する関数（引数にWaitGroup追加）
func sub(val string, wg *sync.WaitGroup) {
	// ④ 処理終了時にDoneメソッド実行するようにする
	defer wg.Done()
	fmt.Println("## sub() " + val)
}
