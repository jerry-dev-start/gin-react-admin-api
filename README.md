


#### 项目目录结构
```text
.
├── cmd/
│   └── admin-api/
│       └── main.go       # 入口文件：在这里按顺序调用初始化函数
├── configs/
│   └── config.yaml       # 配置文件
├── internal/
│   ├── bootstrap/        # 核心：项目启动初始化逻辑
│   │   ├── database.go   # GORM/DB 初始化
│   │   ├── logger.go     # 日志库初始化 (Zap/Logrus)
│   │   ├── redis.go      # Redis 初始化
│   │   └── config.go     # Viper 配置读取
│   ├── conf/             # 配置对应的映射结构体 (Struct)
│   ├── data/             # 持久层（使用初始化好的 DB 实例）
│   ├── ...
├── pkg/                  # 抽离出的公共工具
└── go.mod
```