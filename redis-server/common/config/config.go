package config

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "sync"
)

const DevMode = true
const Version = "1.0.0"

type Config struct {
    configInfo map[string]string
}

var config *Config
var mutex sync.Mutex

func Load(configPath string) {
    if config != nil {
        return
    }
    mutex.Lock()
    defer mutex.Unlock()
    if config != nil {
        return
    }
    config = new(Config)
    config.init(configPath)
}

func (config *Config) init(configPath string) {
    if len(configPath) == 0 {
        configPath = "/etc/myreids.conf"
    }
    file, err := os.Open(configPath)
    if file != nil {
        defer func() {
            _ = file.Close()
        }()
    }
    if err != nil {
        fmt.Println("load config ", configPath, " failed ", err)
    }
    config.configInfo = make(map[string]string)
    reader := bufio.NewReader(file)
    for {
        line, _, err := reader.ReadLine()
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println("read config file error", err)
            break
        }
        lineStr := strings.TrimSpace(string(line))
        if strings.HasPrefix(lineStr, "#") {
            continue
        }
        lineArr := strings.Split(lineStr, "=")
        if len(lineArr) != 2 {
            fmt.Println("config file format error", lineStr)
            continue
        }
        config.configInfo[strings.TrimSpace(lineArr[0])] = strings.TrimSpace(lineArr[1])
    }
}

func ServerPort() string {
    return config.stringValue("server_port", "6379")
}

func StringValue(key string, defaultValue string) string {
    return config.stringValue(key, defaultValue)
}

func (config *Config) stringValue(key string, defaultValue string) string {
    value, ok := config.configInfo[key]
    if ok {
        return value
    } else {
        return defaultValue
    }
}
