// Package command contains cli commands.
package command

import (
	"os"
	"path/filepath"

	"github.com/iqoption/nap"
	"github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/antelman107/metrics/cli"

	"github.com/antelman107/metrics/definition/database"
	"github.com/antelman107/metrics/definition/logger"
)

var (
	// Migrations path.
	migrationsPath string

	// Migrate command group.
	migrateDatabaseCmd = &cobra.Command{
		Use:   "database [command]",
		Short: "Database migrations",
	}

	// Migrate up command.
	migrateDatabaseUpCmd = &cobra.Command{
		Use:   "up",
		Short: "Execute a migration to latest available version",
		Args:  cobra.ExactArgs(0),
		RunE:  migrateUpCmdHandler,
	}

	// Migrate down command.
	migrateDatabaseDownCmd = &cobra.Command{
		Use:   "down",
		Short: "Down migrations",
		Args:  cobra.ExactArgs(0),
		RunE:  migrateDownCmdHandler,
	}
)

// Command init function.
func init() {
	migrateCmd.AddCommand(migrateDatabaseCmd)
	migrateDatabaseCmd.AddCommand(migrateDatabaseUpCmd)
	migrateDatabaseCmd.AddCommand(migrateDatabaseDownCmd)

	var err error
	var appPath string

	// Save application path
	if appPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		cli.StopWithError(err)
	}

	migrateDatabaseCmd.PersistentFlags().StringVarP(&migrationsPath, "migrationsPath", "m", appPath+"/migrations", "Path to migrations")

	// Set custom migration table
	migrate.SetTable("migration")
}

// Command handler func.
func migrateUpCmdHandler(_ *cobra.Command, _ []string) (err error) {
	migrationsList := &migrate.FileMigrationSource{
		Dir: migrationsPath,
	}

	var db *nap.DB
	if err = diContainer.Fill(database.DefPostgres, &db); err != nil {
		return err
	}

	var log logger.Logger
	if err = diContainer.Fill(logger.DefLogger, &log); err != nil {
		return err
	}

	var n int
	if n, err = migrate.Exec(db.Master(), "postgres", migrationsList, migrate.Up); err != nil {
		return err
	}

	log.Info("Applied migrations", zap.Int("count", n))

	return nil
}

// Command handler func.
func migrateDownCmdHandler(_ *cobra.Command, _ []string) (err error) {
	migrationsList := &migrate.FileMigrationSource{
		Dir: migrationsPath,
	}

	var db *nap.DB
	if err = diContainer.Fill(database.DefPostgres, &db); err != nil {
		return err
	}

	var log logger.Logger
	if err = diContainer.Fill(logger.DefLogger, &log); err != nil {
		return err
	}

	var n int
	if n, err = migrate.ExecMax(db.Master(), "postgres", migrationsList, migrate.Down, 1); err != nil {
		return err
	}

	log.Info("Down migrations", zap.Int("count", n))

	return nil
}
