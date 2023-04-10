package migrator

import (
	"tdp-aiart/helper/logman"
	"tdp-aiart/module/model/migration"
)

func Deploy() {

	if err := doMigrate(); err != nil {
		logman.Fatal("Migrate database failed", "error", err)
	}

}

func addMigration(k, v string) error {

	_, err := migration.Create(&migration.CreateParam{
		Version: k, Description: v,
	})

	return err

}

func isMigrated(k string) bool {

	rq := &migration.FetchParam{Version: k}

	if rs, err := migration.Fetch(rq); err == nil {
		return rs.Id > 0
	}

	return false

}

func doMigrate() error {

	funcs := []func() error{
		v100000,
		v100001,
		v100002,
		v100003,
		v100004,
	}

	for _, fn := range funcs {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil

}
