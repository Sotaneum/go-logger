package logger

import (
	"encoding/json"

	file "github.com/Sotaneum/go-json-file"
	ktime "github.com/Sotaneum/go-kst-time"
)

// Logger : 로그를 남기고 불러오고 합니다.
type Logger struct {
	file *file.File
}

// Add : 로그를 추가합니다.
func (logger *Logger) Add(data interface{}) error {
	str, _ := json.Marshal(data)
	now := ktime.GetNowWithSecond().Format("2006-01-02T15:04:05.000")

	objData := map[string]string{}

	json.Unmarshal([]byte(logger.file.Load()), &objData)

	objData[now] = string(str)

	return logger.file.SaveObject(objData)
}

// Get : 모든 로그를 가져옵니다.
func (logger *Logger) Get() *map[string]string {
	objData := map[string]string{}

	json.Unmarshal([]byte(logger.file.Load()), &objData)

	return &objData
}

// Remove : Log를 삭제합니다.
func (logger *Logger) Remove() error {
	return logger.file.Remove()
}

// New : Logger를 생성합니다.
func New(path, name string) *Logger {
	log := Logger{}

	log.file = &file.File{Path: path, Name: name}

	return &log
}
