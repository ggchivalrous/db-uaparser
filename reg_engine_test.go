package dbuaparser_test

import (
	"fmt"
	"testing"

	dbuaparser "github.com/ggchivalrous/db-uaparser"
)

func TestRegEngine(t *testing.T) {
	engineLen := len(engineList)
	errCount := 0
	for _, v := range engineList {
		res := dbuaparser.GetEngine(v.Ua)
		state := res.Engine == v.Expect.Name && res.EngineVer == v.Expect.Version
		if !state {
			errCount++
			fmt.Println("匹配: ", v.Desc, " 失败，结果：name: ", res.Engine, " version: ", res.EngineVer)
		}
	}

	fmt.Println("匹配总数: ", engineLen, " 成功总数:", engineLen-errCount, "失败总数: ", errCount)

	if errCount > 0 {
		t.Error("匹配失败数量: ", errCount)
	}
}

type TestEngine struct {
	Desc   string
	Ua     string
	Expect TestEngineRes
}

type TestEngineRes struct {
	Name    string
	Version string
}

var engineList = []TestEngine{
	{
		Desc: "Blink",
		Ua:   "Mozilla/5.0 (Linux; Android 7.0; SM-G920I Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) OculusBrowser/3.4.9 SamsungBrowser/4.0 Chrome/57.0.2987.146 Mobile VR Safari/537.36",
		Expect: TestEngineRes{
			Name:    "blink",
			Version: "57.0.2987.146",
		},
	},
	{
		Desc: "EdgeHTML",
		Ua:   "Mozilla/5.0 (Windows NT 6.4; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/36.0.1985.143 Safari/537.36 Edge/12.0",
		Expect: TestEngineRes{
			Name:    "edgehtml",
			Version: "12.0",
		},
	},
	{
		Desc: "Flow",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) EkiohFlow/5.7.4.30559 Flow/5.7.4 (like Gecko Firefox/53.0 rv:53.0)",
		Expect: TestEngineRes{
			Name:    "flow",
			Version: "5.7.4.30559",
		},
	},
	{
		Desc: "Gecko",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64; rv:2.0b9pre) Gecko/20110111 Firefox/4.0b9pre",
		Expect: TestEngineRes{
			Name:    "gecko",
			Version: "2.0b9pre",
		},
	},
	{
		Desc: "Goanna",
		Ua:   "Mozilla/5.0 (Windows NT 5.1; rv:38.9) Gecko/20100101 Goanna/2.2 Firefox/38.9 PaleMoon/26.5.0",
		Expect: TestEngineRes{
			Name:    "goanna",
			Version: "2.2",
		},
	},
	{
		Desc: "KHTML",
		Ua:   "Mozilla/5.0 (compatible; Konqueror/4.5; FreeBSD) KHTML/4.5.4 (like Gecko)",
		Expect: TestEngineRes{
			Name:    "khtml",
			Version: "4.5.4",
		},
	},
	{
		Desc: "NetFront",
		Ua:   "Mozilla/4.0 (PDA; Windows CE/1.0.1) NetFront/3.0",
		Expect: TestEngineRes{
			Name:    "netfront",
			Version: "3.0",
		},
	},
	{
		Desc: "Presto",
		Ua:   "Opera/9.80 (Windows NT 6.1; Opera Tablet/15165; U; en) Presto/2.8.149 Version/11.1",
		Expect: TestEngineRes{
			Name:    "presto",
			Version: "2.8.149",
		},
	},
	{
		Desc: "Tasman",
		Ua:   "Mozilla/4.0 (compatible; MSIE 6.0; PPC Mac OS X 10.4.7; Tasman 1.0)",
		Expect: TestEngineRes{
			Name:    "tasman",
			Version: "1.0",
		},
	},
	{
		Desc: "Trident",
		Ua:   "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
		Expect: TestEngineRes{
			Name:    "trident",
			Version: "6.0",
		},
	},
	{
		Desc: "WebKit",
		Ua:   "Mozilla/5.0 (Windows; U; Windows NT 6.1; sv-SE) AppleWebKit/533.19.4 (KHTML, like Gecko) Version/5.0.3 Safari/533.19.4",
		Expect: TestEngineRes{
			Name:    "webkit",
			Version: "533.19.4",
		},
	},
	{
		Desc: "WebKit",
		Ua:   "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/27.0.1453.110 Safari/537.36",
		Expect: TestEngineRes{
			Name:    "webkit",
			Version: "537.36",
		},
	},
	{
		Desc: "WebOS TV 5.x",
		Ua:   "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36 WebAppManager",
		Expect: TestEngineRes{
			Name:    "blink",
			Version: "68.0.3440.106",
		},
	},
	{
		Desc: "WebOS TV 4.x",
		Ua:   "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.34 Safari/537.36 WebAppManager",
		Expect: TestEngineRes{
			Name:    "blink",
			Version: "53.0.2785.34",
		},
	},
	{
		Desc: "WebOS TV 3.x",
		Ua:   "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.36 (KHTML, like Gecko) QtWebEngine/5.2.1 Chrome/38.0.2125.122 Safari/537.36 WebAppManager",
		Expect: TestEngineRes{
			Name:    "blink",
			Version: "38.0.2125.122",
		},
	},
	{
		Desc: "WebOS TV 2.x",
		Ua:   "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/538.2 (KHTML, like Gecko) Large Screen WebAppManager Safari/538.2",
		Expect: TestEngineRes{
			Name:    "webkit",
			Version: "538.2",
		},
	},
	{
		Desc: "WebOS TV 1.x",
		Ua:   "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.41 (KHTML, like Gecko) Large Screen WebAppManager Safari/537.41",
		Expect: TestEngineRes{
			Name:    "webkit",
			Version: "537.41",
		},
	},
}
