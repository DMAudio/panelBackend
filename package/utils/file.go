package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func FileStat(path string) (string, os.FileInfo, error) {
	var err error
	path, err = filepath.Abs(strings.TrimSpace(path))
	if err != nil {
		return "", nil, err
	}
	var info os.FileInfo
	if info, err = os.Stat(path); err != nil {
		return "", nil, err
	}
	return path, info, nil
}

func FileContent(path string) ([]byte, error) {
	var err error
	var file *os.File
	file, err = os.OpenFile(path, os.O_RDONLY, 0)
	defer func() { _ = file.Close() }()
	if err != nil {
		return nil, err
	}
	var content []byte
	content, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func DirList(path string, pattern *regexp.Regexp, recursion bool) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var list []string
	for _, f := range files {
		filePath := filepath.Join(path, f.Name())
		if f.IsDir() && recursion {
			var sublist []string
			sublist, err = DirList(filePath, pattern, recursion)
			if err != nil {
				return nil, err
			}
			list = append(list, sublist...)
			continue
		}
		if pattern != nil && !pattern.MatchString(filePath) {
			continue
		}
		list = append(list, filePath)
	}
	return list, nil
}
