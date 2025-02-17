package util

import "go-crack/types"

var SelectedOptions = types.Options{
	IsPooled:          false,
	TotalThreads:      4,
	MinimumCharacters: 1,
	MaximumCharacters: 8,
}

var Store = types.StoreStruct{
	CompletedOps: 0,
}
