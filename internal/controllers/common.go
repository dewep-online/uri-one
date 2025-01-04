/*
 *  Copyright (c) 2020-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE.md file.
 */

package controllers

import (
	"time"

	"go.osspkg.com/badges"
	"go.osspkg.com/goppy/v2/plugins"
)

var Plugins = plugins.Inject(
	plugins.Plugin{Config: &Config{}},
	NewController,
	badges.New,
)

const (
	captchaCookie = "_cc"
	cookieTtl     = 30 * time.Minute
)
