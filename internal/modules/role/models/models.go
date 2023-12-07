package models

type Permissions struct {
	CreateNews   bool `json:"create_news"`
	CreateEvents bool `json:"create_events"`
	CreatePosts  bool `json:"create_posts"`
}
