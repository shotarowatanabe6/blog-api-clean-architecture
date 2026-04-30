# ブログAPIで学ぶ Clean Architecture

## 宣言

- このリポジトリでは、Clean Architecture について学ぶため、基本的に手作業でコーディングします。
- AIによるコーディング支援は、このREADME作成と、コードレビューのみに限定します。

## コンセプト

架空のブログ記事に対するAPI操作や認証等を実装します。この実装を通じて、保守性の高いコードについて学びます。

## 技術スタック

- 言語: Go 1.26.2
- Webフレームワーク: Gin
- pre-commit v4.6.0
- ドキュメンテーション：Swaggo v1.16.6

### pre-commit

- go-lint: golangci-lintを使用してlinting
- swag: Swaggoを使用してSwaggerドキュメントを作成
  - [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
