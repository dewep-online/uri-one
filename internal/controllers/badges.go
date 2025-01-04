/*
 *  Copyright (c) 2020-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE.md file.
 */

package controllers

import (
	"html"

	"go.osspkg.com/badges"
	"go.osspkg.com/goppy/v2/web"
	"go.osspkg.com/logx"
)

var colors = map[string]badges.Color{
	"primary":   badges.ColorPrimary,
	"secondary": badges.ColorSecondary,
	"success":   badges.ColorSuccess,
	"danger":    badges.ColorDanger,
	"warning":   badges.ColorWarning,
	"info":      badges.ColorInfo,
	"light":     badges.ColorLight,
}

func (v *Controller) BadgeDraw(ctx web.Context) {
	title, err := ctx.Param("title").String()
	if err != nil {
		ctx.String(400, "Invalid `title`")
		logx.Error("Invalid badge key", "err", err, "key", "title", "value", title)
		return
	}

	data, err := ctx.Param("data").String()
	if err != nil {
		ctx.String(400, "Invalid `data`")
		logx.Error("Invalid badge key", "err", err, "key", "data", "value", data)
		return
	}

	color, err := ctx.Param("color").String()
	if err != nil {
		ctx.String(400, "Invalid `color`")
		logx.Error("Invalid badge key", "err", err, "key", "color", "value", color)
		return
	}

	colored, ok := colors[color]
	if !ok {
		colored = badges.ColorPrimary
	}

	err = v.badge.WriteResponse(ctx.Response(), colored, html.EscapeString(title), html.EscapeString(data))
	if err != nil {
		logx.Error("Invalid badge response", "err", err)
	}
}
