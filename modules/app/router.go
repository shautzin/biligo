package app

import (
	"biligo/modules/app/api"
	"biligo/modules/app/note"

	"github.com/gin-gonic/gin"
)

/// app 模块路由设置文件

// 在这里注册你的 app 所有 api 路由
func Route(r *gin.RouterGroup) {

	r.GET("/test", api.TestIndex)

	r.GET("/note/", note.NoteList)
	r.GET("/label/", note.LabelList)
}
