# hello

https://golang.org/doc/tutorial/getting-started

## モジュールの初期設定

例えば `example` というグループに属する `hello` というモジュールを作成する場合は以下のようにする。

```.sh
go mod init example/hello
```

基本的には、 `[User名/Organization名]/[repository名]` という感じで付ければ良さそう。

ここはNode.jsだと `package.json` を、JAVAだと `pom.xml` や `build.gradle` を生成するものと捉えれば良さそう。

## モジュールの作成

main関数が実行される。

また、 `import` でビルドインモジュールなどの外部モジュールを読み込むこともできる。

## モジュールの実行

モジュールの実行は以下のコマンドで行うことができる。

```.sh
go run .
```

## ビルドインモジュール以外の公開モジュールを読み込む

以下のサイトから公開モジュールの検索ができる。  
https://pkg.go.dev/

読み込みたいモジュールが決まったら、それをGoファイル内に記載する。  
その後、以下のコマンドを叩くことで、依存関係を記載するファイルが自動生成される。

```.sh
go mod tidy
```
