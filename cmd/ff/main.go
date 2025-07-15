package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"unicode"
)

func printUsage() {
	fmt.Println("Usage: ff [--js=<true|false>] [--title=<true|false>] [page|comp] [Name]")
	os.Exit(1)
}
func dirOrFileExistsError(path string) {
	errorMsg := fmt.Sprintf("File or Directory %s already exists", path)
	fmt.Println(errorMsg)
	os.Exit(1)
}
func createFiles(path, name, ext string) {
	doesDirExist := dirExists(path)
	if doesDirExist {
		dirOrFileExistsError(path)
	}
	createDir(path)

	cssFile := "styles.css"
	cssPath := filepath.Join(path, cssFile)
	cssPathExist := dirExists(cssPath)
	if cssPathExist {
		dirOrFileExistsError(cssPath)
	}

	createFile(cssPath)
	filePath := filepath.Join(path, name+ext+"x")
	filePathExsts := dirExists(filePath)
	if filePathExsts {
		dirOrFileExistsError(filePath)
	}
	componentFile := createFile(filePath)
	err := writeComponentFile(componentFile, name)
	if err != nil {

	}
	indexPath := filepath.Join(path, "index"+ext)
	indexPathExists := dirExists(indexPath)
	if indexPathExists {
		dirOrFileExistsError(indexPath)
	}
	indexFile := createFile(indexPath)
	indexErr := writeIndexFile(indexFile, name)
	if indexErr != nil {

	}

}
func writeComponentFile(file *os.File, name string) error {
	defer file.Close()
	componentContent := fmt.Sprintf(`import React from 'react'

const %s = () => {
  return (
    <div>%s</div>
  )
}

export default %s
`, name, name, name)
	if _, err := file.WriteString(componentContent); err != nil {
		log.Fatal(err)
	}
	return nil
}
func writeIndexFile(file *os.File, name string) error {
	defer file.Close()
	indexContent := fmt.Sprintf(`import %s from "./%s";
export default %s;
`, name, name, name)

	// Write to index file
	if _, err := file.WriteString(indexContent); err != nil {
		log.Fatal(err)
	}
	return nil
}
func createDir(path string) {

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating:", err)
		os.Exit(1)
	}
	fmt.Println("✅ Created:", path)
}
func createFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating:", err)
		os.Exit(1)
	}
	fmt.Println("✅ Created:", path)
	return file
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
func createBaseDir(candidates []string) string {
	for _, c := range candidates {
		path := filepath.Join("src", c)
		fmt.Println(path, "path")
		if dirExists(path) {
			fmt.Println(path, "path exists")
			return path
		}
	}
	return fmt.Sprintf("./src/%s", candidates[0])
}
func getTitleCase(s string) string {
	if s == "" {
		return s
	}
	// Get first rune and uppercase it
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
func main() {
	isTitleCase := flag.Bool("title", true, "Title cases the name by default, --title=false to keep the casing as is")
	isJs := flag.Bool("js", false, ".ts files by default, add --js=true to switch to js")

	// Parse flags
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		printUsage()
	}

	kind := args[0]
	name := args[1]

	baseDir := ""

	switch kind {
	case "page":
		{
			candidates := []string{"Pages", "pages"}
			baseDir = createBaseDir(candidates)
			fmt.Println(baseDir, "base")

		}
	case "comp":
		{
			candidates := []string{"Components", "components"}
			baseDir = createBaseDir(candidates)
			fmt.Println(baseDir, "base")
		}

	default:
		fmt.Println("Unknown type:", kind)
		os.Exit(1)
	}
	ext := ""
	if *isJs {
		ext = ".js"
	} else {
		ext = ".ts"
	}
	if *isTitleCase {
		name = getTitleCase(name)
	}
	finalPath := filepath.Join(baseDir, name)

	createFiles(finalPath, name, ext)

}
