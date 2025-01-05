/*
 *  Copyright (c) 2020-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE.md file.
 */

package main

import (
	"go.osspkg.com/goppy/v2"
	"go.osspkg.com/goppy/v2/orm"
	"go.osspkg.com/goppy/v2/web"

	"uri-one/internal/controllers"
)

var Version = "v0.0.0-dev"

func main() {
	app := goppy.New("uri.one", Version, "")
	app.Plugins(
		web.WithServer(),
		orm.WithORM(),
		orm.WithMysql(),
	)
	app.Plugins(
		controllers.Plugins...,
	)
	app.Run()
}
