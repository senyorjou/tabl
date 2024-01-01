package main

import (
	"fmt"
	"tabl"
)

var lipsum = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum`

func main() {
	// err := tabl.CreateTabl("foo")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err := tabl.CreateCol("foo", "title")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	row, _ := tabl.InsertRow("foo", "title", "Some content")
	fmt.Println("Created:", row.ID)
	row, _ = tabl.UpdateRow("foo", "title", row.ID, "Some more content")
	fmt.Println("Updated:", row.ID)

}
