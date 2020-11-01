package config

import (
    "bufio"
    "io"
    "log"
    "os"
    "strconv"
    "strings"
    "sync"
)

var config *Config
var configLock sync.Mutex

func InitConfig(configPath string) {
    if config != nil {
        return
    }
    configLock.Lock()
    defer configLock.Unlock()
    if config != nil {
        return
    }
    config = &Config{}
    config.init(configPath)
}

func GetConfigInstance() *Config {
    if config != nil {
        return config
    }
    InitConfig("")
    return config
}

type Config struct {
    configInfo map[string]string
}

func (config *Config) init(configPath string) {
    if len(configPath) == 0 {
        configPath = "/etc/myreids.conf"
    }
    file, err := os.Open(configPath)
    if file != nil {
        defer file.Close()
    }
    if err != nil {
        log.Fatalln("加载配置:", configPath, " 错误", err)
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

func (config *Config) GetServerPort() int {
    return config.getIntValue("server_port", 6379)
}

func (config *Config) getIntValue(key string, defaultValue int) int {
    value, ok := config.configInfo[key]
    if ok {
        iv, err := strconv.Atoi(value)
        if err == nil {
            return iv
        } else {
            return defaultValue
        }
    } else {
        return defaultValue
    }
}
