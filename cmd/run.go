package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [project folder]",
	Short: "Run the Spring Boot app",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("❌ Please specify the project folder. Example: springman run myapp")
			return
		}

		projectDir := args[0]

		// Detect if gradlew or mvnw exists
		gradlew := filepath.Join(projectDir, "gradlew")
		mvnw := filepath.Join(projectDir, "mvnw")

		var runCmd *exec.Cmd

		if fileExists(gradlew) {
			fmt.Println("🚀 Running with Gradle...")
			runCmd = exec.Command(gradlew, "bootRun")
		} else if fileExists(mvnw) {
			fmt.Println("🚀 Running with Maven...")
			runCmd = exec.Command(mvnw, "spring-boot:run")
		} else {
			log.Fatalf("❌ Neither gradlew nor mvnw found in %s", projectDir)
		}

		runCmd.Dir = projectDir
		runCmd.Stdout = os.Stdout
		runCmd.Stderr = os.Stderr
		runCmd.Stdin = os.Stdin

		err := runCmd.Run()
		if err != nil {
			log.Fatalf("❌ Failed to run project: %v", err)
		}
	},
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func init() {
	rootCmd.AddCommand(runCmd)
}
