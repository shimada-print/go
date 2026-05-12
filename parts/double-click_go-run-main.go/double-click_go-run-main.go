package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// 1. 実行したい go ファイルを指定
	sourceFile := "main.go"

	// go run を実行するコマンドを準備
	// 注意: go run 自体もプロセスを作るため、丁寧に扱います
	cmd := exec.Command("go", "run", sourceFile)

	// 標準入出力を現在のターミナルに繋ぐ
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println(">>> Goプログラムを起動します...")
	fmt.Println("--------------------------------")

	// 実行開始
	err := cmd.Run()

	if err != nil {
		fmt.Printf("\n[!] エラーが発生しました: %v\n", err)
	} else {
		fmt.Println("\n--------------------------------")
		fmt.Println(">>> プログラムは正常に終了しました。")
	}

	// 完了後に画面を止めさせる（ここが重要）
	fmt.Print("\n終了するには Enter キーを押してください...")
	bufio.NewScanner(os.Stdin).Scan()
}
