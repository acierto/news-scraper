package model

type InputElement struct {
	Source         string
	Find           string
	Link           []string
	Title          []string
	Charset        string
	ContentSelector string
}

func (self *InputElement) Initialize() { if self.Charset == "" { self.Charset = "UTF-8" } }

type Article struct {
	Link      string
	Title     string
}

type SourceArticle struct {
	Source   string
	ContentSelector string
	Articles []Article
}
