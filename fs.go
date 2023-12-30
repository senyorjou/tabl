package tabl

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

var dbPathfn = func() string {
	return "DB"
}

func CreateTabl(tabl string) error {

	tablName := filepath.Join("./", dbPathfn(), tabl+".tabl")
	_, err := os.Stat(tablName)
	if err == nil {
		return fmt.Errorf("cannot create %s here", tablName)
	}

	// Directory doesn't exist, create it
	err = os.MkdirAll(tablName, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating %s table", tablName)
	}

	// Directory created, create the file inside
	filename := filepath.Join(tablName, "meta")
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating meta file for %s", filename)
	}
	defer file.Close()

	// add all fields needed for default metas
	data := make(map[string]interface{})
	data["init"] = time.Now()
	data["id"] = "uuid"

	newYamlData, err := yaml.Marshal(&data)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return fmt.Errorf("error marshaling meta file for %s", filename)

	}
	_, err = file.Write(newYamlData)
	if err != nil {
		return fmt.Errorf("error writing to meta file for %s", tablName)
	}

	return nil
}

func cleanFilename(filename string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9\.\-_]`)
	return strings.ToLower(strings.Trim(reg.ReplaceAllString(filename, ""), "."))
}

// create a function that returns the full path of a col file
func ColPath(tabl, col string) string {
	return filepath.Join("./", dbPathfn(), tabl+".tabl", cleanFilename(col)+".col")
}

func CreateCol(tabl, name string) error {
	colName := filepath.Join("./", dbPathfn(), tabl+".tabl")

	filename := filepath.Join(colName, cleanFilename(name)+".col")
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating col file for %s", filename)
	}
	defer file.Close()
	return nil
}
