package dbuaparser

type CpuInfo struct {
	Cpu string // CPU架构
}

var cpuMap = []RegexpMap{
	{
		Regs: []string{
			`(?i)(?:(amd|x(?:(?:86|64)[-_])?|wow|win)64)[;\)]`, // AMD64 (x64)
		},
		Props: Props{
			Cpu: PropMap{
				Index: 1,
				Value: "amd64",
			},
		},
	},
	{
		Regs: []string{
			`(?i)(ia32(?=;))`, // IA32 (quicktime)
		},
		Props: Props{
			Cpu: PropMap{
				Index:  1,
				Handle: toLower{},
			},
		},
	},
	{
		Regs: []string{
			`(?i)((?:i[346]|x)86)[;\)]`, // IA32 (x86)
		},
		Props: Props{
			Cpu: PropMap{
				Index: 1,
				Value: "ia32",
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(aarch64|arm(v?8e?l?|_?64))\b`, // ARM64
		},
		Props: Props{
			Cpu: PropMap{
				Index: 1,
				Value: "arm64",
			},
		},
	},
	{
		Regs: []string{
			`(?i)\b(arm(?:v[67])?ht?n?[fl]p?)\b`, // ARMHF
		},
		Props: Props{
			Cpu: PropMap{
				Index: 1,
				Value: "armhf",
			},
		},
	},
	{
		Regs: []string{
			`(?i)windows (ce|mobile); ppc;`,
		},
		Props: Props{
			Cpu: PropMap{
				Index: 1,
				Value: "arm",
			},
		},
	},
	{
		Regs: []string{
			`(?i)((?:ppc|powerpc)(?:64)?)(?: mac|;|\))`, // PowerPC
		},
		Props: Props{
			Cpu: PropMap{
				Index:   1,
				Reg:     "(?i)power",
				Replace: "p",
				Handle:  toLower{},
			},
		},
	},
	{
		Regs: []string{
			`(?i)(sun4\w)[;\)]`, // SPARC
		},
		Props: Props{
			Cpu: PropMap{
				Index: 1,
				Value: "sparc",
			},
		},
	},
	{
		Regs: []string{
			`(?i)((?:avr32|ia64(?=;))|68k(?=\))|\barm(?=v(?:[1-7]|[5-7]1)l?|;|eabi)|(?=atmel )avr|(?:irix|mips|sparc)(?:64)?\b|pa-risc)`, // IA64, 68K, ARM/64, AVR/32, IRIX/64, MIPS/64, SPARC/64, PA-RISC
		},
		Props: Props{
			Cpu: PropMap{
				Index:  1,
				Handle: toLower{},
			},
		},
	},
}
