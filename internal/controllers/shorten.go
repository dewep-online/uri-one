/*
 *  Copyright (c) 2020-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE.md file.
 */

package controllers

//go:generate easyjson

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	enc "go.osspkg.com/encrypt/hash"
	"go.osspkg.com/goppy/v2/web"
	"go.osspkg.com/logx"
)

func (v *Controller) ShortenGet(ctx web.Context) {
	code, err := ctx.Param("code").String()
	if err != nil {
		ctx.String(404, "404")
		logx.Error("Invalid shorten key", "err", err, "key", "code", "value", code)
		return
	}

	id := v.codec.Decode(code)
	uri, err := v.db.GetShorten(ctx.Context(), id)
	if err != nil {
		ctx.String(404, "404")
	} else {
		go v.stats.Add(id)
		ctx.Redirect(uri)
	}
}

func (v *Controller) ShortenAdd(ctx web.Context) {
	request := &ShortenRequestModel{}
	if err := ctx.BindJSON(request); err != nil {
		ctx.ErrorJSON(400, fmt.Errorf("invalid body: %w", err), nil)
		return
	}

	if err := request.Validate(v.conf.ExcludeDomains); err != nil {
		ctx.ErrorJSON(400, err, nil)
		return
	}

	addr, err := url.Parse(v.conf.Contacts.Address)
	if err != nil {
		ctx.ErrorJSON(500, fmt.Errorf("internal error"), nil)
		return
	}

	if v.conf.Captcha.Use {
		c := ctx.Cookie().Get(captchaCookie)
		capOk := c != nil && len(c.Value) > 0 && v.capCache.Has(c.Value)
		if !capOk {
			if !v.checkCaptcha(ctx.Context(), request.Token) {
				ctx.ErrorJSON(400, fmt.Errorf("fail validate captcha"), nil)
				return
			}
			tHash := enc.SHA256(request.Token)
			v.capCache.Set(tHash, struct{}{})
			ctx.Cookie().Set(&http.Cookie{
				Name:     captchaCookie,
				Value:    tHash,
				Path:     "/",
				Domain:   addr.Hostname(),
				Expires:  time.Now().Add(cookieTtl - 1*time.Minute),
				Secure:   addr.Scheme == "https",
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			})
		}
	}

	hash := enc.SHA256(request.Source)

	id, err := v.db.GetShortenIdByHash(ctx.Context(), hash)
	if err != nil {
		id, err = v.db.AddShorten(ctx.Context(), request.Source, request.Domain, hash)
	}

	if err != nil {
		ctx.ErrorJSON(400, fmt.Errorf("fail save"), nil)
		return
	}

	response := &ShortenResponseModel{
		URL:    fmt.Sprintf("%s/%s", v.conf.Contacts.Address, v.codec.Encode(id)),
		Source: request.Source,
	}
	ctx.JSON(200, response)
}

func (v *Controller) ShortenUpdateStats(ctx context.Context) {
	data := v.stats.Get()
	err := v.db.UpdateStatsShorten(ctx, data)
	if err != nil {
		logx.Error("Update stats shorten", "err", err)
	}
}

func (v *Controller) ShortenRemoveUnused(ctx context.Context) {
	err := v.db.DeleteUnusedShorten(ctx)
	if err != nil {
		logx.Error("Delete unused shorten", "err", err)
	}
}
