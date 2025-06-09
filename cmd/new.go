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

		// Map build type
		buildType := "maven-project"
		if buildTool == "gradle" {
			buildType = "gradle-project"
		}

		fmt.Printf("📦 Generating %s project: %s\n", buildTool, name)

		url := fmt.Sprintf("https://start.spring.io/starter.zip?type=%s&language=java&bootVersion=3.2.5&baseDir=%s&dependencies=web", buildType, name)

		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("❌ Failed to download: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("❌ Failed to read zip content: %v", err)
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
				log.Fatalf("❌ Failed to open zip file content: %v", err)
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
