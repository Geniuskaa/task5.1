package main

import (
	"fmt"
	"sort"
	"github.com/Geniuskaa/task5.1/pkg/transactions"
)

func main() {
	const users = 40
	const transactionsPerUser = 10
	const transactionsAmount = 1_00


	transactionsOfService := make([]transactions.TransactionInfo, users * transactionsPerUser) //всего 400

	transactions.DateAdding(0,80,transactionsOfService,1) //80
	transactions.DateAdding(80,140,transactionsOfService,8) //60
	transactions.DateAdding(140,240,transactionsOfService,2) // 100
	transactions.DateAdding(240,300,transactionsOfService,11) // 60
	transactions.DateAdding(300,400,transactionsOfService,6)  // 100


	sort.SliceStable(transactionsOfService, func(i, j int) bool {
		return transactionsOfService[i].Date < transactionsOfService[j].Date })
	sum := transactions.SumConcurrently(transactionsOfService,5)
	fmt.Println(sum)



}


