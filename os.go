package main

import (
	"fmt"
	"strings"
)

type Os struct {
	Name    string
	Version string
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
	"windows":       "Windows",
	"android":       "Android",
}

func (ua *UserAgent) guessOS() {
	os := Os{}

	osToken := ""
	for token, _ := range ua.clients {
		if val, ok := possibleOs[strings.ToLower(token)]; ok {
			os.Name = val
			os.Version = ua.clients[token]
			osToken = token
			break
		}
	}

	// if ua.exists(osToken) {
	// 	version := ua.clients[osToken]

	// 	os.Version = nVersion
	// }
	if osToken == "" {
		for token := range ua.clients {
			switch {
			case strings.HasPrefix(token, "CPU"):
				os.Name, os.Version = parseIOS(token)
			case strings.HasPrefix(token, "PPC"):
				os.Name, os.Version = parseMacOS(token)
			case ua.exists("windows"):
				os.Name, os.Version = parseWindows(ua)
			}
		}
	}
	// parse windows
	if osToken == "windows" || osToken == "windows nt" {
		os.Name, os.Version = parseWindows(ua)
	}
	ua.OS = os
}

func parseIOS(version string) (string, string) {
	// iPhone OS 5_1_1 like Mac OS X
	name := "MacOS"
	versionSlice := strings.Split(version, " ")
	if versionSlice[1] == "iPhone" || versionSlice[1] == "iPad" {
		name = "iOS"
	}
	newVersion := strings.ReplaceAll(versionSlice[3], "_", ".")
	return name, newVersion
}

func parseMacOS(version string) (string, string) {
	// PPC Mac OS X 10.5	name := "MacOS"
	name := "MacOS"
	versionSlice := strings.Split(version, " ")
	return name, versionSlice[4]
}

func parseWindows(ua *UserAgent) (string, string) {
	name := "Windows"
	version := ""
	windowsMapping := map[string]string{
		"4.90":   "ME",
		"NT3.51": "NT 3.11",
		"NT4.0":  "NT 4.0",
		"NT5.0":  "2000",
		"NT5.1":  "XP",
		"NT5.2":  "XP",
		"NT6.1":  "7",
		"NT6.2":  "8",
		"NT6.3":  "8.1",
		"NT6.4":  "10",
		"NT10.0": "10",
		"ARM":    "RT",
	}

	if ua.exists("windows nt") {
		ntString := fmt.Sprintf("NT%s", ua.getValue("windows nt"))
		version = windowsMapping[ntString]
	} else {
		version = "RT"
	}
	return name, version
}
