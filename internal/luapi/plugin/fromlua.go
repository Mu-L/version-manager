package plugin

import (
	"github.com/gvcgo/version-manager/internal/luapi/lua_global"
	lua "github.com/yuin/gopher-lua"
)

type LuaConfItem string

const (
	SDKName         LuaConfItem = "sdk_name"
	PluginName      LuaConfItem = "plugin_name"
	PluginVersion   LuaConfItem = "plugin_version"
	Prequisite      LuaConfItem = "prequisite"
	Homepage        LuaConfItem = "homepage"
	Crawler         LuaConfItem = "crawl"
	PostInstall     LuaConfItem = "postInstall" // optional
	CustomedInstall LuaConfItem = "install"     // optional
)

var InstallerConfig LuaConfItem = LuaConfItem(lua_global.InstallerConfigName)

func GetLuaConfItemString(L *lua.LState, item LuaConfItem) (result string) {
	v := L.GetGlobal(string(item))
	if v == nil {
		return
	}
	if v.Type() == lua.LTString {
		return v.String()
	} else if v.Type() == lua.LTFunction {
		if err := L.CallByParam(lua.P{
			Fn:      v,
			NRet:    1,
			Protect: true,
		}); err != nil {
			panic(err)
		}
		result := L.Get(-1)
		return result.String()
	}
	return
}

func DoesLuaItemExist(L *lua.LState, item LuaConfItem) bool {
	v := L.GetGlobal(string(item))
	return v.String() != "nil"
}
