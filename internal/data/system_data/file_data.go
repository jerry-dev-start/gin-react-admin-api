package system_data

import (
	"apis/global"
	"apis/internal/model/system_model"
	"errors"

	"gorm.io/gorm"
)

type FileData struct {
}

func NewFileData() *FileData {
	return &FileData{}
}

func (d FileData) GetFileByMd5(md5 string) (*system_model.FileModel, error) {
	var fileModel system_model.FileModel
	if err := global.Db.Where("file_md5 = ?", md5).First(&fileModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &fileModel, nil
}
