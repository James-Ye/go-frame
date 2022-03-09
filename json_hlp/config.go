package json_hlp

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/James-Ye/go-frame/logger"
	"github.com/James-Ye/go-frame/path_mgr"
)

/////////////////////////////////////////////////////////////////////////////////////////
type Config struct {
	ConfigMain
}

func (c *Config) InitConfig() {
	c.load(0)
}

type ConfigMain struct {
	m_mutex    sync.Mutex
	m_doc      interface{}
	m_default  interface{} //store default value, no changed beside init()
	m_filename string
	m_modify   bool
}

type ConfigNode struct {
	m_main          *ConfigMain
	m_value         interface{}
	m_default_value interface{}
}

func (base *ConfigNode) Define(name string, def interface{}) (interface{}, bool) {
	if len(name) == 0 || base.m_main == nil || base.m_main.m_doc == nil {
		return nil, false
	}

	base.m_main.m_mutex.Lock()
	defer base.m_main.m_mutex.Unlock()

	if _, ok := base.m_main.m_doc.(map[string]interface{})[name]; !ok {
		if def != nil {
			base.m_value.(map[string]interface{})[name] = def
			base.m_default_value.(map[string]interface{})[name] = def
			base.m_main.setmodify()
			return nil, true
		} else {
			node := make(map[string]interface{})
			node_default := make(map[string]interface{})
			base.m_value.(map[string]interface{})[name] = node
			base.m_default_value.(map[string]interface{})[name] = node_default
			base.m_main.setmodify()
			base.m_value = node
			base.m_default_value = node_default
			return base, true
		}
	}

	return nil, false
}

func (base *ConfigNode) isinit(sub *ConfigNode) bool {
	return sub != nil && sub.m_main != nil
}

func (base *ConfigNode) initsub(sub *ConfigNode, name string) {
	if sub == nil || len(name) == 0 || sub.m_main == nil || sub.m_value == nil {
		return
	}

	base.m_main.m_mutex.Lock()
	defer base.m_main.m_mutex.Unlock()

	if sub.m_main != nil {
		return
	}
	sub.m_main = base.m_main
	if v, ok := base.m_value.(map[string]interface{})[name]; ok {
		sub.m_value = v
	}
}

func (main *ConfigMain) GetRoot() *ConfigNode {
	return main.getRoot()
}

func (main *ConfigMain) getRoot() *ConfigNode {
	if main.m_doc == nil {
		return nil
	}

	base := new(ConfigNode)
	base.m_main = main
	base.m_value = main.m_doc
	base.m_default_value = main.m_default

	return base
}

func (base *ConfigNode) GetNode(name string) *ConfigNode {
	return base.getNode(name)
}

func (base *ConfigNode) getNode(name string) *ConfigNode {

	if v, ok := base.get(name); ok {
		return v.(*ConfigNode)
	} else {
		//此处应该返回默认值
		return nil
	}
}

func (base *ConfigNode) GetValue(name string) interface{} {
	return base.getValue(name)
}

func (base *ConfigNode) getValue(name string) interface{} {
	if v, ok := base.get(name); ok {
		return v
	} else {
		//此处应该返回默认值
		return nil
	}
}

func (base *ConfigNode) get(name string) (interface{}, bool) {
	if len(name) == 0 || base.m_value == nil || base.m_main == nil {
		return nil, false
	}

	base.m_main.m_mutex.Lock()
	defer base.m_main.m_mutex.Unlock()

	if v, ok := base.m_value.(map[string]interface{})[name]; ok {
		if strings.Compare(reflect.TypeOf(v).Name(), "string") == 0 || strings.Compare(reflect.TypeOf(v).Name(), "int") == 0 {
			return v, true
		} else {
			base.m_value = v
			if v, ok := base.m_default_value.(map[string]interface{})[name]; ok {
				base.m_default_value = v
			}
			return base, true
		}
	}
	//由于初始化的时候已经把m_default中的值全部赋值给了m_doc,所以理论上不存在取不到值需要取默认值的情况

	return nil, false
}

func (base *ConfigNode) SetPassword(name string, value interface{}) interface{} {
	//此处作加密处理
	svalue := value
	return base.set(name, svalue)
}

func (base *ConfigNode) Set(name string, value interface{}) interface{} {
	return base.set(name, value)
}

func (base *ConfigNode) set(name string, value interface{}) interface{} {
	if len(name) == 0 || base.m_value == nil {
		return nil
	}

	base.m_main.m_mutex.Lock()
	defer base.m_main.m_mutex.Unlock()

	if v, ok := base.m_value.(map[string]interface{})[name]; ok {
		if strings.Compare(reflect.TypeOf(v).Name(), reflect.TypeOf(value).Name()) == 0 {
			if v != value {
				base.m_value.(map[string]interface{})[name] = value
				base.m_main.setmodify()
			}
		}
		return v
	} else {
		return nil
	}
}

func (base *ConfigNode) ToStr() string {
	if base.m_main == nil || base.m_value == nil {
		return "{}"
	}

	base.m_main.m_mutex.Lock()
	defer base.m_main.m_mutex.Unlock()

	return MapToJson(base.m_value.(map[string]interface{}))

}
func (base *ConfigNode) DupValue() (interface{}, bool) {
	if base.m_main == nil || base.m_value == nil {
		return nil, false
	}

	base.m_main.m_mutex.Lock()
	defer base.m_main.m_mutex.Unlock()

	return base.m_value, true
}

func (base *ConfigNode) CopyFrom(cfg *ConfigNode) bool {
	if base.m_main == nil || base.m_value == nil {
		return false
	}

	base.m_main.m_mutex.Lock()
	defer base.m_main.m_mutex.Unlock()

	base.m_value = cfg.m_value
	base.m_main.setmodify()
	return true
}

func (main *ConfigMain) setmodify() {
	main.m_modify = true
}

func (main *ConfigMain) load(size int) {
	main.m_mutex.Lock()
	defer main.m_mutex.Unlock()

	if m, ok := main.loadbin(size); ok {
		fmt.Println("load ok")
		fmt.Println(main)
		doc := main.m_doc.(map[string]interface{})
		MergeMap(&(doc), m.(map[string]interface{}))
		fmt.Println(main)
	} else {
		fmt.Println("load fail")
		main.m_doc = make(map[string]interface{})
		main.m_default = make(map[string]interface{})
	}
}

func (main *ConfigMain) Save() {
	if main.m_modify || main.m_doc == nil {
		return
	}

	main.m_mutex.Lock()
	defer main.m_mutex.Unlock()

	main.savebin()
}

func (main *ConfigMain) Ismodify() bool {
	return main.m_modify
}

func (cf *ConfigMain) savebin() bool {
	return Savejson(cf.m_filename, cf.m_doc)
}
func (cf *ConfigMain) loadbin(size int) (interface{}, bool) {
	return Loadjson(cf.m_filename, size)
}

func (cf *ConfigMain) Loadfile(filename string, path *string, size int) {
	cf.loadfile(filename, path, size)
}

func (cf *ConfigMain) loadfile(filename string, path *string, size int) {
	strfilename := ""
	if path == nil {
		strfilename = path_mgr.GetAppPath()
	} else {
		strfilename = *path
	}

	strfilename += "\\"
	strfilename += filename

	cf.m_filename = strfilename
	logger.Debug("loadfile :%s", cf.m_filename)
	cf.load(size)
}
