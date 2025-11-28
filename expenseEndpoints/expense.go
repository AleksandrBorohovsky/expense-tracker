package expenseendpoints

import (
	"encoding/json"
	"expense-tracker/structures"
	"expense-tracker/utils"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

var path = "./storageFiles/expenses.json"

func CreateExpense(description string, amount int) {
	if description == "" || amount < 0 {
		fmt.Println("Invalid data")
		return
	}
	expense := structures.Expense{
		Id:          1,
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	utils.Check(err, "Cannot open the file")
	defer func() {
		if err = file.Close(); err != nil {
			utils.Check(err, "Cannot close file connection")
		}
	}()

	fileData := readFile(path)
	if len(fileData) != 0 {
		expense.Id = fileData[len(fileData)-1].Id + 1
	}
	fileData = append(fileData, expense)
	writeFile(path, fileData)

	fmt.Printf("Expense added successfully (ID:%d)\n", expense.Id)
}

func GetAllExpenses() {
	expenses := readFile(path)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprint(w, "ID\tDate\tDescription\tAmount\n")
	for _, expense := range expenses {
		year, month, day := expense.Date.Date()
		fmt.Fprintf(w, "%v\t%v-%v-%v\t%v\t%v$\n", expense.Id, year, int(month), day, expense.Description, expense.Amount)
	}
	w.Flush()
	// litter.Config.StripPackageNames = true
	// litter.Dump(expenses)
}

func GetSummary() {
	expenses := readFile(path)
	sum := 0
	for _, expense := range expenses {
		sum += expense.Amount
	}

	fmt.Printf("Total expense: %d$\n", sum)
}

func GetSummaryByMonth(month int) {
	if month < 1 || month > 12 {
		fmt.Println("Wrong month number")
		return
	}
	expenses := readFile(path)
	sum := 0
	var Month time.Month
	for _, expense := range expenses {
		if int(expense.Date.Month()) == month {
			Month = expense.Date.Month()
			sum += expense.Amount
		}
	}

	fmt.Printf("Total expenses for %s: %d$\n", Month, sum)
}

func DeleteExpense(id int) {
	expenses := readFile(path)
	result := []structures.Expense{}
	flag := false
	for _, expense := range expenses {
		if expense.Id != id {
			result = append(result, expense)
		} else {
			flag = true
		}
	}
	if !flag {
		fmt.Println("There is no expense with this id")
		return
	}
	writeFile(path, result)
	fmt.Println("Expense deleted successfully")
}

func readFile(path string) []structures.Expense {
	result := []structures.Expense{}
	data, err := os.ReadFile(path)
	utils.Check(err, "Cannot read the file")
	err = json.Unmarshal(data, &result)
	if len(result) != 0 {
		utils.Check(err, "Cannot convert file data")
	}
	return result
}

func writeFile(path string, data []structures.Expense) {
	fileData, err := json.Marshal(data)
	utils.Check(err, "Cannot convert data to json")
	err = os.WriteFile(path, fileData, os.ModePerm)
	utils.Check(err, "Cannot write data to the file")
}
