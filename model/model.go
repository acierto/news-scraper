package model


type InputElement struct {
	Source  string
	Find    string
	Link    []string
	Title   []string
	Charset string
}

func (self *InputElement) Initialize() { if self.Charset == "" { self.Charset = "UTF-8" } }

type Article struct {
	Source string
	Link  string
	Title string
}
