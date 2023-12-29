package tabl

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const dbPath = "DB"

func CreateTabl(name string) error {

	tablName := filepath.Join("./", dbPath, name+".tabl")
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

	// Write text to the file
	now := time.Now()
	init := fmt.Sprintf("init: %s\n", now.Format("2006-01-02T15:04:05Z"))

	_, err = file.WriteString(init)
	if err != nil {
		return fmt.Errorf("error writing to meta file for %s", tablName)
	}

	// fmt.Println("File created successfully!")
	return nil
}
