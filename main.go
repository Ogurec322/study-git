package main

import (
	"fmt"
	"postgr/feature1"
	"postgr/feature2"
	"postgr/feature_postgres/simple_connection"
)

func main() {
	fmt.Println("Hello git hubby")

	feature1.Feature1()
	feature2.Feature2()

	simple_connection.CheckConnection()

}
