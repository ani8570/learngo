package proceed

import (
	"encoding/json"

	"github.com/ani8570/learngo/010.file/checkfunc"
)

// Enjson : encoding json
func Enjson(m checkfunc.Member) []byte {
	jsonBytes, err := json.Marshal(m)
	checkfunc.CheckErr(err)
	return jsonBytes
}

// Dejson : decoding json
func Dejson(jsonBytes []byte) checkfunc.Member {
	var m checkfunc.Member
	err := json.Unmarshal(jsonBytes, &m)
	checkfunc.CheckErr(err)
	return m
}

// EnXML : encoding XML
func EnXML() {

}

// DeXML : decoding XML
func DeXML() {

}
