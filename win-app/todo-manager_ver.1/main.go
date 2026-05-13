package main

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
)

const fileName = "todo.txt"

func main() {
	ensureFile()

	// PowerShellスクリプト：GUI内で全ての処理を完結させる
	psCode := []string{
		"Add-Type -AssemblyName PresentationFramework",
		// XAML（UI設計図）
		"$xaml = '<Window xmlns=\"http://schemas.microsoft.com/winfx/2006/xaml/presentation\" Title=\"ToDoマネージャー ver.1\" Height=\"450\" Width=\"375\" Topmost=\"True\" WindowStartupLocation=\"CenterScreen\"><Grid Margin=\"10\"><Grid.RowDefinitions><RowDefinition Height=\"Auto\"/><RowDefinition Height=\"*\"/><RowDefinition Height=\"Auto\"/><RowDefinition Height=\"Auto\"/></Grid.RowDefinitions><TextBlock Grid.Row=\"0\" Text=\"ToDo一覧\" FontSize=\"18\" Margin=\"0,0,0,10\"/><ListBox Grid.Row=\"1\" Name=\"TodoBox\" Margin=\"0,0,0,10\"/><StackPanel Grid.Row=\"2\" Orientation=\"Horizontal\" Margin=\"0,0,0,10\"><TextBox Name=\"InputBox\" Width=\"240\" Margin=\"0,0,5,0\"/><Button Name=\"AddBtn\" Content=\"追加\" Width=\"75\"/></StackPanel><Button Grid.Row=\"3\" Name=\"DelBtn\" Content=\"削除\" Height=\"30\" Background=\"MistyRose\"/></Grid></Window>'",
		"$reader = [System.Xml.XmlReader]::Create([System.IO.StringReader]$xaml)",
		"$window = [Windows.Markup.XamlReader]::Load($reader)",
		"$todoBox = $window.FindName(\"TodoBox\")",
		"$inputBox = $window.FindName(\"InputBox\")",
		
		// リストを更新する関数（BOMを考慮して読み込み）
		"function Update-List { $todoBox.Items.Clear(); if(Test-Path \"" + fileName + "\"){ $content = Get-Content \"" + fileName + "\" -Encoding UTF8; foreach($line in $content){ if($line.Trim()){ [void]$todoBox.Items.Add($line.Trim().TrimStart([char]0xfeff)) } } } }",
		"Update-List",

		// 追加ボタンの動作
		"$window.FindName(\"AddBtn\").Add_Click({ if($inputBox.Text -ne ''){ Add-Content \"" + fileName + "\" ($inputBox.Text + \"`r`n\") -Encoding UTF8; $inputBox.Text=''; Update-List } })",
		
		// 削除ボタンの動作
		"$window.FindName(\"DelBtn\").Add_Click({ if($todoBox.SelectedItem){ $target = $todoBox.SelectedItem; $lines = (Get-Content \"" + fileName + "\" -Encoding UTF8) | Where-Object { $_.Trim().TrimStart([char]0xfeff) -ne $target }; if($lines){ $lines | Set-Content \"" + fileName + "\" -Encoding UTF8 } else { Clear-Content \"" + fileName + "\" }; Update-List } })",
		
		"$window.ShowDialog() | Out-Null",
	}

	// 隠しウィンドウでPowerShellを起動
	cmd := exec.Command("powershell", "-NoProfile", "-WindowStyle", "Hidden", "-Command", strings.Join(psCode, ";"))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	
	// PowerShellが終了するまで実行を維持
	cmd.Run()
}

func ensureFile() {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// UTF-8 BOM付きで空ファイルを作成
		os.WriteFile(fileName, []byte{0xEF, 0xBB, 0xBF}, 0644)
	}
}
