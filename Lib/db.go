package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // comment
)

// Test function
func Test() {

	dbString := "root:dieppe@/balances"

	conn, err := sql.Open("mysql", dbString)
	if err != nil {
		fmt.Println("Database failed to open")
		fmt.Println(err.Error())

	} else {
		// 	execute the query
		// Company name and Address, VAT number, Due Date, period year, period month, box numbers,
		var sqlcode string
		sqlcode = "SELECT DueDate, Month(DueDate), Year(DueDate), VATBoxes, Company.Company, Concat(Address1, ',', Address2, ',', Address3, ',', Address4) "
		sqlcode = sqlcode + "FROM ((payments inner join expensecode on Payments.Nominal = ExpenseCOde.ExpenseCategoryID) inner join zMetaNominals on ExpenseCode.MetaNominal = zMetaNominals.Id) inner join company on Payments.Company = Company.Compid "
		sqlcode = sqlcode + "where zMetaNominals.MetaCode = 'VAT' "
		sqlcode = sqlcode + "and Payments.DueDate > DATE_SUB(Now(), INTERVAL 180 DAY) "
		sqlcode = sqlcode + "and ManualVatReceiptID is Null "

		rows, err := conn.Query(sqlcode)
		if err != nil {
			fmt.Println(err.Error())
		}

		var DueDate string
		var Month string
		var Year string
		var VatBoxes string
		var CompanyName string
		var Address string

		for rows.Next() {
			// Scan the value to []byte
			err = rows.Scan(&DueDate, &Month, &Year, &VatBoxes, &CompanyName, &Address)

			rows.Scan()
			if err != nil {
				fmt.Println(err.Error())
			}

			// Use the string value
			fmt.Println(DueDate, Month, Year, VatBoxes, CompanyName, Address)
		}
		conn.Close()
	}
	// CheckError(err)

}
