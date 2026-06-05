package models

import "time"

// Comment represents a comment on a collection item
type Comment struct {
	ID      string `json:"id"`
	User    string `json:"user"`
	Avatar  string `json:"avatar"`
	Role    string `json:"role,omitempty"`
	Time    string `json:"time"`
	Content string `json:"content"`
}

// Collection represents a treasure collection item
type Collection struct {
	ID           string    `json:"id"`
	TitleCN      string    `json:"titleCN"`
	TitleEN      string    `json:"titleEN"`
	Category     string    `json:"category"`
	Image        string    `json:"image"`
	DetailImages []string  `json:"detailImages,omitempty"`
	Views        int       `json:"views"`
	Likes        int       `json:"likes"`
	CommentCount int       `json:"commentsCount"`
	BadgeCN      string    `json:"badgeCN,omitempty"`
	BadgeEN      string    `json:"badgeEN,omitempty"`
	DateStrCN    string    `json:"dateStrCN"`
	DateStrEN    string    `json:"dateStrEN"`
	DescriptionCN string   `json:"descriptionCN"`
	DescriptionEN string   `json:"descriptionEN"`
	DetailDescCN  string   `json:"detailDescCN"`
	DetailDescEN  string   `json:"detailDescEN"`
	Comments     []Comment `json:"comments,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// CreateCollectionRequest is the request body for creating a collection
type CreateCollectionRequest struct {
	ID            string    `json:"id"`
	TitleCN       string    `json:"titleCN"`
	TitleEN       string    `json:"titleEN"`
	Category      string    `json:"category"`
	Image         string    `json:"image"`
	DetailImages  []string  `json:"detailImages"`
	Views         int       `json:"views"`
	Likes         int       `json:"likes"`
	CommentCount  int       `json:"commentsCount"`
	BadgeCN       string    `json:"badgeCN"`
	BadgeEN       string    `json:"badgeEN"`
	DateStrCN     string    `json:"dateStrCN"`
	DateStrEN     string    `json:"dateStrEN"`
	DescriptionCN string    `json:"descriptionCN"`
	DescriptionEN string    `json:"descriptionEN"`
	DetailDescCN  string    `json:"detailDescCN"`
	DetailDescEN  string    `json:"detailDescEN"`
	Comments      []Comment `json:"comments"`
}

// UpdateCollectionRequest is the request body for updating a collection
type UpdateCollectionRequest struct {
	TitleCN       string    `json:"titleCN"`
	TitleEN       string    `json:"titleEN"`
	Category      string    `json:"category"`
	Image         string    `json:"image"`
	DetailImages  []string  `json:"detailImages"`
	Views         int       `json:"views"`
	Likes         int       `json:"likes"`
	CommentCount  int       `json:"commentsCount"`
	BadgeCN       string    `json:"badgeCN"`
	BadgeEN       string    `json:"badgeEN"`
	DateStrCN     string    `json:"dateStrCN"`
	DateStrEN     string    `json:"dateStrEN"`
	DescriptionCN string    `json:"descriptionCN"`
	DescriptionEN string    `json:"descriptionEN"`
	DetailDescCN  string    `json:"detailDescCN"`
	DetailDescEN  string    `json:"detailDescEN"`
	Comments      []Comment `json:"comments"`
}

// ListCollectionsResponse is the response for listing collections
type ListCollectionsResponse struct {
	Collections []Collection `json:"collections"`
	Total       int          `json:"total"`
}