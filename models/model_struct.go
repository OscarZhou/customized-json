package models

type ModelStruct struct {
	ModelName string
	KeyTypes  []KeyType
}

type KeyType struct {
	Key  string
	Type string
}
