package structs

type RegKeyValue struct {
	KeyName string
	Value   string
}

type RegKeyValues struct {
	Values []RegKeyValue
}
