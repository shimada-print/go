# Go言語のバージョンを表示するWebサーバーソフト
本体書類を起動後はファイヤーウォールの問い合わせをOKとし、WebブラウザでURL「localhost:8080」（ローカルホストをプライベートIPにすると他のPCというLANでも閲覧可）などでGoのバージョンが閲覧できれば成功。

### プロジェクト構成
go-ver-server_simple-1  
　|- main.go　本体テキスト書類（VSコードなどで編集する書類）  
　|- go run main.go.ps1　クリックで本体テキスト書類を起動できる  
　|- go build main.go.ps1　クリックで本体テキスト書類を実行ファイルにできる  
　|- main.exe　クリックでも実行ファイルをコマンドプロンプトで起動できる（go bulidで作成後に出来る書類）  
　|- start_main.exe.ps1　クリックで実行ファイルを起動できる  
　|- README.md　仕様説明書（本書類）  
<p>拡張子.ps1はWindowPowerShellScriptだが、PowerShellCoreを設定すればLinuxやMacでも使用可かもしれない。</p>

## シンプル版以外でWebブラウザーも同時に起動することもできる
## バージョンだけでなく他にgo envで環境変数なども表示可能
