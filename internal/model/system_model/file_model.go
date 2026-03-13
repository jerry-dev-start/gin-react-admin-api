package system_model

type FileModel struct {
	Id       uint64 `json:"id" gorm:"primaryKey"`
	FileMd5  string `json:"fileMd5" gorm:"column:file_md5"`
	FileName string `json:"fileName" gorm:"column:file_name"`
	FilePath string `json:"filePath" gorm:"column:file_path"`
	FileSize int64  `json:"fileSize" gorm:"column:file_size"`
	FileExt  string `json:"fileExt" gorm:"column:file_ext"`
	Status   int8   `json:"status" gorm:"column:status"`
	// 返回给前端时直接就是数字格式
	CreateTime int64 `json:"createTime" gorm:"column:create_time"`
	UpdateTime int64 `json:"updateTime" gorm:"column:update_time"`
}

// TableName 指定 GORM 使用的表名
func (FileModel) TableName() string {
	return "files"
}
