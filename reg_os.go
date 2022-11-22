package dbuaparser

type OsInfo struct {
	Os    string // 操作系统
	OsVer string // 操作系统版本
}

var windowsVersionMap = handleMaps{
	"4.90":    "me",
	"NT3.51":  "nt 3.11",
	"NT4.0":   "nt 4.0",
	"NT 5.0":  "2000",
	"NT 5.1":  "xp",
	"NT 5.2":  "xp",
	"NT 6.0":  "vista",
	"NT 6.1":  "7",
	"NT 6.2":  "8",
	"NT 6.3":  "8.1",
	"NT 6.4":  "10",
	"NT 10.0": "10",
	"ARM":     "rt",
}

var osMap = []RegexpMap{
	// Windows
	{
		Regs: []string{
			`(?i)microsoft (windows) (vista|xp)`, // Windows (iTunes)
		},
		Props: Props{
			Os: PropMap{
				Index: 1,
			},
			OsVer: PropMap{
				Index: 2,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(windows) nt 6\.2; (arm)`,                             // Windows RT
			`(?i)(windows (?:phone(?: os)?|mobile))[\/ ]?([\d\.\w ]*)`, // Windows Phone
			`(?i)(windows)[\/ ]?([ntce\d\. ]+\w)(?!.+xbox)`,
		},
		Props: Props{
			Os: PropMap{
				Index: 1,
			},
			OsVer: PropMap{
				Index:     2,
				Handle:    strMapper{},
				HandleMap: windowsVersionMap,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(win(?=3|9|n)|win 9x )([nt\d\.]+)`,
		},
		Props: Props{
			Os: PropMap{
				Value: "windows",
			},
			OsVer: PropMap{
				Index:     2,
				Handle:    strMapper{},
				HandleMap: windowsVersionMap,
			},
		},
	},

	// iOS/macOS
	{
		Regs: []string{
			`(?i)ip[honead]{2,4}\b(?:.*os ([\w]+) like mac|; opera)`, // IOS
			`(?i)cfnetwork\/.+darwin`,
		},
		Props: Props{
			Os: PropMap{
				Value: "ios",
			},
			OsVer: PropMap{
				Index:   1,
				Reg:     "_",
				Replace: ".",
			},
		},
	},
	{
		Regs: []string{
			`(?i)(mac os x) ?([\w\. ]*)`, // Mac OS
			`(?i)(macintosh|mac_powerpc\b)(?!.+haiku)`,
		},
		Props: Props{
			Os: PropMap{
				Value: "mac os",
			},
			OsVer: PropMap{
				Index:   2,
				Reg:     "_",
				Replace: ".",
			},
		},
	},

	// Mobile OSes
	{
		Regs: []string{
			`(?i)droid ([\w\.]+)\b.+(android[- ]x86|harmonyos)`, // Android-x86/HarmonyOS
			`(?i)(macintosh|mac_powerpc\b)(?!.+haiku)`,
		},
		Props: Props{
			Os: PropMap{
				Index: 2,
			},
			OsVer: PropMap{
				Index: 1,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(android|webos|qnx|bada|rim tablet os|maemo|meego|sailfish)[-\/ ]?([\w\.]*)`, // Android/WebOS/QNX/Bada/RIM/Maemo/MeeGo/Sailfish OS
			`(?i)(blackberry)\w*\/([\w\.]*)`,  // Blackberry
			`(?i)(tizen|kaios)[\/ ]([\w\.]+)`, // Tizen/KaiOS
			`(?i)\((series40);`,               // Series 40
		},
		Props: Props{
			Os: PropMap{
				Index: 1,
			},
			OsVer: PropMap{
				Index: 2,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\(bb(10);`, // BlackBerry 10
		},
		Props: Props{
			Os: PropMap{
				Value: _BLACKBERRY,
			},
			OsVer: PropMap{
				Index: 1,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(?:symbian ?os|symbos|s60(?=;)|series60)[-\/ ]?([\w\.]*)`, // Symbian
		},
		Props: Props{
			Os: PropMap{
				Value: "symbian",
			},
			OsVer: PropMap{
				Index: 1,
			},
		},
	},
	{
		Regs: []string{
			`(?i)mozilla\/[\d\.]+ \((?:mobile|tablet|tv|mobile; [\w ]+); rv:.+ gecko\/([\w\.]+)`, // Firefox OS
		},
		Props: Props{
			Os: PropMap{
				Value: _FIREFOX + " os",
			},
			OsVer: PropMap{
				Index: 1,
			},
		},
	},
	{
		Regs: []string{
			`(?i)web0s;.+rt(tv)`, // WebOS
			`(?i)\b(?:hp)?wos(?:browser)?\/([\w\.]+)`,
		},
		Props: Props{
			Os: PropMap{
				Value: "webos",
			},
			OsVer: PropMap{
				Index: 1,
			},
		},
	},

	// Google Chromecast
	{
		Regs: []string{
			`(?i)crkey\/([\d\.]+)`, // Google Chromecast
		},
		Props: Props{
			Os: PropMap{
				Value: _CHROME + "cast",
			},
			OsVer: PropMap{
				Index: 1,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(cros) [\w]+ ([\w\.]+\w)`, // Chromium OS
		},
		Props: Props{
			Os: PropMap{
				Value: "chromium os",
			},
			OsVer: PropMap{
				Index: 2,
			},
		},
	},

	// Console
	{
		Regs: []string{
			`(?i)(nintendo|playstation) ([wids345portablevuch]+)`, // Nintendo/Playstation
			`(?i)(xbox); +xbox ([^\);]+)`,                         // Microsoft Xbox (360, One, X, S, Series X, Series S)

			// Other
			`(?i)\b(joli|palm)\b ?(?:os)?\/?([\w\.]*)`, // Joli/Palm
			`(?i)(mint)[\/\(\) ]?(\w*)`,                // Mint
			`(?i)(mageia|vectorlinux)[; ]`,             // Mageia/VectorLinux
			// Ubuntu/Debian/SUSE/Gentoo/Arch/Slackware/Fedora/Mandriva/CentOS/PCLinuxOS/RedHat/Zenwalk/Linpus/Raspbian/Plan9/Minix/RISCOS/Contiki/Deepin/Manjaro/elementary/Sabayon/Linspire
			`(?i)([kxln]?ubuntu|debian|suse|opensuse|gentoo|arch(?= linux)|slackware|fedora|mandriva|centos|pclinuxos|red ?hat|zenwalk|linpus|raspbian|plan 9|minix|risc os|contiki|deepin|manjaro|elementary os|sabayon|linspire)(?: gnu\/linux)?(?: enterprise)?(?:[- ]linux)?(?:-gnu)?[-\/ ]?(?!chrom|package)([-\w\.]*)`,
			`(?i)(hurd|linux) ?([\w\.]*)`, // Hurd/Linux
			`(?i)(gnu) ?([\w\.]*)`,        // GNU
			`(?i)\b([-frentopcghs]{0,5}bsd|dragonfly)[\/ ]?(?!amd|[ix346]{1,2}86)([\w\.]*)`, // FreeBSD/NetBSD/OpenBSD/PC-BSD/GhostBSD/DragonFly
			`(?i)(haiku) (\w+)`, // Haiku
		},
		Props: Props{
			Os: PropMap{
				Index: 1,
			},
			OsVer: PropMap{
				Index: 2,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(sunos) ?([\w\.\d]*)`, // Solaris
		},
		Props: Props{
			Os: PropMap{
				Value: "solaris",
			},
			OsVer: PropMap{
				Index: 2,
			},
		},
	},
	{
		Regs: []string{
			`(?i)((?:open)?solaris)[-\/ ]?([\w\.]*)`,                   // Solaris
			`(?i)(aix) ((\d)(?=\.|\)| )[\w\.])*`,                       // AIX
			`(?i)\b(beos|os\/2|amigaos|morphos|openvms|fuchsia|hp-ux)`, // BeOS/OS2/AmigaOS/MorphOS/OpenVMS/Fuchsia/HP-UX
			`(?i)(unix) ?([\w\.]*)`,                                    // UNIX
		},
		Props: Props{
			Os: PropMap{
				Index: 1,
			},
			OsVer: PropMap{
				Index: 2,
			},
		},
	},
}
