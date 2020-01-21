# LINE 風チャット Web アプリ

## 概要

Qiita 記事: [LINE 風チャット Web アプリを作ってみる(初心者向け)](https://qiita.com/Ryoma0413/items/84777a0af1191ab2696d)を写経させていただき作ったチャットアプリ。(本当に写経させていただいただけです)

## ローカルでの実行方法

```
$ go run cmd/main.go
```

ブラウザでタブを二つ開き、それぞれで`http://localhost:8080/`を開き、同名の部屋にアクセス。

## 学んだこと

- Golang で Gin + Melody でサーバーサイドで WebSocket 通信を実現する方法([参考](https://github.com/olahol/melody/blob/master/README.md))
- JS でブラウザ側で WebSocket 通信を実現する方法。([参考](https://www.sejuku.net/blog/70583))
- ブラウザでローカルストレージでデータを保存できる

## これから改善できるところ

以下を直すあたりからこのアプリをさらに発展させて行けそうである。(いつかやるかも)

- 実行者と同一名のユーザが存在した場合、その人物のチャットは右側(自分が発言したもの)として表示されてしまう
- ユーザ名はブラウザのローカルストレージに保存しているため、別のタブでトークを開くとその度にローカルストレージが書き換わってしまう(あまり問題ないが)
- トークルームに「JoinRoom」ページを介さずにアクセスできてしまう。この場合、ユーザ名は"null"になる。(ローカルストレージに user 名が保存されていればその名前になる)
- 発言者のプロフィール画像が現段階では登録できず NotFound になっている
- トークのデータ履歴がサーバーサイドで保持されていない
