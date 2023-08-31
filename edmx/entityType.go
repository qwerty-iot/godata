package edmx

type EntityTypeProperty struct {
	Name      string  `xml:"Name,attr"`
	Type      EdmType `xml:"Type,attr"`
	Nullable  string  `xml:"Nullable,attr,omitempty"`
	MaxLength string  `xml:"MaxLength,attr,omitempty"`
	Partner   string  `xml:"Partner,attr,omitempty"`
}

type EntityType struct {
	m        *Model
	Name     string `xml:"Name,attr"`
	BaseType string `xml:"BaseType,attr,omitempty"`
	OpenType string `xml:"OpenType,attr,omitempty"`
	Key      struct {
		PropertyRef struct {
			Name string `xml:"Name,attr"`
		} `xml:"PropertyRef"`
	} `xml:"Key"`
	Property           []*EntityTypeProperty `xml:"Property"`
	NavigationProperty []*EntityTypeProperty `xml:"NavigationProperty"`
}

func (m *Model) AddEntityType(name string, options *Options) *EntityType {
	o := &EntityType{m: m, Name: name, OpenType: options.getOpen()}
	m.DataServices.Schema.EntityType = append(m.DataServices.Schema.EntityType, o)
	return o
}

func (o *EntityType) SetKey(name string) {
	o.Key.PropertyRef.Name = name
	return
}

func (o *EntityType) AddProperty(name string, edmType EdmType, options *Options) {
	o.Property = append(o.Property, &EntityTypeProperty{
		Name:      name,
		Type:      edmType,
		Nullable:  options.getNullable(),
		MaxLength: options.getMaxLength(),
	})
	return
}

func (o *EntityType) AddNavigationProperty(name string, entity string, collection bool, partner string) {
	if collection {
		entity = "Collection(" + o.m.DataServices.Schema.Namespace + "." + entity + ")"
	} else {
		entity = o.m.DataServices.Schema.Namespace + "." + entity
	}
	o.NavigationProperty = append(o.NavigationProperty, &EntityTypeProperty{
		Name:    name,
		Type:    EdmType(entity),
		Partner: partner,
	})
	return
}
