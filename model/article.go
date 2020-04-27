package model

import (
	"in-world-server/pkg/setting"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	UserID int  `json:"user_id" gorm:"index"`
	User   User `json:"user"`

	Title    string `json:"title"`
	Summary  string `json:"summary"`
	CoverUrl string `json:"coverUrl"`
}

func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetArticleTotal() (int, error) {
	var count int
	if err := db.Model(&Article{}).Where(GetMaps()).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetArticles(pageNum int) ([]*Article, error) {
	var articles []*Article
	pageSize := setting.AppSetting.PageSize
	err := db.Preload("User").Where(GetMaps()).Offset(pageNum).Limit(pageSize).Order("id desc").Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

type CompleteArticle struct {
	ID         int              `json:"id"`
	CreatedOn  int              `json:"createdOn"`
	ModifiedOn int              `json:"modifiedOn"`
	DeletedOn  int              `json:"deletedOn"`
	Title      string           `json:"title"`
	CoverUrl   string           `json:"coverUrl"`
	User       User             `json:"user"`
	Contents   []*ArticleDetail `json:"contents"`
}

func GetArticle(id int) (*CompleteArticle, error) {
	var article Article
	err := db.Preload("User").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	details, err := GetArticleDetails(id)
	if err != nil {
		return nil, err
	}

	completeArticle := CompleteArticle{
		ID:         article.ID,
		CreatedOn:  article.CreatedOn,
		ModifiedOn: article.ModifiedOn,
		DeletedOn:  article.DeletedOn,
		Title:      article.Title,
		CoverUrl:   article.CoverUrl,
		User:       article.User,
		Contents:   details,
	}

	return &completeArticle, nil
}

func (a *Article) AddArticle(details []*ArticleDetail) error {
	article := Article{
		UserID:   a.UserID,
		Title:    a.Title,
		Summary:  a.Summary,
		CoverUrl: a.CoverUrl,
	}
	if err := db.Create(&article).Error; err != nil {
		return err
	}

	for i := range details {
		details[i].ArticleID = article.ID
	}
	err := AddArticleDetails(details)
	if err != nil {
		return err
	}

	return nil
}

func EditArticle(id int, data interface{}) error {
	if err := db.Model(&Article{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func DeleteArticle(id int) error {
	if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllArticle() error {
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&Article{}).Error; err != nil {
		return err
	}

	return nil
}

func GetMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	return maps
}
