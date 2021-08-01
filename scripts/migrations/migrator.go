package migrations

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"sort"
	"time"
	"webarticles/scripts/migrations/db"
)

// Migration ..
type Migration struct {
	Version int64
	Up      func(*db.Persistence) error
	Down    func(*db.Persistence) error

	done bool
}

// Migrator ..
type Migrator struct {
	p *db.Persistence
	// Versions   []string
	Migrations map[int64]*Migration
}

var migrator = &Migrator{
	//Versions:   []string{},
	Migrations: map[int64]*Migration{},
}

// AddMigration ..
func (m *Migrator) AddMigration(mg *Migration) {
	// Add the migration to the hash with version as key
	m.Migrations[mg.Version] = mg
}

// Init ..
func Init(p *db.Persistence) (*Migrator, error) {
	migrator.p = p

	result, err := db.InitSchemaMigrations(p)
	if err != nil {
		return migrator, err
	}

	for _, v := range result {
		if migrator.Migrations[v] != nil {
			migrator.Migrations[v].done = true
		}
	}

	return migrator, err
}

// Up ..
func (m *Migrator) Up(version int64) error {
	keys := make([]int64, 0, len(m.Migrations))
	for k := range m.Migrations {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	if version == 0 && len(keys) > 0 {
		version = keys[len(keys)-1]
	}

	for _, v := range keys {
		if v == version {
			mg := m.Migrations[v]
			if mg.done {
				return fmt.Errorf("This version %v already migrated", version)
			}

			fmt.Println("Running migration", mg.Version)
			if err := mg.Up(m.p); err != nil {
				return err
			}

			if err := db.InsertSchemaMigrations(m.p, mg.Version); err != nil {
				return err
			}

			fmt.Println("Finished running migration", mg.Version)

			return nil
		}
	}

	return fmt.Errorf("No matched version for %v", version)
}

// Down ..
func (m *Migrator) Down(version int64) error {
	keys := make([]int64, 0, len(m.Migrations))
	for k := range m.Migrations {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	if version == 0 {
		return errors.New("Please specify the -version or -v")
	}

	for _, v := range keys {
		if v == version {
			mg := m.Migrations[v]
			if !mg.done {
				return fmt.Errorf("This version %v wasn't migrated yet", version)
			}

			fmt.Println("Reverting Migration", mg.Version)
			if err := mg.Down(m.p); err != nil {
				return err
			}

			if err := db.DeleteSchemaMigrations(m.p, mg.Version); err != nil {
				return err
			}
			fmt.Println("Finished reverting migration", mg.Version)

			return nil
		}
	}

	return fmt.Errorf("No matched version for %v", version)
}

// MigrationStatus ..
func (m *Migrator) MigrationStatus() error {
	keys := make([]int64, 0, len(m.Migrations))
	for k := range m.Migrations {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, v := range keys {
		mg := m.Migrations[v]

		if mg.done {
			fmt.Printf("Migration %v... completed\n", v)
		} else {
			fmt.Printf("Migration %v... pending\n", v)
		}
	}

	return nil
}

// Create ..
func Create(name string) error {
	version := time.Now().Format("20060102150405")

	in := struct {
		Version string
		Name    string
	}{
		Version: version,
		Name:    name,
	}

	var out bytes.Buffer

	t := template.Must(template.ParseFiles("./scripts/migrations/template.txt"))
	err := t.Execute(&out, in)
	if err != nil {
		return errors.New("Unable to execute template:" + err.Error())
	}

	f, err := os.Create(fmt.Sprintf("./scripts/migrations/%s_%s.go", version, name))
	if err != nil {
		return errors.New("Unable to create migration file:" + err.Error())
	}
	defer f.Close()

	if _, err := f.WriteString(out.String()); err != nil {
		return errors.New("Unable to write to migration file:" + err.Error())
	}

	fmt.Println("Generated new migration files...", f.Name())
	return nil
}
