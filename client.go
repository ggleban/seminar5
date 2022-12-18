package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Note struct {
	name    string
	surname string
	note    string
}

var httpClient = &http.Client{}

func main() {
	f := Input()
	f.PostNote()

	fl := true
	for fl == true {
		fl = Choose()
	}
}

func Choose() bool {

	fmt.Printf("\nWhat's next?\n")
	fmt.Printf("\ny - proceed, n - finish, p - show all\n")

	var choice string
	_, err := fmt.Scanf("%s\n", &choice)
	if err != nil {
		return false
	}

	if choice == "n" {
		return false
	}

	if choice == "y" {
		f := Input()
		f.PostNote()
		return true
	}

	if choice == "p" {
		PrintNotes()
		return true
	}

	CallClear()
	fmt.Printf("Unidentified command, please try again\n")
	return true
}

func Input() Note {
	CallClear()

	Scan := bufio.NewScanner(os.Stdin)

	fmt.Println("Name")
	Scan.Scan()
	name := Scan.Text()

	fmt.Println("Surname")
	Scan.Scan()
	surname := Scan.Text()

	fmt.Println("Info")
	Scan.Scan()
	note := Scan.Text()

	f := Note{name, surname, note}

	return f
}

func (f *Note) PostNote() {
	marshal, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}

	bb := bytes.Buffer{}
	bb.Write(marshal)

	req, err := http.NewRequest("POST", "http://127.0.0.1:4000/save_note", &bb)
	if err != nil {
		log.Fatal(err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == 200 {
		fmt.Printf("\nEntered data:")
		fmt.Printf("\nName - %s, Surname - %s, Info - %s\n", f.name, f.surname, f.note)
	} else {
		log.Fatal(err)
	}
}

func PrintNotes() {
	var notes []Note

	res, err := http.Get("http://127.0.0.1:4000/get_notes")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if json.Unmarshal(body, &notes) != nil {
		log.Fatal(err)
		return
	}

	CallClear()
	for i, note := range notes {
		fmt.Printf("Note № %d\n", i+1)
		fmt.Printf("Name - %s\nSurname - %s\nInfo - %s\n\n", note.name, note.surname, note.note)
	}
}

func CallClear() {
	fmt.Print("\033[H\033[2J")
}
