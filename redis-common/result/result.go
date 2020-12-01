package result

import "encoding/json"

const (
    SUCCESS   uint32 = 200
    ERROR     uint32 = 500
    NOT_FOUND uint32 = 404
)

type Result struct {
    Code uint32      `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

func New() *Result {
    return &Result{}
}

func NewOfCode(code uint32, msg string) *Result {
    return &Result{Code: code, Msg: msg}
}

func NewOfData(data interface{}) *Result {
    return &Result{Code: SUCCESS, Msg: "OK", Data: data}
}

func ToJson(result *Result) (string, error) {
    buf, err := json.Marshal(result)
    if err != nil {
        return "", err
    }
    return string(buf), nil
}

func FromJson(data string) (*Result, error) {
    result := New()
    err := json.Unmarshal([]byte(data), result)
    if err != nil {
        return nil, err
    }
    return result, err
}
