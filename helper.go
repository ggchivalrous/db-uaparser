package dbuaparser

import (
	"regexp"
	"strings"

	"github.com/dlclark/regexp2"
)

type PropMap struct {
	Index     int
	Reg       string
	Replace   string
	Handle    Handle
	Value     string
	HandleMap map[string]string
}

type Props struct {
	Name      PropMap
	Version   PropMap
	Device    PropMap
	Vendor    PropMap
	Type      PropMap
	Cpu       PropMap
	Engine    PropMap
	EngineVer PropMap
	Os        PropMap
	OsVer     PropMap
}

type RegexpMap struct {
	Regs []string
	Props
}

type HandleRegexpMap struct {
	Regs []IRegexp
	Props
}

type IRegexp struct {
	reg2 *regexp2.Regexp
	reg  *regexp.Regexp
}

func (i *IRegexp) FindStringMatch(s string) []string {
	if i.reg != nil {
		return i.reg.FindStringSubmatch(s)
	} else {
		match, _ := i.reg2.FindStringMatch(s)
		if match == nil {
			return []string{}
		}
		matchList := match.Groups()
		matchLen := match.GroupCount()
		strList := make([]string, matchLen)
		for i := 0; i < matchLen; i++ {
			strList[i] = matchList[i].String()
		}
		return strList
	}
}

type Handle interface {
	handle(matchStr []byte, maps map[string]string) string
}

type UaInfo struct {
	Browser    string // 浏览器名称
	BrowserVer string // 浏览器版本
	Cpu        string // CPU架构
	Device     string // 设备名称
	Vendor     string // 供应商
	Type       string // 设备类型
	Engine     string // 引擎名称
	EngineVer  string // 引擎版本
	Os         string // 操作系统
	OsVer      string // 操作系统版本
}

type handleMaps = map[string]string

func has(str1 string, str2 string) bool {
	return strings.Contains(lowerize(str1), lowerize(str2))
}

func lowerize(str string) string {
	return strings.ToLower(str)
}

func trim(str string, length int) string {
	regS, _ := regexp.Compile(`/^\s\s*/`)
	regE, _ := regexp.Compile(`/\s\s*$/`)

	byteList := regS.ReplaceAll([]byte(str), []byte(""))
	byteList = regE.ReplaceAll(byteList, []byte(""))

	if length == 0 {
		return string(byteList)
	} else {
		byteListCopy := make([]byte, _UA_MAX_LENGTH)
		copy(byteListCopy, byteList)
		return string(byteListCopy)
	}
}

func execMap(ua string, arrays []HandleRegexpMap) UaInfo {
	uaInfo := &UaInfo{}
	arrLen := len(arrays)
	matches := false

	for i := 0; i < arrLen && !matches; i++ {
		regs := arrays[i].Regs
		regLen := len(regs)
		props := arrays[i].Props

		for j := 0; j < regLen && !matches; j++ {
			strList := regs[j].FindStringMatch(ua)
			strListLen := len(strList)
			matches = strListLen > 0

			if matches {
				handlePropsValue(props, uaInfo)

				for k := 1; k < strListLen; k++ {
					v := []byte(strings.ToLower(strList[k]))
					handlePropsIndex(props, k, v, uaInfo)
				}
			}
		}
	}

	return *uaInfo
}

func handlePropsValue(props Props, uaInfo *UaInfo) {
	if props.Name.Index == 0 {
		uaInfo.Browser = handlePropMap(props.Name, nil)
	}

	if props.Version.Index == 0 {
		uaInfo.BrowserVer = handlePropMap(props.Version, nil)
	}

	if props.Device.Index == 0 {
		uaInfo.Device = handlePropMap(props.Device, nil)
	}

	if props.Vendor.Index == 0 {
		uaInfo.Vendor = handlePropMap(props.Vendor, nil)
	}

	if props.Type.Index == 0 {
		uaInfo.Type = handlePropMap(props.Type, nil)
	}

	if props.Cpu.Index == 0 {
		uaInfo.Cpu = handlePropMap(props.Cpu, nil)
	}

	if props.Engine.Index == 0 {
		uaInfo.Engine = handlePropMap(props.Engine, nil)
	}

	if props.EngineVer.Index == 0 {
		uaInfo.EngineVer = handlePropMap(props.EngineVer, nil)
	}

	if props.Os.Index == 0 {
		uaInfo.Os = handlePropMap(props.Os, nil)
	}

	if props.OsVer.Index == 0 {
		uaInfo.OsVer = handlePropMap(props.OsVer, nil)
	}
}

func handlePropsIndex(props Props, index int, match []byte, uaInfo *UaInfo) {
	if props.Name.Index == index {
		uaInfo.Browser = handlePropMap(props.Name, match)
	} else if props.Version.Index == index {
		uaInfo.BrowserVer = handlePropMap(props.Version, match)
	} else if props.Vendor.Index == index {
		uaInfo.Vendor = handlePropMap(props.Vendor, match)
	} else if props.Type.Index == index {
		uaInfo.Type = handlePropMap(props.Type, match)
	} else if props.Device.Index == index {
		uaInfo.Device = handlePropMap(props.Device, match)
	} else if props.Cpu.Index == index {
		uaInfo.Cpu = handlePropMap(props.Cpu, match)
	} else if props.Engine.Index == index {
		uaInfo.Engine = handlePropMap(props.Engine, match)
	} else if props.EngineVer.Index == index {
		uaInfo.EngineVer = handlePropMap(props.EngineVer, match)
	} else if props.Os.Index == index {
		uaInfo.Os = handlePropMap(props.Os, match)
	} else if props.OsVer.Index == index {
		uaInfo.OsVer = handlePropMap(props.OsVer, match)
	}
}

func handlePropMap(props PropMap, match []byte) string {
	returnStr := ""

	if props.Handle != nil {
		if props.Reg == "" {
			returnStr = props.Handle.handle(match, props.HandleMap)
		} else if props.Replace != "" {
			r, _ := regexp.Compile(props.Reg)
			str := r.ReplaceAll(match, []byte(props.Replace))
			returnStr = props.Handle.handle(str, props.HandleMap)
		} else {
			returnStr = props.Handle.handle(match, props.HandleMap)
		}
	} else if props.Reg != "" {
		r, _ := regexp.Compile(props.Reg)
		str := r.ReplaceAll(match, []byte(props.Replace))
		returnStr = string(str)
	} else if props.Value != "" {
		returnStr = props.Value
	} else {
		returnStr = string(match)
	}

	return returnStr
}
