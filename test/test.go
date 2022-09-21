/**

Filename: 		conf.go
Author: 		alvin - yishuwen.zb@ccbft.com
Description:	loongrpc config logic
Create:			2022-06-18 10:09:01
Last Modified:	2022-07-01 11:32:23

*/

package conf

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// conf is a command-line and config file parser.
// It is a golang implement of co/flag.
// See https://github.com/idealvin/cocoyaxi for details.
// By Alvin, 2022.06.17.

// internal config items
func init() {
	String("conf", "", ".path of config file")
	Bool("mkconf", false, ".generate config file")
}

type configType int

const (
	typeBool configType = iota + 1
	typeInt
	typeString
	typeFloat64
)

type configItem struct {
	name  string
	defv  string
	help  string
	pkg   string
	value interface{}
	t     configType
}

var configMap = map[string]*configItem{}

func newConfigItem(name, value, help string, v interface{}, t configType) *configItem {
	pkg := packageName()
	return &configItem{
		name:  name,
		defv:  value,
		help:  help,
		pkg:   pkg,
		value: v,
		t:     t,
	}
}

// 定义 bool 类型的配置项
func Bool(name string, val bool, help string) {
	_, ok := configMap[name]
	if ok {
		panic("config already exists: " + name)
	}
	if val {
		configMap[name] = newConfigItem(name, "true", help, val, typeBool)
	} else {
		configMap[name] = newConfigItem(name, "false", help, val, typeBool)
	}
}

// 定义 int 类型的配置项
func Int(name string, val int, help string) {
	_, ok := configMap[name]
	if ok {
		panic("config already exists: " + name)
	}
	configMap[name] = newConfigItem(name, int2str(val), help, val, typeInt)
}

// 定义 float64 类型的配置项
func Float64(name string, val float64, help string) {
	_, ok := configMap[name]
	if ok {
		panic("config already exists: " + name)
	}
	configMap[name] = newConfigItem(name, strconv.FormatFloat(val, 'g', 6, 64), help, val, typeFloat64)
}

// 定义 string 类型的配置项
func String(name string, val string, help string) {
	_, ok := configMap[name]
	if ok {
		panic("config already exists: " + name)
	}
	configMap[name] = newConfigItem(name, val, help, val, typeString)
}

// 获取 bool 类型配置项的值
func GetBool(key string) bool {
	x, ok := configMap[key]
	if !ok {
		panic("config not exists: " + key)
	}

	if x.t != typeBool {
		panic("config is not bool type: " + key)
	}

	return x.value.(bool)
}

// 获取 int 类型配置项的值
func GetInt(key string) int {
	x, ok := configMap[key]
	if !ok {
		panic("config not exists: " + key)
	}

	if x.t != typeInt {
		panic("config is not int type: " + key)
	}

	return x.value.(int)
}

// 获取 string 类型配置项的值
func GetString(key string) string {
	x, ok := configMap[key]
	if !ok {
		panic("config not exists: " + key)
	}

	if x.t != typeString {
		panic("config is not string type: " + key)
	}

	return x.value.(string)
}

// 获取 float64 类型配置项的值
func GetFloat64(key string) float64 {
	x, ok := configMap[key]
	if !ok {
		panic("config not exists: " + key)
	}

	if x.t != typeFloat64 {
		panic("config is not float64 type: " + key)
	}

	return x.value.(float64)
}

func Parse() []string {
	args := os.Args
	n := len(args)
	if n <= 1 {
		return []string{}
	}

	// -- & --help
	if n == 2 && strings.HasPrefix(args[1], "--") {
		if args[1] == "--help" {
			showHelp()
			os.Exit(0)
		}

		if args[1] == "--" {
			showAllFlags()
			os.Exit(0)
		}
	}

	kv := map[string]string{}
	bools := []string{}
	v := analyze(args, &kv, &bools)

	// check -conf
	x, ok := kv["conf"]
	if ok {
		configMap["conf"].value = x
	} else if len(v) > 0 {
		if strings.HasSuffix(v[0], ".conf") || strings.HasSuffix(v[0], "config") {
			info, err := os.Stat(v[0])
			if err == nil && !info.IsDir() {
				configMap["conf"].value = v[0]
			}
		}
	}

	// parse config file if -conf is set
	path := configMap["conf"].value.(string)
	if len(path) > 0 {
		parseConfigFile(path)
	}

	for key, val := range kv {
		err := setValue(key, val)
		if len(err) > 0 {
			fmt.Println(err)
			os.Exit(0)
		}
	}

	for i := 0; i < len(bools); i++ {
		err := setBool(bools[i])
		if len(err) > 0 {
			fmt.Println(err)
			os.Exit(0)
		}
	}

	if GetBool("mkconf") {
		mkconf(args[0])
		os.Exit(0)
	}

	return v
}

func findConfig(key string) *configItem {
	x, ok := configMap[key]
	if ok {
		return x
	}
	return nil
}

func setValue(key, val string) string {
	x, ok := configMap[key]
	if !ok {
		return "config not defined: " + key
	}

	switch x.t {
	case typeString:
		x.value = val
	case typeBool:
		if val == "true" || val == "1" {
			x.value = true
		} else if val == "false" || val == "0" {
			x.value = false
		} else {
			return "invalid value for bool: " + val
		}
	case typeInt:
		r, e := str2int(val)
		if e != nil {
			return e.Error() + ": " + val
		}
		x.value = r
	case typeFloat64:
		r, e := strconv.ParseFloat(val, 64)
		if e != nil {
			return "invalid value for float64: " + val
		}
		x.value = r
	default:
		return "unknown config type"
	}

	return ""
}

func setBool(key string) string {
	x, ok := configMap[key]
	if ok {
		if x.t == typeBool {
			x.value = true
			return ""
		} else {
			return "value not set for non-bool config: " + key
		}
	}

	if len(key) == 1 {
		return "undefined bool flag: " + key
	}

	for i := 0; i < len(key); i++ {
		s := key[i : i+1]
		x, ok := configMap[s]
		if !ok {
			return "undefined bool config -" + s + " in -" + key
		} else if x.t != typeBool {
			return "-" + s + " is not bool in -" + key
		} else {
			x.value = true
		}
	}

	return ""
}

func analyze(args []string, kv *map[string]string, bools *[]string) []string {
	res := []string{}

	name := ""
	next := ""
	var cfg *configItem
	for i := 1; i < len(args); i++ {
		arg := args[i]
		name = strings.TrimLeft(arg, "-")
		if len(name) == len(arg) || ('0' <= arg[1] && arg[1] <= '9') {
			res = append(res, name)
			continue
		}

		if len(name) == 0 {
			fmt.Println("invalid parameter: " + arg)
			os.Exit(0)
		}

		if i+1 == len(args) {
			goto no_value
		}

		next = args[i+1]
		if strings.HasPrefix(next, "-") {
			x := strings.TrimLeft(next, "-")
			if x != "" && !('0' <= next[1] && next[1] <= '9') {
				goto no_value
			}
		}

		cfg = findConfig(name)
		if cfg == nil {
			goto no_value
		}
		if cfg.t != typeBool {
			goto has_value
		}

		if next == "0" || next == "1" || next == "false" || next == "true" {
			goto has_value
		}

	no_value:
		*bools = append(*bools, name)
		continue
	has_value:
		(*kv)[name] = next
		continue
	}

	return res
}

func (f *configItem) typeStr() string {
	switch f.t {
	case typeBool:
		return "bool"
	case typeInt:
		return "int"
	case typeString:
		return "string"
	case typeFloat64:
		return "float64"
	default:
		return "unknown"
	}
}
func (f *configItem) print() {
	color.Set(color.FgGreen)
	fmt.Print("    -", f.name)
	color.Set(color.FgBlue)
	fmt.Print("  ", f.help, "\n")
	color.Unset()
	fmt.Print("\ttype: ", f.typeStr())
	fmt.Print("\t  default: ", f.defv)
	fmt.Print("\n\tfrom package: ", f.pkg, "\n")
}

func showHelp() {
	fmt.Print("usage:  ")
	color.Blue("$exe [-flag] [value]\n")
	fmt.Print("\t", "$exe -x -i 23 -s ok        # x=true, i=23, s=\"ok\"\n")
	fmt.Print("\t", "$exe --                    # print all flags\n")
	fmt.Print("\t", "$exe -mkconf               # generate config file\n")
	fmt.Print("\t", "$exe -conf xx.conf         # run with config file\n\n")

	first := true
	for _, val := range configMap {
		if val.help != "" {
			if first {
				first = false
				fmt.Println("flags:")
			}
			val.print()
		}
	}
}

func showAllFlags() {
	fmt.Print("flags:\n")
	for _, val := range configMap {
		if val.help != "" {
			val.print()
		}
	}
}

func mkconf(exe string) {
	flags := map[string]map[string]*configItem{}
	for key, val := range configMap {
		if val.help == "" || val.help[0] == '.' {
			continue
		}
		_, ok := flags[val.pkg]
		if !ok {
			flags[val.pkg] = make(map[string]*configItem, 0)
		}
		flags[val.pkg][key] = val
	}

	var fname string
	var keys []string
	for key, _ := range flags {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	if strings.HasSuffix(exe, ".exe") {
		fname = exe[0:len(exe)-4] + ".conf"
	} else {
		fname = exe + ".conf"
	}

	//	ioutil.WriteFile()
	buf := bytes.NewBuffer(nil)
	buf.WriteString(strings.Repeat("#", 72))
	buf.WriteByte('\n')
	buf.WriteString("###  > # for comments\n")
	buf.WriteString("###  > k,m,g,t,p (case insensitive, 1k for 1024, etc.)\n")
	buf.WriteString(strings.Repeat("#", 72))
	buf.WriteString("\n\n\n")

	for _, key := range keys {
		m := flags[key]
		var xx []string
		for x, _ := range m {
			xx = append(xx, x)
		}
		sort.Strings(xx)

		buf.WriteString("# >> ")
		buf.WriteString(key)
		buf.WriteByte('\n')
		buf.WriteByte('#')
		buf.WriteString(strings.Repeat("=", 71))
		buf.WriteByte('\n')

		for _, x := range xx {
			cfg := m[x]
			buf.WriteString("# ")
			buf.WriteString(strings.ReplaceAll(cfg.help, "\n", "\n# "))
			buf.WriteByte('\n')
			buf.WriteString(cfg.name)
			buf.WriteString(" = ")
			switch cfg.t {
			case typeString:
				buf.WriteString(formatStr(cfg.value.(string)))
			case typeInt:
				buf.WriteString(int2str(cfg.value.(int)))
			case typeBool:
				if cfg.value.(bool) {
					buf.WriteString("true")
				} else {
					buf.WriteString("false")
				}
			case typeFloat64:
				buf.WriteString(strconv.FormatFloat(cfg.value.(float64), 'g', 6, 64))
			default:
				panic("unknown config type")
			}
			buf.WriteString("\n\n")
		}
		buf.WriteByte('\n')
	}

	ioutil.WriteFile(fname, buf.Bytes(), 0644)
}

// 配置文件格式:
//   # 注释, 支持单行注释及行尾注释
//   key = value
func parseConfigFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("open file failed: ", path)
		os.Exit(0)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	n := 0
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		n++
		s := strings.ReplaceAll(string(b), "　", " ")
		s = strings.Trim(s, " \t\r\n")

		// # 开头的行是注释, 目前只支持单行注释，不支持行尾注释
		if s == "" || s[0] == '#' {
			continue
		}

		index := strings.Index(s, "=")
		if index <= 0 {
			fmt.Println("invalid config: ", s, ", at ", path, ":", n)
			os.Exit(0)
		}

		key := strings.Trim(s[:index], " \t")
		if len(key) == 0 {
			fmt.Println("invalid config: ", s, ", at ", path, ":", n)
			os.Exit(0)
		}

		val := strings.TrimLeft(s[index+1:], " \t")
		val = removeQuotesAndComments(val)

		x := setValue(key, val)
		if x != "" {
			if strings.HasPrefix(x, "config not defined") {
				fmt.Println("WARNING: ", x, ", at ", path, ":", n)
			} else {
				fmt.Println(err, ", at ", path, ":", n)
				os.Exit(0)
			}
		}
	}
}

func removeQuotesAndComments(s string) string {
	if s == "" {
		return ""
	}

	r := ""
	c := s[0]
	p := 0
	l := 0

	if c == '"' || c == '\'' || c == '`' {
		if strings.HasPrefix(s, "```") {
			p = strings.LastIndex(s, "```")
			l = 3
		} else {
			p = strings.LastIndex(s, s[0:1])
			l = 1
		}

		if p == 0 {
			goto no_quotes
		}

		r = strings.TrimLeft(s[p+l:], " \t")
		if r == "" {
			r = strings.TrimRight(s, " \t")
		} else if r[0] == '#' {
			r = strings.TrimRight(s[0:p+l], " \t")
		} else {
			goto no_quotes
		}

		return r[l:p]
	}

no_quotes:
	p = strings.Index(s, "#")
	if p < 0 {
		return s
	}
	return strings.TrimRight(s[0:p], " \t")
}

func packageName() string {
	pc, _, _, _ := runtime.Caller(3)
	path := runtime.FuncForPC(pc).Name() // package.func
	v := strings.Split(path, ".")
	n := len(v)

	s := ""
	if v[n-2][0] == '(' {
		s = strings.Join(v[0:n-2], ".")
	} else {
		s = strings.Join(v[0:n-1], ".")
	}

	return strings.TrimSuffix(s, ".init")
}

func int2str(v int) string {
	if (0 <= v && v <= 8192) || (v < 0 && v >= -8192) {
		return strconv.Itoa(v)
	}

	u := "kmgtp"
	i := -1
	for {
		if v == 0 || (v&1023) != 0 {
			break
		}
		v = v >> 10
		i++
		if i >= 4 {
			break
		}
	}

	s := strconv.Itoa(v)
	if i >= 0 {
		s = s + string(u[i])
	}
	return s
}

type convError struct {
	err string
}

func (e *convError) Error() string {
	return e.err
}

func str2int(s string) (int, error) {
	if len(s) == 0 {
		return 0, nil
	}

	off := 0
	c := s[len(s)-1]
	switch c {
	case 'k', 'K':
		off = 10
	case 'm', 'M':
		off = 20
	case 'g', 'G':
		off = 30
	case 't', 'T':
		off = 40
	case 'p', 'P':
		off = 50
	}

	if off == 0 {
		return strconv.Atoi(s)
	}

	s = s[0 : len(s)-1]
	if len(s) == 0 {
		return 1 << off, nil
	}

	r, e := strconv.Atoi(s)
	if e != nil {
		return 0, e
	}

	if r == 0 {
		return 0, nil
	}

	if r < (math.MinInt64>>off) || r > (math.MaxInt64>>off) {
		return 0, &convError{err: "out of range"}
	}

	return r << off, nil
}

// add quotes to the string if necessary
func formatStr(s string) string {
	x := strings.Index(s, "\"")
	y := strings.Index(s, "'")
	z := strings.Index(s, "`")
	if x < 0 && y < 0 && z < 0 {
		return s
	}

	if x < 0 {
		return "\"" + s + "\""
	}

	if y < 0 {
		return "'" + s + "'"
	}

	if z < 0 {
		return "`" + s + "`"
	}

	return "```" + s + "```"
}
