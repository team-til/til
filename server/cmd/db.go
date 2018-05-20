// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	"unicode"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq" // database/sql Postgres driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(migrateCmd)
	dbCmd.AddCommand(createCmd)
	dbCmd.AddCommand(dropCmd)
	dbCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(migrationCmd)
}

type DbConfig struct {
	Username string
	Password string
	Hostname string
	Database string
	Port     int
	Sslmode  string
}

var dbCmd = &cobra.Command{
	Use: "db",
}

var createCmd = &cobra.Command{
	Use: "create",
	Run: dbCreate,
}

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Run: dbMigrate,
}

var dropCmd = &cobra.Command{
	Use: "drop",
	Run: dbDrop,
}

var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate yo migration",
	Aliases: []string{"g"},
}

var migrationCmd = &cobra.Command{
	Use: "migration",
	Run: generateMigration,
}

// BuildDbConnectionStr returns a postgres compliant connection string
func (dbConfig DbConfig) BuildDbConnectionStr() string {
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		dbConfig.Username, dbConfig.Password, dbConfig.Hostname, dbConfig.Port, dbConfig.Database, dbConfig.Sslmode)
	return connectionStr
}

func generateMigration(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		log.Fatal("Must include exactly one argument")
	}

	if _, err := os.Stat("migrations/"); os.IsNotExist(err) {
		os.Mkdir("migrations/", 0777)
	}

	migrationName := fmt.Sprintf("%s_%s", getNowTimestamp(), toSnake(args[0]))

	log.Infof("Generating migration %+v", migrationName)

	if err := generateSQLFiles(migrationName); err != nil {
		log.Fatalf("Error generating SQL files. Err: %+v", err.Error())
	}

}

func toSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

func getNowTimestamp() string {
	now := time.Now()

	return fmt.Sprintf("%d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

func generateSQLFiles(migrationName string) error {
	_, err := os.Create(fmt.Sprintf("migrations/%s.up.sql", migrationName))
	if err != nil {
		return err
	}
	_, err = os.Create(fmt.Sprintf("migrations/%s.down.sql", migrationName))
	if err != nil {
		return err
	}
	return nil
}

func dbCreate(cmd *cobra.Command, args []string) {
	var dbConfig DbConfig
	viper.UnmarshalKey("datastore", &dbConfig)

	log.Infof("Creating database - %s", dbConfig.Database)

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d sslmode=disable dbname=postgres",
		dbConfig.Username, dbConfig.Password, dbConfig.Hostname, dbConfig.Port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to DB. Err: %+v", err.Error())
	}

	defer db.Close()

	query := fmt.Sprintf("create database %s", dbConfig.Database)
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Unable to create database. Err: %+v", err.Error())
	}

	log.Info("Database created")
}

func dbDrop(cmd *cobra.Command, args []string) {
	var dbConfig DbConfig
	viper.UnmarshalKey("datastore", &dbConfig)

	log.Infof("Dropping database - %s", dbConfig.Database)

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d sslmode=disable dbname=postgres",
		dbConfig.Username, dbConfig.Password, dbConfig.Hostname, dbConfig.Port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to DB. Err: %+v", err.Error())
	}

	defer db.Close()

	query := fmt.Sprintf("drop database %s", dbConfig.Database)
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Unable to drop database. Err: %+v", err.Error())
	}

	log.Info("Database dropped")
}

func dbMigrate(cmd *cobra.Command, args []string) {
	var dbConfig DbConfig
	viper.UnmarshalKey("datastore", &dbConfig)

	log.Infof("Running up migrations on %s", dbConfig.Database)

	connStr := dbConfig.BuildDbConnectionStr()
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatalf("Unable to connect to DB. Err: %+v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Unable to locate postgres driver. Err: %+v", err)
	}

	_, err = (&file.File{}).Open("file://migrations")
	if err != nil {
		log.Fatalf("Unable to open migrations. Err: %+v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		log.Fatalf("Migrations failed to run. Err: %+v", err)
	}
	log.Info("Migrations ran successfully")
}
