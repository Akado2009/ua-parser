package main

import (
	"strings"
)

type Device struct {
	Name   string
	Model  string
	Mobile bool
	Tablet bool
}

// DeviceConstants
var possibleDevices = map[string]string{
	"ipad":            "iPad",
	"playbook":        "PlayBook",
	"applecoremedia":  "iPad",
	"appletv":         "Apple TV",
	"apple tv":        "Apple TV",
	"archos":          "Archos",
	"hp touchpad":     "HP TouchPad",
	"hp tablet":       "HP Tablet",
	"hp":              "HP",
	"kindle":          "Kindle",
	"nook":            "Nook",
	" nook":           "Nook",
	"dell strea":      "Dell Streak",
	"kf":              "Kindle Fire HD",
	"sd":              "Fire Phone",
	"iphone":          "iPhone",
	"ipod":            "iPod",
	"blackberry":      "Blackberry",
	"benq":            "BenQ",
	"sonyericsson":    "Sony Ericsson",
	"acer":            "Acer",
	"asus":            "Asus",
	"dell":            "Dell",
	"meizu":           "Meizu",
	"motorola":        "Motorola",
	"polytron":        "Polytron",
	"bb10":            "Blackberry 10",
	"sony":            "Sony",
	" ouya":           "Ouya",
	"ouya":            "Ouya",
	"nintendo":        "Nintendo",
	"playstation":     "Playstation",
	"sprint":          "Sprint Phones",
	"htc":             "HTC",
	"zte":             "ZTE",
	"huawei":          "Huawei",
	"nexus":           "Nexus",
	"vog-l29":         "Nexus",
	"ane-lx1":         "Nexus",
	"eml-l29":         "Nexus",
	"ele-l29":         "Nexus",
	"microsoft lumia": "Microsoft Lumia",
	"lumia":           "Microsoft Lumia",
	"xbox":            "Xbox",
	" xbox":           "Xbox",
	"kin":             "Microsoft Kin",
	"hbbtv":           "HbbTV",
	"dtv":             "Sharp",
	"smart-tv":        "Samsung SmartTV",
	"sie":             "Siemens",
	"maemo":           "Nokia",
	"nokia":           "Nokia",
	"lg":              "LG",
	"lenovo":          "Lenovo",
	"pebble":          "Pebble",
	"crkey":           "Google Chromecast",
	"mz":              "Meizu",
	"milestone":       "Motorola",
	"alcatel":         "Alcatel",
	"geeksphone":      "GeeksPhone",
	"nexian":          "Nexian",
	"panasonic":       "Panasonic",
	"nexus 9":         "Nexus 9",
	"samsung":         "Samsung",
}

var mobileDeviceTypes = map[string]bool{
	"ipad":            false,
	"playbook":        false,
	"applecoremedia":  false,
	"appletv":         false,
	"apple tv":        false,
	"archos":          false,
	"hp touchpad":     false,
	"hp tablet":       false,
	"hp":              false,
	"kindle":          false,
	"nook":            false,
	" nook":           false,
	"dell strea":      false,
	"kf":              false,
	"sd":              false,
	"iphone":          true,
	"ipod":            true,
	"blackberry":      true,
	"benq":            true,
	"sonyericsson":    true,
	"acer":            true,
	"asus":            true,
	"dell":            true,
	"meizu":           true,
	"motorola":        true,
	"polytron":        true,
	"bb10":            true,
	"sony":            false,
	" ouya":           false,
	"ouya":            false,
	"nintendo":        false,
	"playstation":     false,
	"sprint":          true,
	"htc":             true,
	"zte":             true,
	"huawei":          true,
	"nexus":           true,
	"vog-l29":         true,
	"ane-lx1":         true,
	"eml-l29":         true,
	"ele-l29":         true,
	"microsoft lumia": true,
	"lumia":           true,
	"xbox":            false,
	" xbox":           false,
	"kin":             false,
	"hbbtv":           false,
	"dtv":             false,
	"smart-tv":        false,
	"sie":             true,
	"maemo":           true,
	"nokia":           true,
	"lg":              true,
	"lenovo":          false,
	"pebble":          false,
	"crkey":           false,
	"mz":              true,
	"milestone":       true,
	"alcatel":         true,
	"geeksphone":      true,
	"nexian":          true,
	"panasonic":       true,
	"nexus 9":         true,
	"samsung":         true,
}

func (ua *UserAgent) guessDevice() {
	device := Device{}
	deviceToken := ""
	for token, _ := range ua.clients {
		if val, ok := possibleDevices[token]; ok {
			deviceToken = token
			device.Name = val
			device.Model = strings.ToUpper(ua.clients[deviceToken])
			device.Mobile = mobileDeviceTypes[token]
			device.Tablet = !device.Mobile
		}
	}
	if deviceToken == "" {
		// possibly android or palmsource
		if ua.exists("palmsource") {
			device.Name = "Palm"
			device.Model = parsePalm(ua.getValue("PalmSource"))
			device.Tablet = false
			device.Mobile = true
		}
		for token, _ := range ua.clients {
			if strings.HasPrefix("android", token) {
				device = parseAndroid(token, ua.clients[token])
			}
		}
	}
	if device.Name == "" {
		device.Tablet = false
		device.Mobile = false
	}
	ua.Device = device
}

func parsePalm(palmString string) string {
	s := strings.Split(palmString, "-")
	if len(s) == 2 {
		return s[1]
	}
	return ""
}

func parseAndroid(token, tokenValue string) Device {
	device := Device{}

	// switch {
	// 	case strings.Contains(token)
	// }
	return device
}
