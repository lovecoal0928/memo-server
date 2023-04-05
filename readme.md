# Golangのバージョンが古かったので今更バージョン変更した

```
$ apt info golang
Package: golang
Version: 2:1.13~1ubuntu2
```
初めて触った時のGolangがv1.13だったことが発覚  
ubuntuを使ってサーバー起動とかしてたのでWindowsPowerShellと勝手が違ったので調べながら更新しました

## 現行バージョン確認
```
$ apt info golang
```

## PPAを使って最新バージョンまで読み込み
```
$ sudo add-apt-repository ppa:longsleep/golang-backports
```
23/3/16現在v1.20までずらーっと出てきた

## 最新バージョンのインストール
```
$ sudo apt update
```
ちゃんとv1.20がダウンロードされました

## なんで古いってわかったの？
```
# command-line-arguments
./main.go:17:15: undefined: os.ReadFile
```
これ  
v1.15以前のgolangではos.ReadFileが使えないらしくてそれでエラーが出て気づいた
早めに気づけて良かったね