package etherscan

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

var (
	Url = "https://api.etherscan.io/api?"
)

func ParseUrl(url string, params... string) string {
	var u strings.Builder
	u.WriteString(url)
	for k, v := range params {
		u.WriteString(v)
		if k % 2 == 0 {
			u.WriteByte('=')
		} else {
			u.WriteRune('&')
		}
	}
	return u.String()
}

// Http Get Request
func get(url string, v any) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}

	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, v)
	
}