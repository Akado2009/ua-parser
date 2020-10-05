package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Browser struct {
	Name  string
	Major int32
	Minor int32
}

// BrowserConstants

var ignoredBrowser = map[string]string{
	"mozilla": "Mozilla",
}
var possibleBrowsers = map[string]string{
	"crios":        "Chrome",
	"opera":        "Opera",
	"opera mini":   "Opera Mini",
	"opios":        "Opera Mini",
	"opr":          "Opera Webkit",
	"kindle":       "Kindle",
	"lunascape":    "LunaScape",
	"maxthon":      "Maxthon",
	"netfron":      "NetFront",
	"jasmine":      "Jasmine",
	"blazer":       "Blazer",
	"avant":        "Avant",
	"iemobile":     "IEMobile",
	"slim":         "SlimBrowser",
	"bidubrowser":  "Baidu Browser",
	"baidubrowser": "Baidu Browser",
	"rekonq":       "Rekonq",
	"chromium":     "Chromium",
	"flock":        "Flock",
	"rockmelt":     "RockMelt",
	"midori":       "Midori",
	"epiphany":     "Epiphany",
	"silk":         "Silk",
	"skyfire":      "Skyfire",
	"bolt":         "Bolt",
	"iron":         "Iron",
	"vivaldi":      "Vivaldi",
	"iridium":      "Iridium",
	"phantomjs":    "PhantomJS",
	"bowser":       "Bowser",
	"qupzilla":     "QupZilla",
	"falkon":       "Falkon",
	"konqueror":    "Konqueror",
	"trident":      "IE11",
	// edge|edgios|edga|edg
	"edge":                 "Microsoft Edge",
	"edgios":               "Microsoft Edge",
	"edga":                 "Microsoft Edge",
	"edg":                  "Microsoft Edge",
	"yabrowser":            "Yandex",
	"avast":                "Avast Secure Browser",
	"avg":                  "AVG Secure Browser",
	"puffin":               "Puffin",
	"focus":                "Firefox Focus",
	"opt":                  " Opera Touch",
	"browser":              "UCBrowser",
	"comodo_dragon":        "Comodo Dragon",
	"windowswechat qbcore": "WeChat Desktop for Windows Built-in Browse",
	"micromessenger":       "WeChat",
	"brave":                "Brave browser",
	"qqbrowserlite":        "QQBrowserLite",
	"qq":                   "QQ, aka ShouQ",
	"qqbrowser":            "QQBrowser",
	"mqqbrowser":           "QQBrowser",
	"baiduboxapp":          "Baidu App",
	"2345explorer":         "2345 Browser",
	"metasr":               "SouGouBrowser",
	"lbbrowser":            "LieBao Browser",
	"xiaomi":               "MIUI Browser",
	"fbav":                 "Facebook App for iOS & Android",
	"headlesschrome":       "Chrome Headless",
	"oculusbrowser":        "Oculus Browser",
	"sailfishbrowser":      "Sailfish Browser",
	"chrome":               "Chrome",
	"omniweb":              "OmniWeb",
	"arora":                "Arora",
	"tizenoka":             "Tizenoka",
	"dolfin":               "Dolphin",
	"samsungbrowser":       "Samsung Browser",
	"netscape":             "Netscape",
	"qihu":                 "360",
	"coast":                "Opera Coast",
	"fxios":                "Firefox iOS",
	"qhbrowser":            "360",
	"qihoobrowser":         "360",
	"360browser":           "360",
	"icedragon":            "IceDragon",
	"iceweasel":            "Iceweasel",
	"camino":               "Camino",
	"chimera":              "Chimera",
	"fennec":               "Fennec",
	"maemo":                "Maemo",
	"minimo":               "Minimo",
	"conkeror":             "Conkeror",
	"firefox":              "Firefox",
	"seamonkey":            "SeaMonkey",
	"k-meleon":             "K-Meleon",
	"icecat":               "IceCat",
	"iceape":               "IceApe",
	"firebird":             "Firebird",
	"phoenix":              "Phoneix",
	"palemoon":             "Palemoon",
	"basilisk":             "Basilisk",
	"waterfox":             "Waterfox",
	"polaris":              "Polaris",
	"lynx":                 "Lynx",
	"dillo":                "Dillo",
	"icab":                 "Icab",
	"doris":                "Doris",
	"amaya":                "Amaya",
	"w3m":                  "W3M",
	"netsurf":              "Netsurf",
	"sleipnir":             "Sleipnir",
	"links":                "Links",
	"gobrowser":            "GoBrowser",
	"ice":                  "ICE Browser",
	"mosaic":               "Mosaic",
	"android":              "Android Browser",
}

func (ua *UserAgent) guessBrowser() {
	browser := Browser{}

	browserToken := ""
	for token, _ := range ua.clients {
		if val, ok := possibleBrowsers[token]; ok {
			browserToken = token
			browser.Name = val
		}
	}
	if browserToken == "" {
		for token, _ := range ua.clients {
			if val, ok := ignoredBrowser[token]; ok {
				browserToken = token
				browser.Name = val
			}
		}
	}
	// check whether there is a version field
	if ua.exists("version") {
		version := strings.ReplaceAll(ua.clients["version"], "/", ".") // replace / to a dot to parse properly
		major, minor := parseVersion(version)
		browser.Major = major
		browser.Minor = minor
	}

	// check if there is a version nearby
	if ua.exists(browserToken) {
		version := ua.clients[browserToken]
		major, minor := parseVersion(version)
		browser.Major = major
		browser.Minor = minor
	}

	ua.Browser = browser
}

func parseVersion(version string) (int32, int32) {
	versionSlice := strings.Split(version, ".")
	major, err := strconv.Atoi(versionSlice[0])
	if err != nil {
		fmt.Println("Error parsing major: ", err)
		return int32(0), int32(0)
	}
	minor, err := strconv.Atoi(string(versionSlice[1][0]))
	if err != nil {
		fmt.Println("Error parsing minor: ", err)
		return int32(0), int32(0)
	}
	return int32(major), int32(minor)
}
