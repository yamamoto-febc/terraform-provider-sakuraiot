# プロジェクト(sakuraiot_project)

### 設定例

```
resource "sakuraiot_project" "project01" {
    name = "your_project_name"
}
```

### パラメーター

|パラメーター         |必須  |名称                |初期値     |設定値                    |補足                                          |
|-------------------|:---:|--------------------|:--------:|------------------------|----------------------------------------------|
| `name`            | ◯   | プロジェクト名           | -        | 文字列                  | - |

### 属性

|属性名                | 名称                    | 補足                                        |
|---------------------|------------------------|--------------------------------------------|
| `id`                | ID               | -                                          |
| `name`              | プロジェクト名               | -                                          |