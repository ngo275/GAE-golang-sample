# GAE sample

## 環境

- go1.8.3
- Google App Engine(GAE)

## 実行方法

ローカルで実行するにはGAEのローカル開発サーバーを利用する。gcloudコマンドのインストールが必要なのでインストールしてください。

```bash
dev_appserver.py --port=9999 app/app.yaml
```

[http://localhost:9999](http://localhost:9999)でアクセス可能になる。

## Deployment

```bash
gcloud app deploy app/app.yaml
```
