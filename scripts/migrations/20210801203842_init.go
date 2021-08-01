package migrations

import (
	"webarticles/scripts/migrations/db"
)

func init() {
	migrator.AddMigration(&Migration{
		Version: 20210801203842,
		Up:      mig_20210801203842_init_up,
		Down:    mig_20210801203842_init_down,
	})
}

func mig_20210801203842_init_up(p *db.Persistence) error {
	_, err := p.SQLDB.Exec("CREATE TABLE article (" +
		"id INTEGER AUTO_INCREMENT PRIMARY KEY, " +
		"author TEXT NOT NULL, " +
		"title TEXT NOT NULL, " +
		"body TEXT DEFAULT NULL, " +
		"created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, " +
		"is_deleted BOOL DEFAULT 0" +
		");")
	if err != nil {
		return err
	}

	_, err = p.SQLDB.Exec("CREATE INDEX article_is_deleted_idx ON article(is_deleted);")
	return err
}

func mig_20210801203842_init_down(p *db.Persistence) error {
	_, err := p.SQLDB.Exec("DROP TABLE IF EXISTS article;")
	if err != nil {
		return err
	}

	return nil
}
