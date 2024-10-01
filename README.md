# Go（Go言語）は他の言語と比べると早い・メジャーになりつつある・簡単
<p>Go言語とはGoのことですが、分かりやすくするためにGo言語と呼ばれています。本説明はWindows版Go向けです。  
GoはWindowsなどでもデスクトップアプリや、スマホのアプリや、Webサーバーソフトなども作れます。</p>
<p>プログラミング言語の中ではC++に続き２番手ぐらいに動作などが早いです。C++よりも簡単で新しいです。  
検索・スマホOSの最大手のGoogleで開発されているなどで、派遣案件などの報酬も高単価な傾向です。  
ただしベテラン派遣案件などでは、昔から多くあるJavaというVMで遅くなる言語の補修作業が多いのが実情です。
しかし機械はいずれ壊れ、情報量は激増しているので、早い最新なモダン言語が求められています。</p>

## インストールが他の言語などよりも簡単
他の言語などと比べるとGoの実行ファイルなどのインストールはかなり簡単です。ただしライブラリーなどは多少異なります。
### インストーラーでインストールするだけ
初心者はこの方法が簡単です。また別バージョンを設置＝長期運用する予定がなく、すぐに動かすだけをしたいのでしたら、この方法の方が良いかもしれません。しかし下記のような現状があるので、圧縮書類を設置の方がお勧めですし、難しくもありません。
### 圧縮書類なら解凍して置いて環境変数を指定するだけ
別バージョンを共存するのも、これなら環境変数のパス設定なども混乱しづらいです。  
#### フォルダ構成例  
Z:/usr/bin/go/  
 　|- go-ver-1/　大昔の作成書類を動作させるために残す（go getでインストール時代な物など）  
 　|- go-ver-2/　現在の作成書類を動作させるために残す（環境変数はコレがお勧め）  
 　|- go-ver-3/　最新バージョンのために最近の作成書類が動かない場合もある

<p>バージョンは実存しない分かりやすく説明する例ですので、実際はZIP書類内のフォルダ名「go」をバージョンやarm（386と表記のZIP書類が存在するかは不明）などが付いた「ZIP書類名」にすると仕様を正確に把握しやすいです。Z:とはWindowsのZドライブで、節約的にはパーテーションを区切った論理ドライブで可能です。これがCドライブだとOSが故障したら接続できずに、復元ができない場合もあります。 </p>

<p>C:\Program Files (x86)などにインストールするとリモートメンテナンスがしづらいと思うので、上記のようにUnixのディレクトリ構成と類似したフォルダ構成にした方が、他の言語であるPerlやPHPなどもメンテナンスもしやすいと思います。</p>p
 
<p>なおarmとは最近のPCやモバイル機器などが対応し、x86(386)とは古いPCなどです。ただし古いといってもCPUがCore iなどなので、armが最新製品という感じです。この２つの規格が違っても、AIの回答によるとインストーラーならインストールができなく、私見ですがZIP書類で違っていてもProgram Files (x86)からして動作はするかもしれません。</p>
