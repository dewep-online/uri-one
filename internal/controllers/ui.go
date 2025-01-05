/*
 *  Copyright (c) 2020-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE.md file.
 */

package controllers

import "go.osspkg.com/static"

//go:generate static ./../../ui/dist/ui/browser ui

var ui static.Reader
