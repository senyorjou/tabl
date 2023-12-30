package tabl

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type Row struct {
	ID   string // UUID
	Init int64  // Unix epoch time in milliseconds
	Body string // Body of the row
}

func NewRow(body string) *Row {
	return &Row{
		ID:   uuid.New().String(),
		Init: time.Now().UnixNano() / int64(time.Millisecond),
		Body: body,
	}
}

func OldRow(id string, body string) *Row {
	return &Row{
		ID:   id,
		Init: time.Now().UnixNano() / int64(time.Millisecond),
		Body: body,
	}
}

func (r *Row) String() string {
	return fmt.Sprintf("%s,%d,%s\n", r.ID, r.Init, r.Body)
}

// display only id and date in a compact format
func (r *Row) Compact() string {
	return fmt.Sprintf("id:%s init:%d", r.ID, r.Init)
}

// display only row date in human readable format
func (r *Row) Date() string {
	return time.Unix(r.Init/1000, 0).UTC().String()
}

func InsertRow(tabl, col, body string) (*Row, error) {
	row := NewRow(body)
	res, err := addRow(tabl, col, row)
	return res, err
}

func UpdateRow(tabl, col, id, body string) (*Row, error) {
	row := OldRow(id, body)
	res, err := addRow(tabl, col, row)
	return res, err
}

func addRow(tabl, col string, row *Row) (*Row, error) {
	file := ColPath(tabl, col)
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
		return row, err
	}
	defer f.Close()

	// Write the string to the file
	bytes, err := f.WriteString(row.String())
	if err != nil {
		log.Fatalf("Failed writing to file: %s", err)
		return row, err
	}
	log.Printf("Wrote %d bytes to %s\n", bytes, file)

	return row, nil
}
