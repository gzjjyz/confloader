package confloader

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
)

type Loader struct {
	loadFile string
	conf     interface{}
}

func NewLoader(loadConfigFile string, conf interface{}) *Loader {
	loader := &Loader{
		loadFile: loadConfigFile,
		conf:     conf,
	}
	return loader
}

func (l *Loader) Load() error {
	buf, err := os.ReadFile(l.loadFile)
	if err != nil {
		return err
	}

	newCfgType := reflect.New(reflect.ValueOf(l.conf).Elem().Type())
	newestCfg := newCfgType.Interface()

	switch filepath.Ext(l.loadFile) {
	case ".json":
		err = json.Unmarshal(buf, newestCfg)
		if err != nil {
			return err
		}
	default:
		return errors.New("only JSON is supported")
	}

	reflect.ValueOf(l.conf).Elem().Set(newCfgType.Elem())

	return nil
}
