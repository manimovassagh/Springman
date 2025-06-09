package cmd

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type Dependency struct {
	XMLName    xml.Name `xml:"dependency"`
	GroupID    string   `xml:"groupId"`
	ArtifactID string   `xml:"artifactId"`
	Version    string   `xml:"version,omitempty"`
}

type Project struct {
	XMLName      xml.Name      `xml:"project"`
	Dependencies *Dependencies `xml:"dependencies"`
}

type Dependencies struct {
	XMLName    xml.Name     `xml:"dependencies"`
	Dependency []Dependency `xml:"dependency"`
}

var addCmd = &cobra.Command{
	Use:   "add [project folder] [groupId:artifactId[:version]]",
	Short: "Add a Maven dependency to pom.xml",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("❌ Usage: springman add <project-folder> <groupId:artifactId[:version]>")
			return
		}

		projectDir := args[0]
		coordinates := args[1]
		parts := strings.Split(coordinates, ":")

		if len(parts) < 2 || len(parts) > 3 {
			fmt.Println("❌ Invalid coordinates. Use groupId:artifactId[:version] format.")
			return
		}

		groupId := parts[0]
		artifactId := parts[1]
		version := ""
		if len(parts) == 3 {
			version = parts[2]
		}

		pomPath := filepath.Join(projectDir, "pom.xml")
		data, err := os.ReadFile(pomPath)
		if err != nil {
			fmt.Printf("❌ Failed to read pom.xml: %v\n", err)
			return
		}

		var project Project
		err = xml.Unmarshal(data, &project)
		if err != nil {
			fmt.Printf("❌ Failed to parse pom.xml: %v\n", err)
			return
		}

		if project.Dependencies == nil {
			project.Dependencies = &Dependencies{}
		}

		// Check for duplicates
		for _, dep := range project.Dependencies.Dependency {
			if dep.GroupID == groupId && dep.ArtifactID == artifactId {
				fmt.Println("⚠️ Dependency already exists.")
				return
			}
		}

		newDep := Dependency{
			GroupID:    groupId,
			ArtifactID: artifactId,
			Version:    version,
		}

		project.Dependencies.Dependency = append(project.Dependencies.Dependency, newDep)

		output, err := xml.MarshalIndent(project, "", "  ")
		if err != nil {
			fmt.Printf("❌ Failed to marshal updated pom.xml: %v\n", err)
			return
		}

		err = os.WriteFile(pomPath, output, 0644)
		if err != nil {
			fmt.Printf("❌ Failed to write updated pom.xml: %v\n", err)
			return
		}

		fmt.Println("✅ Dependency added successfully.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}