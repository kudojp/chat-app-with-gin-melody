package chat

import (
    "github.com/gin-gonic/gin"
    melody "gopkg.in/olahol/melody.v1"
    "net/http"
)

func Run() {
    r := gin.Default()
    m := melody.New()

    r.Static("/static", "./view/static")
    r.LoadHTMLGlob("view/*.html")

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{})
    })

	r.GET("/room/:name", func(c *gin.Context) {
        c.HTML(http.StatusOK, "room.html", gin.H{"Name": c.Param("name")})
    })

    // websocket通信はGETでとれ、データはContextに入るらしい
    r.GET("/room/:name/ws", func(c *gin.Context){
        // Upgrade HTTP request to WebSocket connection
        // and dispatch
        // m.HandleMessageを呼んでいる、、、？
        m.HandleRequest(c.Writer, c.Request)
    })

	// クライアントから文字列を受信した時に行う処理
	m.HandleMessage(func(s *melody.Session, msg []byte){
		// 送られてきたmsgをサーバとつながっているクライアントに送信
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			//　あるセッションqが現在のクライアントとセッションsと同じ部屋かのboolを返す
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	r.Run(":8080")
}