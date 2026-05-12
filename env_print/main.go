package main

import (
	"fmt"
	"os"
)

func main() {
	// すべての環境変数を取得
	envs := os.Environ()

	// 1つずつ取り出して表示
	for _, env := range envs {
		fmt.Println(env)
	}
}
