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

## 作業ログ

本レポジトリのレポを残す。
方針は以下である。

- バックエンドを最適化するように設定する。つまりログイン認証、DB 設計、API 設計など。フロントエンドは JS の DOM 操作で賄う。

## 1. トーク履歴の永続化

ルーム名をユニークにし、新しくその部屋に入った人はそのトークルームの全ての履歴が見えるようにする。一回のメッセージは１文字以上２００文字以内とし、空白を許可しない。user名は1文字以上１５文字以内とし、空白を許可しない。

# 　現段階でのテーブル設計

Users table
| カラム名 | データ型 | Key | Extra |
| ---- | ---- | ---- | ---- |
| id | int | FK | NOT NULL |
| name | char | ---- | NOT NULL |
| icon | ??? | ---- | ---- |
| created_at | datetime | ---- | NOT NULL |
| deleted_at | datetime | ---- | ---- |

Rooms table
| カラム名 | データ型 | Key | Extra |
| ---- | ---- | ---- | ---- |
| id | ---- | PK | NOT NULL |
| created_at | datetime | ---- | NOT NULL |
| deleted_at | datetime | ---- | ---- |

Messages table
| カラム名 | データ型 | Key | Extra |
| ---- | ---- | ---- | ---- |
| id | int | PK | NOT NULL |
| user_id | int | FK | NOT NULL |
| room_id | int | FK | NOT NULL |
| message | text | ---- | NOT NULL, 1-200文字 |
| created_at | datetime | ---- | NOT NULL |
| deleted_at | datetime | ---- | ---- |

Room と User の Many-to-many は後ほど。

## ローカル環境
MySQLにgo-chatというデータベースを作る。

マイグレーションにはgoosego get bitbucket.org/liamstask/goose/cmd/gooseを使った。