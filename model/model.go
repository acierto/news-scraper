package model

import "time"

type InputElement struct {
	Source         string
	Find           string
	Link           []string
	Title          []string
	Time		   []string
	Charset        string
	ContentSelector string
}

func (self *InputElement) Initialize() { if self.Charset == "" { self.Charset = "UTF-8" } }

type Article struct {
	Link      string
	Title     string
	Time	  time.Time
}

type SourceArticle struct {
	Source   string
	ContentSelector string
	Articles []Article
}
