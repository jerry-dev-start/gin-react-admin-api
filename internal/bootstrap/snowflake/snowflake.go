package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

// InitSnowflake 初始化雪花算法节点
func InitSnowflake(startTime string, machineID int64) (err error) {
	var st time.Time
	// 设定一个起始时间（通常是项目启动日期），可以延长算法的使用年限
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}

	snowflake.Epoch = st.UnixNano() / 1000000
	// 创建节点，machineID 在单体应用中通常传 1 即可
	node, err = snowflake.NewNode(machineID)
	return err
}

// GenID 生成 64 位整型 ID
func GenID() uint64 {
	return uint64(node.Generate().Int64())
}

// GenIDString 生成字符串形式的 ID（前端 JavaScript 处理大整数建议用字符串）
func GenIDString() string {
	return node.Generate().String()
}
