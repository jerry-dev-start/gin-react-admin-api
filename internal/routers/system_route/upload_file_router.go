package system_route

import (
	"apis/internal/handler/system"

	"github.com/gin-gonic/gin"
)

type UploadFileRouter struct {
}

func (u *UploadFileRouter) Register(public, private *gin.RouterGroup) {
	fileGroup := public.Group("/file")
	fileHandler := system.NewUploadFileHandler()
	{
		fileGroup.POST("/uploadSimple", fileHandler.UploadSimple)
		fileGroup.POST("/chunkInit", fileHandler.UploadChunkInit)
		fileGroup.POST("/uploadChunk", fileHandler.UploadFileChunk)
		fileGroup.POST("/chunkComplete", fileHandler.UploadChunkComplete)
	}
}
