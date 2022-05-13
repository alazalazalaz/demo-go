package main

import (
	lua "github.com/yuin/gopher-lua"
	"log"
)

func main() {
	//callLuaFunc()

	callLuaFuncReg()
}

func callLuaFuncReg() {
	errorLog := "[DataName]:user_new_chat_pf,[Data]:{\"game_id\":\"b2\",\"server_id\":\"0\",\"message\":\" \u200f ‏ ‣HHHHHH这个是有错的\",\"translate_cn\":\"很好，谢谢你\"}\n"
	l := lua.NewState()
	defer l.Close()
	if err := l.DoFile("./go_call_lua/be_call.lua"); err != nil {
		panic(err)
	}

	err := l.CallByParam(lua.P{
		Fn:      l.GetGlobal("regText"),
		NRet:    0,
		Protect: false,
	}, lua.LString(errorLog))

	if err != nil {
		panic(err)
	}
}

func callLuaFunc() {
	l := lua.NewState()
	defer l.Close()
	if err := l.DoFile("./go_call_lua/be_call.lua"); err != nil {
		panic(err)
	}

	err := l.CallByParam(lua.P{
		Fn:      l.GetGlobal("printString"),
		NRet:    0,
		Protect: false,
	}, lua.LString("\u2023"), lua.LString("\\u200f"))

	if err != nil {
		panic(err)
	}

	log.Println("over")
}
