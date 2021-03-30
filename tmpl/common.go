package tmpl

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/config"
	"github.com/taoshihan1991/imaptool/models"
	"github.com/taoshihan1991/imaptool/tools"
	"html"
	"html/template"
	"net/http"
)

type CommonHtml struct {
	Header template.HTML
	Nav    template.HTML
	Left   template.HTML
	Bottom template.HTML
	Rw     http.ResponseWriter
}

func NewRender(rw http.ResponseWriter) *CommonHtml {
	obj := new(CommonHtml)
	obj.Rw = rw
	header := tools.FileGetContent("html/header.html")
	nav := tools.FileGetContent("html/nav.html")
	obj.Header = template.HTML(header)
	obj.Nav = template.HTML(nav)
	return obj
}
func (obj *CommonHtml) SetLeft(file string) {
	leftStr := tools.FileGetContent("html/" + file + ".html")
	obj.Left = template.HTML(leftStr)
}
func (obj *CommonHtml) SetBottom(file string) {
	str := tools.FileGetContent("html/" + file + ".html")
	obj.Bottom = template.HTML(str)
}
func (obj *CommonHtml) Display(file string, data interface{}) {
	if data == nil {
		data = obj
	}
	main := tools.FileGetContent("html/" + file + ".html")
	t, _ := template.New(file).Parse(main)
	t.Execute(obj.Rw, data)
}

//首页
func PageIndex(c *gin.Context) {
	if c.Request.RequestURI == "/favicon.ico" {
		return
	}

	lang, _ := c.Get("lang")
	language := config.CreateLanguage(lang.(string))
	about := models.FindAboutByPageLanguage("index", lang.(string))
	cssJs := html.UnescapeString(about.CssJs)
	title := about.TitleCn
	keywords := about.KeywordsCn
	desc := html.UnescapeString(about.DescCn)
	content := html.UnescapeString(about.HtmlCn)
	if lang == "en" {
		title = about.TitleEn
		keywords = about.KeywordsEn
		desc = html.UnescapeString(about.DescEn)
		content = html.UnescapeString(about.HtmlEn)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"OnlineChat": language.IndexOnlineChat,
		"Notice":     language.Notice,
		"NowAsk":     language.NowAsk,
		"LaterAsk":   language.LaterAsk,
		"Lang":       lang,
		"Title":      title,
		"Keywords":   keywords,
		"Desc":       desc,
		"Content":    template.HTML(content),
		"CssJs":      template.HTML(cssJs),
	})
}

//登陆界面
func PageMain(c *gin.Context) {
	nav := tools.FileGetContent("html/nav.html")
	c.HTML(http.StatusOK, "main.html", gin.H{
		"Nav": template.HTML(nav),
	})
}

//客服界面
func PageChatMain(c *gin.Context) {
	c.HTML(http.StatusOK, "chat_main.html", nil)
}
