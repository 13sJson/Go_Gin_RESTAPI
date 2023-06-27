package controller

import (
	"book/model"
	"book/service"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*** golangのAPIの基礎 ***/
// gin.Contextはhttpから送られてきた(フロント)データを格納している構造体
// c(gin.Context).JSON/JSONP/Stringでデータ送信元に結果を返す
// nil判定でエラー処理を行い、[正常・異常] 判定を行う
// 大体の流れはmain.go->controller->service(init.go->book.go)->controller

/*** insert ***/
func BookAdd(c *gin.Context) {

	// model.Book構造体の宣言
	book := model.Book{}

	// gin.Contextのデータ(をbook(model.Book)へ挿入
	err := c.Bind(&book)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	// service.BookService構造体の宣言（Serviceのメソッドを使用するため（クラス宣言的な感じ））
	bookService := service.BookService{}
	err = bookService.SetBook(&book)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	// 上記データinsertが完了したら
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

/*** Select ***/
func BookList(c *gin.Context) {
	// service.BookService構造体の宣言
	bookService := service.BookService{}

	// GetBookListで取得してきたデータをBookListに格納
	BookList := bookService.GetBookList()

	// c.JSONPでBookListデータを送る
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    BookList,
	})
}

/*** Update ***/
func BookUpdate(c *gin.Context) {
	// model.Book構造体の宣言
	book := model.Book{}

	// gin.ContextからUpdate用のBookデータをbookに格納
	err := c.Bind(&book)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Insert同様のクラス宣言的な構造体宣言
	bookService := service.BookService{}
	err = bookService.UpdateBook(&book)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

/*** Delete ***/
func BookDelete(c *gin.Context) {
	// Delete対象のidを取得
	id := c.PostForm("id")

	// delete対象のIDをチェック
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	bookService := service.BookService{}
	err = bookService.DeleteBook(int(intId))
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
