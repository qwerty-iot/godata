package edmx

import (
	"encoding/xml"
	"github.com/qwerty-iot/godata/result"
)

type Function struct {
	Text          string `xml:",chardata"`
	Name          string `xml:"Name,attr"`
	IsBound       string `xml:"IsBound,attr"`
	EntitySetPath string `xml:"EntitySetPath,attr"`
	ReturnType    struct {
		Text string `xml:",chardata"`
		Type string `xml:"Type,attr"`
	} `xml:"ReturnType"`
	Parameter []struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"Name,attr"`
		Type     string `xml:"Type,attr"`
		Nullable string `xml:"Nullable,attr"`
		Unicode  string `xml:"Unicode,attr"`
	} `xml:"Parameter"`
}

type Action struct {
	Text      string `xml:",chardata"`
	Name      string `xml:"Name,attr"`
	IsBound   string `xml:"IsBound,attr"`
	Parameter []struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"Name,attr"`
		Type     string `xml:"Type,attr"`
		Nullable string `xml:"Nullable,attr"`
		Unicode  string `xml:"Unicode,attr"`
	} `xml:"Parameter"`
	ReturnType struct {
		Text     string `xml:",chardata"`
		Type     string `xml:"Type,attr"`
		Nullable string `xml:"Nullable,attr"`
	} `xml:"ReturnType"`
}

type Schema struct {
	Namespace       string             `xml:"Namespace,attr"`
	Xmlns           string             `xml:"xmlns,attr"`
	EntityType      []*EntityType      `xml:"EntityType"`
	ComplexType     []*ComplexType     `xml:"ComplexType"`
	EnumType        []*EnumType        `xml:"EnumType"`
	Function        []*Function        `xml:"Function"`
	Action          []*Action          `xml:"Action"`
	EntityContainer []*EntityContainer `xml:"EntityContainer"`
}

type DataServices struct {
	Schema Schema `xml:"Schema"`
}

type Model struct {
	context      string
	XMLName      xml.Name     `xml:"edmx:Edmx"`
	Version      string       `xml:"Version,attr"`
	Edmx         string       `xml:"xmlns:edmx,attr"`
	DataServices DataServices `xml:"edmx:DataServices"`
}

func New(context string, namespace string) *Model {
	return &Model{
		context: context,
		Version: "4.0",
		Edmx:    "http://docs.oasis-open.org/odata/ns/edmx",
		DataServices: DataServices{
			Schema: Schema{
				Namespace: namespace,
				Xmlns:     "http://docs.oasis-open.org/odata/ns/edm",
			},
		},
	}
}

func (m *Model) ToXml() ([]byte, error) {
	return xml.Marshal(m)
}

func (m *Model) ToSimpleJson(scope string) ([]byte, error) {
	ret := result.New(m.context + "/$metadata")
	for _, ec := range m.DataServices.Schema.EntityContainer {
		for _, es := range ec.EntitySet {
			ret.Add(es.ToObject())
		}
	}
	return ret.ToJson()
}
