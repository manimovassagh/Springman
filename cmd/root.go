package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command
var rootCmd = &cobra.Command{
	Use:   "springman",
	Short: "🚀 Springman: CLI tool for Spring Boot project automation",
	Long: `Springman helps generate, manage, and scaffold Spring Boot apps easily.

Usage:
  springman new <project-name>`,
}

// Execute runs the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(newCmd)
}
