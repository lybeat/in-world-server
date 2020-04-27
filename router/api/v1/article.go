package v1

import (
	"in-world-server/model"
	"in-world-server/pkg/e"
	"in-world-server/pkg/net"
	"in-world-server/pkg/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetArticle(c *gin.Context) {
	netG := net.Gin{c}

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		net.MarkErrors(valid.Errors)
		netG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	exist, err := model.ExistArticleByID(id)
	if err != nil {
		netG.Response(http.StatusOK, e.ERROR_CHECK_DATA_FAIL, nil)
		return
	}
	if !exist {
		netG.Response(http.StatusOK, e.ERROR_NOT_EXIST_DATA, nil)
		return
	}

	article, err := model.GetArticle(id)
	if err != nil {
		netG.Response(http.StatusOK, e.ERROR_GET_DATA_FAIL, nil)
		return
	}

	netG.Response(http.StatusOK, e.SUCCESS, article)
}

func GetArticles(c *gin.Context) {
	netG := net.Gin{c}
	valid := validation.Validation{}

	if valid.HasErrors() {
		net.MarkErrors(valid.Errors)
		netG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	total, err := model.GetArticleTotal()
	if err != nil {
		netG.Response(http.StatusInternalServerError, e.ERROR_GET_DATA_FAIL, nil)
		return
	}

	articles, err := model.GetArticles(util.GetPage(c))
	if err != nil {
		netG.Response(http.StatusInternalServerError, e.ERROR_GET_DATA_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["total"] = total

	netG.Response(http.StatusOK, e.SUCCESS, data)
}

type AddArticleForm struct {
	UserID   int                    `form:"userId" valid:"Required;Min(1)"`
	Title    string                 `form:"title" valid:"Required;MaxSize(100)"`
	Summary  string                 `form:"summary" valid:"Required;MaxSize(65535)"`
	CoverUrl string                 `form:"coverUrl" valid:"Required;MaxSize(255)"`
	Contents []*model.ArticleDetail `form:"contents" valid:"Required"`
}

func AddArticle(c *gin.Context) {
	var (
		netG = net.Gin{c}
		form AddArticleForm
	)

	httpCode, errCode := net.BindJsonAndValid(c, &form)
	if errCode != e.SUCCESS {
		netG.Response(httpCode, errCode, nil)
		return
	}

	article := model.Article{
		UserID:   form.UserID,
		Title:    form.Title,
		Summary:  form.Summary,
		CoverUrl: form.CoverUrl,
	}

	if err := article.AddArticle(form.Contents); err != nil {
		netG.Response(http.StatusInternalServerError, e.ERROR_ADD_DATA_FAIL, nil)
		return
	}

	netG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditArticleForm struct {
	ID       int    `form:"id" valid:"Required;Min(1)"`
	Title    string `form:"title" valid:"Required;MaxSize(100)"`
	Content  string `form:"content" valid:"Required;MaxSize(65535)"`
	Author   string `form:"author" valid:"Required;MaxSize(100)"`
	CoverUrl string `form:"coverUrl" valid:"Required;MaxSize(255)"`
}

func EditArticle(c *gin.Context) {
	var (
		netG = net.Gin{c}
		form = EditArticleForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := net.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		netG.Response(httpCode, errCode, nil)
		return
	}

	article := map[string]interface{}{
		"title":    form.Title,
		"content":  form.Content,
		"author":   form.Author,
		"coverUrl": form.CoverUrl,
	}
	if err := model.EditArticle(form.ID, article); err != nil {
		netG.Response(http.StatusInternalServerError, e.ERROR_EDIT_DATA_FAIL, nil)
		return
	}

	netG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteArticle(c *gin.Context) {
	netG := net.Gin{c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		net.MarkErrors(valid.Errors)
		netG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	exist, err := model.ExistArticleByID(id)
	if err != nil {
		netG.Response(http.StatusBadRequest, e.ERROR_CHECK_DATA_FAIL, nil)
		return
	}
	if !exist {
		netG.Response(http.StatusBadRequest, e.ERROR_NOT_EXIST_DATA, nil)
		return
	}

	err = model.DeleteArticle(id)
	if err != nil {
		netG.Response(http.StatusBadRequest, e.ERROR_DELETE_DATA_FAIL, nil)
		return
	}

	netG.Response(http.StatusOK, e.SUCCESS, nil)
}
