package proceed

import (
	"encoding/json"

	"github.com/ani8570/learngo/010.file/f"
)

// Enjson : encoding json
func Enjson(m f.Member) []byte {
	jsonBytes, err := json.Marshal(m)
	f.CheckErr(err)
	return jsonBytes
}

// Dejson : decoding json
func Dejson(jsonBytes []byte) f.Member {
	var m f.Member
	err := json.Unmarshal(jsonBytes, &m)
	f.CheckErr(err)
	return m
}
