package dbuaparser

type EngineInfo struct {
	Engine    string // 引擎
	EngineVer string // 引擎版本
}

var engineMap = []RegexpMap{
	{
		Regs: []string{
			`(?i)windows.+ edge\/([\w\.]+)`, // EdgeHTML
		},
		Props: Props{
			Engine: PropMap{
				Value: _EDGE + "html",
			},
			EngineVer: PropMap{
				Index: 1,
			},
		},
	},
	{
		Regs: []string{
			`(?i)webkit\/537\.36.+chrome\/(?!27)([\w\.]+)`, // Blink
		},
		Props: Props{
			Engine: PropMap{
				Value: "blink",
			},
			EngineVer: PropMap{
				Index: 1,
			},
		},
	},
	{
		Regs: []string{
			`(?i)(presto)\/([\w\.]+)`, // Presto
			`(?i)(webkit|trident|netfront|netsurf|amaya|lynx|w3m|goanna)\/([\w\.]+)`, // WebKit/Trident/NetFront/NetSurf/Amaya/Lynx/w3m/Goanna
			`(?i)ekioh(flow)\/([\w\.]+)`,                // Flow
			`(?i)(khtml|tasman|links)[\/ ]\(?([\w\.]+)`, // KHTML/Tasman/Links
			`(?i)(icab)[\/ ]([23]\.[\d\.]+)`,            // iCab
		},
		Props: Props{
			Engine: PropMap{
				Index: 1,
			},
			EngineVer: PropMap{
				Index: 2,
			},
		},
	},
	{
		Regs: []string{
			`(?i)rv\:([\w\.]{1,9})\b.+(gecko)`, // Gecko
		},
		Props: Props{
			Engine: PropMap{
				Index: 2,
			},
			EngineVer: PropMap{
				Index: 1,
			},
		},
	},
}
