# Go_Gin_RESTAPI
 Go×GinでRESTAPIを実装してみた(バックエンド側のみ)

 【API流れ】
 
➀main.go起動時にmiddleware.info()でlogger準備(ログ指定)
  service.init()でDB(MySQL)接続をする->その後httpリクエスト待機

➁httpリクエストlocalhost:8080/directory/v1/book ->リクエストどれか[POST/GET/PUT/DELETE]が来ると
  controllerのメソッドを起動してリクエストと送られてきたデータをgin.Context構造体で受取り、
  そのデータをmodelに記載した構造体に格納する

③controllerでserviceメソッドを起動（serviceメソッドは主にDBとのやり取り）
  戻り値によりエラー判定し、その結果をレスポンスとして返す

④またレスポンスが来るまで待機


【実装テスト(POSTMAN)】

<main.go実行>
![screenshot 22](https://github.com/13sJson/Go_Gin_RESTAPI/assets/115130634/68ac72c6-001b-4dd1-bf7c-cc04cc4753f6)

<POSTMANリクエスト送信>
![screenshot 24](https://github.com/13sJson/Go_Gin_RESTAPI/assets/115130634/b929680c-0b95-4eea-ab0c-14797faf8258)

<リクエスト受付log>
![screenshot 23](https://github.com/13sJson/Go_Gin_RESTAPI/assets/115130634/0f11f68e-3a4d-4196-8527-e71f49aa8475)

<DB確認>

![screenshot 26](https://github.com/13sJson/Go_Gin_RESTAPI/assets/115130634/738b1e47-86eb-4c31-913a-b1bc902b0770)
