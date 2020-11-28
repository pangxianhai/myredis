package config

import (
    "bufio"
    "io"
    "log"
    "os"
    "strings"
    "sync"
)

type Config struct {
    configInfo map[string]string
    mutex      sync.Mutex
}

var config *Config

func Load(configPath string) {
    if config != nil {
        return
    }
    config = new(Config)
    config.mutex.Lock()
    defer config.mutex.Unlock()
    if config != nil {
        return
    }
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
        log.Fatalln("加载配置:", configPath, " 失败", err)
    }
    config.configInfo = make(map[string]string)
    reader := bufio.NewReader(file)
    for {
        line, _, err := reader.ReadLine()
        if err != nil {
            if err == io.EOF {
                break
            }
            log.Println("读取配置文件错误", err)
        }
        lineStr := strings.TrimSpace(string(line))
        if strings.HasPrefix(lineStr, "#") {
            continue
        }
        lineArr := strings.Split(lineStr, "=")
        if len(lineArr) != 2 {
            log.Println("读取配置项格式错误", lineStr)
        }
        config.configInfo[strings.TrimSpace(lineArr[0])] = strings.TrimSpace(lineArr[1])
    }
}

func ServerPort() string {
    return config.stringValue("server_port", "6379")
}

func (config *Config) stringValue(key string, defaultValue string) string {
    value, ok := config.configInfo[key]
    if ok {
        return value
    } else {
        return defaultValue
    }
}
