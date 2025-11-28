package main

import (
	expenseendpoints "expense-tracker/expenseEndpoints"
	"flag"
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "add":
		{
			addFlag := flag.NewFlagSet("add", flag.ExitOnError)
			description := addFlag.String("description", "", "Expense description")
			amount := addFlag.Int("amount", 0, "Expense amount")

			addFlag.Parse(os.Args[2:])
			expenseendpoints.CreateExpense(*description, *amount)
		}
	case "list":
		{
			expenseendpoints.GetAllExpenses()
		}
	case "summary":
		{
			if len(os.Args) == 2 {
				expenseendpoints.GetSummary()
				return
			}
			summaryFlag := flag.NewFlagSet("summary", flag.ExitOnError)

			month := summaryFlag.Int("month", 0, "Month for summary")

			summaryFlag.Parse(os.Args[2:])

			expenseendpoints.GetSummaryByMonth(*month)
		}
	case "delete":
		{
			deleteFlag := flag.NewFlagSet("delete", flag.ExitOnError)

			id := deleteFlag.Int("id", 0, "Expense id")
			deleteFlag.Parse(os.Args[2:])
			expenseendpoints.DeleteExpense(*id)
		}
	default:
		{
			fmt.Println("There is no such a command")
		}
	}

}
