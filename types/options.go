package types

type Options struct {
	IsPooled          bool
	IsLinear          bool
	TotalThreads      int
	FileLocation      string
	Strategy          string
	AvailableBinaries []string
	MinimumCharacters int
	MaximumCharacters int
}
