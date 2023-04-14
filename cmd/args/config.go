package args

// 调试模式

var Debug bool

// 数据库参数

var Database struct {
	Type   string
	Host   string
	User   string
	Passwd string
	Name   string
	Option string
}

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
