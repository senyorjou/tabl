package tabl

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Row struct {
	ID     string // UUID
	Millis int64  // Unix epoch time in milliseconds
	Body   string // Body of the row
}

func New(id, body string) *Row {
	if id == "" {
		id = uuid.New().String()
	}
	return &Row{
		ID:     id,
		Millis: time.Now().UnixMilli(),
		Body:   body,
	}
}

func (r *Row) String() string {
	return fmt.Sprintf("%s,%d,%s\n", r.ID, r.Millis, r.Body)
}

// display only id and date in a compact format
func (r *Row) Compact() string {
	return fmt.Sprintf("id:%s init:%d", r.ID, r.Millis)
}

// display only row date in human readable format
func (r *Row) Date() string {
	return time.Unix(r.Millis, 0).UTC().String()
}

func InsertRow(tabl, col, body string) (*Row, error) {
	row := New("", body)
	res, err := writeRow(tabl, col, row)
	return res, err
}

func UpdateRow(tabl, col, id, body string) (*Row, error) {
	row := New(id, body)
	res, err := writeRow(tabl, col, row)
	return res, err
}
