package args

import (
	"github.com/open-tdp/go-helper/dborm"
)

// 调试模式

var Debug bool

// 数据库参数

var Database = dborm.Config{}

// 数据集参数

var Dataset struct {
	Dir    string
	Secret string
}

// 日志参数

var Logger struct {
	Dir    string
	Level  string
	Target string
}

// 主节点参数

var Server struct {
	Listen string
	JwtKey string
}
