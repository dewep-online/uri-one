/*
 *  Copyright (c) 2020-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE.md file.
 */

package pkg

import (
	"context"
	"fmt"
	"time"

	"go.osspkg.com/goppy/v2/orm"
)

type Database struct {
	db orm.ORM
}

func NewDatabase(db orm.ORM) *Database {
	return &Database{
		db: db,
	}
}

func (v *Database) Master() orm.Stmt {
	return v.db.Tag("master")
}

func (v *Database) Slave() orm.Stmt {
	return v.db.Tag("slave")
}

func (v *Database) GetShorten(ctx context.Context, id uint64) (string, error) {
	var source string
	err := v.Slave().Query(ctx, "select_shorten", func(q orm.Querier) {
		q.SQL("SELECT `source` FROM `shorten_url` WHERE `id` = ? AND `lock` = 0 LIMIT 1;", id)
		q.Bind(func(bind orm.Scanner) error {
			return bind.Scan(&source)
		})
	})
	if err != nil {
		return "", err
	}
	if len(source) == 0 {
		return "", fmt.Errorf("not found")
	}
	return source, nil
}

func (v *Database) GetShortenIdByHash(ctx context.Context, hash string) (uint64, error) {
	var id uint64
	err := v.Slave().Query(ctx, "select_shorten", func(q orm.Querier) {
		q.SQL("SELECT `id` FROM `shorten_url` WHERE `hash` = ? LIMIT 1;", hash)
		q.Bind(func(bind orm.Scanner) error {
			return bind.Scan(&id)
		})
	})
	if err != nil {
		return 0, err
	}
	if id <= 0 {
		return 0, fmt.Errorf("not found")
	}
	return id, nil
}

func (v *Database) AddShorten(ctx context.Context, source, domain, hash string) (uint64, error) {
	var id uint64
	err := v.Master().Exec(ctx, "insert_new_shorten", func(q orm.Executor) {
		q.SQL("INSERT INTO `shorten_url` (`source`, `domain`, `hash`, `lock`, `request_count`, `created_at`, `updated_at`) VALUES (?, ?, ?, 0, 0, now(), now());")
		q.Params(source, domain, hash)
		q.Bind(func(rowsAffected, lastInsertId int64) error {
			id = uint64(lastInsertId)
			if id == 0 {
				return fmt.Errorf("invalid insert")
			}
			return nil
		})
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (v *Database) DeleteUnusedShorten(ctx context.Context) error {
	return v.Master().Exec(ctx, "delete_unused_shorten", func(q orm.Executor) {
		q.SQL("DELETE FROM `shorten_url` WHERE updated_at < ?;")
		q.Params(time.Now().Add(-3 * 30 * 24 * 60 * time.Minute))
	})
}

func (v *Database) UpdateStatsShorten(ctx context.Context, data map[uint64]uint64) error {
	return v.Master().Tx(ctx, "update_stats_shorten", func(tx orm.Tx) {
		for id, count := range data {
			tx.Exec(func(e orm.Executor) {
				e.SQL("UPDATE `shorten_url` SET `request_count` = `request_count`+? WHERE `id`=?;")
				e.Params(count, id)
			})
		}
	})
}
