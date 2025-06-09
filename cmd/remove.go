package cmd

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type RemoveProject struct {
	XMLName      xml.Name    `xml:"project"`
	Dependencies *RemoveDeps `xml:"dependencies"`
}

type RemoveDeps struct {
	XMLName    xml.Name    `xml:"dependencies"`
	Dependency []RemoveDep `xml:"dependency"`
}

type RemoveDep struct {
	XMLName    xml.Name `xml:"dependency"`
	GroupID    string   `xml:"groupId"`
	ArtifactID string   `xml:"artifactId"`
	Version    string   `xml:"version,omitempty"`
}

var removeCmd = &cobra.Command{
	Use:   "remove [project folder] [groupId:artifactId]",
	Short: "Remove a Maven dependency from pom.xml",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("❌ Usage: springman remove <project-folder> <groupId:artifactId>")
			return
		}

		projectDir := args[0]
		coordinates := args[1]
		parts := strings.Split(coordinates, ":")

		if len(parts) != 2 {
			fmt.Println("❌ Invalid coordinates. Use groupId:artifactId format.")
			return
		}

		groupId := parts[0]
		artifactId := parts[1]

		pomPath := filepath.Join(projectDir, "pom.xml")
		data, err := os.ReadFile(pomPath)
		if err != nil {
			fmt.Printf("❌ Failed to read pom.xml: %v\n", err)
			return
		}

		var project RemoveProject
		err = xml.Unmarshal(data, &project)
		if err != nil {
			fmt.Printf("❌ Failed to parse pom.xml: %v\n", err)
			return
		}

		if project.Dependencies == nil {
			fmt.Println("⚠️ No dependencies found.")
			return
		}

		updatedDeps := make([]RemoveDep, 0)
		found := false
		for _, dep := range project.Dependencies.Dependency {
			if dep.GroupID == groupId && dep.ArtifactID == artifactId {
				found = true
				continue
			}
			updatedDeps = append(updatedDeps, dep)
		}

		if !found {
			fmt.Println("⚠️ Dependency not found.")
			return
		}

		project.Dependencies.Dependency = updatedDeps

		output, err := xml.MarshalIndent(project, "", "  ")
		if err != nil {
			fmt.Printf("❌ Failed to marshal updated pom.xml: %v\n", err)
			return
		}

		outputWithHeader := []byte(xml.Header + string(output) + "\n")
		err = os.WriteFile(pomPath, outputWithHeader, 0644)
		if err != nil {
			fmt.Printf("❌ Failed to write updated pom.xml: %v\n", err)
			return
		}

		fmt.Println("✅ Dependency removed successfully.")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
