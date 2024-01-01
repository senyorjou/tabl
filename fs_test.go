package tabl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateTabl(t *testing.T) {
	dbPathfn = func() string {
		return "TestDB"
	}
	// Define a test table name
	testTabl := "testTabl"

	// Call CreateTabl
	err := CreateTabl(testTabl)
	if err != nil {
		t.Fatalf("CreateTabl failed: %v", err)
	}

	// Check if the table directory was created
	tablDir := filepath.Join("./", dbPathfn(), testTabl+".tabl")
	_, err = os.Stat(tablDir)
	if os.IsNotExist(err) {
		t.Fatalf("Table directory was not created: %v", err)
	}

	// Check if the meta file was created
	metaFile := filepath.Join(tablDir, "meta")
	_, err = os.Stat(metaFile)
	if os.IsNotExist(err) {
		t.Fatalf("Meta file was not created: %v", err)
	}

	// Clean up
	os.RemoveAll(tablDir)
	os.RemoveAll(dbPathfn())

}
func TestCreateCol(t *testing.T) {
	dbPathfn = func() string {
		return "TestDB"
	}

	// Define a test table and column name
	testTabl := "testTabl"
	tablDir := filepath.Join("./", dbPathfn(), testTabl+".tabl")

	testCol := "testCol"
	err := CreateTabl(testTabl)
	if err != nil {
		t.Fatalf("CreateTabl failed: %v", err)
	}
	// Call CreateCol
	err = CreateCol(testTabl, testCol)
	if err != nil {
		t.Fatalf("CreateCol failed: %v", err)
	}

	// Check if the column directory was created
	colDir := filepath.Join("./", dbPathfn(), testTabl+".tabl", testCol+".col")
	_, err = os.Stat(colDir)
	if err == nil {
		t.Fatalf("Column directory was not created: %v", err)
	}

	// Clean up
	os.RemoveAll(tablDir)
	os.RemoveAll(dbPathfn())

}
func TestCleanFilename(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test 1", "Hello, World!", "helloworld"},
		{"Test 2", "a/b\\c:d*e?f\"g<h>i|j", "abcdefghij"},
		{"Test 3", ".....Hello.....", "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := cleanFilename(tt.input)
			if output != tt.expected {
				t.Errorf("got %q, want %q", output, tt.expected)
			}
		})
	}
}
