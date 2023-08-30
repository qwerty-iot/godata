package edmx

type EdmType string

const (
	TypeString   EdmType = "Edm.String"
	TypeInt32    EdmType = "Edm.Int32"
	TypeInt64    EdmType = "Edm.Int64"
	TypeFloat    EdmType = "Edm.Float"
	TypeDouble   EdmType = "Edm.Double"
	TypeBool     EdmType = "Edm.Boolean"
	TypeTime     EdmType = "Edm.DateTimeOffset"
	TypeBinary   EdmType = "Edm.Binary"
	TypeGeoPoint EdmType = "Edm.GeographyPoint"
)

func (m *Model) TypeCollectionOf(t EdmType) EdmType {
	return "Collection(" + t + ")"
}

func (m *Model) TypeEntity(entityName string) EdmType {
	return EdmType(m.DataServices.Schema.Namespace + "." + entityName)
}
