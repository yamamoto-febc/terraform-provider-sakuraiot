# 連携サービス:WebSocket(sakuraiot_service_websocket)

### 設定例

```
resource "sakuraiot_project" "project01" {
    name = "project01"
}

resource "sakuraiot_service_websocket" "websocket01" {
    project_id = "${sakuraiot_project.project01.id}"
    name = "websocket01"
}
```

### パラメーター

|パラメーター         |必須  |名称                |初期値     |設定値                    |補足                                          |
|-------------------|:---:|--------------------|:--------:|------------------------|----------------------------------------------|
| `project_id`      | ◯   | プロジェクトID           | -        | 文字列                  | - |
| `name`            | ◯   | WebSocket名称             | -        | 文字列                  | - |

### 属性

|属性名                | 名称                    | 補足                                        |
|---------------------|------------------------|--------------------------------------------|
| `id`                | ID                         | -                                          |
| `project_id`        | プロジェクトID              | -                                          |
| `name`              | WebSocket名称                | -                                          |
| `url`               | WebSocketのエンドポイントURL | -                                          |
