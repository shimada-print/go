package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Windows標準の「start」コマンドを介して、システム既定のパスを強制的に開きます。
	// /C は実行後にコマンドプロンプトを閉じるための引数です。
	cmd := exec.Command("cmd", "/C", "start", "", "shell:startup")

	// プロセス残留防止：OSレベルでウィンドウを完全に隠し、独立させます。
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000, // CREATE_NO_WINDOW (ゴースト化阻止)
	}

	// 実行（OSに命令を投げる）
	err := cmd.Start()

	if err != nil {
		// 失敗した場合は、直接フルパスを叩く予備手段を実行
		fallback := exec.Command("explorer", os.Getenv("USERPROFILE")+`\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup`)
		fallback.SysProcAttr = cmd.SysProcAttr
		fallback.Start()
	}

	// 鉄則：命令が完了した瞬間に、このGoプログラム自身は即座に終了します。
	// これにより、プロセスがタスクマネージャーに残ることは物理的に不可能です。
	os.Exit(0)
}
