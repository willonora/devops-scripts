package helpers

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupslot/config"
)

func GenerateRandomString(length int) (string, error) {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	for i, v := range b {
		b[i] = letterBytes[v%byte(len(letterBytes))]
	}
	return string(b), nil
}

func ErrMsg(err error) string {
	return fmt.Sprintf("Error: %v", err)
}

func DEDuplicateStringSlice(slice []string) []string {
	uniqueSlice := make([]string, 0)
	seen := make(map[string]bool)
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			uniqueSlice = append(uniqueSlice, item)
		}
	}
	return uniqueSlice
}

func ReadFileContent(filename string) (string, error) {
	b, err := os.ReadFile(filepath.Clean(filename))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func SaveFileContent(filename string, content string) error {
	return os.WriteFile(filepath.Clean(filename), []byte(content), 0644)
}

func absPath(path string) string {
	if !filepath.IsAbs(path) {
		cwd, err := os.Getwd()
		if err != nil {
			return ""
		}
		return filepath.Join(cwd, path)
	}
	return path
}

func GetLatestFileInDir(directory string) (string, error) {
	files, err := filepath.Glob(filepath.Join(directory, "*"))
	if err != nil {
		return "", err
	}
	var latest string
	var latestTimestamp int64
	for _, file := range files {
		f, err := os.Stat(file)
		if err != nil {
			continue
		}
		if f.ModTime().Unix() > latestTimestamp {
			latest = file
			latestTimestamp = f.ModTime().Unix()
		}
	}
	return latest, nil
}

func timeDiffString(start time.Time, now time.Time) string {
	diff := now.Sub(start)
	hours := int(diff.Hours())
	minutes := int(diff.Minutes())
	seconds := int(diff.Seconds())
	return fmt.Sprintf("%d hours, %d minutes, %d seconds", hours, minutes, seconds)
}

func GetSortedSlice[T any](slice []T, less func(x, y T) bool) []T {
	var result []T
	for _, value := range slice {
		for i, v := range result {
			if less(value, v) {
				result = append(result[:i], append([]T{value}, result[i:]...)...)
				break
			}
		}
		if len(result) == 0 || less(value, result[len(result)-1]) {
			result = append(result, value)
		}
	}
	return result
}

func ScanStringForPattern(pattern string, str string) []string {
	result := strings.Split(strings.ReplaceAll(str, pattern, ""), "")
	var patternMatched []string
	for _, v := range result {
		if strings.Contains(v, pattern) {
			patternMatched = append(patternMatched, v)
		}
	}
	return patternMatched
}

func TrimPrefix(s string, prefix string) string {
	if !strings.HasPrefix(s, prefix) {
		return s
	}
	return s[len(prefix):]
}

func getLock() *sync.Mutex {
	return &sync.Mutex{}
}

func IsEmptyString(s string) bool {
	return strings.TrimSpace(s) == ""
}

func GetConfigString(key string) (string, error) {
	value, err := config.Get(key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func GetConfigInt(key string) (int, error) {
	value, err := config.Get(key)
	if err != nil {
		return 0, err
	}
	var i int
	if n, err := fmt.Sscan(value, &i); n != 1 || err != nil {
		return 0, errors.New("invalid value for int key")
	}
	return i, nil
}

func GetConfigInt64(key string) (int64, error) {
	value, err := config.Get(key)
	if err != nil {
		return 0, err
	}
	var i int64
	if n, err := fmt.Sscan(value, &i); n != 1 || err != nil {
		return 0, errors.New("invalid value for int64 key")
	}
	return i, nil
}

func GetConfigFloat64(key string) (float64, error) {
	value, err := config.Get(key)
	if err != nil {
		return 0, err
	}
	var f float64
	if n, err := fmt.Sscan(value, &f); n != 1 || err != nil {
		return 0, errors.New("invalid value for float64 key")
	}
	return f, nil
}

func GetConfigBool(key string) (bool, error) {
	value, err := config.Get(key)
	if err != nil {
		return false, err
	}
	var b bool
	if n, err := fmt.Sscan(value, &b); n != 1 || err != nil {
		return false, errors.New("invalid value for bool key")
	}
	return b, nil
}

func GetBase64String() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}