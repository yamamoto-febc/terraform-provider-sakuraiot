# モジュール(sakuraiot_module)

すでにさくらのIoT Platform上に存在しているモジュールを参照するためのデータソース

### 設定例

```
data "sakuraiot_module" "foobar" {
    project_id = "xxxxxx"
    serial_number = "nnnnnnnn10"
    model = "SCM-LTE-BETA"
    sort = "-id"
}
```

注意：複数ヒットする場合は先頭のリソースを選択します。  
複数ヒットするような条件指定を行う場合は`sort`の指定をお薦めします。

### パラメーター

|パラメーター         |必須  |名称                |初期値     |設定値                    |補足                                          |
|-------------------|:---:|--------------------|:--------:|------------------------|----------------------------------------------|
| `project_id`      | -   | プロジェクトID       | -        | 文字列                  | - |
| `serial_number`   | -   | シリアルナンバー      | -        | 文字列                  | 通信モジュールに記載されているシリアル番号 |
| `model`           | -   | モデル              | -        | 文字列                  | 例：`SCM-LTE-BETA` |
| `sort`            | -   | 並び順              | -        | `id`<br />`-id`<br />`name`<br />`-name`<br />`project`<br />`-project`<br />`serial_number`<br />`-serial_number`                  | 先頭に`-`がない値は昇順<br />先頭に`-`がある場合は降順|

### 属性

|属性名                | 名称                    | 補足                                        |
|---------------------|------------------------|--------------------------------------------|
| `id`                | ID                     | -                                          |
| `project_id`        | プロジェクトID           | -                                          |
| `name`              | モジュール名             | -                                          |
| `is_online`         | オンライン           | -                                          |
