package expense

import (
	"expense/pkg/graphModels"
)
func AddExpense(userID, note, amount, title, date, status string ) (int, error) {
	stmt := `INSERT INTO expense (note, amount, userId, date, title, status)
				VALUES(?, ?, ?, ?, ?, ?)`

	row, err := database.Db.Exec(stmt, note, amount, userID, date, title, status)
	if err != nil {
		return 0, err
	}
	expenseId, err := row.LastInsertId()
	return int(expenseId), nil
}

