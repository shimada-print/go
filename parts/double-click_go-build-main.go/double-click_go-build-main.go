package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// コンパイル対象と出力名の設定
	source := "main.go"
	output := "main.exe"

	// 1. go build コマンドを構築
	// -ldflags="-H windowsgui" により、生成された exe が黒い画面を出さないように設定
	cmd := exec.Command("go", "build", "-ldflags=-H windowsgui", "-o", output, source)

	// 2. このツール実行中に黒い画面（コンソール）を出さないための設定
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000, // CREATE_NO_WINDOW: プロセス作成レベルでウィンドウを抑制
	}

	// 3. コンパイルを実行し、終了するまで「待機」する
	// Run() を使うことで、ビルドが終わるまで Go が責任を持ってプロセスを管理します
	err := cmd.Run()

	// 4. 失敗した場合のみエラーログを吐く
	if err != nil {
		os.WriteFile("build_error.log", []byte(err.Error()), 0644)
		return
	}

	// 5. 成功時はここで明示的に終了
	// これにより、バックグラウンドに「残骸」が残るのを防ぎます
	os.Exit(0)
}
