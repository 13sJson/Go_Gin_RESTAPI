package main

import (
	"book/controller"
	"book/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// engineのinit処理
	engine := gin.Default()

	/* ミドルウェアでログ管理 >> middleware/bookMiddleware.go
	engine.use(middleware.***)で全てのルーティングに適用している。
	もし個別にログを指定したい場合は、処理後にログ出力メソッドを指定したあげればいい。
	*/
	engine.Use(middleware.Logging)

	// CRUD (ルーターグループ化 = /book/v1)
	/* ルーターグループ化により、ルーティングをグループ化する機能 */
	bookEngine := engine.Group("/directory")
	{
		v1 := bookEngine.Group("/v1")
		{
			// パッケージ:controllerへ処理httpアドレスによって処理以降
			/*
				v1.POST("/add", controller.BookAdd)
				v1.GET("/list", controller.BookList)
				v1.PUT("/update", controller.BookUpdate)
				v1.DELETE("/delete", controller.BookDelete)
			*/

			/*** REST API (Path遷移なし) ***/
			v1.POST("/book", controller.BookAdd)
			v1.GET("/book", controller.BookList)
			v1.PUT("/book", controller.BookUpdate)
			v1.DELETE("/book", controller.BookDelete)
		}
	}

	engine.Run(":8080")

}
