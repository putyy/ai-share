package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type AppConfig struct {
	// 版本号
	Version string `env:"VERSION"`
	// 运行模式
	RunMode string `env:"RUN_MODE"`
	// 服务端口
	HTTPPort int `env:"HTTP_PORT"`
	// 读写超时
	ReadTimeout  int64 `env:"READ_TIMEOUT"`
	WriteTimeout int64 `env:"WRITE_TIMEOUT"`
	// 分页每一页数量
	PageSize int `env:"PAGE_SIZE" default:"10"`
	// 默认图片
	DefaultImg string `default:"https://img0.baidu.com/it/u=3751066216,564345018&fm=253&fmt=auto&app=138&f=JPEG?w=400&h=400"`
}

var (
	App    AppConfig
	Mysql  MysqlConfig
	Redis  RedisConfig
	Jwt    JwtConfig
	Qiniu  QiniuConfig
	Wechat WechatConfig
)

var config = make(map[string]string)

func init() {
	inputFile, inputError := os.Open("./.env")
	if inputError != nil {
		fmt.Printf("open env file error: %s!", inputError.Error())
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	var reg *regexp.Regexp
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		}

		reg = regexp.MustCompile("\\s+")
		if reg.ReplaceAllString(inputString, "") == "" {
			continue
		}
		s := strings.Split(inputString, "=")
		s[1] = strings.Trim(s[1], "\r\n")
		s[1] = strings.Trim(s[1], "\"")
		config[s[0]] = s[1]
	}
	reflectParse(&App)
	reflectParse(&Mysql)
	reflectParse(&Redis)
	reflectParse(&Jwt)
	reflectParse(&Qiniu)
	reflectParse(&Wechat)
}

func reflectParse(c interface{}) {
	rValue := reflect.ValueOf(c)
	rType := reflect.TypeOf(c)
	if rType.Kind() != reflect.Ptr {
		fmt.Print("Can only be pointer type")
		return
	}
	rValue = rValue.Elem()
	rType = rType.Elem()
	for i := 0; i < rType.NumField(); i++ {
		envValue := ""
		tField := rType.Field(i)
		vField := rValue.Field(i)

		env := tField.Tag.Get("env")
		dValue := tField.Tag.Get("default")

		if env != "" {
			configValue, ok := config[env]
			if ok {
				envValue = configValue
			}
		} else if dValue != "" {
			envValue = dValue
		}

		rFieldType := tField.Type.Kind()
		if rFieldType == reflect.String {
			vField.Set(reflect.ValueOf(envValue))
		} else if rFieldType == reflect.Int {
			envValue, _ := strconv.Atoi(envValue)
			vField.Set(reflect.ValueOf(envValue))
		} else if rFieldType == reflect.Int64 {
			envValue, _ := strconv.ParseInt(envValue, 10, 64)
			vField.Set(reflect.ValueOf(envValue))

		}
	}
}
