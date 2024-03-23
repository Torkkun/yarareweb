# OC_Summer_2022

2022年オープンキャンパス時に作成したやられサイト

## プロキシ
gitconfig
windows SETが上手く働かないのでenvでhttp_proxyとhttps_proxyを設定
docker engine fordesctopであればポータルのsettingからproxy
コンテナ --env HTTP_PROXYで設定compose.ymlに書いたので恐らく設定せずでも行けるはず

## Buildコマンド
docker-compose build
docker-compose up

## 魔法の言葉
docker-compose down --rmi all --volumes --remove-orphans

## やること
ユーザーを見つけてもらうサイドチャネル攻撃をする
その結果見つかったユーザー名を憶えておいてもらう
guestユーザーのデータが見れるだけなのに他のも見れるようにする

## 時間あれば
今ある物は、バックエンドのみなので出来れば、クリックジャッキング的なものも欲しい押したらエビルポータルにサイトに繋がるような感じ←これは結構時間かかりそうなのでいけそうだったら（出来たらおもしろいかも）

## DBユーザーテーブルパスワード
|user|password|
|----|----|
|Bill|Microsoft|
|Satoshi|Bitcoin|
|John|personalcomputer|
|Steve|iPhone|
|Alan|enigma|
|Linus|linux|
|Yukihiro|ruby|
|Lisa|ryzen|
|Kevin|shimomura|
|Peter|spiderman|
|----|----|
|guest|guest|
-----------------

## ログインフォームサイドチャネル
参考リンク
https://blog.flatt.tech/entry/login_logic_security#%E8%A6%B3%E7%82%B91-%E3%83%AC%E3%82%B9%E3%83%9D%E3%83%B3%E3%82%B9%E3%81%AE%E5%B7%AE%E5%88%86%E3%81%8B%E3%82%89%E3%83%A1%E3%83%BC%E3%83%AB%E3%82%A2%E3%83%89%E3%83%AC%E3%82%B9%E3%81%AE%E5%AD%98%E5%9C%A8%E3%81%8C%E7%A2%BA%E8%AA%8D%E5%8F%AF%E8%83%BD

ログインの差分からユーザーの存在確認ができる
1. メールアドレスであればアドレス所有者がサービスを使っているということが分かってしまう
2. パスワードリスト攻撃の手がかりとして利用されてしまう

* エラーメッセージに差異がある
* レスポンスヘッダに差異がある
* 処理時間に差異がある

## SQL Injection

通常はguestというユーザーの情報しか出ないがテーブル内の全データを表示させる。
 1' or '1' = '1';-- (←空白いるはずだけどなくても通る)を入力することで全部出る
 goの所為で;が文字ではなくURLセパレーターとして解釈される（使われてはないが）自分でURLエンコード処理をする必要がある。
 1' or '1' = '1'%3B-- 
 とすれば全部出る。
 ややこしいので、分かりやすくエラーをconsole.log()で出すようにわざとする？

 なぜこうなるのか
 製品名が'1'か'1'が'1'なら真という条件なので絶対に真となる。
 --はコメントアウトするものなので前の文までが実行される

 ## 簡単なXSSデモ
掲示板のようなシステムを模した入力フォームにjavascriptが実行可能であるので任意コード実行できる
