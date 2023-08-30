package edmx

import "github.com/qwerty-iot/tox"

type EntitySetNavigationPropertyBinding struct {
	Path   string `xml:"Path,attr"`
	Target string `xml:"Target,attr"`
}

type EntitySetAnnotation struct {
	Term       string `xml:"Term,attr"`
	Collection struct {
		PropertyPath string `xml:"PropertyPath"`
	} `xml:"Collection"`
}

type EntitySet struct {
	Name                      string                                `xml:"Name,attr"`
	EntityType                string                                `xml:"EntityType,attr"`
	NavigationPropertyBinding []*EntitySetNavigationPropertyBinding `xml:"NavigationPropertyBinding"`
	Annotation                []*EntitySetAnnotation                `xml:"Annotation"`
}

func (o *EntitySet) ToObject() tox.Object {
	r := tox.NewObject(nil)
	r.Set("kind", "EntitySet")
	r.Set("name", o.Name)
	r.Set("url", o.Name)
	return r
}

type Singleton struct {
	Name                      string                                `xml:"Name,attr"`
	Type                      string                                `xml:"Type,attr"`
	NavigationPropertyBinding []*EntitySetNavigationPropertyBinding `xml:"NavigationPropertyBinding"`
}

type FunctionImport struct {
	Name      string `xml:"Name,attr"`
	Function  string `xml:"Function,attr"`
	EntitySet string `xml:"EntitySet,attr"`
}

type ActionImport struct {
	Name   string `xml:"Name,attr"`
	Action string `xml:"Action,attr"`
}

type EntityContainer struct {
	m              *Model
	Name           string            `xml:"Name,attr"`
	EntitySet      []*EntitySet      `xml:"EntitySet"`
	Singleton      []*Singleton      `xml:"Singleton"`
	FunctionImport []*FunctionImport `xml:"FunctionImport"`
	ActionImport   []*ActionImport   `xml:"ActionImport"`
}

func (m *Model) AddEntityContainer(name string) *EntityContainer {
	o := &EntityContainer{m: m, Name: name}
	m.DataServices.Schema.EntityContainer = append(m.DataServices.Schema.EntityContainer, o)
	return o
}

func (o *EntityContainer) AddEntitySet(name string, entityType string) *EntitySet {
	e := &EntitySet{Name: name, EntityType: o.m.DataServices.Schema.Namespace + "." + entityType}
	o.EntitySet = append(o.EntitySet, e)
	return e
}
