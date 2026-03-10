package gin

import (
	"apis/configs"
	"apis/global"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine     *gin.Engine
	httpServer *http.Server
	Conf       *configs.Config
	Close      func()
}

// InitGin 初始化 Gin 引擎
// 设置运行模式、全局中间件（Recovery、请求日志），并将引擎存入 global.Engine
func NewServer(conf *configs.Config) *Server {
	if conf.Server == nil {
		panic("Server configuration not found.")
	}
	gin.SetMode(conf.Server.GetModel())

	ginEngine := gin.Default()

	return &Server{
		Engine: ginEngine,
		Conf:   conf,
	}
}

// Run 根据 global.Config.Server 的 host、port 启动 HTTP 服务（阻塞）
// StartWeb 启动Web服务
func (s *Server) StartWeb() {
	// 拼接地址
	address := fmt.Sprintf("%s:%d", s.Conf.Server.GetHost(), s.Conf.Server.GetPort())

	s.httpServer = &http.Server{
		Addr:    address,
		Handler: s.Engine,
	}

	go func() {
		global.Log.Info("Start web server at ", address)
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// ---- 优雅启停逻辑 ----
	quit := make(chan os.Signal, 1)
	// 如果程序捕获到 Ctrl+C 或者是系统关闭指令，则把他发送到quit Channel
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 在没有把 syscall.SIGINT,syscall.SIGTERM 信息发送到quit 之前会一直阻塞到这里
	<-quit
	log.Printf("Shutting down server...")
	if s.Close != nil {
		s.Close()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server exited gracefully.")
}
