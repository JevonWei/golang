package controllers

import (
	"github.com/imsilence/gocmdb/server/controllers/auth"
)

type LayoutController struct {
	auth.LoginRequiredController
}

func (c *LayoutController) Prepare() {
	c.LoginRequiredController.Prepare()
	c.Layout = "layouts/base.html"

	// c.LayoutSections = make(map[string]string)
	// c.LayoutSections["LayoutStyle"] = ""
	// c.LayoutSections["LayoutScript"] = ""

	c.LayoutSections = map[string]string{
		"LayoutStyle":  "",
		"LayoutScript": "",
	}

	c.Data["menu"] = ""
	c.Data["expand"] = ""
}
