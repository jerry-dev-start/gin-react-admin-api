package res

type InitUploadResponse struct {
	UploadId       string `json:"uploadId"`
	UploadedChunks []int  `json:"uploadedChunks"` // 已成功的分片索引列表
	IsExisted      bool   `json:"isExisted"`      // 是否秒传
	FileUrl        string `json:"fileUrl,omitempty"`
}
