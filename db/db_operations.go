package db

import (
	"database/sql"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/hokageCV/gotrack/utils"
)

func DisplayTasksFromDB(db *sql.DB) {
	// using query because it returns rows, unlike exec which returns metadata
	rows, err := db.Query("SELECT * FROM tasks")
	utils.CheckNilErr(err)
	defer rows.Close()

	// tabwriter for table structure
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Printf("\n")
	fmt.Fprintln(tw, "----------------------")
	fmt.Fprintf(tw, "Task ID\tTitle\tIs Done\n")

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.IsDone)
		utils.CheckNilErr(err)

		fmt.Fprintf(tw, "%d\t%s\t%v\n", task.ID, task.Title, task.IsDone)
	}

	fmt.Fprintln(tw, "----------------------")
	fmt.Printf("\n")

	tw.Flush()

}

func CreateTaskInDB(db *sql.DB, task Task) error {

	// instead of writing values into query, binding it using ? to avoid sql injection

	// prepare
	stmt, err := db.Prepare("INSERT INTO tasks(Title, IsDone) VALUES(?, ?)")
	utils.CheckNilErr(err)
	defer stmt.Close()

	// Execute statement with user values
	_, err = stmt.Exec(task.Title, false)
	utils.CheckNilErr(err)

	return nil
}

func DoneTaskInDB(db *sql.DB, taskID int) {
	stmt, err := db.Prepare("UPDATE tasks SET IsDone=true WHERE ID=?")
	utils.CheckNilErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(taskID)
	utils.CheckNilErr(err)
}

func EditTaskInDB(db *sql.DB, editedTask string, taskID int) {
	stmt, err := db.Prepare("UPDATE tasks SET title=? WHERE ID=?")
	utils.CheckNilErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(editedTask, taskID)
	utils.CheckNilErr(err)
}

func DeleteTaskFromDB(db *sql.DB, taskID int) {
	stmt, err := db.Prepare("DELETE FROM tasks WHERE ID=?")
	utils.CheckNilErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(taskID)
	utils.CheckNilErr(err)
}
