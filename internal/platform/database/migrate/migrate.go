package migrate

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/uptrace/bun/migrate"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "database migrations",
	}
}

func NewMigrateCommand(rootCmd *cobra.Command, migrator *migrate.Migrator) {
	rootCmd.AddCommand(&cobra.Command{
		Short: "create migration tables",
		Use:   "init",
		Run: func(cmd *cobra.Command, args []string) {
			migrator.Init(cmd.Context())
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Short: "migrate database",
		Use:   "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			err := migrator.Init(cmd.Context())
			if err != nil {
				return
			}
			if err := migrator.Lock(cmd.Context()); err != nil {
				return
			}
			defer migrator.Unlock(cmd.Context()) //nolint:errcheck
			group, err := migrator.Migrate(cmd.Context())
			if err != nil {
				return
			}
			if group.IsZero() {
				fmt.Printf("there are no new migrations to run (database is up to date)\n")
				return
			}
			fmt.Printf("migrated to %s\n", group)
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Short: "rollback the last migration group",
		Use:   "rollback",
		Run: func(cmd *cobra.Command, args []string) {
			if err := migrator.Lock(cmd.Context()); err != nil {
				return
			}
			defer migrator.Unlock(cmd.Context()) //nolint:errcheck

			group, err := migrator.Rollback(cmd.Context())
			if err != nil {
				return
			}
			if group.IsZero() {
				fmt.Printf("there are no groups to roll back\n")
				return
			}
			fmt.Printf("rolled back %s\n", group)

		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Short: "lock migrations",
		Use:   "lock",
		Run: func(cmd *cobra.Command, args []string) {
			migrator.Lock(cmd.Context())
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Short: "unlock migrations",
		Use:   "unlock",
		Run: func(cmd *cobra.Command, args []string) {
			migrator.Unlock(cmd.Context())
		},
	})
}
