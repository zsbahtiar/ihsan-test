package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/zsbahtiar/ihsan-test/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "bankapp - @zsbahtiar",
	Short: "bank application with account management",
}

var cfg *config.Config

func init() {
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	rootCmd.AddCommand(migrateCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
