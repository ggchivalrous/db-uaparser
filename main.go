package dbuaparser

import (
	"regexp"

	"github.com/dlclark/regexp2"
)

var browserHandleMap []HandleRegexpMap
var cpuHandleMap []HandleRegexpMap
var deviceHandleMap []HandleRegexpMap
var engineHandleMap []HandleRegexpMap
var osHandleMap []HandleRegexpMap

func init() {
	browserHandleMap = initRegexpMap(browserMap)
	cpuHandleMap = initRegexpMap(cpuMap)
	deviceHandleMap = initRegexpMap(deviceMap)
	engineHandleMap = initRegexpMap(engineMap)
	osHandleMap = initRegexpMap(osMap)
}

func initRegexpMap(maps []RegexpMap) []HandleRegexpMap {
	newMaps := make([]HandleRegexpMap, len(maps))

	for _, v := range maps {
		regs := make([]IRegexp, len(v.Regs))

		for i, v := range v.Regs {
			ireg := IRegexp{}

			reg, err := regexp.Compile(v)

			if err != nil {
				reg2, err := regexp2.Compile(v, regexp2.IgnoreCase)

				if err != nil {
					panic(err)
				}

				ireg.reg2 = reg2
			} else {
				ireg.reg = reg
			}
			regs[i] = ireg
		}
		newMaps = append(newMaps, HandleRegexpMap{
			Regs:  regs,
			Props: v.Props,
		})
	}

	return newMaps
}

func GetBrowser(ua string) BrowserInfo {
	uaInfo := execMap(ua, browserHandleMap)
	return BrowserInfo{
		Browser:    uaInfo.Browser,
		BrowserVer: uaInfo.BrowserVer,
	}
}

func GetCpu(ua string) CpuInfo {
	uaInfo := execMap(ua, cpuHandleMap)
	return CpuInfo{
		Cpu: uaInfo.Cpu,
	}
}

func GetDevice(ua string) DeviceInfo {
	uaInfo := execMap(ua, deviceHandleMap)
	return DeviceInfo{
		Device: uaInfo.Device,
		Vendor: uaInfo.Vendor,
		Type:   uaInfo.Type,
	}
}

func GetEngine(ua string) EngineInfo {
	uaInfo := execMap(ua, engineHandleMap)
	return EngineInfo{
		Engine:    uaInfo.Engine,
		EngineVer: uaInfo.EngineVer,
	}
}

func GetOs(ua string) OsInfo {
	uaInfo := execMap(ua, osHandleMap)
	return OsInfo{
		Os:    uaInfo.Os,
		OsVer: uaInfo.OsVer,
	}
}

func GetInfo(ua string) UaInfo {
	browser := GetBrowser(ua)
	cpu := GetCpu(ua)
	device := GetDevice(ua)
	engine := GetEngine(ua)
	os := GetOs(ua)

	return UaInfo{
		Browser:    browser.Browser,
		BrowserVer: browser.BrowserVer,
		Cpu:        cpu.Cpu,
		Device:     device.Device,
		Vendor:     device.Vendor,
		Type:       device.Type,
		Engine:     engine.Engine,
		EngineVer:  engine.EngineVer,
		Os:         os.Os,
		OsVer:      os.OsVer,
	}
}
