package dbuaparser

var oldSafariMap = map[string]string{
	"/8":   "1.0",
	"/1":   "1.2",
	"/3":   "1.3",
	"/412": "2.0",
	"/416": "2.0.2",
	"/417": "2.0.3",
	"/419": "2.0.4",
	"/":    "?",
}

type BrowserInfo struct {
	Browser    string // 浏览器名称
	BrowserVer string // 浏览器版本
}

var browserMap = []RegexpMap{
	{
		Regs: []string{
			`(?i)\b(?:crmo|crios)\/([\w\.]+)`, // Chrome for Android/iOS
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: _CHROME,
			},
		},
	},
	{
		Regs: []string{
			`(?i)edg(?:e|ios|a)?\/([\w\.]+)`, // Microsoft Edge
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: _EDGE,
			},
		},
	},

	// Presto based
	{
		Regs: []string{
			`(?i)(opera mini)\/([-\w\.]+)`,                        // Opera Mini
			`(?i)(opera [mobiletab]{3,6})\b.+version\/([-\w\.]+)`, // Opera Mobi/Tablet
			`(?i)(opera)(?:.+version\/|[\/ ]+)([\w\.]+)`,          // Opera
		},
		Props: Props{
			Version: PropMap{
				Index: 2,
			},
			Name: PropMap{
				Index: 1,
			},
		},
	},
	{
		Regs: []string{
			`(?i)opios[\/ ]+([\w\.]+)`, // Opera mini on iphone >= 8.0
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: _OPERA + " mini",
			},
		},
	},
	{
		Regs: []string{
			`(?i)\bopr\/([\w\.]+)`, // Opera Webkit
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: _OPERA,
			},
		},
	},

	{
		Regs: []string{
			// Mixed
			`(?i)(kindle)\/([\w\.]+)`, // Kindle
			`(?i)(lunascape|maxthon|netfront|jasmine|blazer)[\/ ]?([\w\.]*)`, // Lunascape/Maxthon/Netfront/Jasmine/Blazer

			// Trident based
			`(?i)(avant |iemobile|slim)(?:browser)?[\/ ]?([\w\.]*)`, // Avant/IEMobile/SlimBrowser
			`(?i)(ba?idubrowser)[\/ ]?([\w\.]+)`,                    // Baidu Browser
			`(?i)(?:ms|\()(ie) ([\w\.]+)`,                           // Internet Explorer

			// Webkit/KHTML based
			`(?i)(flock|rockmelt|midori|epiphany|silk|skyfire|ovibrowser|bolt|iron|vivaldi|iridium|phantomjs|bowser|quark|qupzilla|falkon|rekonq|puffin|brave|whale|qqbrowserlite|qq|duckduckgo)\/([-\w\.]+)`, // Rekonq/Puffin/Brave/Whale/QQBrowserLite/QQ, aka ShouQ
			`(?i)(weibo)__([\d\.]+)`, // Weibo
		},
		Props: Props{
			Version: PropMap{
				Index: 2,
			},
			Name: PropMap{
				Index: 1,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(?:\buc? ?browser|(?:juc.+)ucweb)[\/ ]?([\w\.]+)`, // UCBrowser
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: "uc " + _BROWSER,
			},
		},
	},
	{
		Regs: []string{
			`(?i)microm.+\bqbcore\/([\w\.]+)`, // WeChat Desktop for Windows Built-in Browser
			`(?i)\bqbcore\/([\w\.]+).+microm`,
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: "wechat(win) desktop",
			},
		},
	},
	{
		Regs: []string{
			`(?i)micromessenger\/([\w\.]+)`, // WeChat
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: "wechat",
			},
		},
	},
	{
		Regs: []string{
			`(?i)konqueror\/([\w\.]+)`, // Konqueror
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: "konqueror",
			},
		},
	},
	{
		Regs: []string{
			`(?i)trident.+rv[: ]([\w\.]{1,9})\b.+like gecko`, // IE11
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: "ie",
			},
		},
	},
	{
		Regs: []string{
			`(?i)yabrowser\/([\w\.]+)`, // Yandex
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: "yandex",
			},
		},
	},
	{
		Regs: []string{
			`(?i)(avast|avg)\/([\w\.]+)`, // Avast/AVG Secure Browser
		},
		Props: Props{
			Version: PropMap{
				Index: 2,
			},
			Name: PropMap{
				Index:   1,
				Reg:     "(.+)",
				Replace: "$1 secure " + _BROWSER,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\bfocus\/([\w\.]+)`, // Firefox Focus
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: _FIREFOX + " focus",
			},
		},
	},
	{
		Regs: []string{
			`(?i)\bopt\/([\w\.]+)`, // Opera Touch
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: _OPERA + " touch",
			},
		},
	},
	{
		Regs: []string{
			`(?i)coc_coc\w+\/([\w\.]+)`, // Coc Coc Browser
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: "coc coc",
			},
		},
	},
	{
		Regs: []string{
			`(?i)dolfin\/([\w\.]+)`, // Dolphin
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: "dolphin",
			},
		},
	},
	{
		Regs: []string{
			`(?i)coast\/([\w\.]+)`, // Opera Coast
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: _OPERA + " coast",
			},
		},
	},
	{
		Regs: []string{
			`(?i)miuibrowser\/([\w\.]+)`, // MIUI Browser
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: "miui " + _BROWSER,
			},
		},
	},
	{
		Regs: []string{
			`(?i)fxios\/([-\w\.]+)`, // Firefox for iOS
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: _FIREFOX,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\bqihu|(qi?ho?o?|360)browser`, // 360
		},
		Props: Props{
			Name: PropMap{
				Value: "360 " + _BROWSER,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(oculus|samsung|sailfish|huawei)browser\/([\w\.]+)`, // Oculus/Samsung/Sailfish/Huawei Browser
		},
		Props: Props{
			Version: PropMap{
				Index: 2,
			},
			Name: PropMap{
				Index:   1,
				Reg:     "(.+)",
				Replace: "$1 " + _BROWSER,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(comodo_dragon)\/([\w\.]+)`, // Comodo Dragon
		},
		Props: Props{
			Version: PropMap{
				Index: 2,
			},
			Name: PropMap{
				Index:   1,
				Reg:     "(?g)_",
				Replace: " ",
			},
		},
	},
	{
		Regs: []string{
			`(?i)(electron)\/([\w\.]+) safari`,                          // Electron-based App
			`(?i)(tesla)(?: qtcarbrowser|\/(20\d\d\.[-\w\.]+))`,         // Tesla
			`(?i)m?(qqbrowser|baiduboxapp|2345Explorer)[\/ ]?([\w\.]+)`, // QQBrowser/Baidu App/2345 Browser
		},
		Props: Props{
			Version: PropMap{
				Index: 2,
			},
			Name: PropMap{
				Index: 1,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(metasr)[\/ ]?([\w\.]+)`, // SouGouBrowser
			`(?i)(lbbrowser)`,             // LieBao Browser
			`(?i)\[(linkedin)app\]`,       // LinkedIn App for iOS & Android
		},
		Props: Props{
			Name: PropMap{
				Index: 1,
			},
		},
	},

	// WebView
	{
		Regs: []string{
			`(?i)((?:fban\/fbios|fb_iab\/fb4a)(?:!.+fbav)|;fbav\/([\w\.]+);)`, // Facebook App for iOS & Android
			`(?i)((?:fban\/fbios|fb_iab\/fb4a)|;fbav\/([\w\.]+);)`,            // Facebook App for iOS & Android
		},
		Props: Props{
			Version: PropMap{
				Index: 2,
			},
			Name: PropMap{
				Value: _FACEBOOK,
			},
		},
	},
	{
		Regs: []string{
			`(?i)safari (line)\/([\w\.]+)`,            // Line App for iOS
			`(?i)\b(line)\/([\w\.]+)\/iab`,            // Line App for Android
			`(?i)(chromium|instagram)[\/ ]([-\w\.]+)`, // Chromium/Instagram
		},
		Props: Props{
			Version: PropMap{
				Index: 2,
			},
			Name: PropMap{
				Index: 1,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\bgsa\/([\w\.]+) .*safari\/`, // Google Search Appliance on iOS
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: "gsa",
			},
		},
	},
	{
		Regs: []string{
			`(?i)headlesschrome(?:\/([\w\.]+)| )`, // Chrome Headless
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: _CHROME + " headless",
			},
		},
	},
	{
		Regs: []string{
			`(?i) wv\).+(chrome)\/([\w\.]+)`, // Chrome WebView
		},
		Props: Props{
			Version: PropMap{
				Index: 2,
			},
			Name: PropMap{
				Value: _CHROME + " webview",
			},
		},
	},
	{
		Regs: []string{
			`(?i)droid.+ version\/([\w\.]+)\b.+(?:mobile safari|safari)`, // Android Browser
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: "android " + _BROWSER,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(chrome|omniweb|arora|[tizenoka]{5} ?browser)\/v?([\w\.]+)`, // Chrome/OmniWeb/Arora/Tizen/Nokia
		},
		Props: Props{
			Version: PropMap{
				Index: 2,
			},
			Name: PropMap{
				Index: 1,
			},
		},
	},
	{
		Regs: []string{
			`(?i)version\/([\w\.\,]+) .*mobile\/\w+ (safari)`, // Mobile Safari
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: "mobile safari",
			},
		},
	},
	{
		Regs: []string{
			`(?i)version\/([\w(\.|\,)]+) .*(mobile ?safari|safari)`, // Safari & Safari Mobile
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Index: 2,
			},
		},
	},
	{
		Regs: []string{
			`(?i)webkit.+?(mobile ?safari|safari)(\/[\w\.]+)`, // Safari < 3.0
		},
		Props: Props{
			Version: PropMap{
				Index:     2,
				Handle:    strMapper{},
				HandleMap: oldSafariMap,
			},
			Name: PropMap{
				Index: 1,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(webkit|khtml)\/([\w\.]+)`,
		},
		Props: Props{
			Version: PropMap{
				Index: 2,
			},
			Name: PropMap{
				Index: 1,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(navigator|netscape\d?)\/([-\w\.]+)`, // Netscape
		},
		Props: Props{
			Version: PropMap{
				Index: 2,
			},
			Name: PropMap{
				Value: "netscape",
			},
		},
	},
	{
		Regs: []string{
			`(?i)mobile vr; rv:([\w\.]+)\).+firefox`, // Firefox Reality
		},
		Props: Props{
			Version: PropMap{
				Index: 1,
			},
			Name: PropMap{
				Value: _FIREFOX + " reality",
			},
		},
	},
	{
		Regs: []string{
			`(?i)ekiohf.+(flow)\/([\w\.]+)`, // Flow
			`(?i)(swiftfox)`,                // Swiftfox
			`(?i)(icedragon|iceweasel|camino|chimera|fennec|maemo browser|minimo|conkeror|klar)[\/ ]?([\w\.\+]+)`, // IceDragon/Iceweasel/Camino/Chimera/Fennec/Maemo/Minimo/Conkeror/Klar
			`(?i)(seamonkey|k-meleon|icecat|iceape|firebird|phoenix|palemoon|basilisk|waterfox)\/([-\w\.]+)$`,     // Firefox/SeaMonkey/K-Meleon/IceCat/IceApe/Firebird/Phoenix
			`(?i)(firefox)\/([\w\.]+)`,                    // Other Firefox-based
			`(?i)(mozilla)\/([\w\.]+) .+rv\:.+gecko\/\d+`, // Mozilla
			`(?i)(polaris|lynx|dillo|icab|doris|amaya|w3m|netsurf|sleipnir|obigo|mosaic|(?:go|ice|up)[\. ]?browser)[-\/ ]?v?([\w\.]+)`, // Polaris/Lynx/Dillo/iCab/Doris/Amaya/w3m/NetSurf/Sleipnir/Obigo/Mosaic/Go/ICE/UP.Browser
			`(?i)(links) \(([\w\.]+)`, // Links
		},
		Props: Props{
			Version: PropMap{
				Index: 2,
			},
			Name: PropMap{
				Index: 1,
			},
		},
	},
}
