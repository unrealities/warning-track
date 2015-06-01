package models

type baseRunners struct {
	First  bool
	Second bool
	Third  bool
}

type BaseOut struct {
	Outs         int
	Base_Runners baseRunners
}
