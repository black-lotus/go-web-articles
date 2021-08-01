package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/globalsign/mgo"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionSchema string = "schema_migrations"
)

type (
	// Persistence contains all persistence database
	Persistence struct {
		SQLDB   *sql.DB
		MongoDB *mgo.Database
	}

	// MongoSchemaVersion domain of the data
	MongoSchemaVersion struct {
		Version int64 `bson:"version"`
	}
)

// NewPersistence create persitence database connections
func NewPersistence() *Persistence {
	// load env
	// load main .env and additional .env in app
	if err := godotenv.Load("cmd/migration/.env"); err != nil {
		fmt.Println("System cannot load env file in project. System will try read config from env in os variable ...")
	}

	sqlHostWrite, ok := os.LookupEnv("SQL_HOST_WRITE")
	if !ok {
		fmt.Println("missing SQL_HOST_WRITE environment")
	}

	mongoHostWrite, ok := os.LookupEnv("MONGODB_HOST_WRITE")
	if !ok {
		fmt.Println("missing MONGODB_HOST_WRITE environment")
	}

	mongoDBName, ok := os.LookupEnv("MONGODB_DATABASE_NAME")
	if !ok {
		fmt.Println("missing MONGODB_DATABASE_NAME environment")
	}

	return &Persistence{
		SQLDB:   newSQLConnection(&sqlHostWrite),
		MongoDB: newMongoConnection(&mongoHostWrite, &mongoDBName),
	}
}

// ClosePersistence close all the persistence connections
func ClosePersistence(p *Persistence) {
	if p.SQLDB != nil {
		p.SQLDB.Close()
	}

	if p.MongoDB != nil {
		p.MongoDB.Session.Close()
	}
}

// InitSchemaMigrations close all the persistence connections
func InitSchemaMigrations(p *Persistence) ([]int64, error) {
	if p.SQLDB != nil {
		return initMigrationsSQL(p.SQLDB)
	}

	if p.MongoDB != nil {
		return initMigrationsMongoDB(p.MongoDB)
	}

	return nil, errors.New("No connections found")
}

// InsertSchemaMigrations close all the persistence connections
func InsertSchemaMigrations(p *Persistence, version int64) error {
	if p.SQLDB != nil {
		query := fmt.Sprintf("INSERT INTO `%s` VALUES(%v)", collectionSchema, version)
		if _, err := p.SQLDB.ExecContext(context.TODO(), query); err != nil {
			return err
		}

		return nil
	}

	if p.MongoDB != nil {
		newRecord := MongoSchemaVersion{
			Version: version,
		}

		err := p.MongoDB.C(collectionSchema).Insert(newRecord)
		return err
	}

	return errors.New("No connections found")
}

// DeleteSchemaMigrations close all the persistence connections
func DeleteSchemaMigrations(p *Persistence, version int64) error {
	if p.SQLDB != nil {
		query := fmt.Sprintf("DELETE FROM `%s` WHERE version = %v", collectionSchema, version)
		if _, err := p.SQLDB.ExecContext(context.TODO(), query); err != nil {
			return err
		}

		return nil
	}

	if p.MongoDB != nil {
		err := p.MongoDB.C(collectionSchema).Remove(bson.M{"version": version})
		return err
	}

	return errors.New("No connections found")
}

func newSQLConnection(host *string) *sql.DB {
	if *host == "" {
		return nil
	}

	fmt.Println("Connecting to MySQL database...")

	db, err := sql.Open("mysql", *host)
	if err != nil {
		fmt.Println("Unable to connect to database", err.Error())
		return nil
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Unable to connect to database", err.Error())
		return nil
	}

	fmt.Println("Database connected!")

	return db
}

func newMongoConnection(host *string, db *string) *mgo.Database {
	if *host == "" {
		return nil
	}

	session, err := mgo.Dial(*host)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer session.Close()
	return session.DB(*db)
}

func initMigrationsSQL(db *sql.DB) ([]int64, error) {
	// Create `schema_migrations` table to remember which migrations were executed.
	if _, err := db.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		version varchar(255)
	);`, collectionSchema)); err != nil {
		fmt.Printf("Unable to create `%s` table %v\n", collectionSchema, err)
		return nil, err
	}

	// Find out all the executed migrations
	rows, err := db.Query(fmt.Sprintf("SELECT version FROM `%s`;", collectionSchema))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Mark the migrations as Done if it is already executed
	result := []int64{}
	for rows.Next() {
		var version int64
		err := rows.Scan(&version)
		if err != nil {
			return nil, err
		}

		result = append(result, version)
	}

	return result, nil
}

func initMigrationsMongoDB(db *mgo.Database) ([]int64, error) {
	var result []MongoSchemaVersion

	err := db.C(collectionSchema).Find(bson.M{}).All(&result)
	if err != nil {
		return nil, err
	}

	var res []int64 = make([]int64, 0, len(result))
	for _, v := range result {
		res = append(res, v.Version)
	}

	return res, nil
}
