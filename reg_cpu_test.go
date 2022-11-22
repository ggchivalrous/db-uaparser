package dbuaparser_test

import (
	"fmt"
	"testing"

	dbuaparser "github.com/ggchivalrous/db-uaparser"
)

func TestRegCpu(t *testing.T) {
	cpuLen := len(cpuList)
	errCount := 0
	for _, v := range cpuList {
		res := dbuaparser.GetCpu(v.Ua)
		state := res.Cpu == v.Expect.Name
		if !state {
			errCount++
			fmt.Println("匹配: ", v.Desc, " 失败，结果：name: ", res.Cpu)
		}
	}

	fmt.Println("匹配总数: ", cpuLen, " 成功总数:", cpuLen-errCount, "失败总数: ", errCount)

	if errCount > 0 {
		t.Error("匹配失败数量: ", errCount)
	}
}

type TestCpu struct {
	Desc   string
	Ua     string
	Expect TestCpuRes
}

type TestCpuRes struct {
	Name string
}

var cpuList = []TestCpu{
	{
		Desc: "i686",
		Ua:   "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:19.0) Gecko/20100101 Firefox/19.0",
		Expect: TestCpuRes{
			Name: "ia32",
		},
	},
	{
		Desc: "i386",
		Ua:   "Mozilla/5.0 (X11; U; FreeBSD i386; en-US; rv:1.7) Gecko/20040628 Epiphany/1.2.6",
		Expect: TestCpuRes{
			Name: "ia32",
		},
	},
	{
		Desc: "x86-64",
		Ua:   "Opera/9.80 (X11; Linux x86_64; U; Linux Mint; en) Presto/2.2.15 Version/10.10",
		Expect: TestCpuRes{
			Name: "amd64",
		},
	},
	{
		Desc: "win64",
		Ua:   "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.2; Win64; x64; Trident/6.0; .NET4.0E; .NET4.0C)",
		Expect: TestCpuRes{
			Name: "amd64",
		},
	},
	{
		Desc: "WOW64",
		Ua:   "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0)",
		Expect: TestCpuRes{
			Name: "amd64",
		},
	},
	{
		Desc: "ARM",
		Ua:   "Mozilla/5.0 (Mobile; Windows Phone 8.1; Android 4.0; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 635) like iPhone OS 7_0_3 Mac OS X AppleWebKit/537 (KHTML, like Gecko) Mobile Safari/537",
		Expect: TestCpuRes{
			Name: "arm",
		},
	},
	{
		Desc: "ARMv61",
		Ua:   "Mozilla/5.0 (X11; U; Linux armv61; en-US; rv:1.9.1b2pre) Gecko/20081015 Fennec/1.0a1",
		Expect: TestCpuRes{
			Name: "arm",
		},
	},
	{
		Desc: "ARMv7",
		Ua:   "Mozilla/5.0 (Linux ARMv7) WebKitGTK+/3.4.9 vimprobable2",
		Expect: TestCpuRes{
			Name: "arm",
		},
	},
	{
		Desc: "ARMv7l",
		Ua:   "Mozilla/5.0 (SMART-TV; X11; Linux armv7l) AppleWebKit/537.42 (KHTML, like Gecko) Chromium/25.0.1349.2 Chrome/25.0.1349.2 Safari/537.42",
		Expect: TestCpuRes{
			Name: "arm",
		},
	},
	{
		Desc: "ARMv7l",
		Ua:   "Mozilla/5.0 (X11; CrOS armv7l 9765.85.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.123 Safari/537.36",
		Expect: TestCpuRes{
			Name: "arm",
		},
	},
	{
		Desc: "Nokia N900 Linux mobile",
		Ua:   "Mozilla/5.0 (Maemo; Linux armv7l; rv:10.0) Gecko/20100101 Firefox/10.0 Fennec/10.0",
		Expect: TestCpuRes{
			Name: "arm",
		},
	},
	{
		Desc: "ARMEABI",
		Ua:   "[FBAN/FB4A;FBAV/237.0.0.44.120;FBBV/170693408;FBDM/{density=1.75,width=720,height=1280};FBLC/en_US;FBRV/172067074;FBCR/ ;FBMF/samsung;FBBD/samsung;FBPN/com.facebook.katana;FBDV/SM-S367VL;FBSV/9;FBBK/1;FBOP/19;FBCA/armeabi-v7a:armeabi;]",
		Expect: TestCpuRes{
			Name: "arm",
		},
	},
	{
		Desc: "ARMv8",
		Ua:   "Mozilla/5.0 (X11; Linux armv8l; rv:45.0) Gecko/20100101 Firefox/45.0",
		Expect: TestCpuRes{
			Name: "arm64",
		},
	},
	{
		Desc: "AARCH64",
		Ua:   "Mozilla/5.0 (X11; CrOS aarch64 13310.93.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.133 Safari/537.36",
		Expect: TestCpuRes{
			Name: "arm64",
		},
	},
	{
		Desc: "ARM64",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; ARM64; RM-1096) AppleWebKit/537.36 (KHTML like Gecko) Chrome/51.0.2704.79 Safari/537.36 Edge/14.14393",
		Expect: TestCpuRes{
			Name: "arm64",
		},
	},
	{
		Desc: "ARM64",
		Ua:   "Mozilla/5.0 (Linux; arm_64; Android 9; HRY-LX1T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 YaBrowser/19.12.1.121.00 Mobile Safari/537.36",
		Expect: TestCpuRes{
			Name: "arm64",
		},
	},
	{
		Desc: "Pocket PC",
		Ua:   "Opera/9.7 (Windows Mobile; PPC; Opera Mobi/35166; U; en) Presto/2.2.1",
		Expect: TestCpuRes{
			Name: "arm",
		},
	},
	{
		Desc: "Mac PowerPC",
		Ua:   "Mozilla/4.0 (compatible; MSIE 4.5; Mac_PowerPC)",
		Expect: TestCpuRes{
			Name: "ppc",
		},
	},
	{
		Desc: "Mac PowerPC",
		Ua:   "Mozilla/4.0 (compatible; MSIE 5.17; Mac_PowerPC Mac OS; en)",
		Expect: TestCpuRes{
			Name: "ppc",
		},
	},
	{
		Desc: "Mac PowerPC",
		Ua:   "iCab/2.9.5 (Macintosh; U; PPC; Mac OS X)",
		Expect: TestCpuRes{
			Name: "ppc",
		},
	},
	{
		Desc: "Mac OS X on PowerPC using Firefox",
		Ua:   "Mozilla/5.0 (Macintosh; PPC Mac OS X x.y; rv:10.0) Gecko/20100101 Firefox/10.0",
		Expect: TestCpuRes{
			Name: "ppc",
		},
	},
	{
		Desc: "UltraSPARC",
		Ua:   "Mozilla/5.0 (X11; U; SunOS sun4u; en-US; rv:1.9b5) Gecko/2008032620 Firefox/3.0b5",
		Expect: TestCpuRes{
			Name: "sparc",
		},
	},
	{
		Desc: "sparc64",
		Ua:   "ELinks (0.4.3; NetBSD 3.0.2PATCH sparc64; 141x19)",
		Expect: TestCpuRes{
			Name: "sparc64",
		},
	},
	{
		Desc: "QuickTime",
		Ua:   "QuickTime/7.5.6 (qtver=7.5.6;cpu=IA32;os=Mac 10.5.8)",
		Expect: TestCpuRes{
			Name: "ia32",
		},
	},
	{
		Desc: "XBMC",
		Ua:   "XBMC/12.0 Git:20130127-fb595f2 (Windows NT 6.1;WOW64;Win64;x64; http://www.xbmc.org)",
		Expect: TestCpuRes{
			Name: "amd64",
		},
	},
	{
		Desc: "IRIX64",
		Ua:   "Mozilla/4.8C-SGI [en] (X11; U; IRIX64 6.5 IP27",
		Expect: TestCpuRes{
			Name: "irix64",
		},
	},
}
