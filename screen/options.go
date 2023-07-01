package screen

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/hokageCV/gotrack/db"
	"github.com/hokageCV/gotrack/utils"
)

func DisplayOptions() {
	fmt.Printf("\n")
	fmt.Println("1. Show All Tasks")
	fmt.Println("2. Create Task")
	fmt.Println("3. Mark a Task Done")
	fmt.Println("4. Edit Task Title")
	fmt.Println("5. Delete Task")
	fmt.Println("6. Exit")

}

func GetTasks(dbInstance *sql.DB) {
	db.DisplayTasksFromDB(dbInstance)
}

func AddTask(dbInstance *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter task title: ")
	title, err := reader.ReadString('\n')
	utils.CheckNilErr(err)
	title = strings.TrimSpace(title)

	db.CreateTaskInDB(dbInstance, title)
	fmt.Println("Task added successfully!")
}

func MarkTaskDone(dbInstance *sql.DB) {

	var taskID int
	fmt.Println("Enter task ID:")
	_, err := fmt.Scanf("%d", &taskID)
	utils.CheckNilErr(err)

	db.DoneTaskInDB(dbInstance, taskID)
	fmt.Println("Task done successfully!")
}

func EditTask(dbInstance *sql.DB) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter new task title: ")
	editedTask, err := reader.ReadString('\n')
	utils.CheckNilErr(err)

	var taskID int
	fmt.Println("Enter task ID:")
	_, error := fmt.Scanf("%d", &taskID)
	utils.CheckNilErr(error)
	editedTask = strings.TrimSpace(editedTask)

	db.EditTaskInDB(dbInstance, editedTask, taskID)
	fmt.Println("Task edited successfully!")
}

func DeleteTask(dbInstance *sql.DB) {

	var taskID int
	fmt.Println("Enter task ID:")
	_, err := fmt.Scanf("%d", &taskID)
	utils.CheckNilErr(err)

	db.DeleteTaskFromDB(dbInstance, taskID)
	fmt.Println("Task deleted successfully!")
}

func Exit() {
	fmt.Println("Exiting...")
	os.Exit(0)
}

func HandleChoice(dbInstance *sql.DB, choice int) {
	switch choice {
	case 1:
		GetTasks(dbInstance)
	case 2:
		AddTask(dbInstance)
	case 3:
		MarkTaskDone(dbInstance)
	case 4:
		EditTask(dbInstance)
	case 5:
		DeleteTask(dbInstance)
	default:
		fmt.Println("Invalid choice!")
	}
}
