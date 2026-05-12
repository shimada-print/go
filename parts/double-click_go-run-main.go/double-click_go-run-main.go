package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// 1. 実行したい go ファイルをフルパス（または相対パス）で指定
	// 同じフォルダにあるなら "main.go" でOK
	cmd := exec.Command("go", "run", "main.go")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("プログラムを起動します...")
	err := cmd.Run()

	if err != nil {
		fmt.Printf("エラーが発生しました: %v\n", err)
	}

	// 2. 完了後に画面を止めさせる（これがないと一瞬で消えます）
	fmt.Println("\n終了するには Enter キーを押してください...")
	bufio.NewScanner(os.Stdin).Scan()
}
