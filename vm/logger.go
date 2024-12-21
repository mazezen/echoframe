package vm

import (
	"github.com/mazezen/echoframe/utils"
	"github.com/mazezen/itools"
	"go.uber.org/zap"
)

func newLogger(t int) *zap.Logger {
	pr, err := utils.FindProjectRoot()
	if err != nil {
		panic(err)
	}
	m := itools.Gm.Get("logger").(map[string]interface{})
	var (
		p  string
		ok bool
	)
	switch t {
	case 1: // Log
		p, ok = m["path"].(string)
		if !ok {
			panic(ok)
		}
	case 2: // OrmLog
		p, ok = m["ormlog"].(string)
		if !ok {
			panic(ok)
		}
	}

	p = itools.CompactStr(pr, p)

	mx, ok := m["max"].(int)
	if !ok {
		panic(ok)
	}

	li, ok := m["live"].(int)
	if !ok {
		panic(ok)
	}
	co, ok := m["compress"].(bool)
	if !ok {
		panic(ok)
	}
	lo, ok := m["localtime"].(bool)
	if !ok {
		panic(ok)
	}
	return itools.NewLogger(p, mx, li, lo, co)
}
