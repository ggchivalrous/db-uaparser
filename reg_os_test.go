package dbuaparser_test

import (
	"fmt"
	"testing"

	dbuaparser "github.com/ggchivalrous/db-uaparser"
)

func TestRegOs(t *testing.T) {
	osLen := len(osList)
	errCount := 0
	for _, v := range osList {
		res := dbuaparser.GetOs(v.Ua)
		state := res.Os == v.Expect.Name && res.OsVer == v.Expect.Version
		if !state {
			errCount++
			fmt.Println("匹配: ", v.Desc, " 失败，结果：name: ", res.Os, " version: ", res.OsVer)
		}
	}

	fmt.Println("匹配总数: ", osLen, " 成功总数:", osLen-errCount, "失败总数: ", errCount)

	if errCount > 0 {
		t.Error("匹配失败数量: ", errCount)
	}
}

type TestOs struct {
	Desc   string
	Ua     string
	Expect TestOsRes
}

type TestOsRes struct {
	Name    string
	Version string
}

var osList = []TestOs{
	{
		Desc: "Windows 95",
		Ua:   "Mozilla/1.22 (compatible; MSIE 2.0; Windows 95)",
		Expect: TestOsRes{
			Name:    "windows",
			Version: "95",
		},
	},
	{
		Desc: "Windows 98",
		Ua:   "Mozilla/4.0 (compatible; MSIE 4.01; Windows 98)",
		Expect: TestOsRes{
			Name:    "windows",
			Version: "98",
		},
	},
	{
		Desc: "Windows ME",
		Ua:   "Mozilla/5.0 (Windows; U; Win 9x 4.90) Gecko/20020502 CS 2000 7.0/7.0",
		Expect: TestOsRes{
			Name:    "windows",
			Version: "me",
		},
	},
	{
		Desc: "Windows 2000",
		Ua:   "Mozilla/3.0 (compatible; MSIE 3.0; Windows NT 5.0)",
		Expect: TestOsRes{
			Name:    "windows",
			Version: "2000",
		},
	},
	{
		Desc: "Windows XP",
		Ua:   "Mozilla/5.0 (Windows; U; MSIE 7.0; Windows NT 5.2)",
		Expect: TestOsRes{
			Name:    "windows",
			Version: "xp",
		},
	},
	{
		Desc: "Windows Vista",
		Ua:   "Mozilla/5.0 (compatible; MSIE 7.0; Windows NT 6.0; fr-FR)",
		Expect: TestOsRes{
			Name:    "windows",
			Version: "vista",
		},
	},
	{
		Desc: "Windows 7",
		Ua:   "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)",
		Expect: TestOsRes{
			Name:    "windows",
			Version: "7",
		},
	},
	{
		Desc: "Windows 8",
		Ua:   "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.2; Win64; x64; Trident/6.0; .NET4.0E; .NET4.0C)",
		Expect: TestOsRes{
			Name:    "windows",
			Version: "8",
		},
	},
	{
		Desc: "Windows 10",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36 Edge/12.0",
		Expect: TestOsRes{
			Name:    "windows",
			Version: "10",
		},
	},
	{
		Desc: "Windows RT",
		Ua:   "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0)",
		Expect: TestOsRes{
			Name:    "windows",
			Version: "rt",
		},
	},
	{
		Desc: "Windows CE",
		Ua:   "Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.11)",
		Expect: TestOsRes{
			Name:    "windows",
			Version: "ce",
		},
	},
	{
		Desc: "Windows Mobile",
		Ua:   "Mozilla/5.0 (ZTE-E_N72/N72V1.0.0B02;U;Windows Mobile/6.1;Profile/MIDP-2.0 Configuration/CLDC-1.1;320*240;CTC/2.0) IE/6.0 (compatible; MSIE 4.01; Windows CE; PPC)/UC Browser7.7.1.88",
		Expect: TestOsRes{
			Name:    "windows mobile",
			Version: "6.1",
		},
	},
	{
		Desc: "Windows Mobile",
		Ua:   "Opera/9.80 (Windows Mobile; WCE; Opera Mobi/WMD-50433; U; en) Presto/2.4.13 Version/10.00",
		Expect: TestOsRes{
			Name:    "windows mobile",
			Version: "",
		},
	},
	{
		Desc: "Windows Phone",
		Ua:   "Opera/9.80 (Windows Phone; Opera Mini/7.6.8/35.7518; U; ru) Presto/2.8.119 Version/11.10",
		Expect: TestOsRes{
			Name:    "windows phone",
			Version: "",
		},
	},
	{
		Desc: "Windows Phone OS",
		Ua:   "Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.0; Trident/3.1; IEMobile/7.0; DELL; Venue Pro)",
		Expect: TestOsRes{
			Name:    "windows phone os",
			Version: "7.0",
		},
	},
	{
		Desc: "Windows Phone 8",
		Ua:   "Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; HTC; Windows Phone 8X by HTC)",
		Expect: TestOsRes{
			Name:    "windows phone",
			Version: "8.0",
		},
	},
	{
		Desc: "Windows NT on x86 or aarch64 CPU using Firefox",
		Ua:   "Mozilla/5.0 (Windows NT x.y; rv:10.0) Gecko/20100101 Firefox/10.0",
		Expect: TestOsRes{
			Name:    "windows",
			Version: "nt x",
		},
	},
	{
		Desc: "Windows NT on x64 CPU using Firefox",
		Ua:   "Mozilla/5.0 (Windows NT x.y; Win64; x64; rv:10.0) Gecko/20100101 Firefox/10.0",
		Expect: TestOsRes{
			Name:    "windows",
			Version: "nt x",
		},
	},
	{
		Desc: "BlackBerry",
		Ua:   "BlackBerry9300/5.0.0.912 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/378",
		Expect: TestOsRes{
			Name:    "blackberry",
			Version: "5.0.0.912",
		},
	},
	{
		Desc: "BlackBerry 10",
		Ua:   "Mozilla/5.0 (BB10; Touch) AppleWebKit/537.3+ (KHTML, like Gecko) Version/10.0.9.386 Mobile Safari/537.3+",
		Expect: TestOsRes{
			Name:    "blackberry",
			Version: "10",
		},
	},
	{
		Desc: "Tizen",
		Ua:   "Mozilla/5.0 (SMART-TV; Linux; Tizen 2.3) AppleWebkit/538.1 (KHTML, like Gecko) SamsungBrowser/1.0 TV Safari/538.1",
		Expect: TestOsRes{
			Name:    "tizen",
			Version: "2.3",
		},
	},
	{
		Desc: "Tizen",
		Ua:   "Mozilla/5.0 (Linux; Tizen 2.3; SAMSUNG SM-Z130H) AppleWebKit/537.3 (KHTML, like Gecko) Version/2.3 Mobile Safari/537.3",
		Expect: TestOsRes{
			Name:    "tizen",
			Version: "2.3",
		},
	},
	{
		Desc: "Android",
		Ua:   "Mozilla/5.0 (Linux; U; Android 2.2.2; en-us; VM670 Build/FRG83G) AppleWebKit/533.1 (KHTML, like Gecko)",
		Expect: TestOsRes{
			Name:    "android",
			Version: "2.2.2",
		},
	},
	{
		Desc: "HarmonyOS",
		Ua:   "Mozilla/5.0 (Linux; Android 10; HarmonyOS; YAL-AL10; HMSCore 6.3.0.327; GMSCore 21.48.15) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.105 HuaweiBrowser/12.0.3.310 Mobile Safari/537.36",
		Expect: TestOsRes{
			Name:    "harmonyos",
			Version: "10",
		},
	},
	{
		Desc: "Sailfish",
		Ua:   "Mozilla/5.0 (Linux; U; Sailfish 3.0; Mobile; rv:45.0) Gecko/45.0 Firefox/45.0 SailfishBrowser/1.0",
		Expect: TestOsRes{
			Name:    "sailfish",
			Version: "3.0",
		},
	},
	{
		Desc: "WebOS",
		Ua:   "Mozilla/5.0 (hp-tablet; Linux; hpwOS/3.0.5; U; en-US) AppleWebKit/534.6 (KHTML, like Gecko) wOSBrowser/234.83 Safari/534.6 TouchPad/1.0",
		Expect: TestOsRes{
			Name:    "webos",
			Version: "3.0.5",
		},
	},
	{
		Desc: "WebOS",
		Ua:   "Mozilla/5.0 (webOS/1.4.5; U; en-US) AppleWebKit/532.2 (KHTML, like Gecko) Version/1.0 Safari/532.2 Pre/1.0",
		Expect: TestOsRes{
			Name:    "webos",
			Version: "1.4.5",
		},
	},
	{
		Desc: "WebOS TV 5.x",
		Ua:   "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36 WebAppManager",
		Expect: TestOsRes{
			Name:    "webos",
			Version: "tv",
		},
	},
	{
		Desc: "WebOS TV 4.x",
		Ua:   "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.34 Safari/537.36 WebAppManager",
		Expect: TestOsRes{
			Name:    "webos",
			Version: "tv",
		},
	},
	{
		Desc: "WebOS TV 3.x",
		Ua:   "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.36 (KHTML, like Gecko) QtWebEngine/5.2.1 Chrome/38.0.2125.122 Safari/537.36 WebAppManager",
		Expect: TestOsRes{
			Name:    "webos",
			Version: "tv",
		},
	},
	{
		Desc: "WebOS TV 2.x",
		Ua:   "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/538.2 (KHTML, like Gecko) Large Screen WebAppManager Safari/538.2",
		Expect: TestOsRes{
			Name:    "webos",
			Version: "tv",
		},
	},
	{
		Desc: "WebOS TV 1.x",
		Ua:   "Mozilla/5.0 (Web0S; Linux/SmartTV) AppleWebKit/537.41 (KHTML, like Gecko) Large Screen WebAppManager Safari/537.41",
		Expect: TestOsRes{
			Name:    "webos",
			Version: "tv",
		},
	},
	{
		Desc: "QNX",
		Ua:   "Mozilla/5.0 (Photon; U; QNX x86pc; en-US; rv:1.8.1.20) Gecko/20090127 BonEcho/2.0.0.20",
		Expect: TestOsRes{
			Name:    "qnx",
			Version: "x86pc",
		},
	},
	{
		Desc: "Bada",
		Ua:   "Mozilla/5.0 (SAMSUNG; SAMSUNG-GT-S5253/S5253DDKC1; U; Bada/1.0; en-us) AppleWebKit/533.1 (KHTML, like Gecko) Dolfin/2.0 Mobile WQVGA SMM-MMS/1.2.0 OPN-B",
		Expect: TestOsRes{
			Name:    "bada",
			Version: "1.0",
		},
	},
	{
		Desc: "RIM Tablet OS",
		Ua:   "Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML like Gecko) Version/7.2.1.0 Safari/536.2+",
		Expect: TestOsRes{
			Name:    "rim tablet os",
			Version: "2.1.0",
		},
	},
	{
		Desc: "Nokia N900 Linux mobile, on the Fennec browser",
		Ua:   "Mozilla/5.0 (Maemo; Linux armv7l; rv:10.0) Gecko/20100101 Firefox/10.0 Fennec/10.0",
		Expect: TestOsRes{
			Name:    "maemo",
			Version: "",
		},
	},
	{
		Desc: "MeeGo",
		Ua:   "Mozilla/5.0 (MeeGo; NokiaN9) AppleWebKit/534.13 (KHTML, like Gecko) NokiaBrowser/8.5.0 Mobile Safari/534.13",
		Expect: TestOsRes{
			Name:    "meego",
			Version: "",
		},
	},
	{
		Desc: "Symbian",
		Ua:   "Nokia5250/10.0.011 (SymbianOS/9.4; U; Series60/5.0 Mozilla/5.0; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/525 (KHTML, like Gecko) Safari/525 3gpp-gba",
		Expect: TestOsRes{
			Name:    "symbian",
			Version: "9.4",
		},
	},
	{
		Desc: "Symbian",
		Ua:   "Mozilla/5.0 (Symbian/3; Series60/5.2 NokiaC7-00/024.001; Profile/MIDP-2.1 Configuration/CLDC-1.1 ) AppleWebKit/533.4 (KHTML, like Gecko) NokiaBrowser/7.3.1.37 Mobile Safari/533.4 3gpp-gba",
		Expect: TestOsRes{
			Name:    "symbian",
			Version: "5.2",
		},
	},
	{
		Desc: "Series40",
		Ua:   "Mozilla/5.0 (Series40; Nokia2055/03.20; Profile/MIDP-2.1 Configuration/CLDC-1.1) Gecko/20100401 S40OviBrowser/2.2.0.0.34",
		Expect: TestOsRes{
			Name:    "series40",
			Version: "",
		},
	},
	{
		Desc: "Firefox OS",
		Ua:   "Mozilla/5.0 (Mobile; rv:14.0) Gecko/14.0 Firefox/14.0",
		Expect: TestOsRes{
			Name:    "firefox os",
			Version: "14.0",
		},
	},
	{
		Desc: "Firefox OS on Tablet",
		Ua:   "Mozilla/5.0 (Tablet; rv:26.0) Gecko/26.0 Firefox/26.0",
		Expect: TestOsRes{
			Name:    "firefox os",
			Version: "26.0",
		},
	},
	{
		Desc: "Firefox OS on TV",
		Ua:   "Mozilla/5.0 (TV; rv:44.0) Gecko/44.0 Firefox/44.0",
		Expect: TestOsRes{
			Name:    "firefox os",
			Version: "44.0",
		},
	},
	{
		Desc: "Google Chromecast",
		Ua:   "Mozilla/5.0 (X11; Linux aarch64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.81 Safari/537.36 CrKey/1.42.183786",
		Expect: TestOsRes{
			Name:    "chromecast",
			Version: "1.42.183786",
		},
	},
	{
		Desc: "Nintendo Switch",
		Ua:   "Mozilla/5.0 (Nintendo Switch; WifiWebAuthApplet) AppleWebKit/606.4 (KHTML, like Gecko) NF/6.0.1.15.4 NintendoBrowser/5.1.0.20393",
		Expect: TestOsRes{
			Name:    "nintendo",
			Version: "switch",
		},
	},
	{
		Desc: "PlayStation 4",
		Ua:   "Mozilla/5.0 (PlayStation 4 3.00) AppleWebKit/537.73 (KHTML, like Gecko)",
		Expect: TestOsRes{
			Name:    "playstation",
			Version: "4",
		},
	},
	{
		Desc: "Xbox 360",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox 360) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36",
		Expect: TestOsRes{
			Name:    "xbox",
			Version: "360",
		},
	},
	{
		Desc: "Xbox One",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One; WebView/3.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.19041",
		Expect: TestOsRes{
			Name:    "xbox",
			Version: "one",
		},
	},
	{
		Desc: "Xbox X",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.82 Safari/537.36 Edge/20.02",
		Expect: TestOsRes{
			Name:    "xbox",
			Version: "x",
		},
	},
	{
		Desc: "Xbox Series X",
		Ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox Series X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.82 Safari/537.36 Edge/20.02 ",
		Expect: TestOsRes{
			Name:    "xbox",
			Version: "series x",
		},
	},
	{
		Desc: "Mint",
		Ua:   "Opera/9.80 (X11; Linux x86_64; Edition Linux Mint) Presto/2.12.388 Version/12.16",
		Expect: TestOsRes{
			Name:    "mint",
			Version: "",
		},
	},
	{
		Desc: "Mint",
		Ua:   "Opera/9.64 (X11; Linux i686; U; Linux Mint; nb) Presto/2.1.1",
		Expect: TestOsRes{
			Name:    "mint",
			Version: "",
		},
	},
	{
		Desc: "Mint",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.5) Gecko/2008121622 Linux Mint/6 (Felicia) Firefox/3.0.4",
		Expect: TestOsRes{
			Name:    "mint",
			Version: "6",
		},
	},
	{
		Desc: "Ubuntu",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.22+ (KHTML, like Gecko) Chromium/17.0.963.56 Chrome/17.0.963.56 Safari/535.22+ Ubuntu/12.04 (3.4.1-0ubuntu1) Epiphany/3.4.1",
		Expect: TestOsRes{
			Name:    "ubuntu",
			Version: "12.04",
		},
	},
	{
		Desc: "Ubuntu",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/31.0.1650.63 Chrome/31.0.1650.63 Safari/537.36",
		Expect: TestOsRes{
			Name:    "ubuntu",
			Version: "",
		},
	},
	{
		Desc: "Kubuntu",
		Ua:   "Mozilla/5.0 (compatible; Konqueror/4.4; Linux 2.6.32-22-generic; X11; en_US) KHTML/4.4.3 (like Gecko) Kubuntu",
		Expect: TestOsRes{
			Name:    "kubuntu",
			Version: "",
		},
	},
	{
		Desc: "Debian",
		Ua:   "Mozilla/5.0 (compatible; Konqueror/3.5; Linux) KHTML/3.5.7 (like Gecko) (Debian)",
		Expect: TestOsRes{
			Name:    "debian",
			Version: "",
		},
	},
	{
		Desc: "Debian",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64; Debian GNU/Linux 8.1 (jessie)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.0 Maxthon/1.0.5.3 Safari/537.36",
		Expect: TestOsRes{
			Name:    "debian",
			Version: "8.1",
		},
	},
	{
		Desc: "Debian",
		Ua:   "ELinks/0.12~pre5-4 (textmode; Debian; Linux 3.2.0-4-amd64 x86_64 192x47-2)",
		Expect: TestOsRes{
			Name:    "debian",
			Version: "",
		},
	},
	{
		Desc: "Debian",
		Ua:   "w3m/0.5.3+debian-19",
		Expect: TestOsRes{
			Name:    "debian",
			Version: "19",
		},
	},
	{
		Desc: "Debian",
		Ua:   "Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.0.3) Gecko/2008092814 (Debian-3.0.1-1)",
		Expect: TestOsRes{
			Name:    "debian",
			Version: "3.0.1-1",
		},
	},
	{
		Desc: "Debian",
		Ua:   "Mozilla/5.0 (compatible; Konqueror/3.5; Linux 2.6.24.4; X11) KHTML/3.5.9 (like Gecko) (Debian package 4:3.5.9.dfsg.1-2+b1)",
		Expect: TestOsRes{
			Name:    "debian",
			Version: "",
		},
	},
	{
		Desc: "OpenSUSE",
		Ua:   "Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.17) Gecko/20110420 SUSE/3.6.17-0.2.1 Firefox/3.6.17",
		Expect: TestOsRes{
			Name:    "suse",
			Version: "3.6.17-0.2.1",
		},
	},
	{
		Desc: "Gentoo",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.1.16) Gecko/20080716 (Gentoo) Galeon/2.0.6",
		Expect: TestOsRes{
			Name:    "gentoo",
			Version: "",
		},
	},
	{
		Desc: "Gentoo",
		Ua:   "Xombrero (X11; U; Gentoo Linux amd64; en-US) Webkit/2.8.5",
		Expect: TestOsRes{
			Name:    "gentoo",
			Version: "amd64",
		},
	},
	{
		Desc: "Gentoo",
		Ua:   "Xombrero/1.6.4 (Linux amd64; en; Gentoo)",
		Expect: TestOsRes{
			Name:    "gentoo",
			Version: "",
		},
	},
	{
		Desc: "Gentoo",
		Ua:   "Links (2.8; Linux 3.17.2-gentoo-x86 i686; GNU C 4.8.2; x)",
		Expect: TestOsRes{
			Name:    "gentoo",
			Version: "x86",
		},
	},
	{
		Desc: "Arch",
		Ua:   "Uzbl (Webkit 1.1.10) (Arch Linux)",
		Expect: TestOsRes{
			Name:    "arch",
			Version: "",
		},
	},
	{
		Desc: "Slackware",
		Ua:   "Mozilla/5.0 Slackware/13.37 (X11; U; Linux x86_64; en-US) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/13.0.782.41",
		Expect: TestOsRes{
			Name:    "slackware",
			Version: "13.37",
		},
	},
	{
		Desc: "Fedora",
		Ua:   "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:40.0) Gecko/20100101 Firefox/40.0",
		Expect: TestOsRes{
			Name:    "fedora",
			Version: "",
		},
	},
	{
		Desc: "Fedora",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; en-GB; rv:2.0) Gecko/20110404 Fedora/16-dev Firefox/4.0",
		Expect: TestOsRes{
			Name:    "fedora",
			Version: "16-dev",
		},
	},
	{
		Desc: "Fedora",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; sk; rv:1.9.0.4) Gecko/2008111217 Fedora/3.0.4-1.fc10 Firefox/3.0.4",
		Expect: TestOsRes{
			Name:    "fedora",
			Version: "3.0.4-1.fc10",
		},
	},
	{
		Desc: "Mandriva",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.2.22) Gecko/20110907 Mandriva Linux/1.9.2.22-0.1mdv2010.2 (2010.2) Firefox/3.6.22",
		Expect: TestOsRes{
			Name:    "mandriva",
			Version: "1.9.2.22-0.1mdv2010.2",
		},
	},
	{
		Desc: "Chromium OS",
		Ua:   "Mozilla/5.0 (X11; CrOS x86_64 10575.58.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36",
		Expect: TestOsRes{
			Name:    "chromium os",
			Version: "10575.58.0",
		},
	},
	{
		Desc: "Fuchsia",
		Ua:   "Mozilla/5.0 (X11; Fuchsia x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3557.0 Safari/537.36",
		Expect: TestOsRes{
			Name:    "fuchsia",
			Version: "",
		},
	},
	{
		Desc: "Solaris",
		Ua:   "Mozilla/5.0 (X11; U; SunOS sun4u; en-US; rv:1.7) Gecko/20070606",
		Expect: TestOsRes{
			Name:    "solaris",
			Version: "sun4u",
		},
	},
	{
		Desc: "FreeBSD",
		Ua:   "Mozilla/5.0 (X11; U; FreeBSD x86_64; en-US) AppleWebKit/534.16 (KHTML, like Gecko) Chrome/10.0.648.204 Safari/534.16",
		Expect: TestOsRes{
			Name:    "freebsd",
			Version: "",
		},
	},
	{
		Desc: "OpenBSD",
		Ua:   "Mozilla/5.0 (X11; U; OpenBSD i386; en-US; rv:1.9.1) Gecko/20090702 Firefox/3.5",
		Expect: TestOsRes{
			Name:    "openbsd",
			Version: "",
		},
	},
	{
		Desc: "NetBSD",
		Ua:   "ELinks (0.4.3; NetBSD 3.0.2PATCH sparc64; 141x19)",
		Expect: TestOsRes{
			Name:    "netbsd",
			Version: "3.0.2patch",
		},
	},
	{
		Desc: "DragonFly",
		Ua:   "Mozilla/5.0 (X11; U; DragonFly i386; de; rv:1.9.1) Gecko/20090720 Firefox/3.5.1",
		Expect: TestOsRes{
			Name:    "dragonfly",
			Version: "",
		},
	},
	{
		Desc: "iOS in App",
		Ua:   "AppName/version CFNetwork/version Darwin/version",
		Expect: TestOsRes{
			Name:    "ios",
			Version: "",
		},
	},
	{
		Desc: "iOS with Chrome",
		Ua:   "Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; en) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3",
		Expect: TestOsRes{
			Name:    "ios",
			Version: "5.1.1",
		},
	},
	{
		Desc: "iOS with Opera Mini",
		Ua:   "Opera/9.80 (iPhone; Opera Mini/7.1.32694/27.1407; U; en) Presto/2.8.119 Version/11.10",
		Expect: TestOsRes{
			Name:    "ios",
			Version: "",
		},
	},
	{
		Desc: "Mac OS on PowerPC",
		Ua:   "Mozilla/4.0 (compatible; MSIE 5.0b1; Mac_PowerPC)",
		Expect: TestOsRes{
			Name:    "mac os",
			Version: "",
		},
	},
	{
		Desc: "Mac OS X on x86, x86_64, or aarch64 using Firefox",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:10.0) Gecko/20100101 Firefox/10.0",
		Expect: TestOsRes{
			Name:    "mac os",
			Version: "x.y",
		},
	},
	{
		Desc: "Mac OS X on PowerPC using Firefox",
		Ua:   "Mozilla/5.0 (Macintosh; PPC Mac OS X x.y; rv:10.0) Gecko/20100101 Firefox/10.0",
		Expect: TestOsRes{
			Name:    "mac os",
			Version: "x.y",
		},
	},
	{
		Desc: "Mac OS",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.95 Safari/537.36",
		Expect: TestOsRes{
			Name:    "mac os",
			Version: "10.6.8",
		},
	},
	{
		Desc: "Haiku",
		Ua:   "Mozilla/5.0 (Macintosh; Intel Haiku R1 x86) AppleWebKit/602.1.1 (KHTML, like Gecko) WebPositive/1.2 Version/8.0 Safari/602.1.1",
		Expect: TestOsRes{
			Name:    "haiku",
			Version: "r1",
		},
	},
	{
		Desc: "KaiOS",
		Ua:   "Mozilla/5.0 (Mobile; Nokia_8110_4G; rv:48.0) Gecko/48.0 Firefox/48.0 KAIOS/2.5",
		Expect: TestOsRes{
			Name:    "kaios",
			Version: "2.5",
		},
	},
	{
		Desc: "iTunes Windows Vista",
		Ua:   "iTunes/10.7 (Windows; Microsoft Windows Vista Home Premium Edition Service Pack 1 (Build 6001)) AppleWebKit/536.26.9",
		Expect: TestOsRes{
			Name:    "windows",
			Version: "vista",
		},
	},
	{
		Desc: "iOS BE App",
		Ua:   "APP-BE Test/1.0 (iPad; Apple; CPU iPhone OS 7_0_2 like Mac OS X)",
		Expect: TestOsRes{
			Name:    "ios",
			Version: "7.0.2",
		},
	},
	{
		Desc: "KTB-Nexus 5",
		Ua:   "APP-My App/1.0 (Linux; Android 4.2.1; Nexus 5 Build/JOP40D)",
		Expect: TestOsRes{
			Name:    "android",
			Version: "4.2.1",
		},
	},
	{
		Desc: "Solaris",
		Ua:   "NCSA Mosaic/1.0 (X11;SunOS 4.1.4 sun4m)",
		Expect: TestOsRes{
			Name:    "solaris",
			Version: "4.1.4",
		},
	},
	{
		Desc: "Raspbian",
		Ua:   "Mozilla/5.0 (X11; Linux armv7l) AppleWebKit/537.36 (KHTML, like Gecko) Raspbian Chromium/72.0.3626.121 HeadlessChrome/72.0.3626.121 Safari/537.36",
		Expect: TestOsRes{
			Name:    "raspbian",
			Version: "",
		},
	},
	{
		Desc: "Raspbian",
		Ua:   "Mozilla/5.0 (X11; Linux armv7l) AppleWebKit/538.15 (KHTML, like Gecko) Version/8.0 Safari/538.15 Raspbian/9.0 (1:3.8.2.0-0rpi28) Epiphany/3.8.2",
		Expect: TestOsRes{
			Name:    "raspbian",
			Version: "9.0",
		},
	},
	{
		Desc: "AIX",
		Ua:   "Mozilla/5.0 (X11; U; AIX 000138384C00; en-US; rv:1.0.1) Gecko/20030213 Netscape/7.0",
		Expect: TestOsRes{
			Name:    "aix",
			Version: "",
		},
	},
	{
		Desc: "Plan9",
		Ua:   "NCSA_Mosaic/5.0 (X11;Plan 9 4.0)",
		Expect: TestOsRes{
			Name:    "plan 9",
			Version: "4.0",
		},
	},
	{
		Desc: "Minix",
		Ua:   "Mozilla/5.0 (X11; Original ; Minix 3.3 ; rv:3.0)",
		Expect: TestOsRes{
			Name:    "minix",
			Version: "3.3",
		},
	},
	{
		Desc: "BeOS",
		Ua:   "Mozilla/5.0 (BeOS; U; BeOS BePC; en-US; rv:1.8.1.8pre) Gecko/20070926 SeaMonkey/1.1.5pre",
		Expect: TestOsRes{
			Name:    "beos",
			Version: "",
		},
	},
	{
		Desc: "OS/2",
		Ua:   "Links (2.1pre14; OS/2 1 i386; 80x33)",
		Expect: TestOsRes{
			Name:    "os/2",
			Version: "",
		},
	},
	{
		Desc: "AmigaOS",
		Ua:   "Mozilla/4.0 (compatible; AWEB 3.4 SE; AmigaOS)",
		Expect: TestOsRes{
			Name:    "amigaos",
			Version: "",
		},
	},
	{
		Desc: "MorphOS",
		Ua:   "AmigaVoyager/3.4.4 (MorphOS/PPC native)",
		Expect: TestOsRes{
			Name:    "morphos",
			Version: "",
		},
	},
	{
		Desc: "UNIX",
		Ua:   "Surf/0.4.1 (X11; U; Unix; en-US) AppleWebKit/531.2+ Compatible (Safari)",
		Expect: TestOsRes{
			Name:    "unix",
			Version: "",
		},
	},
	{
		Desc: "Joli",
		Ua:   "Mozilla/5.0 (X11; Jolicloud Linux i686) AppleWebKit/537.6 (KHTML, like Gecko) Joli OS/1.2 Chromium/23.0.1240.0 Chrome/23.0.1240.0 Safari/537.6",
		Expect: TestOsRes{
			Name:    "joli",
			Version: "1.2",
		},
	},
	{
		Desc: "CentOS",
		Ua:   "Konqueror/15.13 (CentOS Linux 7.4; cs-CZ;)",
		Expect: TestOsRes{
			Name:    "centos",
			Version: "7.4",
		},
	},
	{
		Desc: "PCLinuxOS",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.2.13) Gecko/20101209 PCLinuxOS/1.9.2.13-1pclos2010 (2010) Firefox/3.6.13",
		Expect: TestOsRes{
			Name:    "pclinuxos",
			Version: "1.9.2.13-1pclos2010",
		},
	},
	{
		Desc: "RedHat",
		Ua:   "Mozilla/5.0 (compatible; Konqueror/4.3; Linux) KHTML/4.3.4 (like Gecko) Red Hat Enterprise Linux/4.3.4-11.el6_1.4",
		Expect: TestOsRes{
			Name:    "red hat",
			Version: "4.3.4-11.el6_1.4",
		},
	},
	{
		Desc: "RedHat",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.0.13pre) Gecko/20070717 Red Hat/1.0.9-4.el4 SeaMonkey/1.0.9",
		Expect: TestOsRes{
			Name:    "red hat",
			Version: "1.0.9-4.el4",
		},
	},
	{
		Desc: "RedHat",
		Ua:   "iTunes/4.7.1 (Linux; N; Red Hat; x86_64-linux; EN; utf8) SqueezeCenter, Squeezebox Server, Logitech Media Server/7.9.1/1522157629",
		Expect: TestOsRes{
			Name:    "red hat",
			Version: "",
		},
	},
	{
		Desc: "RedHat",
		Ua:   "curl/7.20.0 (x86_64-redhat-linux-gnu) libcurl/7.20.0 OpenSSL/0.9.8b zlib/1.2.3 libidn/0.6.5",
		Expect: TestOsRes{
			Name:    "redhat",
			Version: "",
		},
	},
	{
		Desc: "RISC OS",
		Ua:   "Mozilla/1.10 [en] (Compatible; RISC OS 3.70; Oregano 1.10)",
		Expect: TestOsRes{
			Name:    "risc os",
			Version: "3.70",
		},
	},
	{
		Desc: "Zenwalk",
		Ua:   "Flock/2.16 (Zenwalk 7.3; es_PR;)",
		Expect: TestOsRes{
			Name:    "zenwalk",
			Version: "7.3",
		},
	},
	{
		Desc: "Hurd",
		Ua:   "Mozilla/5.0 (X11; Hurd 0.9 i386; en-US) libwww-FM/2.14 SSL-MM/1.4.1 GNUTLS/3.7.0 Safari/696.96",
		Expect: TestOsRes{
			Name:    "hurd",
			Version: "0.9",
		},
	},
	{
		Desc: "Linux",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36",
		Expect: TestOsRes{
			Name:    "linux",
			Version: "x86_64",
		},
	},
	{
		Desc: "Deepin",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64; Deepin 15.5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.75 Safari/537.36 NFSBrowser/5.0.0.1886",
		Expect: TestOsRes{
			Name:    "deepin",
			Version: "15.5",
		},
	},
	{
		Desc: "Palm OS",
		Ua:   "Mozilla/4.76 [en] (PalmOS; U; WebPro3.0; Palm-Arz1)",
		Expect: TestOsRes{
			Name:    "palm",
			Version: "",
		},
	},
	{
		Desc: "HP-UX",
		Ua:   "Mozilla/5.0 (X11; U; HP-UX 9000/785; es-ES; rv:1.0.1) Gecko/20020827 Netscape/7.0",
		Expect: TestOsRes{
			Name:    "hp-ux",
			Version: "",
		},
	},
	{
		Desc: "Contiki",
		Ua:   "Contiki/1.0 (Commodore 64; http://dunkels.com/adam/contiki/)",
		Expect: TestOsRes{
			Name:    "contiki",
			Version: "1.0",
		},
	},
	{
		Desc: "Linpus",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9b5pre) Gecko/2008032619 Linpus/3.0-0.49",
		Expect: TestOsRes{
			Name:    "linpus",
			Version: "3.0-0.49",
		},
	},
	{
		Desc: "Manjaro",
		Ua:   "Mozilla/5.0 (X11; Manjaro 19.0.2; Arch; x64; rv:84.0) Gecko/20100101 Firefox/84.0",
		Expect: TestOsRes{
			Name:    "manjaro",
			Version: "19.0.2",
		},
	},
	{
		Desc: "elementary OS",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/604.1 (KHTML, like Gecko) Version/11.0 Safari/604.1 elementary OS/0.4 (Loki) Epiphany/3.18.11",
		Expect: TestOsRes{
			Name:    "elementary os",
			Version: "0.4",
		},
	},
	{
		Desc: "GhostBSD",
		Ua:   "Mozilla/5.0 (X11; GhostBSD/10.3; x86_64; rv:50.0.1) Gecko/20100101 Firefox/50.0.1",
		Expect: TestOsRes{
			Name:    "ghostbsd",
			Version: "10.3",
		},
	},
	{
		Desc: "Android-x86",
		Ua:   "Mozilla/5.0 (Linux; Android 7.1.2; Generic Android-x86) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36 OPR/61.2.3076.56749",
		Expect: TestOsRes{
			Name:    "android-x86",
			Version: "7.1.2",
		},
	},
	{
		Desc: "Sabayon",
		Ua:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/536.5 (KHTML, like Gecko) Sabayon Chrome/19.0.1084.46 Safari/536.5",
		Expect: TestOsRes{
			Name:    "sabayon",
			Version: "",
		},
	},
	{
		Desc: "Linspire",
		Ua:   "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.0.4) Gecko/20060803 Firefox/1.5.0.4 Linspire/1.5.0.4",
		Expect: TestOsRes{
			Name:    "linspire",
			Version: "1.5.0.4",
		},
	},
}
