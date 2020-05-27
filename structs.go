package main

type post struct {
	author  string
	content string
}

type thread struct {
	tid      int64
	URL      string
	title    string
	authorID int64
	isGood   bool // nice tanslation there Baidu
	posts    []post
}

type threadDataField struct {
	Tid            int64  `json:"id"`
	IsGood         bool   `json:"is_good"`
	AuthorName     string `json:"author_name"`
	AuthorNickname string `json:"author_nickname"`
}

type user struct {
	name     string
	nickname string
	uid      string
}

type authorIconDataField struct {
	UID int64 `json:"user_id"`
}
