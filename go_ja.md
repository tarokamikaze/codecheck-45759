# コマンドラインアプリケーション(CLI アプリ)作成用テンプレート(Go)

[main.go](main.go)を編集して、CLIアプリを実装してください。  
チャレンジ内でファイルの作成が許可されていれば、可読性等のためにファイルを分割する事も可能です。

## コマンドライン引数の取得方法
[main.go](main.go)ファイルに定義されている、`run`という関数内で `args` の名前で取得可能です。

``` go
func run(args []string) {
  // code to run
}
```

ここでの `args` は、同ファイルの `main` 関数内で渡された `os.Args` の一部で、
「スクリプト名を除いたコマンドライン引数」が渡されています。

## コマンド実行結果の標準出力への出力
標準出力への出力は `fmt.Println` 等のメソッドで可能です。 Cs

``` go
fmt.Println(args)
```

## 外部ライブラリの追加方法
外部ライブラリを使用する場合は以下の手順で実施してください。

- [codecheck.yml](codecheck.yml)に以下の内容を `go build` の前に追加  
(複数のライブラリのインストールも行を追加していく事で可能です)

``` yaml
build:
  - go get namespace.of/some/library
```
