package filekv

const Separator = ";;;"

type Options struct {
	Path    string
	Dedupe  bool
	Cleanup bool
}

type Stats struct {
	NumberOfAddedItems uint
	NumberOfDupedItems uint
	NumberOfItems      uint
}

var DefaultOptions Options = Options{
	Dedupe:  true,
	Cleanup: true,
}
