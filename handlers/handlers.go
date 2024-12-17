package handlers

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"noteapp/entities"
	"noteapp/services"
)

func Start() {
	notes, _ := services.LoadNotes()

	for {
		fmt.Println("\n1. Add Note")
		fmt.Println("2. List Note")
		fmt.Println("3. Exit")
		fmt.Println("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addNote(&notes)
		case 2:
			listNotes(notes)
		case 3:
			services.SaveNotes(notes)
			fmt.Println("Notes saved successfully. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addNote(notes *[]entities.Note) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter note title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter note content: ")
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	newNote := entities.Note{
		Title: title,
		Content: content,
	}
	*notes = append(*notes, newNote)
	fmt.Println("Note added successfully!")
}

func listNotes(notes []entities.Note) {
	if len(notes) == 0 {
		fmt.Println("No notes available.")
		return
	}

	for i, note := range notes {
		fmt.Printf("\nNote %d\n", i+1)
		fmt.Printf("Title %s\n", note.Title)
		fmt.Printf("Content: %s\n", note.Content)
	}
}