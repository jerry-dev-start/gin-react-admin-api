package req

type InitUploadRequest struct {
	FileName string `form:"fileName" binding:"required"`
	FileMd5  string `form:"md5" binding:"required"`
	FileSize int64  `form:"fileSize"`
}
