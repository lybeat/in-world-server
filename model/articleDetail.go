package model

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ArticleDetail struct {
	Model

	ArticleID int    `json:"articleId"`
	Type      int    `json:"type"`
	Content   string `json:"content"`
	Desc      string `json:"desc"`
}

func ExistByID(id int) (bool, error) {
	var articleDetail ArticleDetail
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&articleDetail).Error
	if err != nil {
		return false, err
	}

	if articleDetail.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetArticleDetails(articleId int) ([]*ArticleDetail, error) {
	var details []*ArticleDetail
	err := db.Where("article_id = ? AND deleted_on = ?", articleId, 0).Find(&details).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return details, nil
}

type AddArticleDetailForm struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
	Desc    string `json:"desc"`
}

func AddArticleDetails(details []*ArticleDetail) error {
	tx := db.Begin()
	sqlStr := "INSERT INTO `article_detail` (`id`, `article_id`, `type`, `content`, `desc`, `created_on`, `modified_on`, `deleted_on`) VALUES "
	vals := []interface{}{}
	const rowSQL = "(?, ?, ?, ?, ?, ?, ?, ?)"
	var inserts []string
	nowTime := time.Now().Unix()
	for _, elem := range details {
		inserts = append(inserts, rowSQL)
		vals = append(vals, elem.ID, elem.ArticleID, elem.Type, elem.Content, elem.Desc, nowTime, nowTime, 0)
	}
	sqlStr = sqlStr + strings.Join(inserts, ",")
	err := tx.Exec(sqlStr, vals...).Error
	if err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
	}

	return nil
}
