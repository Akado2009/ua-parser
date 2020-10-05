package useragent

import (
	"bytes"
	"strings"
)

// UserAgent ...
type UserAgent struct {
	Browser  Browser
	OS       Os
	Device   Device
	clients  map[string]string
	uaString string
}

var ignore = map[string]struct{}{
	"KHTML, like Gecko": struct{}{},
	"U":                 struct{}{},
	"compatible":        struct{}{},
	"WOW64":             struct{}{},
}

func checkVer(s string) (name, v string) {
	i := strings.LastIndex(s, " ")
	if i == -1 {
		return s, ""
	}
	switch s[:i] {
	case "Linux", "Windows NT", "Windows Phone OS", "MSIE", "Android":
		return s[:i], s[i+1:]
	default:
		return s, ""
	}
}

func NewUAParser(ua string) *UserAgent {
	return &UserAgent{
		uaString: ua,
	}
}

func (ua *UserAgent) Parse() {
	clients := make(map[string]string, 0)
	slash := false
	isURL := false
	var buff, val bytes.Buffer
	addToken := func() {
		if buff.Len() != 0 {
			s := strings.TrimSpace(buff.String())
			if _, ign := ignore[s]; !ign {
				if isURL {
					s = strings.TrimPrefix(s, "+")
				}
				if val.Len() == 0 { // only if value don't exists
					var ver string
					s, ver = checkVer(s) // determin version string and split
					clients[s] = ver
				} else {
					clients[s] = strings.TrimSpace(val.String())
				}
			}
		}
		buff.Reset()
		val.Reset()
		slash = false
		isURL = false
	}

	parOpen := false

	bua := []byte(ua.uaString)
	for i, c := range bua {
		switch {
		case c == 41: // )
			addToken()
			parOpen = false
		case parOpen && c == 59: // ;
			addToken()
		case c == 40: // (
			addToken()
			parOpen = true
		case slash && c == 32:
			addToken()
		case slash:
			val.WriteByte(c)
		case c == 47 && !isURL: //   /
			if i != len(bua)-1 && bua[i+1] == 47 && (bytes.HasSuffix(buff.Bytes(), []byte("http:")) || bytes.HasSuffix(buff.Bytes(), []byte("https:"))) {
				buff.WriteByte(c)
				isURL = true
			} else {
				slash = true
			}
		case c == 32:
			addToken()
		default:
			buff.WriteByte(c)
		}
	}
	addToken()

	ua.clients = clients
	ua.postParseClients()

	ua.guessBrowser()
	ua.guessOS()
	ua.guessDevice()

}

func (ua *UserAgent) postParseClients() {
	for key, value := range ua.clients {
		newKey, newValue := ua.postParseToken(key) //ignoring value)
		if newKey != "" {
			ua.clients[strings.ToLower(strings.TrimSpace(newKey))] = strings.ToLower(newValue)
		}
		ua.clients[strings.ToLower(key)] = strings.ToLower(value)

	}
}

func (ua *UserAgent) postParseToken(token string) (k, v string) {
	separatorIndex := 0
	strBytes := []byte(token)
	for i := 0; i < len(strBytes)-1; i++ {
		if isCharOrCapital(strBytes[i]) && !isCharOrCapital(strBytes[i+1]) {
			separatorIndex = i
			break
		}
	}
	if separatorIndex != 0 {
		return token[0:separatorIndex], token[separatorIndex+1 : len(token)]
	}
	return "", ""
}

func (ua *UserAgent) exists(token string) bool {
	if _, ok := ua.clients[token]; !ok {
		return false
	}
	return true
}

func (ua *UserAgent) getValue(token string) string {
	val, _ := ua.clients[token]
	return val
}

func isCharOrCapital(c byte) bool {
	if c <= '9' && c >= '0' {
		return false
	}
	if c <= 'Z' && c >= 'A' {
		return false
	}
	return true
}
