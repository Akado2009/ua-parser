package useragent

import "testing"

func TestParse(t *testing.T) {
	var testTable = map[string]UserAgent{
		"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.3) Gecko/2008092921 IceCat/3.0.3-g1":                                                                     createUserAgent("IceCat", 3, 0, "Linux", "", "", false, false),
		"Mozilla/5.0 (iPhone; U; CPU iPhone OS 5_1_1 like Mac OS X; en) AppleWebKit/534.46.0 (KHTML, like Gecko) CriOS/19.0.1084.60 Mobile/9B206 Safari/7534.48.3": createUserAgent("Chrome", 19, 0, "MacOS", "iPhone", "", true, false),
		"Dillo/2.2": createUserAgent("Dillo", 2, 2, "", "", "", false, false),
		"Mozilla/5.0 (X11; U; FreeBSD i386; en-US; rv:1.7) Gecko/20040628 Epiphany/1.2.6":                 createUserAgent("Epiphany", 1, 2, "FreeBSD", "", "", false, false),
		"Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10.5; en-US; rv:1.9.0.3) Gecko/2008092414 Firefox/3.0.3": createUserAgent("Firefox", 3, 0, "MacOS", "", "", false, false),
	}

	for test, result := range testTable {
		ua := NewUAParser(test)
		ua.Parse()
		if ua.Browser.Name != result.Browser.Name {
			t.Errorf("error parsing browser name, expected: %s, got: %s", result.Browser.Name, ua.Browser.Name)
		}
		if ua.Browser.Major != result.Browser.Major {
			t.Errorf("error parsing browser major, expected: %d, got: %d", result.Browser.Major, ua.Browser.Major)
		}
		if ua.Browser.Minor != result.Browser.Minor {
			t.Errorf("error parsing browser minor, expected: %d, got: %d", result.Browser.Minor, ua.Browser.Minor)
		}
		if ua.OS.Name != result.OS.Name {
			t.Errorf("error parsing os name, expected: %s, got: %s", result.OS.Name, ua.OS.Name)
		}
		if ua.Device.Name != result.Device.Name {
			t.Errorf("error parsing device name, expected: %s, got: %s", result.Device.Name, ua.Device.Name)
		}
		if ua.Device.Model != result.Device.Model {
			t.Errorf("error parsing device model, expected: %s, got: %s", result.Device.Model, ua.Device.Model)
		}
		if ua.Device.Mobile != result.Device.Mobile {
			t.Errorf("error parsing device mobile, expected: %t, got: %t", result.Device.Mobile, ua.Device.Mobile)
		}
		if ua.Device.Tablet != result.Device.Tablet {
			t.Errorf("error parsing device tablet, expected: %t, got: %t", result.Device.Tablet, ua.Device.Tablet)
		}
	}

}

func createUserAgent(brName string, brMajor, brMinor int32, osName string, devName, devModel string, mobile, tablet bool) UserAgent {
	return UserAgent{
		Browser: Browser{
			Name:  brName,
			Major: brMajor,
			Minor: brMinor,
		},
		OS: Os{
			Name: osName,
		},
		Device: Device{
			Name:   devName,
			Model:  devModel,
			Mobile: mobile,
			Tablet: tablet,
		},
	}
}
