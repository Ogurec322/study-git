package simple_sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) ([]TaskModel, error) {
	sqlQuery := `
	SELECT id , title, description , completed , created_at, completed_at
	FROM tasks;
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]TaskModel, 0)

	for rows.Next() {
		var task TaskModel

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAT,
			&task.CompletedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)

		PrintTask(task)
	}

	return tasks, nil
}

func PrintTask(task TaskModel) {
	fmt.Println("--------------------------------------")
	fmt.Println("id: ", task.ID)
	fmt.Println("title: ", task.Title)
	fmt.Println("description: ", task.Description)
	fmt.Println("completed: ", task.Completed)
	fmt.Println("created at: ", task.CreatedAT)
	fmt.Println("completed at: ", task.CompletedAt)
	fmt.Println("--------------------------------------")

}
