package dbuaparser_test

import (
	"fmt"
	"testing"

	dbuaparser "github.com/ggchivalrous/db-uaparser"
)

func TestRegBrowser(t *testing.T) {
	uaLen := len(uaList)
	errCount := 0
	for _, v := range uaList {
		res := dbuaparser.GetBrowser(v.Ua)
		state := res.Browser == v.Expect.Name && res.BrowserVer == v.Expect.Version
		if !state {
			errCount++
			fmt.Println("匹配: ", v.Desc, " 失败，结果：name: ", res.Browser, " ver: ", res.BrowserVer)
		}
	}

	fmt.Println("匹配总数: ", uaLen, " 成功总数:", uaLen-errCount, "失败总数: ", errCount)

	if errCount > 0 {
		t.Error("匹配失败数量: ", errCount)
	}
}

type TestBrowser struct {
	Desc   string
	Ua     string
	Expect TestBrowserRes
}

type TestBrowserRes struct {
	Name    string
	Version string
	Major   string
}

var uaList = []TestBrowser{
	{
		Desc: "360 Browser on iOS",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 12_4_1 like Mac OS X) AppleWebKit/607.3.9 (KHTML, like Gecko) Mobile/16G102 QHBrowser/317 QihooBrowser/4.0.10",
		Expect: TestBrowserRes{
			Name:    "360 browser",
			Version: "",
			Major:   "",
		},
	},
	{
		Desc: "Android Browser on Galaxy Nexus",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.0.2; en-us; Galaxy Nexus Build/ICL53F) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		Expect: TestBrowserRes{
			Name:    "android browser",
			Version: "4.0",
			Major:   "4",
		},
	},
	{
		Desc: "Android Browser on Galaxy S3",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.4; en-us; SAMSUNG GT-I9300I Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Version/1.5 Chrome/28.0.1500.94 Mobile Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "android browser",
			Version: "1.5",
			Major:   "1",
		},
	},
	{
		Desc: "Android Browser on HTC Flyer (P510E)",
		Ua:   "Mozilla/5.0 (Linux; U; Android 3.2.1; ru-ru; HTC Flyer P510e Build/HTK75C) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
		Expect: TestBrowserRes{
			Name:    "android browser",
			Version: "4.0",
			Major:   "4",
		},
	},
	{
		Desc: "Android Browser on Huawei Honor Glory II (U9508)",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.0.4; ru-by; HUAWEI U9508 Build/HuaweiU9508) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30 ACHEETAHI/2100050044",
		Expect: TestBrowserRes{
			Name:    "android browser",
			Version: "4.0",
			Major:   "4",
		},
	},
	{
		Desc: "Android Browser on Huawei P8 (H891L)",
		Ua:   "Mozilla/5.0 (Linux; Android 4.4.4; HUAWEI H891L Build/HuaweiH891L) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/33.0.0.0 Mobile Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "android browser",
			Version: "4.0",
			Major:   "4",
		},
	},
	{
		Desc: "Android Browser on Samsung S6 (SM-G925F)",
		Ua:   "Mozilla/5.0 (Linux; Android 5.0.2; SAMSUNG SM-G925F Build/LRX22G) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/3.0 Chrome/38.0.2125.102 Mobile Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "samsung browser",
			Version: "3.0",
			Major:   "3",
		},
	},
	{
		Desc: "Sailfish Browser",
		Ua:   "Mozilla/5.0 (Linux; U; Sailfish 3.0; Mobile; rv:45.0) Gecko/45.0 Firefox/45.0 SailfishBrowser/1.0",
		Expect: TestBrowserRes{
			Name:    "sailfish browser",
			Version: "1.0",
			Major:   "1",
		},
	},
	{
		Desc: "Arora",
		Ua:   "Mozilla/5.0 (Windows; U; Windows NT 5.1; de-CH) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.2",
		Expect: TestBrowserRes{
			Name:    "arora",
			Version: "0.2",
			Major:   "0",
		},
	},
	{
		Desc: "Avant",
		Ua:   "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; GTB5; Avant Browser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		Expect: TestBrowserRes{
			Name:    "avant ",
			Version: "",
			Major:   "",
		},
	},
	{
		Desc: "Avast Secure Browser",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36 Avast/72.0.1174.122",
		Expect: TestBrowserRes{
			Name:    "avast secure browser",
			Version: "72.0.1174.122",
			Major:   "72",
		},
	},
	{
		Desc: "AVG Secure Browser",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36 AVG/72.0.719.123",
		Expect: TestBrowserRes{
			Name:    "avg secure browser",
			Version: "72.0.719.123",
			Major:   "72",
		},
	},
	{
		Desc: "Baidu",
		Ua:   "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; baidubrowser 1.x)",
		Expect: TestBrowserRes{
			Name:    "baidubrowser",
			Version: "1.x",
			Major:   "1",
		},
	},
	{
		Desc: "Bolt",
		Ua:   "Mozilla/5.0 (X11; 78; CentOS; US-en) AppleWebKit/527+ (KHTML, like Gecko) Bolt/0.862 Version/3.0 Safari/523.15",
		Expect: TestBrowserRes{
			Name:    "bolt",
			Version: "0.862",
			Major:   "0",
		},
	},
	{
		Desc: "Bowser",
		Ua:   "Mozilla/5.0 (iOS; like Mac OS X) AppleWebKit/536.36 (KHTML, like Gecko) not Chrome/27.0.1500.95 Mobile/10B141 Safari/537.36 Bowser/0.2.1",
		Expect: TestBrowserRes{
			Name:    "bowser",
			Version: "0.2.1",
			Major:   "0",
		},
	},
	{
		Desc: "Camino",
		Ua:   "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10.4; en; rv:1.9.0.19) Gecko/2011091218 Camino/2.0.9 (like Firefox/3.0.19)",
		Expect: TestBrowserRes{
			Name:    "camino",
			Version: "2.0.9",
			Major:   "2",
		},
	},
	{
		Desc: "Camino on Mac",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.5; rv:2.0.1) Gecko/20100101 Firefox/4.0.1 Camino/2.2.1",
		Expect: TestBrowserRes{
			Name:    "camino",
			Version: "2.2.1",
			Major:   "2",
		},
	},
	{
		Desc: "Chimera",
		Ua:   "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; pl-PL; rv:1.0.1) Gecko/20021111 Chimera/0.6",
		Expect: TestBrowserRes{
			Name:    "chimera",
			Version: "0.6",
			Major:   "0",
		},
	},
	{
		Desc: "Chrome",
		Ua:   "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1090.0 Safari/536.6",
		Expect: TestBrowserRes{
			Name:    "chrome",
			Version: "20.0.1090.0",
			Major:   "20",
		},
	},
	{
		Desc: "Chrome",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4758.102 Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "chrome",
			Version: "100.0.4758.102",
			Major:   "100",
		},
	},
	{
		Desc: "Chrome Headless",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "chrome headless",
			Version: "",
			Major:   "",
		},
	},
	{
		Desc: "Chrome Headless",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/60.0.3112.113 Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "chrome headless",
			Version: "60.0.3112.113",
			Major:   "60",
		},
	},
	{
		Desc: "Chrome WebView",
		Ua:   "Mozilla/5.0 (Linux; Android 5.1.1; Nexus 5 Build/LMY48B; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/43.0.2357.65 Mobile Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "chrome webview",
			Version: "43.0.2357.65",
			Major:   "43",
		},
	},
	{
		Desc: "Chrome on iOS",
		Ua:   "Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; en) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3",
		Expect: TestBrowserRes{
			Name:    "chrome",
			Version: "19.0.1084.60",
			Major:   "19",
		},
	},
	{
		Desc: "Chromium",
		Ua:   "Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.7 (KHTML, like Gecko) Ubuntu/11.10 Chromium/16.0.912.21 Chrome/16.0.912.21 Safari/535.7",
		Expect: TestBrowserRes{
			Name:    "chromium",
			Version: "16.0.912.21",
			Major:   "16",
		},
	},
	{
		Desc: "Chrome on Android",
		Ua:   "Mozilla/5.0 (Linux; U; Android-4.0.3; en-us; Galaxy Nexus Build/IML74K) AppleWebKit/535.7 (KHTML, like Gecko) CrMo/16.0.912.75 Mobile Safari/535.7",
		Expect: TestBrowserRes{
			Name:    "chrome",
			Version: "16.0.912.75",
			Major:   "16",
		},
	},
	{
		Desc: "Coc Coc Browser",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) coc_coc_browser/78.0.129 Chrome/72.0.3626.129 Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "coc coc",
			Version: "78.0.129",
			Major:   "78",
		},
	},
	{
		Desc: "Dillo",
		Ua:   "Dillo/2.2",
		Expect: TestBrowserRes{
			Name:    "dillo",
			Version: "2.2",
			Major:   "2",
		},
	},
	{
		Desc: "Dolphin",
		Ua:   "Mozilla/5.0 (SCH-F859/F859DG12;U;NUCLEUS/2.1;Profile/MIDP-2.1 Configuration/CLDC-1.1;480*800;CTC/2.0) Dolfin/2.0",
		Expect: TestBrowserRes{
			Name:    "dolphin",
			Version: "2.0",
			Major:   "2",
		},
	},
	{
		Desc: "Doris",
		Ua:   "Doris/1.15 [en] (Symbian)",
		Expect: TestBrowserRes{
			Name:    "doris",
			Version: "1.15",
			Major:   "1",
		},
	},
	{
		Desc: "DuckDuckGo",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1.0) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.131 Mobile DuckDuckGo/5 Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "duckduckgo",
			Version: "5",
			Major:   "5",
		},
	},
	{
		Desc: "Epiphany",
		Ua:   "Mozilla/5.0 (X11; U; FreeBSD i386; en-US; rv:1.7) Gecko/20040628 Epiphany/1.2.6",
		Expect: TestBrowserRes{
			Name:    "epiphany",
			Version: "1.2.6",
			Major:   "1",
		},
	},
	{
		Desc: "Flow",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) EkiohFlow/5.7.4.30559 Flow/5.7.4 (like Gecko Firefox/53.0 rv:53.0)",
		Expect: TestBrowserRes{
			Name:    "flow",
			Version: "5.7.4",
			Major:   "5",
		},
	},
	{
		Desc: "Waterfox",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64; rv:55.0) Gecko/20100101 Firefox/55.2.2 Waterfox/55.2.2",
		Expect: TestBrowserRes{
			Name:    "waterfox",
			Version: "55.2.2",
			Major:   "55",
		},
	},
	{
		Desc: "PaleMoon",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64; rv:52.9) Gecko/20100101 Goanna/3.4 Firefox/52.9 PaleMoon/27.6.1",
		Expect: TestBrowserRes{
			Name:    "palemoon",
			Version: "27.6.1",
			Major:   "27",
		},
	},
	{
		Desc: "Basilisk",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64; rv:55.0) Gecko/20100101 Goanna/4.0 Firefox/55.0 Basilisk/20171113",
		Expect: TestBrowserRes{
			Name:    "basilisk",
			Version: "20171113",
			Major:   "20171113",
		},
	},
	{
		Desc: "Facebook in-App Browser for Android with version",
		Ua:   "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/43.0.2357.121 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/35.0.0.48.273;]",
		Expect: TestBrowserRes{
			Name:    "facebook",
			Version: "35.0.0.48.273",
			Major:   "35",
		},
	},
	{
		Desc: "Facebook in-App Browser for iOS with version",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Mobile/14E304 [FBAN/FBIOS;FBAV/91.0.0.41.73;FBBV/57050710;FBDV/iPhone8,1;FBMD/iPhone;FBSN/iOS;FBSV/10.3.1;FBSS/2;FBCR/Telekom.de;FBID/phone;FBLC/de_DE;FBOP/5;FBRV/0])",
		Expect: TestBrowserRes{
			Name:    "facebook",
			Version: "91.0.0.41.73",
			Major:   "91",
		},
	},
	{
		Desc: "Facebook in-App Browser for iOS without version",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [FBAN/FBIOS;FBDV/iPhone10,2;FBMD/iPhone;FBSN/iOS;FBSV/13.3.1;FBSS/3;FBID/phone;FBLC/en_US;FBOP/5;FBCR/]",
		Expect: TestBrowserRes{
			Name:    "facebook",
			Version: "",
			Major:   "",
		},
	},
	{
		Desc: "Instagram in-App Browser for iOS",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 14_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Instagram 142.0.0.22.109 (iPhone12,5; iOS 14_1; en_US; en-US; scale=3.00; 1242x2688; 214888322) NW/1",
		Expect: TestBrowserRes{
			Name:    "instagram",
			Version: "142.0.0.22.109",
			Major:   "142",
		},
	},
	{
		Desc: "Falkon",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Falkon/3.0.0 Chrome/61.0.3163.140 Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "falkon",
			Version: "3.0.0",
			Major:   "3",
		},
	},
	{
		Desc: "Firebird",
		Ua:   "Mozilla/5.0 (Windows; U; Win98; en-US; rv:1.5) Gecko/20031007 Firebird/0.7",
		Expect: TestBrowserRes{
			Name:    "firebird",
			Version: "0.7",
			Major:   "0",
		},
	},
	{
		Desc: "Firefox",
		Ua:   "Mozilla/5.0 (Windows NT 6.1; rv:15.0) Gecko/20120716 Firefox/15.0a2",
		Expect: TestBrowserRes{
			Name:    "firefox",
			Version: "15.0a2",
			Major:   "15",
		},
	},
	{
		Desc: "Firefox",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:100.0) Gecko/20100101 Firefox/100.0",
		Expect: TestBrowserRes{
			Name:    "firefox",
			Version: "100.0",
			Major:   "100",
		},
	},
	{
		Desc: "Firefox Reality",
		Ua:   "Mozilla/5.0 (Android 7.1.2; Mobile VR; rv:65.0) Gecko/65.0 Firefox/65.0",
		Expect: TestBrowserRes{
			Name:    "firefox reality",
			Version: "65.0",
			Major:   "65",
		},
	},
	{
		Desc: "Firefox-based browser",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64; rv:80.0) Gecko/20100101 Firefox/80.0 AppName/1.0",
		Expect: TestBrowserRes{
			Name:    "firefox",
			Version: "80.0",
			Major:   "80",
		},
	},
	{
		Desc: "Fennec",
		Ua:   "Mozilla/5.0 (X11; U; Linux armv61; en-US; rv:1.9.1b2pre) Gecko/20081015 Fennec/1.0a1",
		Expect: TestBrowserRes{
			Name:    "fennec",
			Version: "1.0a1",
			Major:   "1",
		},
	},
	{
		Desc: "Firefox for Maemo (Nokia N900)",
		Ua:   "Mozilla/5.0 (Maemo; Linux armv7l; rv:10.0.1) Gecko/20100101 Firefox/10.0.1 Fennec/10.0.1",
		Expect: TestBrowserRes{
			Name:    "fennec",
			Version: "10.0.1",
			Major:   "10",
		},
	},
	{
		Desc: "Firefox Focus",
		Ua:   "Mozilla/5.0 (Linux; Android 7.0) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Focus/6.1.1 Chrome/68.0.3440.91 Mobile Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "firefox focus",
			Version: "6.1.1",
			Major:   "6",
		},
	},
	{
		Desc: "Flock",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.3) Gecko/2008100716 Firefox/3.0.3 Flock/2.0",
		Expect: TestBrowserRes{
			Name:    "flock",
			Version: "2.0",
			Major:   "2",
		},
	},
	{
		Desc: "GoBrowser",
		Ua:   "Nokia5700XpressMusic/GoBrowser/1.6.91",
		Expect: TestBrowserRes{
			Name:    "gobrowser",
			Version: "1.6.91",
			Major:   "1",
		},
	},
	{
		Desc: "HuaweiBrowser",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; LYA-AL00；HMSCore/4.0.0 GMS/10.4 ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.64 HuaweiBrowser/10.0.3.102 Mobile Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "huawei browser",
			Version: "10.0.3.102",
			Major:   "10",
		},
	},
	{
		Desc: "HuaweiBrowser",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; LYA-AL00；HMSCore/4.0.0 ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.64 HuaweiBrowser/10.0.3.102 Mobile Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "huawei browser",
			Version: "10.0.3.102",
			Major:   "10",
		},
	},
	{
		Desc: "HuaweiBrowser",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; LYA-AL00；GMS/10.4 ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.64 HuaweiBrowser/10.0.3.102 Mobile Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "huawei browser",
			Version: "10.0.3.102",
			Major:   "10",
		},
	},
	{
		Desc: "HuaweiBrowser",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; LYA-AL00 ) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.64 HuaweiBrowser/10.0.3.102 Mobile Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "huawei browser",
			Version: "10.0.3.102",
			Major:   "10",
		},
	},
	{
		Desc: "IceApe",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.1.19) Gecko/20110817 Iceape/2.0.14",
		Expect: TestBrowserRes{
			Name:    "iceape",
			Version: "2.0.14",
			Major:   "2",
		},
	},
	{
		Desc: "IceCat",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.3) Gecko/2008092921 IceCat/3.0.3-g1",
		Expect: TestBrowserRes{
			Name:    "icecat",
			Version: "3.0.3-g1",
			Major:   "3",
		},
	},
	{
		Desc: "Iceweasel",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; de; rv:1.9.0.16) Gecko/2009121610 Iceweasel/3.0.6 (Debian-3.0.6-3)",
		Expect: TestBrowserRes{
			Name:    "iceweasel",
			Version: "3.0.6",
			Major:   "3",
		},
	},
	{
		Desc: "iCab",
		Ua:   "iCab/2.9.5 (Macintosh; U; PPC; Mac OS X)",
		Expect: TestBrowserRes{
			Name:    "icab",
			Version: "2.9.5",
			Major:   "2",
		},
	},
	{
		Desc: "IEMobile",
		Ua:   "Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.11) 320x240; VZW; Motorola-Q9c; Windows Mobile 6.1 Standard",
		Expect: TestBrowserRes{
			Name:    "iemobile",
			Version: "7.11",
			Major:   "7",
		},
	},
	{
		Desc: "IE 11 with IE token",
		Ua:   "Mozilla/5.0 (IE 11.0; Windows NT 6.3; WOW64; Trident/7.0; rv:11.0) like Gecko",
		Expect: TestBrowserRes{
			Name:    "ie",
			Version: "11.0",
			Major:   "11",
		},
	},
	{
		Desc: "IE 11 without IE token",
		Ua:   "Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv 11.0) like Gecko",
		Expect: TestBrowserRes{
			Name:    "ie",
			Version: "11.0",
			Major:   "11",
		},
	},
	{
		Desc: "K-Meleon",
		Ua:   "Mozilla/5.0 (Windows; U; Win98; en-US; rv:1.5) Gecko/20031016 K-Meleon/0.8.2",
		Expect: TestBrowserRes{
			Name:    "k-meleon",
			Version: "0.8.2",
			Major:   "0",
		},
	},
	{
		Desc: "Kindle Browser",
		Ua:   "Mozilla/4.0 (compatible; Linux 2.6.22) NetFront/3.4 Kindle/2.5 (screen 600x800; rotate)",
		Expect: TestBrowserRes{
			Name:    "kindle",
			Version: "2.5",
			Major:   "2",
		},
	},
	{
		Desc: "Konqueror",
		Ua:   "Mozilla/5.0 (compatible; Konqueror/3.5; Linux; X11; x86_64) KHTML/3.5.6 (like Gecko) (Kubuntu)",
		Expect: TestBrowserRes{
			Name:    "konqueror",
			Version: "3.5",
			Major:   "3",
		},
	},
	{
		Desc: "Konqueror",
		Ua:   "Mozilla/5.0 (X11; Linux i686) AppleWebKit/534.34 (KHTML, like Gecko) konqueror/5.0.97 Safari/534.34",
		Expect: TestBrowserRes{
			Name:    "konqueror",
			Version: "5.0.97",
			Major:   "5",
		},
	},
	{
		Desc: "LINE on Android",
		Ua:   "Mozilla/5.0 (Linux; Android 5.0; ASUS_Z00AD Build/LRX21V; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/51.0.2704.81 Mobile Safari/537.36 Line/6.5.1/IAB",
		Expect: TestBrowserRes{
			Name:    "line",
			Version: "6.5.1",
			Major:   "6",
		},
	},
	{
		Desc: "LINE on iOS",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 11_2_6 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) Mobile/15D100 Safari Line/8.4.1",
		Expect: TestBrowserRes{
			Name:    "line",
			Version: "8.4.1",
			Major:   "8",
		},
	},
	{
		Desc: "Lunascape",
		Ua:   "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.9.1.2) Gecko/20090804 Firefox/3.5.2 Lunascape/5.1.4.5",
		Expect: TestBrowserRes{
			Name:    "lunascape",
			Version: "5.1.4.5",
			Major:   "5",
		},
	},
	{
		Desc: "Lynx",
		Ua:   "Lynx/2.8.5dev.16 libwww-FM/2.14 SSL-MM/1.4.1 OpenSSL/0.9.6b",
		Expect: TestBrowserRes{
			Name:    "lynx",
			Version: "2.8.5dev.16",
			Major:   "2",
		},
	},
	{
		Desc: "Maemo Browser",
		Ua:   "Mozilla/5.0 (X11; U; Linux armv7l; ru-RU; rv:1.9.2.3pre) Gecko/20100723 Firefox/3.5 Maemo Browser 1.7.4.8 RX-51 N900",
		Expect: TestBrowserRes{
			Name:    "maemo browser",
			Version: "1.7.4.8",
			Major:   "1",
		},
	},
	{
		Desc: "Maxthon",
		Ua:   "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; SV1; Maxthon; .NET CLR 1.1.4322)",
		Expect: TestBrowserRes{
			Name:    "maxthon",
			Version: "",
			Major:   "",
		},
	},
	{
		Desc: "Midori",
		Ua:   "Midori/0.2.2 (X11; Linux i686; U; en-us) WebKit/531.2+",
		Expect: TestBrowserRes{
			Name:    "midori",
			Version: "0.2.2",
			Major:   "0",
		},
	},
	{
		Desc: "Minimo",
		Ua:   "Mozilla/5.0 (X11; U; Linux armv6l; rv 1.8.1.5pre) Gecko/20070619 Minimo/0.020",
		Expect: TestBrowserRes{
			Name:    "minimo",
			Version: "0.020",
			Major:   "0",
		},
	},
	{
		Desc: "MIUI Browser on Xiaomi Hongmi WCDMA (HM2013023)",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.2.2; ru-ru; 2013023 Build/HM2013023) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30 XiaoMi/MiuiBrowser/1.0",
		Expect: TestBrowserRes{
			Name:    "miui browser",
			Version: "1.0",
			Major:   "1",
		},
	},
	{
		Desc: "Mobile Safari",
		Ua:   "Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_0 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8A293 Safari/6531.22.7",
		Expect: TestBrowserRes{
			Name:    "mobile safari",
			Version: "4.0.5",
			Major:   "4",
		},
	},
	{
		Desc: "Mosaic",
		Ua:   "NCSA_Mosaic/2.6 (X11; SunOS 4.1.3 sun4m)",
		Expect: TestBrowserRes{
			Name:    "mosaic",
			Version: "2.6",
			Major:   "2",
		},
	},
	{
		Desc: "Mozilla",
		Ua:   "Mozilla/5.0 (X11; U; SunOS sun4u; en-US; rv:1.7) Gecko/20070606",
		Expect: TestBrowserRes{
			Name:    "mozilla",
			Version: "5.0",
			Major:   "5",
		},
	},
	{
		Desc: "MSIE",
		Ua:   "Mozilla/4.0 (compatible; MSIE 5.0b1; Mac_PowerPC)",
		Expect: TestBrowserRes{
			Name:    "ie",
			Version: "5.0b1",
			Major:   "5",
		},
	},
	{
		Desc: "NetFront",
		Ua:   "Mozilla/4.0 (PDA; Windows CE/1.0.1) NetFront/3.0",
		Expect: TestBrowserRes{
			Name:    "netfront",
			Version: "3.0",
			Major:   "3",
		},
	},
	{
		Desc: "Netscape on Windows ME",
		Ua:   "Mozilla/5.0 (Windows; U; Win 9x 4.90; en-US; rv:1.8.1.8pre) Gecko/20071015 Firefox/2.0.0.7 Navigator/9.0",
		Expect: TestBrowserRes{
			Name:    "netscape",
			Version: "9.0",
			Major:   "9",
		},
	},
	{
		Desc: "Netscape on Windows 2000",
		Ua:   "Mozilla/5.0 (Windows; U; Windows NT 5.0; en-US; rv:1.7.5) Gecko/20050519 Netscape/8.0.1",
		Expect: TestBrowserRes{
			Name:    "netscape",
			Version: "8.0.1",
			Major:   "8",
		},
	},
	{
		Desc: "Netscape 6",
		Ua:   "Mozilla/5.0 (Windows; U; Win95; de-DE; rv:0.9.2) Gecko/20010726 Netscape6/6.1",
		Expect: TestBrowserRes{
			Name:    "netscape",
			Version: "6.1",
			Major:   "6",
		},
	},
	{
		Desc: "Nokia Browser",
		Ua:   "Mozilla/5.0 (Symbian/3; Series60/5.2 NokiaN8-00/025.007; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.37 Mobile Safari/533.4 3gpp-gba",
		Expect: TestBrowserRes{
			Name:    "nokiabrowser",
			Version: "7.3.1.37",
			Major:   "7",
		},
	},
	{
		Desc: "Obigo",
		Ua:   "LG-GS290/V100 Obigo/WAP2.0 Profile/MIDP-2.1 Configuration/CLDC-1.1",
		Expect: TestBrowserRes{
			Name:    "obigo",
			Version: "wap2.0",
			Major:   "2",
		},
	},
	{
		Desc: "Obigo",
		Ua:   "LG/KU990i/v10a Browser/Obigo-Q05A/3.6 MMS/LG-MMS-V1.0/1.2 Java/ASVM/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1",
		Expect: TestBrowserRes{
			Name:    "obigo",
			Version: "q05a",
			Major:   "05",
		},
	},
	{
		Desc: "Oculus Browser",
		Ua:   "Mozilla/5.0 (Linux; Android 7.0; SM-G920I Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) OculusBrowser/3.4.9 SamsungBrowser/4.0 Chrome/57.0.2987.146 Mobile VR Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "oculus browser",
			Version: "3.4.9",
			Major:   "3",
		},
	},
	{
		Desc: "Oculus Browser",
		Ua:   "Mozilla/5.0 (Linux; Android 10; Quest 2) AppleWebKit/537.36 (KHTML, like Gecko) OculusBrowser/15.0.0.0.22.280317669 SamsungBrowser/4.0 Chrome/89.0.4389.90 VR Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "oculus browser",
			Version: "15.0.0.0.22.280317669",
			Major:   "15",
		},
	},
	{
		Desc: "OmniWeb",
		Ua:   "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; en-US) AppleWebKit/85 (KHTML, like Gecko) OmniWeb/v558.48",
		Expect: TestBrowserRes{
			Name:    "omniweb",
			Version: "558.48",
			Major:   "558",
		},
	},
	{
		Desc: "Opera > 9.80",
		Ua:   "Opera/9.80 (X11; Linux x86_64; U; Linux Mint; en) Presto/2.2.15 Version/10.10",
		Expect: TestBrowserRes{
			Name:    "opera",
			Version: "10.10",
			Major:   "10",
		},
	},
	{
		Desc: "Opera < 9.80 on Windows",
		Ua:   "Mozilla/4.0 (compatible; MSIE 5.0; Windows 95) Opera 6.01 [en]",
		Expect: TestBrowserRes{
			Name:    "opera",
			Version: "6.01",
			Major:   "6",
		},
	},
	{
		Desc: "Opera < 9.80 on OSX",
		Ua:   "Opera/8.5 (Macintosh; PPC Mac OS X; U; en)",
		Expect: TestBrowserRes{
			Name:    "opera",
			Version: "8.5",
			Major:   "8",
		},
	},
	{
		Desc: "Opera Mobile",
		Ua:   "Opera/9.80 (Android 2.3.5; Linux; Opera Mobi/ADR-1111101157; U; de) Presto/2.9.201 Version/11.50",
		Expect: TestBrowserRes{
			Name:    "opera mobi",
			Version: "11.50",
			Major:   "11",
		},
	},
	{
		Desc: "Opera Webkit",
		Ua:   "Mozilla/5.0 AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.123 Mobile Safari/537.22 OPR/14.0.1025.52315",
		Expect: TestBrowserRes{
			Name:    "opera",
			Version: "14.0.1025.52315",
			Major:   "14",
		},
	},
	{
		Desc: "Opera Mini",
		Ua:   "Opera/9.80 (J2ME/MIDP; Opera Mini/5.1.21214/19.916; U; en) Presto/2.5.25",
		Expect: TestBrowserRes{
			Name:    "opera mini",
			Version: "5.1.21214",
			Major:   "5",
		},
	},
	{
		Desc: "Opera Mini 8 above on iPhone",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 9_2 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) OPiOS/12.1.1.98980 Mobile/13C75 Safari/9537.53",
		Expect: TestBrowserRes{
			Name:    "opera mini",
			Version: "12.1.1.98980",
			Major:   "12",
		},
	},
	{
		Desc: "Opera Tablet",
		Ua:   "Opera/9.80 (Windows NT 6.1; Opera Tablet/15165; U; en) Presto/2.8.149 Version/11.1",
		Expect: TestBrowserRes{
			Name:    "opera tablet",
			Version: "11.1",
			Major:   "11",
		},
	},
	{
		Desc: "Opera Coast",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 9_3_2 like Mac OS X; en) AppleWebKit/601.1.46 (KHTML, like Gecko) Coast/5.04.110603 Mobile/13F69 Safari/7534.48.3",
		Expect: TestBrowserRes{
			Name:    "opera coast",
			Version: "5.04.110603",
			Major:   "5",
		},
	},
	{
		Desc: "Opera Touch",
		Ua:   "Mozilla/5.0 (Linux; Android 7.0; Lenovo P2a42 Build/NRD90N) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/68.0.3440.91 Mobile Safari/537.36 OPT/1.10.33",
		Expect: TestBrowserRes{
			Name:    "opera touch",
			Version: "1.10.33",
			Major:   "1",
		},
	},
	{
		Desc: "PhantomJS",
		Ua:   "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.34 (KHTML, like Gecko) PhantomJS/1.9.2 Safari/534.34",
		Expect: TestBrowserRes{
			Name:    "phantomjs",
			Version: "1.9.2",
			Major:   "1",
		},
	},
	{
		Desc: "Phoenix",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.2b) Gecko/20021029 Phoenix/0.4",
		Expect: TestBrowserRes{
			Name:    "phoenix",
			Version: "0.4",
			Major:   "0",
		},
	},
	{
		Desc: "Polaris",
		Ua:   "LG-LX600 Polaris/6.0 MMP/2.0 Profile/MIDP-2.1 Configuration/CLDC-1.1",
		Expect: TestBrowserRes{
			Name:    "polaris",
			Version: "6.0",
			Major:   "6",
		},
	},
	{
		Desc: "QQ",
		Ua:   "Mozilla/5.0 (Linux; U; Android 4.4.4; zh-cn; OPPO R7s Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko)Version/4.0 Chrome/37.0.0.0 MQQBrowser/7.1 Mobile Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "qqbrowser",
			Version: "7.1",
			Major:   "7",
		},
	},
	{
		Desc: "QupZilla",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/538.1 (KHTML, like Gecko) QupZilla/1.8.9 Safari/538.1",
		Expect: TestBrowserRes{
			Name:    "qupzilla",
			Version: "1.8.9",
			Major:   "1",
		},
	},
	{
		Desc: "RockMelt",
		Ua:   "Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.7 (KHTML, like Gecko) RockMelt/0.8.36.78 Chrome/7.0.517.44 Safari/534.7",
		Expect: TestBrowserRes{
			Name:    "rockmelt",
			Version: "0.8.36.78",
			Major:   "0",
		},
	},
	{
		Desc: "Safari",
		Ua:   "Mozilla/5.0 (Windows; U; Windows NT 5.2; en-US) AppleWebKit/533.17.8 (KHTML, like Gecko) Version/5.0.1 Safari/533.17.8",
		Expect: TestBrowserRes{
			Name:    "safari",
			Version: "5.0.1",
			Major:   "5",
		},
	},
	{
		Desc: "Safari < 3.0",
		Ua:   "Mozilla/5.0 (Macintosh; U; PPC Mac OS X; sv-se) AppleWebKit/419 (KHTML, like Gecko) Safari/419.3",
		Expect: TestBrowserRes{
			Name:    "safari",
			Version: "2.0.4",
			Major:   "2",
		},
	},
	{
		Desc: "Samsung Browser",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; SAMSUNG-SM-G925A Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/4.0 Chrome/44.0.2403.133 Mobile Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "samsung browser",
			Version: "4.0",
			Major:   "4",
		},
	},
	{
		Desc: "SeaMonkey",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.1b4pre) Gecko/20090405 SeaMonkey/2.0b1pre",
		Expect: TestBrowserRes{
			Name:    "seamonkey",
			Version: "2.0b1pre",
			Major:   "2",
		},
	},
	{
		Desc: "SeaMonkey on Mac",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.5; rv:10.0.1) Gecko/20100101 Firefox/10.0.1 SeaMonkey/2.7.1",
		Expect: TestBrowserRes{
			Name:    "seamonkey",
			Version: "2.7.1",
			Major:   "2",
		},
	},
	{
		Desc: "Silk Browser",
		Ua:   "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_3; en-us; Silk/1.1.0-84)",
		Expect: TestBrowserRes{
			Name:    "silk",
			Version: "1.1.0-84",
			Major:   "1",
		},
	},
	{
		Desc: "Skyfire",
		Ua:   "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_7; en-us) AppleWebKit/530.17 (KHTML, like Gecko) Version/4.0 Safari/530.17 Skyfire/2.0",
		Expect: TestBrowserRes{
			Name:    "skyfire",
			Version: "2.0",
			Major:   "2",
		},
	},
	{
		Desc: "SlimBrowser",
		Ua:   "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SlimBrowser)",
		Expect: TestBrowserRes{
			Name:    "slim",
			Version: "",
			Major:   "",
		},
	},
	{
		Desc: "Swiftfox",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.1) Gecko/20061024 Firefox/2.0 (Swiftfox)",
		Expect: TestBrowserRes{
			Name:    "swiftfox",
			Version: "",
			Major:   "",
		},
	},
	{
		Desc: "Tesla",
		Ua:   "Mozilla/5.0 (X11; GNU/Linux) AppleWebKit/601.1 (KHTML, like Gecko) Tesla QtCarBrowser Safari/601.1",
		Expect: TestBrowserRes{
			Name:    "tesla",
			Version: "",
			Major:   "",
		},
	},
	{
		Desc: "Tesla",
		Ua:   "Mozilla/5.0 (X11; GNU/Linux) AppleWebKit/537.36 (KHTML, like Gecko) Chromium/79.0.3945.130 Chrome/79.0.3945.130 Safari/537.36 Tesla/2020.16.2.1-e99c70fff409",
		Expect: TestBrowserRes{
			Name:    "tesla",
			Version: "2020.16.2.1-e99c70fff409",
			Major:   "2020",
		},
	},
	{
		Desc: "Tizen Browser",
		Ua:   "Mozilla/5.0 (Linux; U; Tizen/1.0 like Android; en-us; AppleWebKit/534.46 (KHTML, like Gecko) Tizen Browser/1.0 Mobile",
		Expect: TestBrowserRes{
			Name:    "tizen browser",
			Version: "1.0",
			Major:   "1",
		},
	},
	{
		Desc: "UC Browser",
		Ua:   "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.99 UBrowser/5.6.12860.7 Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "uc browser",
			Version: "5.6.12860.7",
			Major:   "5",
		},
	},
	{
		Desc: "UC Browser",
		Ua:   "Mozilla/5.0 (Linux; U; Android 6.0.1; en-US; Lenovo P2a42 Build/MMB29M) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 UCBrowser/11.2.0.915 U3/0.8.0 Mobile Safari/534.30",
		Expect: TestBrowserRes{
			Name:    "uc browser",
			Version: "11.2.0.915",
			Major:   "11",
		},
	},
	{
		Desc: "UC Browser on Samsung",
		Ua:   "Mozilla/5.0 (Java; U; Pt-br; samsung-gt-s5620) UCBrowser8.2.1.144/69/352/UCWEB Mobile UNTRUSTED/1.0",
		Expect: TestBrowserRes{
			Name:    "uc browser",
			Version: "8.2.1.144",
			Major:   "8",
		},
	},
	{
		Desc: "UC Browser on Nokia",
		Ua:   "Mozilla/5.0 (S60V3; U; en-in; NokiaN73)/UC Browser8.4.0.159/28/351/UCWEB Mobile",
		Expect: TestBrowserRes{
			Name:    "uc browser",
			Version: "8.4.0.159",
			Major:   "8",
		},
	},
	{
		Desc: "UC Browser J2ME",
		Ua:   "UCWEB/2.0 (MIDP-2.0; U; zh-CN; HTC EVO 3D X515m) U2/1.0.0 UCBrowser/10.4.0.558 U2/1.0.0 Mobile",
		Expect: TestBrowserRes{
			Name:    "uc browser",
			Version: "10.4.0.558",
			Major:   "10",
		},
	},
	{
		Desc: "UC Browser J2ME 2",
		Ua:   "JUC (Linux; U; 2.3.5; zh-cn; GT-I9100; 480*800) UCWEB7.9.0.94/139/800",
		Expect: TestBrowserRes{
			Name:    "uc browser",
			Version: "7.9.0.94",
			Major:   "7",
		},
	},
	{
		Desc: "UPBrowser",
		Ua:   "BenQ-CF61/1.00/WAP2.0/MIDP2.0/CLDC1.0 UP.Browser/6.3.0.4.c.1.102 (GUI) MMP/2.0",
		Expect: TestBrowserRes{
			Name:    "up.browser",
			Version: "6.3.0.4.c.1.102",
			Major:   "6",
		},
	},
	{
		Desc: "WeChat on iOS",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 8_4_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Mobile/12H321 MicroMessenger/6.3.6 NetType/WIFI Language/zh_CN",
		Expect: TestBrowserRes{
			Name:    "wechat",
			Version: "6.3.6",
			Major:   "6",
		},
	},
	{
		Desc: "WeChat on Android",
		Ua:   "Mozilla/5.0 (Linux; U; Android 5.1; zh-cn; Lenovo K50-t5 Build/LMY47D) AppleWebKit/533.1 (KHTML, like Gecko)Version/4.0 MQQBrowser/5.4 TBS/025478 Mobile Safari/533.1 MicroMessenger/6.3.5.50_r1573191.640 NetType/WIFI Language/zh_CN",
		Expect: TestBrowserRes{
			Name:    "wechat",
			Version: "6.3.5.50_r1573191.640",
			Major:   "6",
		},
	},
	{
		Desc: "WeiBo on Android",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16A366 Weibo (iPhone8,2__weibo__8.9.3__iphone__os12.0)",
		Expect: TestBrowserRes{
			Name:    "weibo",
			Version: "8.9.3",
			Major:   "8",
		},
	},
	{
		Desc: "Vivaldi",
		Ua:   "Mozilla/5.0 (Windows NT 6.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.89 Vivaldi/1.0.83.38 Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "vivaldi",
			Version: "1.0.83.38",
			Major:   "1",
		},
	},
	{
		Desc: "Yandex",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_2) AppleWebKit/536.5 (KHTML, like Gecko) YaBrowser/1.0.1084.5402 Chrome/19.0.1084.5402 Safari/536.5",
		Expect: TestBrowserRes{
			Name:    "yandex",
			Version: "1.0.1084.5402",
			Major:   "1",
		},
	},
	{
		Desc: "Puffin",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0.1; Lenovo P2a42 Build/MMB29M; en-us) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Mobile Safari/537.36 Puffin/6.0.8.15804AP",
		Expect: TestBrowserRes{
			Name:    "puffin",
			Version: "6.0.8.15804ap",
			Major:   "6",
		},
	},
	{
		Desc: "Microsoft Edge 0.1",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36 Edge/12.0",
		Expect: TestBrowserRes{
			Name:    "edge",
			Version: "12.0",
			Major:   "12",
		},
	},
	{
		Desc: "Microsoft Edge 42",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/17.17134",
		Expect: TestBrowserRes{
			Name:    "edge",
			Version: "17.17134",
			Major:   "17",
		},
	},
	{
		Desc: "Microsoft Edge 44",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/18.17763",
		Expect: TestBrowserRes{
			Name:    "edge",
			Version: "18.17763",
			Major:   "18",
		},
	},
	{
		Desc: "Microsoft Edge 100",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.1108.55 Safari/537.36 Edg/100.0.1108.55",
		Expect: TestBrowserRes{
			Name:    "edge",
			Version: "100.0.1108.55",
			Major:   "100",
		},
	},
	{
		Desc: "Microsoft Edge on iOS",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 11_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.0 EdgiOS/42.1.1.0 Mobile/15F79 Safari/605.1.15",
		Expect: TestBrowserRes{
			Name:    "edge",
			Version: "42.1.1.0",
			Major:   "42",
		},
	},
	{
		Desc: "Microsoft Edge on Android",
		Ua:   "Mozilla/5.0 (Linux; Android 8.0.0; G8441 Build/47.1.A.12.270) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.123 Mobile Safari/537.36 EdgA/42.0.0.2529",
		Expect: TestBrowserRes{
			Name:    "edge",
			Version: "42.0.0.2529",
			Major:   "42",
		},
	},
	{
		Desc: "Microsoft Edge Chromium",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.48 Safari/537.36 Edg/74.1.96.24",
		Expect: TestBrowserRes{
			Name:    "edge",
			Version: "74.1.96.24",
			Major:   "74",
		},
	},
	{
		Desc: "Iridium",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Iridium/43.8 Safari/537.36 Chrome/43.0.2357.132",
		Expect: TestBrowserRes{
			Name:    "iridium",
			Version: "43.8",
			Major:   "43",
		},
	},
	{
		Desc: "Firefox iOS",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) FxiOS/1.1 Mobile/13B143 Safari/601.1.46",
		Expect: TestBrowserRes{
			Name:    "firefox",
			Version: "1.1",
			Major:   "1",
		},
	},
	{
		Desc: "Firefox iOS using iPad",
		Ua:   "Mozilla/5.0 (iPad; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) FxiOS/1.0 Mobile/12F69 Safari/600.1.4",
		Expect: TestBrowserRes{
			Name:    "firefox",
			Version: "1.0",
			Major:   "1",
		},
	},
	{
		Desc: "QQ on iOS",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 10_0_2 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Mobile/14A456 QQ/6.5.3.410 V1_IPH_SQ_6.5.3_1_APP_A Pixel/1080 Core/UIWebView NetType/WIFI Mem/26",
		Expect: TestBrowserRes{
			Name:    "qq",
			Version: "6.5.3.410",
			Major:   "6",
		},
	},
	{
		Desc: "QQ on Android",
		Ua:   "Mozilla/5.0 (Linux; Android 6.0; PRO 6 Build/MRA58K) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/37.0.0.0 Mobile MQQBrowser/6.8 TBS/036824 Safari/537.36 V1_AND_SQ_6.5.8_422_YYB_D PA QQ/6.5.8.2910 NetType/WIFI WebP/0.3.0 Pixel/1080",
		Expect: TestBrowserRes{
			Name:    "qq",
			Version: "6.5.8.2910",
			Major:   "6",
		},
	},
	{
		Desc: "baidu app on iOS",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16C101 main%2F1.0 baiduboxapp/11.12.0.18 (Baidu; P2 12.1.2)",
		Expect: TestBrowserRes{
			Name:    "baiduboxapp",
			Version: "11.12.0.18",
			Major:   "11",
		},
	},
	{
		Desc: "baidu app on Android",
		Ua:   "Mozilla/5.0 (Linux; Android 8.1.0; BKK-AL10 Build/HONORBKK-AL10; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/63.0.3239.83 Mobile Safari/537.36 T7/11.11 baiduboxapp/11.11.0.0 (Baidu; P1 8.1.0)",
		Expect: TestBrowserRes{
			Name:    "baiduboxapp",
			Version: "11.11.0.0",
			Major:   "11",
		},
	},
	{
		Desc: "WeChat Desktop for Windows Built-in Browser",
		Ua:   "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 MicroMessenger/6.5.2.501 NetType/WIFI WindowsWechat QBCore/3.43.901.400 QQBrowser/9.0.2524.400",
		Expect: TestBrowserRes{
			Name:    "wechat(win) desktop",
			Version: "3.43.901.400",
			Major:   "3",
		},
	},
	{
		Desc: "WeChat Desktop for Windows Built-in Browser major version in 4",
		Ua:   "mozilla/5.0 (windows nt 10.0; wow64) applewebkit/537.36 (khtml, like gecko) chrome/53.0.2785.116 safari/537.36 qbcore/4.0.1301.400 qqbrowser/9.0.2524.400 mozilla/5.0 (windows nt 6.1; wow64) applewebkit/537.36 (khtml, like gecko) chrome/81.0.4044.138 safari/537.36 nettype/wifi micromessenger/7.0.20.1781(0x6700143b) windowswechat",
		Expect: TestBrowserRes{
			Name:    "wechat(win) desktop",
			Version: "4.0.1301.400",
			Major:   "4",
		},
	},
	{
		Desc: "Supposed not to be detected as WeChat",
		Ua:   "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.124 Safari/537.36 qblink wegame.exe WeGame/5.1.1.11100 QBCore/3.70.107.400 QQBrowser/9.0.2524.400",
		Expect: TestBrowserRes{
			Name:    "qqbrowser",
			Version: "9.0.2524.400",
			Major:   "9",
		},
	},
	{
		Desc: "GSA on iOS",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) GSA/30.1.161623614 Mobile/14F89 Safari/602.1",
		Expect: TestBrowserRes{
			Name:    "gsa",
			Version: "30.1.161623614",
			Major:   "30",
		},
	},
	{
		Desc: "Sogou Browser",
		Ua:   "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.221 Safari/537.36 SE 2.X MetaSr 1.0",
		Expect: TestBrowserRes{
			Name: "metasr",
		},
	},
	{
		Desc: "LieBao Browser",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.154 Safari/537.36 LBBROWSER",
		Expect: TestBrowserRes{
			Name: "lbbrowser",
		},
	},
	{
		Desc: "BaiDu Browser",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.106 BIDUBrowser/8.7 Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "bidubrowser",
			Version: "8.7",
			Major:   "8",
		},
	},
	{
		Desc: "2345 Browser",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.90 Safari/537.36 2345Explorer/9.2.1.17116",
		Expect: TestBrowserRes{
			Name:    "2345explorer",
			Version: "9.2.1.17116",
			Major:   "9",
		},
	},
	{
		Desc: "QQBrowserLite",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_1) AppleWebKit/602.2.14 (KHTML, like Gecko) Version/10.0.1 Safari/602.2.14 QQBrowserLite/1.1.0",
		Expect: TestBrowserRes{
			Name:    "qqbrowserlite",
			Version: "1.1.0",
			Major:   "1",
		},
	},
	{
		Desc: "Brave Browser",
		Ua:   "Brave/4.5.16 CFNetwork/893.13.1 Darwin/17.3.0 (x86_64)",
		Expect: TestBrowserRes{
			Name:    "brave",
			Version: "4.5.16",
			Major:   "4",
		},
	},
	{
		Desc: "Whale Browser",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.146 Whale/2.6.90.14 Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "whale",
			Version: "2.6.90.14",
			Major:   "2",
		},
	},
	{
		Desc: "Electron",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Atom/1.41.0 Chrome/69.0.3497.128 Electron/4.2.7 Safari/537.36",
		Expect: TestBrowserRes{
			Name:    "electron",
			Version: "4.2.7",
			Major:   "4",
		},
	},
	{
		Desc: "IE11 on Windows 7 (ua length >255)",
		Ua:   "Mozilla/5.0 (Windows NT 6.1; WOW64; APCPMS=^N201205020840572565478A37A6F9C41BD33F_9975^; Trident/7.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.3; .NET4.0C; .NET4.0E; MARKANYEPS#25118; Zoom 3.6.0; rv:11.0) like Gecko",
		Expect: TestBrowserRes{
			Name:    "ie",
			Version: "11.0",
			Major:   "11",
		},
	},
	{
		Desc: "LinkedIn",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 [LinkedInApp]",
		Expect: TestBrowserRes{
			Name: "linkedin",
		},
	},
	{
		Desc: "Safari including comma in minor version number",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.6,2 Safari/605.1.15",
		Expect: TestBrowserRes{
			Name:    "safari",
			Version: "15.6,2",
			Major:   "15",
		},
	},
	{
		Desc: "Mobile Safari including comma in minor version number",
		Ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 15_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.6,2 Mobile/15E148 Safari/604.1",
		Expect: TestBrowserRes{
			Name:    "mobile safari",
			Version: "15.6,2",
			Major:   "15",
		},
	},
}
