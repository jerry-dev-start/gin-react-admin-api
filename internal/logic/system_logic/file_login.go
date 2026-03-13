package systemlogic

import (
	"apis/internal/data/system_data"
	"apis/internal/model/system_model"
)

type FileLogic struct {
	fileData *system_data.FileData
}

func NewFileLogic() *FileLogic {
	return &FileLogic{
		fileData: system_data.NewFileData(),
	}
}

// GetFileByMd5 通过Md5获取到上传的文件信息
func (f *FileLogic) GetFileByMd5(md5 string) (*system_model.FileModel, error) {
	if result, err := f.fileData.GetFileByMd5(md5); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (f *FileLogic) SaveMergeCompleteFileInfo(record *system_model.FileModel) error {
	return f.fileData.SaveMergeCompleteFileInfo(record)
}
