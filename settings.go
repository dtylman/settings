package settings

import (
	"fmt"
	"strings"
)

//FileFormat configuration file format
type FileFormat int

const (
	//FormatJSON ...
	FormatJSON = iota + 1
)

//View represents a settings view
type View map[string]interface{}

//New creates a new empty setting view
func New() View {
	return make(View)
}

//Load loads a file and populate into a view
func (v View) Load(fileName string, format FileFormat) error {
	switch format {
	case FormatJSON:
		return v.loadJSON(fileName)
	}
	return fmt.Errorf("Format not supported: %v", format)
}

//Print prints the view in the format provided
func (v View) Print(format FileFormat) error {
	switch format {
	case FormatJSON:
		return v.printJSON()
	}
	return fmt.Errorf("Format not supported: %v", format)
}

//Save saves the view to a file
func (v View) Save(fileName string, format FileFormat) error {
	switch format {
	case FormatJSON:
		return v.saveJSON(fileName)
	}
	return fmt.Errorf("Format not supported: %v", format)
}

//Set sets a configuration key
func (v View) Set(key string, value interface{}) {
	subkeys := strings.SplitN(key, ".", 2)
	if len(subkeys) == 1 {
		v[key] = value
	} else {
		name := subkeys[0]
		view, ok := v[name].(View)
		if !ok {
			view, ok = v[name].(map[string]interface{})
			if !ok {
				view = make(View)
			}
		}
		view.Set(subkeys[1], value)
		v[name] = view
	}
}

//Get gets configuration key
func (v View) Get(key string, def interface{}) interface{} {
	subkeys := strings.SplitN(key, ".", 2)
	val, ok := v[subkeys[0]]
	if ok {
		view, ok := val.(View)
		if ok {
			return view.Get(subkeys[1], def)
		}
		m, ok := val.(map[string]interface{})
		if ok {
			return View(m).Get(subkeys[1], def)
		}
		return val
	}
	return def
}

//Has returns true if view has the key
func (v View) Has(key string) bool {
	subkeys := strings.SplitN(key, ".", 2)
	val, ok := v[subkeys[0]]
	if ok {
		view, ok := val.(map[string]interface{})
		if ok {
			return View(view).Has(subkeys[1])
		}
		return true
	}
	return false
}

//GetString gets a configuration item as a string
func (v View) GetString(key string, def string) string {
	val := v.Get(key, def)
	s, ok := val.(string)
	if ok {
		return s
	}
	return fmt.Sprintf("%v", val)
}
