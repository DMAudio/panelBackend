package config

import (
	_ "embed"
	"flag"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/build"

	"github.com/DMAudio/panelBackend/package/errgo/fmt/errors"
	"github.com/DMAudio/panelBackend/package/log"
	"github.com/DMAudio/panelBackend/package/utils"
)

//go:embed template.cue
var innerTemplate string

const innerTemplatePath = "__innerTemplate__"

var (
	configPath               string
	pathPattern              = regexp.MustCompile(`.*\.cue$`)
	ErrUnspecifiedConfigPath = errors.New("unspecified config path")
	ErrConfigNotFound        = errors.New("config not found")
	ErrInvalidConfigPath     = errors.New("invalid config path")
	ErrInvalidConfigFile     = errors.New("invalid config file")
)

func Load() {
	if instance, err := LoadFromPath(configPath); err != nil {
		log.Fatal(err)
	} else {
		configRWLocker.Lock()
		defer configRWLocker.Unlock()
		configInstance = instance
	}
}

func LoadFromPath(path string) (*cue.Instance, error) {
	path = filepath.Clean(strings.TrimSpace(path))
	// empty config path
	if path == "" {
		return nil, ErrUnspecifiedConfigPath
	}
	var err error
	// check config status
	var pathStat os.FileInfo
	if path, pathStat, err = utils.FileStat(path); err != nil {
		if os.IsNotExist(err) {
			return nil, ErrConfigNotFound
		}
		return nil, errors.Because(ErrInvalidConfigPath, err, path)
	}
	// load internal template
	var instance *cue.Instance
	var instanceBuilder = build.NewContext().NewInstance(path, nil)
	if err = instanceBuilder.AddFile(innerTemplatePath, innerTemplate); err != nil {
		return nil, errors.Becausef(ErrInvalidConfigFile, err, "path:%s", innerTemplatePath)
	}
	// load config files
	if pathStat.IsDir() {
		var pathList []string
		if pathList, err = utils.DirList(path, pathPattern, true); err != nil {
			return nil, errors.Becausef(err, ErrInvalidConfigPath, "path:%s", path)
		}
		if len(pathList) == 0 {
			return nil, errors.Becausef(os.ErrNotExist, ErrInvalidConfigPath, "path:%s", path)
		}
		for _, subPath := range pathList {
			if err = instanceBuilder.AddFile(subPath, nil); err != nil {
				return nil, errors.Becausef(ErrInvalidConfigFile, err, "path:%s", subPath)
			}
		}
	} else if err = instanceBuilder.AddFile(path, nil); err != nil {
		return nil, errors.Becausef(ErrInvalidConfigFile, err, "path:%s", path)
	}
	if instance, err = (&cue.Runtime{}).Build(instanceBuilder); err != nil {
		return nil, errors.Becausef(ErrInvalidConfigPath, err, "path:%s", path)
	}
	return instance, nil
}

func RegisterFlags() {
	flag.StringVar(&configPath, "c", "config", "Path to config file or dir")
}
