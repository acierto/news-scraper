package model

type InputElement struct {
	Source         string
	Find           string
	Link           []string
	Title          []string
	FindImg        []string
	Charset        string
}

func (self *InputElement) Initialize() { if self.Charset == "" { self.Charset = "UTF-8" } }

type Article struct {
	Link      string
	Title     string
	Img       string
}

type SourceArticle struct {
	Source   string
	Articles []Article
}
