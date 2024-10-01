# Go言語のバージョンを表示するWebサーバーソフト
拡張子.ps1はWindowPowerScriptだが、LinuxやMacでも使用可。  
本体書類を起動後はファイヤーウォールの問い合わせをOKとし、  
WebブラウザでURL「localhost:8080」などで閲覧。

### プロジェクト構成
go-ver-server_simple-1  
　|- go run main.go.ps1　クリックで本体テキスト書類を起動できる  
　|- go build main.go.ps1　クリックで本体テキスト書類を実行ファイルにできる  
　|- main.exe　クリックでも実行ファイルをコマンドプロンプトで起動できる（go bulidで作成後に出来る書類）  
　|- start_main.exe.ps1　クリックで実行ファイルを起動できる  
　|- README.md　仕様説明書（本書類）
