package config

import (
	"sync"

	"cuelang.org/go/cue"
)

var configInstance *cue.Instance
var configRWLocker sync.RWMutex

func Lookup(fieldPath ...string) cue.Value {
	configRWLocker.RLock()
	defer configRWLocker.RUnlock()
	if configInstance == nil {
		return cue.Value{}
	}
	return configInstance.Lookup(fieldPath...)
}

func Set(value interface{}, fieldPath ...string) error {
	configRWLocker.Lock()
	defer configRWLocker.Unlock()
	instance, err := configInstance.Fill(value, fieldPath...)
	if err != nil {
		return err
	}
	configInstance = instance
	return nil
}
