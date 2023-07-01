package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hokageCV/gotrack/db"
	"github.com/hokageCV/gotrack/screen"
	"github.com/hokageCV/gotrack/utils"
)

func main() {
	dbInstance, err := db.InitializeDB()
	utils.CheckNilErr(err)
	defer dbInstance.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		screen.DisplayOptions()

		fmt.Println("\n🧾 Enter your choice:")
		choiceStr, err := reader.ReadString('\n')
		utils.CheckNilErr(err)

		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid choice! Please enter a number.")
			continue
		}

		if choice == 6 {
			screen.Exit()
		}

		screen.HandleChoice(dbInstance, choice)

	}
}
