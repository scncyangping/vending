package util

import "testing"

func TestSnowflake(t *testing.T) {
	for i := 0; i < 100000; i++ {
		println(NextId())
	}
}
