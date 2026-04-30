# ブログAPIで学ぶ Clean Architecture

## 宣言

- このリポジトリでは、Clean Architecture について学ぶため、基本的に手作業でコーディングします。
- AIによるコーディング支援は、このREADME作成と、コードレビューのみに限定します。

## コンセプト

架空のブログ記事に対するAPI操作や認証等を実装します。この実装を通じて、保守性の高いコードについて学びます。

## 技術スタック

- 言語: Go 1.24+
- Webフレームワーク: Echo
- ORM: GORM
- DB: MySQL
- 認証: JWT (JSON Web Token)
- DI: 手動 DI（Wire 等は使用しない）
- ドキュメンテーション：Swaggo v1.16.5

### Swaggoについて

以下のコマンドを実行してSwagger用のファイル（/docs）を作成する

```bash
swag init -g cmd/main.go
swag fmt
```

以下のURLからSwaggerへアクセスする

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
