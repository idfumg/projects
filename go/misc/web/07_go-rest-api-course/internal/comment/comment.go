package comment

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/jinzhu/gorm"
)

type Service struct {
	DB *gorm.DB
}

type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

type Comments []*Comment

type CommentService interface {
	GetComment(ID int) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID int, newComment Comment) (Comment, error)
	DeleteComment(ID int) error
	GetAllComments() ([]Comment, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (c *Comment) FromJson(r io.Reader) (*Comment, error) {
	comment := Comment{}
	if err := json.NewDecoder(r).Decode(&comment); err != nil {
		return nil, fmt.Errorf("failed to decode JSON body")
	}
	return &comment, nil
}

func (s *Service) GetComment(ID int) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

func (s *Service) GetCommentsBySlug(slug string) (Comments, error) {
	var comments Comments
	if result := s.DB.Find(comments).Where("slug = ?", slug); result.Error != nil {
		return Comments{}, result.Error
	}
	return comments, nil
}

func (s *Service) PostComment(comment *Comment) (*Comment, error) {
	if result := s.DB.Save(comment); result.Error != nil {
		return &Comment{}, result.Error
	}
	return comment, nil
}

func (s *Service) UpdateComment(ID int, newComment *Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}

	if result := s.DB.Model(&comment).Updates(&newComment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

func (s *Service) DeleteComment(ID int) error {
	if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Service) GetAllComments() (Comments, error) {
	var comments Comments
	if result := s.DB.Find(&comments); result.Error != nil {
		return Comments{}, result.Error
	}
	return comments, nil
}
