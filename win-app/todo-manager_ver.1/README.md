# ToDo Manager ver.1

すること（ToDo）を記録し管理するWindowsのアプリです。入力するとtodo.txt（こうすると管理しやすい）を作成します。WebサイトのToDoでサーバーにデータを保存だと、情報流出もあるので、クライアントソフトにしました。スマホ版のアプリも良いですが、デスクワークで使う用を想定しています。  

go buildでmain.goをコンパイルしたらmain.exeを「todo-manager.exe」(Windowsソフト本体の命名の通例)などとリネームし、[open_windows-start-up-folder](../open_windows-start-up-folder).exeをダブルクリックでWindowsのスタートアップフォルダーを開き、ショートカットを入れておくと、OS起動直後に本アプリも起動するので便利です。

Make Time: JST, 2026.5.13.3:00-3:30, about.  
Maker:  [Shimada print](http://shimadaprint.stars.ne.jp/)  
