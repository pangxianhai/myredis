package logger

import (
    "fmt"
    "log"
    "os"
    "redis-server/data/config"
    "runtime"
    "strconv"
    "strings"
    "sync"
    "time"
)

type Logger struct {
    fileLog  *log.Logger
    fileName string
    stdLog   *log.Logger
    isDebug  *bool
    mutex    sync.Mutex
}

var logger *Logger
var mutex sync.Mutex

//日志模块初始化
func Init() {
    mutex.Lock()
    logger = newLog()
    mutex.Unlock()
}

func newLog() *Logger {
    logger := new(Logger)
    fileName := logger.newLogFileName()
    flag := 0
    if config.DevMode {
        flag = log.Lshortfile | log.Lmicroseconds | log.Ldate
    } else {
        flag = log.Lmicroseconds | log.Ldate
    }
    if fileName != "" {
        file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
        if err != nil {
            fmt.Println("load logger failed", err)
        } else {
            logger.fileLog = log.New(file, "", flag)
            logger.fileName = fileName
        }
    }
    console := config.StringValue("log_console", "true")
    if strings.ToLower(console) == "true" {
        logger.stdLog = log.New(os.Stderr, "", flag)
    }
    return logger
}

func (log *Logger) logrotate() {
    if log.fileLog == nil {
        return
    }
    log.mutex.Lock()
    log.mutex.Unlock()
    fileName := logger.newLogFileName()
    if fileName == log.fileName {
        return
    }
    log.fileName = fileName
    logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        fmt.Println("log rotate failed", err)
    }
    log.fileLog.SetOutput(logFile)
}

func (log *Logger) newLogFileName() string {
    fileName := config.StringValue("log_file_path", "")
    if fileName == "" {
        return ""
    } else {
        timeValue := time.Now().Format("2006-01-02")
        return fileName + "." + timeValue
    }
}

func (log *Logger) IsDebug() bool {
    if *log.isDebug {
        return *log.isDebug
    }
    level := config.StringValue("log_level", "info")
    if level == "DEBUG" || level == "debug" {
        *log.isDebug = true
    } else {
        *log.isDebug = false
    }
    return *log.isDebug
}

func (log *Logger) output(s string) {
    if log.fileLog != nil {
        _ = log.fileLog.Output(3, s)
    }
    if log.stdLog != nil {
        _ = log.stdLog.Output(3, s)
    }
}

func Debug(v ...interface{}) {
    if !logger.IsDebug() {
        return
    }
    logger.logrotate()
    s := "[DEBUG] " + fmt.Sprintln(v...)
    logger.output(s)
}

func Info(v ...interface{}) {
    logger.logrotate()
    s := "[INFO] " + fmt.Sprintln(v...)
    logger.output(s)
}

func Warn(v ...interface{}) {
    logger.logrotate()
    s := "[WARN] " + fmt.Sprintln(v...)
    logger.output(s)
}

func Error(v ...interface{}) {
    logger.logrotate()
    s := "[ERROR] " + fmt.Sprintln(v...)
    logger.output(s)
}

func PrintBanner() {
    var bannerArr = []string{"",
        "               _._",
        "          _.-``__ ''-._",
        "     _.-``    `.  `_.  ''-._",
        " .-`` .-```.  ```\\/    _.,_ ''-._     Running in " + runtime.GOOS + " " + runtime.GOARCH,
        "(    '      ,       .-`  | `,    )    Redis " + config.Version,
        "|`-._`-...-` __...-.``-._|'` _.-'|    Port: " + config.ServerPort(),
        "|    `-._   `._    /     _.-'    |    PID: " + strconv.Itoa(os.Getpid()),
        " `-._    `-._  `-./  _.-'    _.-'",
        "|`-._`-._    `-.__.-'    _.-'_.-'|",
        "|    `-._`-._        _.-'_.-'    | ",
        " `-._    `-._`-.__.-'_.-'    _.-'",
        "|`-._`-._    `-.__.-'    _.-'_.-'|",
        "|    `-._`-._        _.-'_.-'    |",
        " `-._    `-._`-.__.-'_.-'    _.-'",
        "     `-._    `-.__.-'    _.-'",
        "          `-._        _.-'",
        "              `-.__.-'",
    }
    banner := strings.Join(bannerArr, "\n")
    Info(banner)
}
