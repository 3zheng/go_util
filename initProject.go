package util

//初始化项目，包括读取配置文件和新建日志文件
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

type DBConfig struct {
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	DB       string `json:"database"`
	UserId   string `json:"user id"`
	Password string `json:"password"`
}

// IP   string `json:"ip"`
type ServerConfig struct {
	Path      string `json:"path"`
	ForceIPv4 int    `json:"force ipv4"`
	IP        string `json:"ip"`
	Port      int    `json:"port"`
	CookieKey string `json:"cookie key"`
}

type Config struct {
	Database  DBConfig     `json:"database config"`
	Server    ServerConfig `json:"server config"`
	MysqlConn string       `json:"mysqlConn"`
	Mode      string       `json:"Mode"`
	None      string       `json:"None"`
}

func Recovermain() {
	if err := recover(); err != nil {
		var buf [9000]byte
		n := runtime.Stack(buf[:], false)
		log.Printf("[panic] err: %v\nstack: %s\n", err, buf[:n])
	}
}

func CreateNewFile(config Config, now time.Time) *os.File {
	var filepath string
	if runtime.GOOS == "windows" {
		filepath = "./log/logfile-"
	} else if runtime.GOOS == "linux" {
		filepath = config.Server.Path + "log/logfile-"
	} else {
		fmt.Println("系统不明")
		os.Exit(0)
	}
	today := fmt.Sprintf("%04d%02d%02d", now.Year(), now.Month(), now.Day())
	filepath = filepath + today
	logFile, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open log file failed.")
		return nil
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ltime)
	log.Println("log file opened.")
	return logFile
}

// 每天新建一个日志文件 使用go util.InitLog新建一个协程来初始化日志文件，因为有select阻塞
func InitLog(config Config) {
	now := time.Now()
	logFile := CreateNewFile(config, now) //创建日志文件
	// 获取第二天凌晨的时间00:01,不精准定位在00:00,以免创建新文件时还在前一天
	nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 1, 0, 0, now.Location())
	// 计算时间差
	duration := nextMidnight.Sub(now)
	// 输出秒数
	log.Printf("距离第二天凌晨还有 %v 秒\n", int(duration.Seconds()))
	//time.Sleep(duration) //第一天的程序启动时间是不确定的，使用Sleep到第二天的凌晨0点0分
	//log.Println("Sleep到凌晨")
	//第一天的程序启动时间是不确定的，先把定时器调整为到第二天凌晨
	tk := time.NewTicker(duration)
	//tk := time.NewTicker(5 * time.Minute)
	for now := range tk.C {
		log.Println("定时器时间到")
		tk.Reset(24 * time.Hour) //重置为24小时
		if logFile != nil {
			logFile.Close()
		} else {
			fmt.Println("日志文件句柄为空")
		}
		log.Println("now:", now.Format(time.DateTime))
		logFile = CreateNewFile(config, now)
	}
	//监听单个channel可以用for range替代for select
	/*
		for {
			select {
			case now := <-tk.C:
				//dosomething
			}
		}
	*/
	log.Println("退出InitLog")
}

// 读取配置文件
func ReadConfigFile() Config {
	args := os.Args //main命令行参数
	log.Println("main args = ", args)

	var content []byte
	var err error
	//带命令行参数 -config config.json
	if len(args) >= 3 {
		if args[1] == "-config" {
			log.Println("读取配置文件：", args[2])
			content, err = os.ReadFile(args[2])
			if err != nil {
				log.Fatal("Error when opening file: ", err)
			}
		}
	} else { //不带参数直接读取当前目录下的config.json文件
		content, err = os.ReadFile("config.json")
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}
	}

	//启动"net/http/pprof"运行监控
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	//tablemiddleware.CreatePanic()
	// Now let's unmarshall the data into `payload`
	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return config
}
