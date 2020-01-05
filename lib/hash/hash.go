package hash

import (
	"crypto/sha1"
	"encoding/json"
)

func Sha1(input interface{}) string {
	buf, _ := json.Marshal(input)
	h := sha1.New()
	h.Write(buf)

	return string(h.Sum(nil))
}
