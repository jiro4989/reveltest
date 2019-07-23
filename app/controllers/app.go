package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "こんにちは"
	foobar := "fooooooooooo"
	return c.Render(greeting, foobar)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("名前必須ですよ")
	c.Validation.MinSize(myName, 3).Message("最低でも3文字は必要ですよ")

	// バリデーションに問題があれば、Indexページにリダイレクトさせる
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}
