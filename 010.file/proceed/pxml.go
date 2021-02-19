package proceed

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	"github.com/ani8570/learngo/010.file/f"
)

// EnXML : encoding XML
func EnXML(m f.Member) []byte {
	jsonBytes, err := xml.Marshal(m)
	f.CheckErr(err)
	return jsonBytes
}

// DeXML : decoding XML
func DeXML(jsonBytes []byte) f.Member {
	var m f.Member
	err := xml.Unmarshal(jsonBytes, &m)
	f.CheckErr(err)
	return m
}

// ReadXML : read XML
func ReadXML(s string) f.Members {
	fp, err := os.Open(s)
	f.CheckErr(err)
	defer fp.Close()

	data, err := ioutil.ReadAll(fp)

	var members f.Members
	xmlerr := xml.Unmarshal(data, &members)
	f.CheckErr(xmlerr)
	return members
}
