package useragent

import (
	"strings"
)

type Os struct {
	Name string
}

// DeviceConstants
var possibleOs = map[string]string{
	"microsoft":     "Windows (iTunes)",
	"windows nt":    "Windows RT",
	"windows phone": "Windows Phone",
	"bb":            "Blackberry 10",
	"blackberry":    "Blackberry",
	"tizen":         "Tizen",
	"kaios":         "KaiOS",
	// android|webos|palm\sos|qnx|bada|rim\stablet\sos|meego|sailfish|contiki)
	"android":       "Android",
	"webos":         "WebOS",
	"Palm":          "PalmOS",
	"qnx":           "QNX",
	"bada":          "BadaOS",
	"rim":           "RIM",
	"meego":         "MeeGo",
	"sailfish":      "SailfishOS",
	"contiki":       "ContikiOS",
	"symbian":       "SymbianOS",
	"symbos":        "SymbianOS",
	"s60":           "SymbianOS",
	"series40":      "Series 40 OS",
	"nintendo":      "Nintendo",
	"playstation":   "Playstation",
	"mint":          "Mint",
	"mageia":        "MageiaOS",
	"vectorlinux":   "VectorLinux",
	"hurd":          "Hurd",
	"linux":         "Linux",
	"gnu":           "GNU",
	"cros":          "ChromiumOS",
	"sunos":         "SolarisOS",
	"haiku":         "HaikuOS",
	"joli":          "Jolie",
	"ubuntu":        "Ubuntu",
	"debian":        "Debian",
	"suse":          "SUSE",
	"opensuse":      "SUSE",
	"gentoo":        "Gentoo",
	"arch":          "Arch",
	"slackware":     "Slackware",
	"fedora":        "Fedora",
	"mandriva":      "Mandriva",
	"centos":        "CentOS",
	"pclinuxos":     "PCLinuxOS",
	"redhat":        "RedHat",
	"zenwalk":       "ZenWalk",
	"linpus":        "Linpus",
	"dragonfly":     "DragonFly",
	"freebsd":       "FreeBSD",
	"netbsd":        "NetBSD",
	"openbsd":       "OpenBSD",
	"pc-bsd":        "PC-BSD",
	"cfnetwork":     "iOS",
	"ip":            "iOS",
	"mac":           "MacOS",
	"solaris":       "Solaris",
	"opensolaris":   "Solaris",
	"aix":           "AIX",
	"plan":          "Plan9",
	"minix":         "Minix",
	"beos":          "BeOS",
	"os2":           "OS2",
	"amigaos":       "AmigaOS",
	"morphos":       "MorphOS",
	"riscos":        "RISCOS",
	"openvms":       "OpenVMS",
	"fuchsia":       "Fuchsia",
	"unix":          "Unix",
	"windows 98":    "Windows 98",
	"windows 95":    "Windows 95",
	"windows me":    "Windows ME",
	"windows xp":    "Windows XP",
	"windows vista": "Windows Vista",
	"windows 7":     "Windows 7",
	"windows 8":     "Windows 8",
	"windows 8.1":   "Windows 8.1",
	"windows 10":    "Windows 10",
	"windows rt":    "Windows RT",
	"windows 2000":  "Windows 2000",
}

func (ua *UserAgent) guessOS() {
	os := Os{}

	for token, _ := range ua.clients {
		if val, ok := possibleOs[strings.ToLower(token)]; ok {
			os.Name = val
		}
	}

	ua.OS = os
}
