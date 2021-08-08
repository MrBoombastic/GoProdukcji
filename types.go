package main

import "time"

type Post struct {
	ID                     string      `json:"id"`
	UUID                   string      `json:"uuid"`
	Title                  string      `json:"title"`
	Slug                   string      `json:"slug"`
	HTML                   string      `json:"html"`
	CommentID              string      `json:"comment_id"`
	FeatureImage           string      `json:"feature_image"`
	Featured               bool        `json:"featured"`
	Visibility             string      `json:"visibility"`
	EmailRecipientFilter   string      `json:"email_recipient_filter"`
	CreatedAt              time.Time   `json:"created_at"`
	UpdatedAt              time.Time   `json:"updated_at"`
	PublishedAt            time.Time   `json:"published_at"`
	CustomExcerpt          interface{} `json:"custom_excerpt"`
	CodeinjectionHead      interface{} `json:"codeinjection_head"`
	CodeinjectionFoot      interface{} `json:"codeinjection_foot"`
	CustomTemplate         interface{} `json:"custom_template"`
	CanonicalURL           interface{} `json:"canonical_url"`
	URL                    string      `json:"url"`
	Excerpt                string      `json:"excerpt"`
	ReadingTime            int         `json:"reading_time"`
	Access                 bool        `json:"access"`
	SendEmailWhenPublished bool        `json:"send_email_when_published"`
	OgImage                interface{} `json:"og_image"`
	OgTitle                interface{} `json:"og_title"`
	OgDescription          interface{} `json:"og_description"`
	TwitterImage           interface{} `json:"twitter_image"`
	TwitterTitle           interface{} `json:"twitter_title"`
	TwitterDescription     interface{} `json:"twitter_description"`
	MetaTitle              interface{} `json:"meta_title"`
	MetaDescription        interface{} `json:"meta_description"`
	EmailSubject           interface{} `json:"email_subject"`
	Frontmatter            interface{} `json:"frontmatter"`
}

type Articles struct {
	Posts []Post `json:"posts"`
	Meta  struct {
		Pagination struct {
			Page  int         `json:"page"`
			Limit int         `json:"limit"`
			Pages int         `json:"pages"`
			Total int         `json:"total"`
			Next  int         `json:"next"`
			Prev  interface{} `json:"prev"`
		} `json:"pagination"`
	} `json:"meta"`
}
