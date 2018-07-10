package models

type ModelStruct struct {
	ModelTitle string
	KeyTypes   []KeyType
}

type KeyType struct {
	Key  string
	Type string
}
