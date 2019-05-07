package window

import (
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/icccm"
	lua "github.com/yuin/gopher-lua"
	"log"
)

func GetWindows(L *lua.LState) int {
	EnsureConnection()
	clientids, err := ewmh.ClientListGet(xgbUtil)
	if err != nil {
		log.Fatal(err)
	}
	table := lua.LTable{}
	for _, clientid := range clientids {
		name, err := ewmh.WmNameGet(xgbUtil, clientid)
		if err != nil || len(name) == 0 {
			name, _ = icccm.WmNameGet(xgbUtil, clientid)
		}
		pid, _ := ewmh.WmPidGet(xgbUtil, clientid)
		entry := lua.LTable{}
		entry.RawSetString("handle", lua.LNumber(clientid))
		entry.RawSetString("name", lua.LString(name))
		entry.RawSetString("pid", lua.LNumber(pid))
		table.Append(&entry)
	}
	L.Push(&table)
	return 1

}