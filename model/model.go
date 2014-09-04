package model

import "time"

type InputElement struct {
	Source         string
	Find           string
	Link           []string
	Title          []string
	Time		   []string
	TimeZone	   int
	Charset        string
	ContentSelector string
}

func (self *InputElement) Initialize() { if self.Charset == "" { self.Charset = "UTF-8" } }

type Article struct {
	Source   string
	ContentSelector string
	Link      string
	Title     string
	Time	  time.Time
	Charset	  string
}
