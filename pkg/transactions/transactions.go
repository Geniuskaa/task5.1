package transactions

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Transaction struct {
	Date int64
	Amount int64
}


func ConvertFromTransactionsSliceToMap(slice []Transaction) (map[string][]int64, [][]int64) {
	mapOfTransactions := make(map[string][]int64)
	//sliceOfSumTransaction := []int64{}

	for _, v := range slice {
		t := time.Unix(v.Date, 0)
		year := strconv.Itoa(t.Year())
		month := t.Month().String()
		mapOfTransactions[year + "," + month] = append(mapOfTransactions[year + "," + month], v.Amount)
	}

	a := [][]int64{}
	for _,l := range mapOfTransactions {
		a = append(a, l)
	}

	return mapOfTransactions, a
}

func SumConcurrently(transactions [][]int64, goroutines int) int64 {
	if goroutines > len(transactions) {
		goroutines = len(transactions)
	}
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	total := int64(0)
	//partSize := len(transactions)
	for i := 0; i < goroutines; i++ {
		// ВАЖНО: просто с partSize не прокатит, вам нужно как-то заранее разделить слайс по месяцам
		// ВАЖНО: этот код - лишь шаблон, который показывает вам как запустить горутину
		part := transactions[i]
		go func() {
			sum := sum(part)
			total += sum
			fmt.Println(sum)
			wg.Done()
		}()
	}

	wg.Wait()
	return total
}


func sum(transactions []int64 ) int64  {
	result := int64(0)
	for _, transaction := range transactions {
		result += transaction
	}
	return result
}

