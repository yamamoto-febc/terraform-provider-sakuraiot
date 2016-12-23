# 連携サービス:Incoming Webhook(sakuraiot_service_incoming_webhook)

### 設定例

```
resource "sakuraiot_project" "project01" {
    name = "project01"
}

resource "sakuraiot_service_incoming_webhook" "webhook01" {
    project_id = "${sakuraiot_project.project01.id}"
    name = "webhook01"
    secret = "secret"
}
```

### パラメーター

|パラメーター         |必須  |名称                |初期値     |設定値                    |補足                                          |
|-------------------|:---:|--------------------|:--------:|------------------------|----------------------------------------------|
| `project_id`      | ◯   | プロジェクトID           | -        | 文字列                  | - |
| `name`            | ◯   | Webhook名称             | -        | 文字列                  | - |
| `secret`          | -   | HMAC-SHA1署名用シークレット | `空文字`        | 文字列                  | - |

### 属性

|属性名                | 名称                    | 補足                                        |
|---------------------|------------------------|--------------------------------------------|
| `id`                | ID                         | -                                          |
| `project_id`        | プロジェクトID              | -                                          |
| `name`              | Webhook名称                | -                                          |
| `secret`            | HMAC-SHA1署名用シークレット                | -                                          |
| `url`               | WebhookのエンドポイントURL | -                                          |
