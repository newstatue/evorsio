package constant

const BuildDir = "bin"
const AppName = "evorsio"
const LogDir = "logs"

type Component string

const (
	ComponentApp   Component = "app"
	ComponentFS    Component = "fs"
	ComponentWails Component = "wails"
)

type LogArg string

const (
	LogArgError     LogArg = "error"
	LogArgComponent LogArg = "component"
)

type ErrMsg string

const (
	ErrParseConfig ErrMsg = "配置解析出错"
)
