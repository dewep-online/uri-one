/*
 *  Copyright (c) 2020-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE.md file.
 */

package controllers

import (
	"context"
	"net/http"
	"time"

	"go.osspkg.com/do"
	"go.osspkg.com/goppy/v2/orm"
	"go.osspkg.com/ioutils/cache"
	"go.osspkg.com/routine"

	"uri-one/internal/pkg"

	"go.osspkg.com/algorithms/encoding/base62"
	"go.osspkg.com/badges"
	"go.osspkg.com/goppy/v2/web"
)

type Controller struct {
	conf     *Config
	route    web.Router
	badge    *badges.Badges
	codec    *base62.Base62
	db       *pkg.Database
	stats    *pkg.Stats
	cli      *http.Client
	capCache cache.TCacheTTL[string, struct{}]
}

func NewController(c *Config, r web.RouterPool, db orm.ORM, b *badges.Badges) *Controller {
	return &Controller{
		conf:  c,
		route: r.Main(),
		badge: b,
		codec: base62.New(c.ShortenAlphabet),
		db:    pkg.NewDatabase(db),
		stats: pkg.NewStats(),
		cli:   http.DefaultClient,
	}
}

func (v *Controller) Up(ctx context.Context) error {
	v.route.NotFoundHandler(v.Page404)

	for _, s := range ui.List() {
		v.route.Get(s, v.PageStatic)
	}
	v.route.Get("/", v.PageStatic)
	v.route.Get("/badges", v.PageStatic)
	v.route.Get("/license", v.PageStatic)

	v.route.Get("/badge/{color}/{title}/{data}/image.svg", v.BadgeDraw)
	v.route.Get("/{code}", v.ShortenGet)

	api := v.route.Collection("/api")
	api.Post("/shorten/add", v.ShortenAdd)
	api.Get("/config.json", v.ApiConfig)

	v.capCache = cache.NewWithTTL[string, struct{}](ctx, cookieTtl)
	routine.Interval(ctx, 15*time.Minute, v.ShortenUpdateStats)
	routine.Interval(ctx, 24*60*time.Minute, v.ShortenRemoveUnused)

	return nil
}

func (v *Controller) Down() error {
	return nil
}

func (v *Controller) PageStatic(ctx web.Context) {
	file := ctx.URL().Path
	switch file {
	case "/", "/badges", "/license":
		file = "/index.html"
	default:
	}

	ui.ResponseWrite(ctx.Response(), file) //nolint: errcheck
}

func (v *Controller) Page404(ctx web.Context) {
	ctx.String(404, "404")
}

func (v *Controller) ApiConfig(ctx web.Context) {
	c := ctx.Cookie().Get(captchaCookie)
	capOk := c != nil && len(c.Value) > 0 && v.capCache.Has(c.Value)

	config := &ConfigModel{
		Address:          v.conf.Contacts.Address,
		OrgName:          v.conf.Contacts.OrgName,
		ServiceName:      v.conf.Contacts.ServiceName,
		Email:            v.conf.Contacts.Email,
		CaptchaUse:       do.IfElse(capOk, false, v.conf.Captcha.Use),
		CaptchaClientKey: do.IfElse(capOk, "", v.conf.Captcha.ClientKey),
	}
	ctx.Header().Set("Cache-Control", "max-age=60, public")
	ctx.JSON(200, config)
}
