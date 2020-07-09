package main

import (
	"fmt"
	"github.com/Geniuskaa/task5.1/pkg/transactions"
)

func main() {

	tranz := []transactions.Transaction{
		{1583020800, 20_396_00},
		{1585699200, 245_00},
		{1588291200, 2_694_00},
		{1588377600, 19_294_00},
		{1593561600, 395_294_00},
		{1598918400, 102_00},
		{1604188800, 295_00},
		{1586563200, 12_496_00},
		{1584662400, 46_094_00},
		{1584748800, 294_00},
		{1577836800, 13_285_00},
}


	mapOfTranz, sliceOfSums := transactions.ConvertFromTransactionsSliceToMap(tranz)



	fmt.Println(sliceOfSums)
	fmt.Println(mapOfTranz)
	fmt.Println(len(mapOfTranz))

	fmt.Println("Total:",transactions.SumConcurrently(sliceOfSums,6))

}
