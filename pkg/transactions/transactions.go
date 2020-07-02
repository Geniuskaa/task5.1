package transactions

import (
	"fmt"
	"sync"
	"time"
)

type TransactionInfo struct {
	Date int64
	SumOfTransaction int64
}

func DateAdding(iStartPosition int64, iMaxPosition int64, slice []TransactionInfo, month time.Month)  {
	for i := iStartPosition; i < iMaxPosition; i++ {
		slice[i] = TransactionInfo{
			Date: dateOfTransaction(month),
			SumOfTransaction: 1_00,
		}
	}
	return
}

func dateOfTransaction(month time.Month) int64 {
	t := time.Date(2020, month, 01, 00, 0, 0, 0, time.Local)
	unix := t.Unix()
	return unix
}

func SumConcurrently(transactions []TransactionInfo, goroutines int) []int64 {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	total := []int64{}
	index := 0
	point := 0
	k := 0
	lengthOfFirstMonth := transactions//[]TransactionInfo
	lengthOfSecondMonth := transactions
	lengthOfThirdMonth := transactions
	lengthOfThourthMonth := transactions
	lengthOfThifthMonth := transactions


	for _, m := range transactions {
		index++

		if k == 0 && transactions[point].Date != m.Date {
			lengthOfFirstMonth = lengthOfFirstMonth[point : index - 1]
			point = index
			k++
		}
		if k == 1 && transactions[point].Date != m.Date {
			lengthOfSecondMonth = lengthOfSecondMonth[point : index]
			point = index
			k++
		}
		if k == 2 && transactions[point].Date != m.Date {
			lengthOfThirdMonth = lengthOfThirdMonth[point : index]
			point = index
			k++
		}
		if k == 3 && transactions[point].Date != m.Date {
			lengthOfThourthMonth = lengthOfThourthMonth[point : index]
			point = index
			k++
		}
		if k == 4 && transactions[point].Date != m.Date {
			lengthOfThifthMonth = lengthOfThifthMonth[point : index]
			point = index
			k++
		}
	}
	for i := 0; i < goroutines; i++ {
		// ВАЖНО: просто с partSize не прокатит, вам нужно как-то заранее разделить слайс по месяцам
		// ВАЖНО: этот код - лишь шаблон, который показывает вам как запустить горутину
		var part []TransactionInfo
		switch i {
		case 0:
			part = lengthOfFirstMonth
		case 1:
			part = lengthOfSecondMonth
		case 2:
			part = lengthOfThirdMonth
		case 3:
			part = lengthOfThourthMonth
		case 4:
			part = lengthOfThifthMonth
		default:
			part = transactions[: len(transactions) / goroutines]
		}

		go func() {
			sum := sum(part)
			total = append(total, sum)
			fmt.Println(sum)
			wg.Done()
		}()
	}

	wg.Wait()
	return total
}

func sum(transactions []TransactionInfo) int64  {
	result := int64(0)
	for _, transaction := range transactions {
		result +=transaction.SumOfTransaction
	}
	return result
}

