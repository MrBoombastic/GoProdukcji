package utils

import (
	"time"
)

type Article struct {
	ID                     string          `json:"id"`
	UUID                   string          `json:"uuid"`
	Title                  string          `json:"title"`
	Slug                   string          `json:"slug"`
	HTML                   string          `json:"html"`
	CommentID              string          `json:"comment_id"`
	FeatureImage           string          `json:"feature_image"`
	Featured               bool            `json:"featured"`
	Visibility             string          `json:"visibility"`
	CreatedAt              time.Time       `json:"created_at"`
	UpdatedAt              time.Time       `json:"updated_at"`
	PublishedAt            time.Time       `json:"published_at"`
	CustomExcerpt          string          `json:"custom_excerpt"`
	CodeinjectionHead      interface{}     `json:"codeinjection_head"`
	CodeinjectionFoot      interface{}     `json:"codeinjection_foot"`
	CustomTemplate         interface{}     `json:"custom_template"`
	CanonicalURL           interface{}     `json:"canonical_url"`
	EmailRecipientFilter   string          `json:"email_recipient_filter"`
	Tags                   []ArticleTag    `json:"tags"`
	Authors                []ArticleAuthor `json:"authors"`
	PrimaryAuthor          ArticleAuthor   `json:"primary_author"`
	PrimaryTag             ArticleTag      `json:"primary_tag"`
	URL                    string          `json:"url"`
	Excerpt                string          `json:"excerpt"`
	ReadingTime            int             `json:"reading_time"`
	Access                 bool            `json:"access"`
	SendEmailWhenPublished bool            `json:"send_email_when_published"`
	OgImage                interface{}     `json:"og_image"`
	OgTitle                interface{}     `json:"og_title"`
	OgDescription          interface{}     `json:"og_description"`
	TwitterImage           interface{}     `json:"twitter_image"`
	TwitterTitle           interface{}     `json:"twitter_title"`
	TwitterDescription     interface{}     `json:"twitter_description"`
	MetaTitle              interface{}     `json:"meta_title"`
	MetaDescription        interface{}     `json:"meta_description"`
	EmailSubject           interface{}     `json:"email_subject"`
	Frontmatter            interface{}     `json:"frontmatter"`
	Plaintext              string          `json:"plaintext,omitempty"`
}

type Articles struct {
	Posts []Article `json:"posts"`
	Meta  struct {
		Pagination struct {
			Page  int         `json:"page"`
			Pages int         `json:"pages"`
			Total int         `json:"total"`
			Next  int         `json:"next"`
			Prev  interface{} `json:"prev"`
		} `json:"pagination"`
	} `json:"meta"`
}

type ArticleAuthor struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	Slug            string      `json:"slug"`
	ProfileImage    string      `json:"profile_image"`
	CoverImage      interface{} `json:"cover_image"`
	Bio             string      `json:"bio"`
	Website         string      `json:"website"`
	Location        interface{} `json:"location"`
	Facebook        string      `json:"facebook"`
	Twitter         string      `json:"twitter"`
	MetaTitle       interface{} `json:"meta_title"`
	MetaDescription interface{} `json:"meta_description"`
	URL             string      `json:"url"`
}

type ArticleTag struct {
	ID                 string      `json:"id"`
	Name               string      `json:"name"`
	Slug               string      `json:"slug"`
	Description        interface{} `json:"description"`
	FeatureImage       interface{} `json:"feature_image"`
	Visibility         string      `json:"visibility"`
	MetaTitle          interface{} `json:"meta_title"`
	MetaDescription    interface{} `json:"meta_description"`
	OgImage            interface{} `json:"og_image"`
	OgTitle            interface{} `json:"og_title"`
	OgDescription      interface{} `json:"og_description"`
	TwitterImage       interface{} `json:"twitter_image"`
	TwitterTitle       interface{} `json:"twitter_title"`
	TwitterDescription interface{} `json:"twitter_description"`
	CodeinjectionHead  interface{} `json:"codeinjection_head"`
	CodeinjectionFoot  interface{} `json:"codeinjection_foot"`
	CanonicalURL       interface{} `json:"canonical_url"`
	AccentColor        interface{} `json:"accent_color"`
	URL                string      `json:"url"`
}
