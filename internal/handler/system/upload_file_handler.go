package system

import (
	"apis/global"
	"apis/internal/handler/system/req"
	res2 "apis/internal/handler/system/res"
	systemlogic "apis/internal/logic/system_logic"
	"apis/model/common/response"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UploadFileHandler struct {
	fileLogic *systemlogic.FileLogic
}

func NewUploadFileHandler() *UploadFileHandler {
	return &UploadFileHandler{
		fileLogic: systemlogic.NewFileLogic(),
	}
}

// UploadSimple 上传小文件
func (u *UploadFileHandler) UploadSimple(c *gin.Context) {
	// 1. 从 Form 表单中获取文件，这里的 "file" 必须和前端 append 的 key 一致
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("文件读取失败", c)
		return
	}

	// 2. 准备保存路径 (按日期分类存储，防止单目录下文件过多)
	uploadDir := "/uploads/" + time.Now().Format("2006/01/02")
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		response.FailWithMessage("创建文件夹失败", c)
		return
	}

	// 3. 构造唯一文件名，防止重名覆盖
	// 获取文件后缀，如 .jpg
	ext := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + ext
	dst := filepath.Join(*global.Config.LocalFileServer.FilePath, uploadDir, newFileName)

	// 4. 保存文件到本地服务器磁盘
	// c.SaveUploadedFile 是 Gin 封装好的流式写入方法
	if err := c.SaveUploadedFile(file, dst); err != nil {
		response.FailWithMessage("保存文件失败", c)
		return
	}
	res := res2.UploadSimpleRes{
		FileUrl:  filepath.ToSlash(filepath.Join(uploadDir, newFileName)),
		FileName: file.Filename,
	}
	response.OkWithData(res, c)
}

// UploadChunkInit 初始化分片上传任务
//
// 该方法在切片上传前执行，主要包含以下逻辑：
// 1. 秒传校验：通过文件整体 MD5 值判断服务端是否已存在该文件。若存在，则直接返回上传成功（秒传）。
// 2. 环境准备：若文件未上传过，则根据文件唯一标识（如 MD5 或 UploadID）创建对应的临时目录，用于存放后续上传的切片文件。
// 3. 断点续传逻辑（可选）：返回已上传的分片索引，告知前端从哪个分片开始继续上传。
func (u *UploadFileHandler) UploadChunkInit(c *gin.Context) {
	var req req.InitUploadRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage("参数缺失", c)
		return
	}
	// 1.秒传检查
	fileModel, err := u.fileLogic.GetFileByMd5(req.FileMd5)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if fileModel != nil {
		resp := res2.InitUploadResponse{
			FileUrl:   fileModel.FilePath,
			IsExisted: true,
		}
		response.OkWithData(resp, c)
	}
	//如果 fileModel 为空则是文件没有上传完整
	uploadId := req.FileMd5
	tmpDir := filepath.Join(*global.Config.LocalFileServer.FilePath, "/uploads/tmp", uploadId)

	// 3. 检查断点：扫描已存在的高质量分片
	uploadedChunks := make([]int, 0)
	if _, err := os.Stat(tmpDir); err == nil {
		files, _ := os.ReadDir(tmpDir)
		for _, f := range files {
			// 关键：排除掉正在上传的 .part 临时文件，只统计完整重命名后的文件
			if !f.IsDir() && !strings.HasSuffix(f.Name(), ".part") {
				idx, err := strconv.Atoi(f.Name())
				if err == nil {
					uploadedChunks = append(uploadedChunks, idx)
				}
			}
		}
	} else {
		// 如果目录不存在，则创建
		os.MkdirAll(tmpDir, os.ModePerm)
	}
	res := res2.InitUploadResponse{
		UploadId:       uploadId,
		UploadedChunks: uploadedChunks,
		IsExisted:      false,
	}
	response.OkWithData(res, c)
}
