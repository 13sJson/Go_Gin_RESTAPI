package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/*  Uberが開発した高速なlogを出力するlogge = zap
    仕様としてkey:valueの関係でセットしログを出力する
	ログの出力の指定をしている
*/

func Logging(c *gin.Context) {
	// ロギングライブラリの作成
	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatal(err.Error())
	}
	// 現在時刻の取得
	oldTime := time.Now()

	// リクエストヘッダーから値を返す
	ua := c.GetHeader("User-Agent")

	/*  ミドルウェアのみ使用可の、保留中のハンドラーを実行
	c.Next()後の記載は処理後、その前に書けば処理前にログ出力
	下記の場合はmain.goでの処理が行われてから処理を出力している
	上記時間取得は処理前の時間（リクエスト受け取ったとき）の時間
	*/
	c.Next()

	// ログ構文
	logger.Info("incoming request",
		zap.String("path", c.Request.URL.Path),
		zap.String("Ua", ua),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
}
