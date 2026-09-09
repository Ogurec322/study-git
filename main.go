package main

import (
	"context"
	"fmt"
	"postgr/feature_postgres/simple_connection"
	"postgr/feature_postgres/simple_sql"
	"time"
)

func main() {

	ctx := context.Background()

	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	tasks, err := simple_sql.SelectRows(ctx, conn)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		if task.ID == 3 {
			task.Title = "Искупаться в бассейне"
			task.Description = "Не намочив трусы!"
			task.Completed = true
			now := time.Now()
			task.CompletedAt = &now

			if err := simple_sql.UpdateTask(ctx, conn, task); err != nil {
				panic(err)
			}

			break
		}
	}

	fmt.Println("УСПЕХ !")

}
