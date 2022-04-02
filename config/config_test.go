package config

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	NewConfig("/Users/yapi/WorkSpace/VscodeWorkSpace/vending/config.yml")
	fmt.Println(Base)
}
