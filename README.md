
レイヤードアーキテクチャ ( + DDD )の勉強のために基本的なAPIを実装した。   
APIの仕様は、
* POST  todo/create   
   リクエストで指定されたタイトルをデータベースに保存する
* GET   todo/get   
   クエリがある場合は文字列に部分一致しているtitleを持つtodoのみ取得する  
   クエリがない場合は作成したTODOの一覧を取得する
* POST  todo/update/:todo_id   
   指定した TODO を更新する
* DELETE  todo/delete/:todo_id   
   指定した TODO を削除する

### 起動手順
1. MySQLサーバ、phpMyAdminの起動  
  `docker compose up`  
   phpMyAdminアクセス: <http://localhost:4001/>

2. 環境変数の設定
```
  export MYSQL_USER=root          
  export MYSQL_PASSWORD=go-lesson1
  export MYSQL_HOST=127.0.0.1
  export MYSQL_PORT=3301
  export MYSQL_DATABASE=go_lesson1_api
```
4. `go run ./cmd/main.go`を実行

### 技術の選定理由
* バックエンド：GO
* フレームワーク: Gin  
  調べた感じ最も人気が高そう、パフォーマンスが良いらしい、使ったことがないので使ってみたい  
  
* DB : MySQL  

* アーキテクチャ: レイヤードアーキテクチャ( +  DDD ? )  
  依存関係の整理、テストの実装を行いたい、勉強のためなどの動機から、何かしらのアーキテクチャは採用したい  
  → 今回は単純な機能の実装のためシンプルそうなレイヤードアーキテクチャにしてみる（ DDDについてまだよく分かってないため理解は後回し )
  
### 参考
* Ginの参考：    
  https://qiita.com/Syoitu/items/8e7e3215fb7ac9dabc3a    
  https://qiita.com/hyo_07/items/59c093dda143325b1859    
* アーキテクチャ :    
   https://tech.yyh-gl.dev/blog/go_web_api/   
   https://qiita.com/tono-maron/items/345c433b86f74d314c8d
* go mod init  :  https://qiita.com/TakanoriVega/items/6d7210147c289b45298a  
* ステータスコード : https://developer.mozilla.org/ja/docs/Web/HTTP/Status  
* RowsAffected()： https://golang.shop/post/go-databasesql-05-modifying-ja/  
* docker-composeの参考:  
  https://qiita.com/A-Kira/items/f401aea261693c395966　 　  
  https://zenn.dev/peishim/articles/f7a76ae6c253e4     
  MySQLでUTCの日付時刻が返る:https://zenn.dev/uepyon9/articles/84b5ac68c71d04  
