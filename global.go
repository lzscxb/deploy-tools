package main

type Parameter struct {
	Branch       string // 分支名称
	Dir          string // 监听目录
	TimeInterval int    // 监听时间间隔
	Start        bool   // 启动后是否立即部署一次
	Language     string // 语言
	ScriptPath   string // 脚本路径
	ProjectName  string // 项目名称
	LogDir       string // 日志文件目录
	FileLog      bool   // 将本程序运行日志保存在日志文件中
	Help         bool   // 查看帮助信息
	Background   bool   // 后台启动
}

var (
	Args   *Parameter
	Deploy = make(chan bool) // 部署
)

const (
	PID_FILE = ".pid" // pid 文件
)