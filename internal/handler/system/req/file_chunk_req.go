package req

type InitUploadRequest struct {
	FileName string `json:"fileName" form:"fileName" binding:"required"`
	FileMd5  string `json:"md5" form:"md5" binding:"required"`
	FileSize int64  `json:"fileSize" form:"fileSize"`
}
