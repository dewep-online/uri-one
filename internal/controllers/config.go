/*
 *  Copyright (c) 2020-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE.md file.
 */

package controllers

import "go.osspkg.com/random"

type (
	Config struct {
		ExcludeDomains  []string `yaml:"exclude_domains"`
		ShortenAlphabet string   `yaml:"shorten_alphabet"`
		Contacts        Contacts `yaml:"contacts"`
		Captcha         Captcha  `yaml:"captcha"`
	}
	Contacts struct {
		Address     string `yaml:"address"`
		ServiceName string `yaml:"service_name"`
		OrgName     string `yaml:"org_name"`
		Email       string `yaml:"email"`
	}
	Captcha struct {
		Use         bool   `yaml:"use"`
		ClientKey   string `yaml:"client_key"`
		ServerKey   string `yaml:"server_key"`
		ValidateUrl string `yaml:"validate_url"`
	}
)

func (v *Config) Default() {
	v.Contacts = Contacts{
		Address:     "http://localhost:8080",
		ServiceName: "UriOne",
		OrgName:     "Company Name",
		Email:       "help@osspkg.com",
	}

	v.Captcha = Captcha{
		Use:         false,
		ClientKey:   "*************",
		ServerKey:   "*************",
		ValidateUrl: "https://smartcaptcha.yandexcloud.net/validate",
	}

	v.ExcludeDomains = []string{
		"localhost",
	}

	data := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	random.Shuffle[byte](data)
	v.ShortenAlphabet = string(data)
}
