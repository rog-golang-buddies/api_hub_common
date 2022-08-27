package apispecdoc

// Type represents type of parsed API
type Type string

const (
	TypeOpenApi Type = "OPEN_API"
)

// SchemaType is the type of field.
type SchemaType string

const (
	Unknown    SchemaType = "UNKNOWN"
	NotDefined SchemaType = "NOT_DEFINED"
	Integer    SchemaType = "INTEGER"
	Boolean    SchemaType = "BOOLEAN"
	Number     SchemaType = "NUMBER"
	String     SchemaType = "STRING"
	Date       SchemaType = "DATE"
	Array      SchemaType = "ARRAY"  //Array type may contain a single nested field with a type to define the full array type
	Map        SchemaType = "MAP"    //Map i.e. { "map_name": {"key1":"value1", "key2":"value2"}}
	OneOf      SchemaType = "ONE_OF" //OneOf is one of the different types (C union) (from nested fields)
	AnyOf      SchemaType = "ANY_OF" //AnyOf defines that the result object can contain any set of sub-schemes
	AllOf      SchemaType = "ALL_OF" //AllOf defines that the result object combines all listed objects/properties.
	Not        SchemaType = "NOT"    //Not represents a type that can't be used. So it's possible to use any of the types except "not"
	Object     SchemaType = "OBJECT" //Object represents the object and contains a set of fields inside
)

func ResolveSchemaType(strType string) SchemaType {
	switch strType {
	case "string":
		return String
	case "number":
		return Number
	case "integer":
		return Integer
	case "boolean":
		return Boolean
	case "array":
		return Array
	case "object":
		return Object
	case "":
		return NotDefined
	default:
		return Unknown
	}
}

// MethodType define type of action of the method
type MethodType string

const (
	MethodConnect MethodType = "CONNECT"
	MethodGet     MethodType = "GET"
	MethodPut     MethodType = "PUT"
	MethodPost    MethodType = "POST"
	MethodDelete  MethodType = "DELETE"
	MethodOptions MethodType = "OPTIONS"
	MethodHead    MethodType = "HEAD"
	MethodPatch   MethodType = "PATCH"
	MethodTrace   MethodType = "TRACE"
	//Unary/bidirectional stream etc. for the rpc
)

// ParameterType represents to what part of request parameter relates
type ParameterType string

const (
	ParameterQuery  ParameterType = "query"
	ParameterHeader ParameterType = "header"
	ParameterPath   ParameterType = "path"
	ParameterCookie ParameterType = "cookie"
)
