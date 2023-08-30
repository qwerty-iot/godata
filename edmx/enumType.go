package edmx

type EnumTypeMember struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

type EnumType struct {
	Name   string           `xml:"Name,attr"`
	Member []EnumTypeMember `xml:"Member"`
}

func (m *Model) AddEnumType(name string) *EnumType {
	o := &EnumType{Name: name}
	m.DataServices.Schema.EnumType = append(m.DataServices.Schema.EnumType, o)
	return o
}

func (o *EnumType) AddMember(name string, value string) {
	o.Member = append(o.Member, EnumTypeMember{
		Name:  name,
		Value: value,
	})
	return
}
