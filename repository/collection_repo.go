package repository

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"treasure-server/database"
	"treasure-server/models"
)

// CollectionRepo handles database operations for collections
type CollectionRepo struct{}

// NewCollectionRepo creates a new CollectionRepo
func NewCollectionRepo() *CollectionRepo {
	return &CollectionRepo{}
}

// Create inserts a new collection into the database
func (r *CollectionRepo) Create(req *models.CreateCollectionRequest) (*models.Collection, error) {
	id := req.ID
	if id == "" {
		id = fmt.Sprintf("CL-%d", time.Now().UnixNano())
	}

	detailImagesJSON, _ := json.Marshal(req.DetailImages)
	if req.DetailImages == nil {
		detailImagesJSON = []byte("[]")
	}

	commentsJSON, _ := json.Marshal(req.Comments)
	if req.Comments == nil {
		commentsJSON = []byte("[]")
	}

	now := time.Now()

	_, err := database.DB.Exec(`
		INSERT INTO collections (
			id, title_cn, title_en, category, image, detail_images,
			views, likes, comment_count, badge_cn, badge_en,
			date_str_cn, date_str_en, description_cn, description_en,
			detail_desc_cn, detail_desc_en, comments, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		id, req.TitleCN, req.TitleEN, req.Category, req.Image, string(detailImagesJSON),
		req.Views, req.Likes, req.CommentCount, req.BadgeCN, req.BadgeEN,
		req.DateStrCN, req.DateStrEN, req.DescriptionCN, req.DescriptionEN,
		req.DetailDescCN, req.DetailDescEN, string(commentsJSON), now, now,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert collection: %w", err)
	}

	return r.GetByID(id)
}

// GetByID retrieves a collection by ID
func (r *CollectionRepo) GetByID(id string) (*models.Collection, error) {
	row := database.DB.QueryRow(`
		SELECT id, title_cn, title_en, category, image, detail_images,
			views, likes, comment_count, badge_cn, badge_en,
			date_str_cn, date_str_en, description_cn, description_en,
			detail_desc_cn, detail_desc_en, comments, created_at, updated_at
		FROM collections WHERE id = ?
	`, id)

	return scanCollection(row)
}

// Update updates a collection by ID
func (r *CollectionRepo) Update(id string, req *models.UpdateCollectionRequest) (*models.Collection, error) {
	detailImagesJSON, _ := json.Marshal(req.DetailImages)
	if req.DetailImages == nil {
		detailImagesJSON = []byte("[]")
	}

	commentsJSON, _ := json.Marshal(req.Comments)
	if req.Comments == nil {
		commentsJSON = []byte("[]")
	}

	now := time.Now()

	_, err := database.DB.Exec(`
		UPDATE collections SET
			title_cn = ?, title_en = ?, category = ?, image = ?, detail_images = ?,
			views = ?, likes = ?, comment_count = ?, badge_cn = ?, badge_en = ?,
			date_str_cn = ?, date_str_en = ?, description_cn = ?, description_en = ?,
			detail_desc_cn = ?, detail_desc_en = ?, comments = ?, updated_at = ?
		WHERE id = ?
	`,
		req.TitleCN, req.TitleEN, req.Category, req.Image, string(detailImagesJSON),
		req.Views, req.Likes, req.CommentCount, req.BadgeCN, req.BadgeEN,
		req.DateStrCN, req.DateStrEN, req.DescriptionCN, req.DescriptionEN,
		req.DetailDescCN, req.DetailDescEN, string(commentsJSON), now, id,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update collection: %w", err)
	}

	return r.GetByID(id)
}

// Delete deletes a collection by ID
func (r *CollectionRepo) Delete(id string) error {
	result, err := database.DB.Exec("DELETE FROM collections WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete collection: %w", err)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("collection not found: %s", id)
	}

	return nil
}

// List retrieves collections with optional filtering and pagination
func (r *CollectionRepo) List(category string, keyword string, page, pageSize int) (*models.ListCollectionsResponse, error) {
	var whereClauses []string
	var args []interface{}

	if category != "" && category != "all" {
		whereClauses = append(whereClauses, "category = ?")
		args = append(args, category)
	}

	if keyword != "" {
		whereClauses = append(whereClauses, "(title_cn LIKE ? OR title_en LIKE ? OR id LIKE ?)")
		kw := "%" + keyword + "%"
		args = append(args, kw, kw, kw)
	}

	whereSQL := ""
	if len(whereClauses) > 0 {
		whereSQL = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Count total
	var total int
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM collections %s", whereSQL)
	err := database.DB.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("failed to count collections: %w", err)
	}

	// Pagination
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize

	// Query rows
	query := fmt.Sprintf(`
		SELECT id, title_cn, title_en, category, image, detail_images,
			views, likes, comment_count, badge_cn, badge_en,
			date_str_cn, date_str_en, description_cn, description_en,
			detail_desc_cn, detail_desc_en, comments, created_at, updated_at
		FROM collections %s
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`, whereSQL)

	queryArgs := append(args, pageSize, offset)
	rows, err := database.DB.Query(query, queryArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to query collections: %w", err)
	}
	defer rows.Close()

	collections := make([]models.Collection, 0)
	for rows.Next() {
		scanRow := &rowScanner{rows: rows}
		item, err := scanCollection(scanRow)
		if err != nil {
			return nil, err
		}
		collections = append(collections, *item)
	}

	if collections == nil {
		collections = []models.Collection{}
	}

	return &models.ListCollectionsResponse{
		Collections: collections,
		Total:       total,
	}, nil
}

// scanner interface abstracts both *sql.Row and *sql.Rows
type scanner interface {
	Scan(dest ...interface{}) error
}

type rowScanner struct {
	rows interface {
		Scan(dest ...interface{}) error
	}
}

func (rs *rowScanner) Scan(dest ...interface{}) error {
	return rs.rows.Scan(dest...)
}

func scanCollection(s scanner) (*models.Collection, error) {
	var (
		id, titleCN, titleEN, category, image string
		detailImagesJSON, commentsJSON        string
		views, likes, commentCount            int
		badgeCN, badgeEN                      string
		dateStrCN, dateStrEN                  string
		descCN, descEN, detailDescCN, detailDescEN string
		createdAt, updatedAt                  time.Time
	)

	err := s.Scan(
		&id, &titleCN, &titleEN, &category, &image, &detailImagesJSON,
		&views, &likes, &commentCount, &badgeCN, &badgeEN,
		&dateStrCN, &dateStrEN, &descCN, &descEN,
		&detailDescCN, &detailDescEN, &commentsJSON, &createdAt, &updatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to scan collection: %w", err)
	}

	var detailImages []string
	if detailImagesJSON != "" {
		json.Unmarshal([]byte(detailImagesJSON), &detailImages)
	}
	if detailImages == nil {
		detailImages = []string{}
	}

	var comments []models.Comment
	if commentsJSON != "" {
		json.Unmarshal([]byte(commentsJSON), &comments)
	}
	if comments == nil {
		comments = []models.Comment{}
	}

	return &models.Collection{
		ID:            id,
		TitleCN:       titleCN,
		TitleEN:       titleEN,
		Category:      category,
		Image:         image,
		DetailImages:  detailImages,
		Views:         views,
		Likes:         likes,
		CommentCount:  commentCount,
		BadgeCN:       badgeCN,
		BadgeEN:       badgeEN,
		DateStrCN:     dateStrCN,
		DateStrEN:     dateStrEN,
		DescriptionCN: descCN,
		DescriptionEN: descEN,
		DetailDescCN:  detailDescCN,
		DetailDescEN:  detailDescEN,
		Comments:      comments,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}, nil
}