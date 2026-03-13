package req

type InitUploadRequest struct {
	FileName string `json:"fileName" form:"fileName" binding:"required"`
	FileMd5  string `json:"md5" form:"md5" binding:"required"`
	FileSize int64  `json:"fileSize" form:"fileSize"`
}

// ChunkUploadRequest 定义接收分片的参数
type ChunkUploadRequest struct {
	UploadId string `form:"uploadId" binding:"required"`
	Index    string `form:"index" binding:"required"`
}

type CompleteRequest struct {
	UploadId    string `json:"uploadId" binding:"required"`
	FileName    string `json:"fileName" binding:"required"`
	TotalChunks int    `json:"totalChunks" binding:"required"`
}
