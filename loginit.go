package journal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

type fileLogWriter struct {
	Filename   string `json:"filename"`
	MaxLines   int    `json:"maxlines"`
	MaxFiles   int    `json:"maxfiles"`
	MaxSize    int    `json:"maxsize"`
	Daily      bool   `json:"daily"`
	MaxDays    int64  `json:"maxdays"`
	Hourly     bool   `json:"hourly"`
	MaxHours   int64  `json:"maxhours"`
	Rotate     bool   `json:"rotate"`
	Level      int    `json:"level"`
	Perm       string `json:"perm"`
	RotatePerm string `json:"rotateperm"`
}

func Start() {
	prefix := "./log"
	dic := time.Now().Format("20060102")
	fileNameOnly, suffix := "info", ".log"
	//文件路径
	filename := path.Join(prefix, dic, fileNameOnly+suffix)
	//判断文件路径是否存在，除文件以外的文件夹，如果不存在就创建
	err := pathDeal(path.Join(prefix, dic))
	if err != nil {
		fmt.Println(err)
		return
	}
	//日志的参数
	var param = fileLogWriter{
		Filename:   filename,
		MaxLines:   1000,
		MaxFiles:   999,
		MaxSize:    1 << 28,
		Daily:      true,
		MaxDays:    7,
		Hourly:     false,
		MaxHours:   168,
		Rotate:     true,
		RotatePerm: "0440",
		Level:      logs.LevelWarning,
		Perm:       "0660",
	}
	//将日志参数转换为json格式
	bytes, err := json.Marshal(param)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2 输出引擎 和输出位置
	err = logs.SetLogger(logs.AdapterFile, string(bytes))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 3 日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号,可以如下设置
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	// 4 为了提升性能，设置异步输出
	logs.Async()

}
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func pathDeal(filePath string) (err error) {
	exist, err := pathExists(filePath)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}
	if exist {
		fmt.Printf("has dir![%v]\n", filePath)
	} else {
		fmt.Printf("no dir![%v]\n", filePath)
		//创建文件夹
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir falied[%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
	return
}

//读取key=value类型的配置文件
func initConfig(path string) map[string]string {
	config := make(map[string]string)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(any(err))
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(any(err))
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}
