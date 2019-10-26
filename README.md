# Average

vimの入門用に作ったバグまみれのコードです。
このコードを修正して聞きながら、vimを学んでいきます。

## 使い方

以下のようにしてテストを実行します。
テストは落ちるので、vimの力を使って直してきましょう！

```
$ git clone https://github.com/kurenaif/average
$ cd average/average
$ go test -v
=== RUN   TestAverage
test case 0:
  in: []
  out: "\n"
test case 1:
  in: [1]
--- FAIL: TestAverage (0.00s)
panic: runtime error: index out of range [recovered]
	panic: runtime error: index out of range
```

## 仕様

配列の `i` 番目の要素を `arr[i]`、求めたい値の `i` 番目の要素を `res[i]` とすると、

* `res[i]` = `arr[i-1] + arr[i] + arr[i+1]`

を求めることで、`i`番目の要素の前後の平均値を求めるコードです。

ただし、0番目の要素では `arr[-1]`、最後の要素では `arr[i+1]` の要素が存在しないので、2つの値の平均値を求めてください。

最終的には、Average()関数では、`res`を出力するところまでやります。 最後に改行を出力してください。

もし、`arr`の要素数が0の場合は、改行のみを出力してください。

## Vim Tips

この講座で使用するvimのキーバインド(大文字と小文字は区別します)

基本的に編集の流れは、

縦移動→横移動→ノーマルモードでコマンドを実行→挿入モードに入る→文字を入力する→ノーマルモードに戻る

といった流れになります。

### 縦移動

* `jk`: ↓↑
* `20G`: 20行目に移動する(数字は好きに変えてね！)
* `5j`: 5行下に移動する(数字は好きに変えてね！)

### 横移動

* `hl`: ←→
* `$`: 行末に移動
* `f(`: いまいるカーソルより右にある、一番近い 「`(`」 に移動する。(「`(`」の記号は好きに変えてね！)
* `F(`: いまいるカーソルより左にある、一番近い 「`(`」 に移動する。(「`(`」の記号は好きに変えてね！)

### ノーマルモードで使うコマンド

* `dd`: 一行削除
* `yy`: 一行コピー(ヤンク)
* `p`: 削除やコピーしたものをペーストする
* `ctrl+a`: いまカーソルの上にある数字を1増やす
* `diw`: 今いるカーソルの上にある単語を削除する(delete inner word)

### ノーマルモードから挿入モードへの変更方法

* `i`: 今いるカーソルの左側で、挿入モードに入る
* `a`: 今いるカーソルの右側で、挿入モードに入る
* `o`: 今いるカーソルの次に下を挿入し、かつ挿入モードに入る
* `O`: 今いるカーソルの前に上を挿入し、かつ挿入モードに入る

### 挿入モードからノーマルモードへ移動する

以下の2つの方法がよく使われるが、どちらも微妙に押しにくいのでよくオリジナルのキーバインドが割り当てられている。

* `esc`: 左手だけで押せるのがいい。 スムーズに押すのに練習がいる
* `ctrl+[`: 両手を使うけど、escよりホームポジションに近い

### 選択系コマンド(ノーマルモード→ビジュアルモードへの変更)

* `ctrl+v`: ビジュアルモードに入る(この後、hjklなどで移動して矩形選択が可能)
* `shift+v`: ビジュアルモードに入り、かつカーソルがある行を一行選択する
* ビジュアルモードで `d` : 今選択している文字をすべて削除する
* ビジュアルモードで `y` : 今選択している文字をすべてコピー(ヤンク)する
