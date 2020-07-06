package main

import (
	"fmt"
	"github.com/Geniuskaa/task5.1/pkg/transactions"
	"time"
)

func main() {



	tranz := []transactions.Transaction{
		{time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local), 20_396_00},
		{time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local), 245_00},
		{time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local), 2_694_00},
		{time.Date(2020, 5, 2, 0, 0, 0, 0, time.Local), 19_294_00},
		{time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local), 395_294_00},
		{time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local), 102_00},
		{time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local), 295_00},
		{time.Date(2020, 4, 11, 0, 0, 0, 0, time.Local), 12_496_00},
		{time.Date(2020, 3, 20, 0, 0, 0, 0, time.Local), 46_094_00},
		{time.Date(2020, 3, 21, 0, 0, 0, 0, time.Local), 294_00},
		{time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local), 13_285_00},
	}

	splitTranz := make(map[string][]transactions.Transaction)

	for _, v := range tranz {
		splitTranz[v.Date.Format("2006-01")] = append(splitTranz[v.Date.Format("2006-01")], v)
	}

	a := transactions.SumConcurrently(splitTranz,12)
	fmt.Println("Total:", a)

}


