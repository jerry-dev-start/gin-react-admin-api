package system

import (
	"apis/global"
	"apis/internal/handler/system/req"
	res2 "apis/internal/handler/system/res"
	systemlogic "apis/internal/logic/system_logic"
	"apis/internal/model/system_model"
	"apis/model/common/response"
	"fmt"
	"io"
	"mime/multipart"
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
	if err := c.ShouldBindJSON(&req); err != nil {
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
		return
	}
	//如果 fileModel 为空则是文件没有上传完整
	uploadId := req.FileMd5
	tmpDir := filepath.Join(*global.Config.LocalFileServer.FileChunk, uploadId)

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

// UploadFileChunk 上传文件分片
func (u *UploadFileHandler) UploadFileChunk(c *gin.Context) {
	// 1. 解析非文件参数
	var req req.ChunkUploadRequest
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage("参数缺失", c)
		return
	}

	// 2. 获取分片文件流
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("获取分片数据失败", c)
		return
	}

	taskDir := filepath.Join(*global.Config.LocalFileServer.FileChunk, req.UploadId)
	// 确保任务目录存在
	if err := os.MkdirAll(taskDir, os.ModePerm); err != nil {
		response.FailWithMessage("创建任务目录失败", c)
		return
	}
	finalChunkPath := filepath.Join(taskDir, req.Index)
	tempChunkPath := finalChunkPath + ".part" // 临时后缀
	// 4. 执行原子写入逻辑
	if err := saveChunkAtomic(file, tempChunkPath, finalChunkPath); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("分片保存成功", c)
}

// saveChunkAtomic 原子化保存分片
func saveChunkAtomic(fileHeader *multipart.FileHeader, tempPath, finalPath string) error {
	// 打开上传的文件流
	src, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("打开分片流失败")
	}
	defer src.Close()

	// 创建临时文件
	out, err := os.Create(tempPath)
	if err != nil {
		return fmt.Errorf("创建临时文件失败")
	}

	// 写入标记
	isSuccess := false
	defer func() {
		out.Close()
		if !isSuccess {
			os.Remove(tempPath) // 如果中途失败（如断网），清理残余的 .part 文件
		}
	}()

	// 流式拷贝，不占用大内存
	if _, err = io.Copy(out, src); err != nil {
		return fmt.Errorf("分片写入中断")
	}

	// 确保数据已同步到磁盘
	out.Sync()
	out.Close() // 重命名前必须先关闭文件句柄（尤其在Windows下）

	// 重命名为正式文件名，标志该分片已完整到达
	if err = os.Rename(tempPath, finalPath); err != nil {
		return fmt.Errorf("分片重命名失败")
	}
	isSuccess = true
	return nil
}

// UploadChunkComplete 上传切片完成
func (u *UploadFileHandler) UploadChunkComplete(c *gin.Context) {
	var req req.CompleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数解析失败", c)
		return
	}
	// 1. 定义路径
	taskDir := filepath.Join(*global.Config.LocalFileServer.FileChunk, req.UploadId)
	// 正式存储路径：按日期分类
	subDir := time.Now().Format("2006/01/02")
	finalDir := filepath.Join(*global.Config.LocalFileServer.FilePath, "/uploads/", subDir)
	os.MkdirAll(finalDir, os.ModePerm)

	// 为了防止文件名冲突，建议使用 UploadId 作为文件名
	saveFileName := req.UploadId + filepath.Ext(req.FileName)
	finalPath := filepath.Join(finalDir, saveFileName)

	// 2. 校验分片数量（简单校验）
	files, _ := os.ReadDir(taskDir)

	// 注意：这里只统计完整分片，排除 .part 文件
	count := 0
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".part") {
			count++
		}
	}

	if count != req.TotalChunks {
		response.FailWithMessage("分片数量不一致，请检查是否全部上传成功", c)
		return
	}

	// 3. 开始合并文件
	destFile, err := os.Create(finalPath)
	if err != nil {
		response.FailWithMessage("创建目标文件失败", c)
		return
	}

	// 合并逻辑封装
	for i := 0; i < req.TotalChunks; i++ {
		chunkPath := filepath.Join(taskDir, strconv.Itoa(i))

		// 打开分片
		sourceFile, err := os.Open(chunkPath)
		if err != nil {
			destFile.Close()
			response.FailWithMessage(fmt.Sprintf("读取分片%d失败", i), c)
			return
		}

		// 将分片拷贝到目标文件
		_, err = io.Copy(destFile, sourceFile)
		sourceFile.Close() // 拷完立即关闭，释放句柄

		if err != nil {
			destFile.Close()
			response.FailWithMessage("合并写入失败", c)
			return
		}
	}
	destFile.Close()

	// 4. 获取文件最终大小
	fileInfo, _ := os.Stat(finalPath)
	// 5. 写入数据库记录 (使用之前定义的 FileModel 和时间戳)
	now := time.Now().UnixNano() / int64(time.Millisecond)
	// 转换 URL 路径为正斜杠返回给前端
	fileUrl := "/" + filepath.ToSlash(filepath.Join("uploads/", subDir, saveFileName))

	record := &system_model.FileModel{
		FileMd5:    req.UploadId,
		FileName:   req.FileName,
		FilePath:   fileUrl,
		FileSize:   fileInfo.Size(),
		FileExt:    filepath.Ext(req.FileName),
		Status:     1,
		CreateTime: now,
		UpdateTime: now,
	}
	if err := u.fileLogic.SaveMergeCompleteFileInfo(record); err != nil {
		response.FailWithMessage("文件信息入库失败", c)
		return
	}
	// 6. 清理临时分片目录 (合并成功后再删)
	os.RemoveAll(taskDir)
	response.OkWithData(record, c)
}
