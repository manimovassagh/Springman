package cmd

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var buildTool string // "maven" or "gradle"

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new Spring Boot project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("❌ Please provide a project name. Example: springman new myapp --build gradle")
			return
		}

		name := args[0]
		fmt.Printf("📦 Generating %s project: %s\n", buildTool, name)

		// Set build type and wrapper param
		buildType := "maven-project"
		wrapperFlag := "withMavenWrapper"
		if buildTool == "gradle" {
			buildType = "gradle-project"
			wrapperFlag = "withGradleWrapper"
		}

		// Construct download URL (no baseDir)
		projectURL := fmt.Sprintf(
			"https://start.spring.io/starter.zip?type=%s&language=java&bootVersion=3.3.0&groupId=com.example&artifactId=%s&name=%s&packageName=com.example.%s&dependencies=web&%s=true",
			buildType, name, name, name, wrapperFlag,
		)

		req, err := http.NewRequest("GET", projectURL, nil)
		if err != nil {
			log.Fatalf("❌ Failed to create request: %v", err)
		}
		req.Header.Set("User-Agent", "SpringmanCLI/1.0")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("❌ HTTP request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("❌ Spring Initializr returned error: %s", resp.Status)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("❌ Failed to read response body: %v", err)
		}

		r, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
		if err != nil {
			log.Fatalf("❌ Failed to read zip archive: %v", err)
		}

		for _, f := range r.File {
			path := filepath.Join(name, f.Name)

			if f.FileInfo().IsDir() {
				os.MkdirAll(path, f.Mode())
				continue
			}

			if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
				log.Fatalf("❌ Failed to create directory %s: %v", filepath.Dir(path), err)
			}

			dst, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				log.Fatalf("❌ Failed to create file %s: %v", path, err)
			}

			src, err := f.Open()
			if err != nil {
				log.Fatalf("❌ Failed to open zip content: %v", err)
			}

			_, err = io.Copy(dst, src)
			dst.Close()
			src.Close()
			if err != nil {
				log.Fatalf("❌ Failed to copy file content: %v", err)
			}
		}

		fmt.Println("✅ Spring Boot project created at:", name)
	},
}

func init() {
	newCmd.Flags().StringVarP(&buildTool, "build", "b", "maven", "Build tool to use (maven or gradle)")
}
