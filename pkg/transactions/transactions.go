package transactions

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Transaction struct {
	Date time.Time
	Amount int64
}

func SumConcurrently(transactions map[string][]Transaction, goroutines int) int64 { //12 goroutines
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	total := int64(0)
	for i := 0; i < goroutines; i++ {
		// ВАЖНО: просто с partSize не прокатит, вам нужно как-то заранее разделить слайс по месяцам
		// ВАЖНО: этот код - лишь шаблон, который показывает вам как запустить горутину
		var part []Transaction
		var pointerOnMonth string

			if i < 9 {
				a := strconv.Itoa(i + 1)
				part = transactions["2020-0" + a]
				pointerOnMonth = "2020-0" + a + ":"
			}

			if i >= 9 {
				a := strconv.Itoa(i + 1)
				part = transactions["2020-" + a]
				pointerOnMonth = "2020-" + a + ":"
			}

		go func() {
			sum := sum(part)
			fmt.Println(pointerOnMonth, sum)
			wg.Done()
			total += sum
		}()
	}

	wg.Wait()
	return total
}


func sum(transactions []Transaction) int64  {
	result := int64(0)
	for _, transaction := range transactions {
		result += transaction.Amount
	}
	return result
}
