# Terraform for さくらのIoT Platform

[![Build Status](https://travis-ci.org/yamamoto-febc/terraform-provider-sakuraiot.svg?branch=master)](https://travis-ci.org/yamamoto-febc/terraform-provider-sakuraiot)

[Terraform](https://www.terraform.io)で[さくらのIoT Platform](https://iot.sakura.ad.jp)を操作するためのプラグイン

## クイックスタート

#### 準備

  - Dockerをインストールしておく
  - さくらのIoT PlatformのAPIキーを取得しておく

Dockerがない場合は[Installation / インストール](docs/installation.md)を参考に
TerraformとTerraform for さくらのIoTを手元のマシンにインストールしてからご利用ください。

さくらのIoT Platform APIキーの取得方法は[こちら](docs/installation.md#さくらのiot-platform-apiキーの取得)を参照してください。

```bash
#################################################
# APIキーを環境変数に設定しておく
#################################################
$ export SAKURAIOT_AUTH_TOKEN="APIキーのトークン"
$ export SAKURAIOT_AUTH_SECRET="APIキーのシークレット"

#################################################
# Terraform定義ファイル作成
#################################################
$ mkdir ~/work; cd ~/work
$ tee sakura.tf <<-'EOF'

# ------------------------------------------------------------
# プロジェクトの登録
# ------------------------------------------------------------
resource "sakuraiot_project" "project01" {
    name = "example project"
}

# ------------------------------------------------------------
# モジュールの登録
# ------------------------------------------------------------
resource "sakuraiot_module" "module01" {
    project_id = "${sakuraiot_project.project01.id}"
    name = "example module"
    register_id = "put-your-register-id"      # (要編集)通信モジュール本体に記載されているモジュール登録用ID
    register_pass = "pur-your-register-pass"  # (要編集)通信モジュール本体に記載されているモジュール登録用パスワード
}

# ------------------------------------------------------------
# 連携サービスとしてOutgoing WebhookとIncoming Webhookを登録
# ------------------------------------------------------------
# Outgoing Webhook
resource "sakuraiot_service_outgoing_webhook" "webhook01" {
    project_id = "${sakuraiot_project.project01.id}"
    name = "example outgoing webhook"
    secret = "secret"                         # HMAC-SHA1署名用のシークレット
    url = "https://your-webhook-server.com/"  # Outgoing Webhookの連携先URLを指定
}

# Incoming Webhook
resource "sakuraiot_service_incoming_webhook" "webhook01" {
    project_id = "${sakuraiot_project.project01.id}"
    name = "example incoming webhook"
    secret = "secret"                          # HMAC-SHA1署名用のシークレット
}
EOF

#TerraformコマンドをDocker上で実行するためにエイリアスを作成(Dockerを使わない場合は不要)
alias terraform="docker run -it --rm -w /work -v $PWD:/work -e SAKURAIOT_AUTH_TOKEN -e SAKURAIOT_AUTH_SECRET yamamotofebc/terraform-for-sakuraiot"

#################################################
# 構築プランの確認
#################################################
$ terraform plan

#################################################
# インフラ構築
#################################################
$ terraform apply

#################################################
# 作成された内容を確認
#################################################
$ terraform show

#################################################
# 削除
#################################################
$ terraform destroy 
# 確認入力を要求されるため、`yes`を入力

```

## インストール

[リリースページ](https://github.com/yamamoto-febc/terraform-provider-sakuraiot/releases/latest)から最新のバイナリを取得し、
Terraformバイナリと同じディレクトリに展開してください。

詳細は[Installation / インストール](docs/installation.md)を参照してください。

## 使い方/各リソースの設定方法

Terraform定義ファイル(tfファイル)を作成してご利用ください。

設定ファイルの記載方法は[設定リファレンス](docs/README.md)を参照ください。

さくらのIoT Platformの以下のリソースをサポートしています。

### サポートしているリソース

  - [プロジェクト](docs/configuration/resources/project.md)
  - [モジュール](docs/configuration/resources/module.md)
  - [サービス:WebSocket](docs/configuration/resources/service_websocket.md)
  - [サービス:Incoming-Webhook](docs/configuration/resources/service_incoming_webhook.md)
  - [サービス:Outgoing-Webhook](docs/configuration/resources/serivce_outgoing_webhook.md)

## License

  This project is published under [Apache 2.0 License](LICENSE).

## Author

  * Kazumichi Yamamoto ([@yamamoto-febc](https://github.com/yamamoto-febc))
