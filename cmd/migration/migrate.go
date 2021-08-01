package main

import (
	"fmt"

	"webarticles/scripts/migrations"
	"webarticles/scripts/migrations/db"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "database migrations tool",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new empty migrations file",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Unable to read flag `name`", err.Error())
			return
		}

		if err := migrations.Create(name); err != nil {
			fmt.Println("Unable to create migration", err.Error())
			return
		}
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "run up migrations",
	Run: func(cmd *cobra.Command, args []string) {

		version, err := cmd.Flags().GetInt64("version")
		if err != nil {
			fmt.Println("Unable to read flag `version`")
			return
		}

		p := db.NewPersistence()
		defer db.ClosePersistence(p)
		migrator, err := migrations.Init(p)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		err = migrator.Up(version)
		if err != nil {
			fmt.Printf("Unable to run `up` migrations with error %v", err.Error())
			return
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "run down migrations",
	Run: func(cmd *cobra.Command, args []string) {

		version, err := cmd.Flags().GetInt64("version")
		if err != nil {
			fmt.Println("Unable to read flag `version`")
			return
		}

		p := db.NewPersistence()
		defer db.ClosePersistence(p)
		migrator, err := migrations.Init(p)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		err = migrator.Down(version)
		if err != nil {
			fmt.Printf("Unable to run `down` migrations with error %v", err.Error())
			return
		}
	},
}

var migrateStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "display status of each migrations",
	Run: func(cmd *cobra.Command, args []string) {
		p := db.NewPersistence()
		defer db.ClosePersistence(p)
		migrator, err := migrations.Init(p)
		if err != nil {
			fmt.Println("Unable to fetch migrator")
			return
		}

		if err := migrator.MigrationStatus(); err != nil {
			fmt.Println("Unable to fetch migration status")
			return
		}
	},
}

func init() {
	fmt.Println("init migration")
	// Add "--name" flag to "create" command
	migrateCreateCmd.Flags().StringP("name", "n", "", "Name for the migration")

	// Add "--version" flag to both "up" and "down" command
	migrateUpCmd.Flags().Int64P("version", "v", 0, "The version of migrations to execute")
	migrateDownCmd.Flags().Int64P("version", "v", 0, "The version of migrations to execute")

	// Add "create", "up" and "down" commands to the "migrate" command
	migrateCmd.AddCommand(migrateUpCmd, migrateDownCmd, migrateCreateCmd, migrateStatusCmd)

	// Add "migrate" command to the root command
	rootCmd.AddCommand(migrateCmd)
}
