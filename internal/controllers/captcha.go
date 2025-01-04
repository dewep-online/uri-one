/*
 *  Copyright (c) 2020-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE.md file.
 */

package controllers

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

func (v *Controller) checkCaptcha(ctx context.Context, token string) bool {
	data := url.Values{}
	data.Set("secret", v.conf.Captcha.ServerKey)
	data.Set("token", token)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		v.conf.Captcha.ValidateUrl,
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return false
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := v.cli.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close() //nolint: errcheck

	return resp.StatusCode == http.StatusOK
}
