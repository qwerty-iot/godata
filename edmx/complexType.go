package edmx

type ComplexType struct {
	m        *Model
	Name     string                `xml:"Name,attr"`
	BaseType string                `xml:"BaseType,attr,omitempty"`
	OpenType string                `xml:"OpenType,attr,omitempty"`
	Property []*EntityTypeProperty `xml:"Property"`
}

func (m *Model) AddComplexType(name string, base string, options *Options) *ComplexType {
	if len(base) != 0 {
		base = m.DataServices.Schema.Namespace + "." + base
	}
	o := &ComplexType{m: m, Name: name, BaseType: base, OpenType: options.getOpen()}
	m.DataServices.Schema.ComplexType = append(m.DataServices.Schema.ComplexType, o)
	return o
}

func (o *ComplexType) AddProperty(name string, edmType EdmType, options *Options) {
	o.Property = append(o.Property, &EntityTypeProperty{
		Name:      name,
		Type:      edmType,
		Nullable:  options.getNullable(),
		MaxLength: options.getMaxLength(),
	})
	return
}
