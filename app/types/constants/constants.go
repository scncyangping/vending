package constants

import "errors"

const (
	ZERO            = iota
	EmptyStr        = ""
	DefaultPage     = 1
	DefaultPageSize = 10

	// DebugMode indicates sg mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates sg mode is release.
	ReleaseMode = "release"
	// TestMode indicates sg mode is test.
	TestMode = "test"
)

var ErrNoDocuments = errors.New("mongo: no documents in result")
