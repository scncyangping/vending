package constants

const (
	ZERO     = iota
	EmptyStr = ""
	// 默认页数
	DefaultPage = 1
	// 默认每页数量
	DefaultPageSize = 10

	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)
