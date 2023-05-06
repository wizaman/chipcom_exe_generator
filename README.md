# chipcom_exe_generator

[wizaman/exe_db](https://github.com/wizaman/exe_db) のJSONデータを加工して、 [massy22/exe](https://github.com/massy22/exe) 向けのCSVテーブルデータを生成するツールです。

Go言語で実装しています。

## 使い方

下記のように各種リポジトリを同階層に配置します。

- chipcom_exe_generator: 本リポジトリ（相対パスで処理するので名前はなんでもいい）
- chipcom_exe: `massy22/exe` リポジトリまたはそのfork
- exe_db: `wizaman/exe_db` リポジトリまたはそのfork

下記コマンドを実行します。

```shell
# Goで都度コンパイルして実行
go run . all

# バイナリをビルドして実行
./build.sh
./chipcom all
```
