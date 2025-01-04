/*
 *  Copyright (c) 2020-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE.md file.
 */

package controllers

//go:generate easyjson

import (
	"fmt"
	"net/url"
	"strings"
)

//easyjson:json
type ConfigModel struct {
	Address          string `json:"address"`
	OrgName          string `json:"orgName"`
	ServiceName      string `json:"serviceName"`
	Email            string `json:"email"`
	CaptchaUse       bool   `json:"captchaUse"`
	CaptchaClientKey string `json:"captchaClientKey"`
}

//easyjson:json
type ShortenRequestModel struct {
	URL    string `json:"url"`
	Source string `json:"source"`
	Token  string `json:"token"`
	Domain string `json:"-"`
}

//easyjson:json
type ShortenResponseModel struct {
	URL    string `json:"url"`
	Source string `json:"source"`
}

func (v *ShortenRequestModel) Validate(excludeDomains []string) error {
	if len(v.Source) == 0 || len(v.Source) > 2048 {
		return fmt.Errorf("invalid `source`: <=0 or >2048")
	}

	u, err := url.Parse(v.Source)
	if err != nil {
		return fmt.Errorf("invalid `source`: %w", err)
	}

	v.Domain = u.Hostname()

	if len(u.Scheme) == 0 {
		return fmt.Errorf("invalid `source`: scheme is empty")
	}

	for _, domain := range excludeDomains {
		if strings.EqualFold(u.Hostname(), domain) {
			return fmt.Errorf("invalid `source`: unsupported domain")
		}
	}

	return nil
}
