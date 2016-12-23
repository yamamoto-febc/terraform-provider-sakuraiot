# プロバイダ設定

### 設定例
```
provider "sakuraiot" {
    token = "your API token"
    secret = "your API secret"
}
```

### パラメーター

以下のパラメーターをサポートしています。

|パラメーター|必須  |名称                |初期値     |設定値 |説明                                          |
|----------|:---:|--------------------|:--------:|------|----------------------------------------------|
|`token`   | ◯  |APIキー<br />(トークン)     | -        |文字列|環境変数`SAKURAIOT_AUTH_TOKEN`での指定も可         |
|`secret`  | ◯  |APIキー<br />(シークレット)  | -        |文字列|環境変数`SAKURAIOT_AUTH_SECRET`での指定も可  |

各パラメータとも環境変数での指定が可能です。

`token`と`secret`を環境変数で指定した場合、プロバイダ設定の記述は不要です。

#### 環境変数の指定例

```bash
$ export SAKURAIOT_AUTH_TOKEN="APIキーのトークン"
$ export SAKURAIOT_AUTH_SECRET="APIキーのシークレット"
```