package dbuaparser

type DeviceInfo struct {
	Device string // 设备名称
	Vendor string // 供应商
	Type   string // 设备类型
}

var deviceMap = []RegexpMap{
	//////////////////////////
	// MOBILES & TABLETS
	// Ordered by popularity
	/////////////////////////

	// Samsung
	{
		Regs: []string{
			`(?i)(?i)\b(sch-i[89]0\d|shw-m380s|sm-[ptx]\w{2,4}|gt-[pn]\d{2,4}|sgh-t8[56]9|nexus 10)`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _SAMSUNG,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(?i)\b((?:s[cgp]h|gt|sm)-\w+|galaxy nexus)`,
			`(?i)(?i)samsung[- ]([-\w]+)`,
			`(?i)(?i)sec-(sgh\w+)`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _SAMSUNG,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// Apple
	{
		Regs: []string{
			`(?i)\((ip(?:hone|od)[\w ]*);`, // iPod/iPhone
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _APPLE,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\((ipad);[-\w\),; ]+apple`, // iPad
			`(?i)applecoremedia\/[\w\.]+ \((ipad)`,
			`(?i)\b(ipad)\d\d?,\d\d?[;\]].+ios`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _APPLE,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},

	// Huawei
	{
		Regs: []string{
			`(?i)\b((?:ag[rs][23]?|bah2?|sht?|btv)-a?[lw]\d{2})\b(?!.+d\/s)`,
			`(?i)applecoremedia\/[\w\.]+ \((ipad)`,
			`(?i)\b(ipad)\d\d?,\d\d?[;\]].+ios`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _HUAWEI,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(?:huawei|honor)([-\w ]+)[;\)]`,
			`(?i)\b(nexus 6p|\w{2,4}e?-[atu]?[ln][\dx][012359c][adn]?)\b(?!.+d\/s)`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _HUAWEI,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// Xiaomi
	{
		Regs: []string{
			`(?i)\b(poco[\w ]+)(?: bui|\))`,                     // Xiaomi POCO
			`(?i)\b; (\w+) build\/hm\1`,                         // Xiaomi Hongmi 'numeric' models
			`(?i)\b(hm[-_ ]?note?[_ ]?(?:\d\w)?) bui`,           // Xiaomi Hongmi
			`(?i)\b(redmi[\-_ ]?(?:note|k)?[\w_ ]+)(?: bui|\))`, // Xiaomi Redmi
			`(?i)\b(mi[-_ ]?(?:a\d|one|one[_ ]plus|note lte|max|cc)?[_ ]?(?:\d?\w?)[_ ]?(?:plus|se|lite)?)(?: bui|\))`, // Xiaomi Mi
		},
		Props: Props{
			Device: PropMap{
				Index:   1,
				Reg:     "_",
				Replace: " ",
				Handle:  toLower{},
			},
			Vendor: PropMap{
				Value: _XIAOMI,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(mi[-_ ]?(?:pad)(?:[\w_ ]+))(?: bui|\))`, // Mi Pad tablets
		},
		Props: Props{
			Device: PropMap{
				Index:   1,
				Reg:     "_",
				Replace: " ",
				Handle:  toLower{},
			},
			Vendor: PropMap{
				Value: _XIAOMI,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},

	// _OPPO
	{
		Regs: []string{
			`(?i); (\w+) bui.+ oppo`,
			`(?i)\b(cph[12]\d{3}|p(?:af|c[al]|d\w|e[ar])[mt]\d0|x9007|a101op)\b`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _OPPO,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// Vivo
	{
		Regs: []string{
			`(?i)vivo (\w+)(?: bui|\))`,
			`(?i)\b(v[12]\d{3}\w?[at])(?: bui|;)`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _VIVO,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// Realme
	{
		Regs: []string{
			`(?i)\b(rmx[12]\d{3})(?: bui|;|\))`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _REALME,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// Motorola
	{
		Regs: []string{
			`(?i)\b(milestone|droid(?:[2-4x]| (?:bionic|x2|pro|razr))?:?( 4g)?)\b[\w ]+build\/`,
			`(?i)\bmot(?:orola)?[- ](\w*)`,
			`(?i)((?:moto[\w\(\) ]+|xt\d{3,4}|nexus 6)(?= bui|\)))`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _MOTOROLA,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(mz60\d|xoom[2 ]{0,2}) build\/`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _MOTOROLA,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},

	// _LG
	{
		Regs: []string{
			`(?i)((?=lg)?[vl]k\-?\d{3}) bui| 3\.[-\w; ]{10}lg?-([06cv9]{3,4})`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _LG,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(lm(?:-?f100[nv]?|-[\w\.]+)(?= bui|\))|nexus [45])`,
			`(?i)\blg[-e;\/ ]+((?!browser|netcast|android tv)\w+)`,
			`(?i)\blg-?([\d\w]+) bui`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _LG,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// Lenovo
	{
		Regs: []string{
			`(?i)(ideatab[-\w ]+)`,
			`(?i)lenovo ?(s[56]000[-\w]+|tab(?:[\w ]+)|yt[-\d\w]{6}|tb[-\d\w]{6})`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _LENOVO,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},

	// Nokia
	{
		Regs: []string{
			`(?i)(?:maemo|nokia).*(n900|lumia \d+)`,
			`(?i)nokia[-_ ]?([-\w\.]*)`,
		},
		Props: Props{
			Device: PropMap{
				Index:   1,
				Reg:     "_",
				Replace: " ",
				Handle:  toLower{},
			},
			Vendor: PropMap{
				Value: _NOKIA,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// Google
	{
		Regs: []string{
			`(?i)(pixel c)\b`, // Google Pixel C
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _GOOGLE,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)droid.+; (pixel[\daxl ]{0,6})(?: bui|\))`, // Google Pixel
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _GOOGLE,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// Sony
	{
		Regs: []string{
			`(?i)droid.+ (a?\d[0-2]{2}so|[c-g]\d{4}|so[-gl]\w+|xq-a\w[4-7][12])(?= bui|\).+chrome\/(?![1-6]{0,1}\d\.))`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _SONY,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)sony tablet [ps]`,
			`(?i)\b(?:sony)?sgp\w+(?: bui|\))`,
		},
		Props: Props{
			Device: PropMap{
				Value: "xperia tablet",
			},
			Vendor: PropMap{
				Value: _SONY,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},

	// OnePlus
	{
		Regs: []string{
			`(?i) (kb2005|in20[12]5|be20[12][59])\b`,
			`(?i)(?:one)?(?:plus)? (a\d0\d\d)(?: b|\))`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _ONEPLUS,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// Amazon
	{
		Regs: []string{
			`(?i)(alexa)webm`,
			`(?i)(kf[a-z]{2}wi)( bui|\))`,     // Kindle Fire without Silk
			`(?i)(kf[a-z]+)( bui|\)).+silk\/`, // Kindle Fire HD
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _AMAZON,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)((?:sd|kf)[0349hijorstuw]+)( bui|\)).+silk\/`, // Fire Phone
		},
		Props: Props{
			Device: PropMap{
				Index:   1,
				Reg:     "(?g)(.+)",
				Replace: "fire phone $1",
				Handle:  toLower{},
			},
			Vendor: PropMap{
				Value: _AMAZON,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// BlackBerry
	{
		Regs: []string{
			`(?i)(playbook);[-\w\),; ]+(rim)`, // BlackBerry PlayBook
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _BLACKBERRY,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b((?:bb[a-f]|st[hv])100-\d)`,
			`(?i)\(bb10; (\w+)`, // BlackBerry 10
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _BLACKBERRY,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// Asus
	{
		Regs: []string{
			`(?i)(?:\b|asus_)(transfo[prime ]{4,10} \w+|eeepc|slider \w+|nexus 7|padfone|p00[cj])`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _ASUS,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i) (z[bes]6[027][012][km][ls]|zenfone \d\w?)\b`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _ASUS,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// _HTC
	{
		Regs: []string{
			`(?i)(nexus 9)`, // _HTC Nexus 9
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _HTC,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(htc)[-;_ ]{1,2}([\w ]+(?=\)| bui)|\w+)`, // _HTC
			`(?i)(zte)[- ]([\w ]+?)(?: bui|\/|\))`,
			`(?i)(alcatel|geeksphone|nexian|panasonic|sony(?!-bra))[-_ ]?([-\w]*)`, // Alcatel/GeeksPhone/Nexian/Panasonic/Sony
		},
		Props: Props{
			Device: PropMap{
				Index:   2,
				Reg:     "_",
				Replace: " ",
				Handle:  toLower{},
			},
			Vendor: PropMap{
				Index: 1,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// Acer
	{
		Regs: []string{
			`(?i)droid.+; ([ab][1-7]-?[0178a]\d\d?)`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _ACER,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},

	// Meizu
	{
		Regs: []string{
			`(?i)droid.+; (m[1-5] note) bui`,
			`(?i)\bmz-([-\w]{2,})`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _MEIZU,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// Sharp
	{
		Regs: []string{
			`(?i)\b(sh-?[altvz]?\d\d[a-ekm]?)`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _SHARP,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	// MIXED
	{
		Regs: []string{
			`(?i)(blackberry|benq|palm(?=\-)|sonyericsson|acer|asus|dell|meizu|motorola|polytron)[-_ ]?([-\w]*)`, // BlackBerry/BenQ/Palm/Sony-Ericsson/Acer/Asus/Dell/Meizu/Motorola/Polytron
			`(?i)(hp) ([\w ]+\w)`,            // HP iPAQ
			`(?i)(asus)-?(\w+)`,              // Asus
			`(?i)(microsoft); (lumia[\w ]+)`, // Microsoft Lumia
			`(?i)(lenovo)[-_ ]?([-\w]+)`,     // Lenovo
			`(?i)(jolla)`,                    // Jolla
			`(?i)(oppo) ?([\w ]+) bui`,       // _OPPO
		},
		Props: Props{
			Device: PropMap{
				Index: 2,
			},
			Vendor: PropMap{
				Index: 1,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	{
		Regs: []string{
			`(?i)(archos) (gamepad2?)`,                // Archos
			`(?i)(hp).+(touchpad(?!.+tablet)|tablet)`, // HP TouchPad
			`(?i)(kindle)\/([\w\.]+)`,                 // Kindle
			`(?i)(nook)[\w ]+build\/(\w+)`,            // Nook
			`(?i)(dell) (strea[kpr\d ]*[\dko])`,       // Dell Streak
			`(?i)(le[- ]+pan)[- ]+(\w{1,9}) bui`,      // Le Pan Tablets
			`(?i)(trinity)[- ]*(t\d{3}) bui`,          // Trinity Tablets
			`(?i)(gigaset)[- ]+(q\w{1,9}) bui`,        // Gigaset Tablets
			`(?i)(vodafone) ([\w ]+)(?:\)| bui)`,      // Vodafone
		},
		Props: Props{
			Device: PropMap{
				Index: 2,
			},
			Vendor: PropMap{
				Index:   1,
				Reg:     " (.+)",
				Replace: "$1",
				Handle:  toLower{},
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(surface duo)`, // Surface Duo
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _MICROSOFT,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)droid [\d\.]+; (fp\du?)(?: b|\))`, // Fairphone
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _FAIRPHONE,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(u304aa)`, // AT&T
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: "at&t",
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\bsie-(\w*)`, // Siemens
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _SIEMENS,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(rct\w+) b`, // _RCA Tablets
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _RCA,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(venue[\d ]{2,7}) b`, // Dell Venue Tablets
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _DELL,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(q(?:mv|ta)\w+) b`, // Verizon Tablets
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _VERIZON,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(?:barnes[& ]+noble |bn[rt])([\w\+ ]*) b`, // Barnes & Noble Tablet
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: "barnes & noble",
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(tm\d{3}\w+) b`, // Barnes & Noble Tablet
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _NUVISION,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(k88) b`, // _ZTE K Series Tablet
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _ZTE,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(nx\d{3}j) b`, // _ZTE Nubia
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _ZTE,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(gen\d{3}) b.+49h`, // Swiss GEN Mobile
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _SWISS,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(zur\d{3}) b`, // Swiss ZUR Tablet
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _SWISS,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b((zeki)?tb.*\b) b`, // Zeki Tablets
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _ZEKI,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b([yr]\d{2}) b`,
			`(?i)\b(dragon[- ]+touch |dt)(\w{5}) b`, // Dragon Touch Tablet
		},
		Props: Props{
			Device: PropMap{
				Index: 2,
			},
			Vendor: PropMap{
				Value: "dragon touch",
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(ns-?\w{0,9}) b`, // Insignia Tablets
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _INSIGNIA,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b((nxa|next)-?\w{0,9}) b`, // NextBook Tablets
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _NEXTBOOK,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(xtreme_)?(v(1[045]|2[015]|[3469]0|7[05])) b`, // Voice Xtreme Phones
		},
		Props: Props{
			Device: PropMap{
				Index: 2,
			},
			Vendor: PropMap{
				Value: _VOICE,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(lvtel\-)?(v1[12]) b`, // LvTel Phones
		},
		Props: Props{
			Device: PropMap{
				Index: 2,
			},
			Vendor: PropMap{
				Value: _LVTEL,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(ph-1) `, // Essential PH-1
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _ESSENTIAL,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(v(100md|700na|7011|917g).*\b) b`, // Envizen Tablets
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _ENVIZEN,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(trio[-\w\. ]+) b`, // MachSpeed Tablets
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _MACHSPEED,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\btu_(1491) b`, // Rotor Tablets
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _ROTOR,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(shield[\w ]+) b`, // Nvidia Shield Tablets
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _NVIDIA,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(sprint) (\w+)`, // Sprint Phones
		},
		Props: Props{
			Device: PropMap{
				Index: 2,
			},
			Vendor: PropMap{
				Index: 1,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(kin\.[onetw]{3})`, // Microsoft Kin
		},
		Props: Props{
			Device: PropMap{
				Index:   1,
				Reg:     `(?i)(?g)\.`,
				Replace: " ",
				Handle:  toLower{},
			},
			Vendor: PropMap{
				Value: _MICROSOFT,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)droid.+; (cc6666?|et5[16]|mc[239][23]x?|vc8[03]x?)\)`, // Zebra
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _ZEBRA,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)droid.+; (ec30|ps20|tc[2-8]\d[kx])\)`, // Zebra
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _ZEBRA,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},

	///////////////////
	// CONSOLES
	///////////////////
	{
		Regs: []string{
			`(?i)(ouya)`,                    // Ouya
			`(?i)(nintendo) ([wids3utch]+)`, // Nintendo
		},
		Props: Props{
			Device: PropMap{
				Index: 2,
			},
			Vendor: PropMap{
				Index: 1,
			},
			Type: PropMap{
				Value: _CONSOLE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)droid.+; (shield) bui`, // Nvidia
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _NVIDIA,
			},
			Type: PropMap{
				Value: _CONSOLE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(playstation [345portablevi]+)`, // Playstation
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _SONY,
			},
			Type: PropMap{
				Value: _CONSOLE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(xbox(?: one)?(?!; xbox))[\); ]`, // Microsoft Xbox
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _MICROSOFT,
			},
			Type: PropMap{
				Value: _CONSOLE,
			},
		},
	},

	///////////////////
	// SMARTTVS
	///////////////////
	{
		Regs: []string{
			`(?i)hbbtv.+maple;(\d+)`, // Samsung
		},
		Props: Props{
			Device: PropMap{
				Index:   1,
				Reg:     "^",
				Replace: "smarttv",
			},
			Vendor: PropMap{
				Value: _SAMSUNG,
			},
			Type: PropMap{
				Value: _SMARTTV,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(apple) ?tv`, // Apple TV
		},
		Props: Props{
			Device: PropMap{
				Value: _APPLE + "TV",
			},
			Vendor: PropMap{
				Index: 1,
			},
			Type: PropMap{
				Value: _SMARTTV,
			},
		},
	},
	{
		Regs: []string{
			`(?i)crkey`, // Google Chromecast
		},
		Props: Props{
			Device: PropMap{
				Value: _CHROME + "cast",
			},
			Vendor: PropMap{
				Value: _GOOGLE,
			},
			Type: PropMap{
				Value: _SMARTTV,
			},
		},
	},
	{
		Regs: []string{
			`(?i)droid.+aft(\w)( bui|\))`, // Fire TV
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _AMAZON,
			},
			Type: PropMap{
				Value: _SMARTTV,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\(dtv[\);].+(aquos)`, // Sharp
			`(?i)(aquos-tv[\w ]+)\)`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _SHARP,
			},
			Type: PropMap{
				Value: _SMARTTV,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(bravia[\w ]+)( bui|\))`, // Sony
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _SONY,
			},
			Type: PropMap{
				Value: _SMARTTV,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(mitv-\w{5}) bui`, // Xiaomi
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _XIAOMI,
			},
			Type: PropMap{
				Value: _SMARTTV,
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(roku)[\dx]*[\)\/]((?:dvp-)?[\d\.]*)`,             // Roku
			`(?i)hbbtv\/\d+\.\d+\.\d+ +\([\w ]*; *(\w[^;]*);([^;]*)`, // HbbTV devices
		},
		Props: Props{
			Device: PropMap{
				Index:  2,
				Handle: trimMapper{},
			},
			Vendor: PropMap{
				Index:  1,
				Handle: trimMapper{},
			},
			Type: PropMap{
				Value: _SMARTTV,
			},
		},
	},

	///////////////////
	// WEARABLES
	///////////////////
	{
		Regs: []string{
			`(?i)((pebble))app`, // Pebble
			`(?i)hbbtv\/\d+\.\d+\.\d+ +\([\w ]*; *(\w[^;]*);([^;]*)`, // HbbTV devices
		},
		Props: Props{
			Device: PropMap{
				Index: 2,
			},
			Vendor: PropMap{
				Index: 1,
			},
			Type: PropMap{
				Value: _WEARABLE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)droid.+; (glass) \d`, // Google Glass
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _GOOGLE,
			},
			Type: PropMap{
				Value: _WEARABLE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)droid.+; (wt63?0{2,3})\)`,
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _ZEBRA,
			},
			Type: PropMap{
				Value: _WEARABLE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(quest( 2)?)`, // Oculus Quest
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _FACEBOOK,
			},
			Type: PropMap{
				Value: _WEARABLE,
			},
		},
	},

	////////////////////
	// MIXED (_GENERIC)
	///////////////////
	{
		Regs: []string{
			`(?i)droid .+?; ([^;]+?)(?: bui|\) applew).+? mobile safari`, // Android Phones from Unidentified Vendors
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Type: PropMap{
				Value: _MOBILE,
			},
		},
	},
	{
		Regs: []string{
			`(?i)droid .+?; ([^;]+?)(?: bui|\) applew).+?(?! mobile) safari`, // Android Tablets from Unidentified Vendors
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Type: PropMap{
				Value: _TABLET,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(android[-\w\. ]{0,9});.+buil`, // Unidentifiable Tablet
		},
		Props: Props{
			Device: PropMap{
				Index: 1,
			},
			Vendor: PropMap{
				Value: _GENERIC,
			},
		},
	},
}
