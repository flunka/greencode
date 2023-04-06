package endpoints

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

const TransationsEndpoint = "/transactions/report"

type Transaction struct {
	DebitAccount  string  `json:"debitAccount"  binding:"required,len=26"`
	CreditAccount string  `json:"creditAccount"  binding:"required,len=26"`
	Amount        float64 `json:"amount"  binding:"required"`
}

type Account struct {
	Account     string  `json:"account" binding:"len=26"`
	DebitCount  int32   `json:"debitCount"`
	CreditCount int32   `json:"creditCount"`
	Balance     float64 `json:"balance"`
}

type ByAccount []Account

func (a ByAccount) Len() int           { return len(a) }
func (a ByAccount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAccount) Less(i, j int) bool { return a[i].Account < a[j].Account }

func (a *Account) withdraw(amount float64) {
	a.Balance -= amount
	a.DebitCount += 1
}

func (a *Account) deposit(amount float64) {
	a.Balance += amount
	a.CreditCount += 1
}

func do_transaction(debitAccount, creditAccout *Account, amount float64) {
	debitAccount.withdraw(amount)
	creditAccout.deposit(amount)
}

func Report(c *gin.Context) {
	var transactions []Transaction
	var accounts = make(map[string]*Account)
	if err := c.ShouldBindJSON(&transactions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, transaction := range transactions {
		fmt.Printf("trans %+v\n", transaction)
		debitAccount, isPresent := accounts[transaction.DebitAccount]
		if !isPresent {
			debitAccount = &Account{
				Account:     transaction.DebitAccount,
				DebitCount:  0,
				CreditCount: 0,
				Balance:     0.0,
			}
			fmt.Printf("new debit account %s\n", transaction.DebitAccount)
			accounts[transaction.DebitAccount] = debitAccount
		}
		creditAccount, isPresent := accounts[transaction.CreditAccount]
		if !isPresent {
			creditAccount = &Account{
				Account:     transaction.CreditAccount,
				DebitCount:  0,
				CreditCount: 0,
				Balance:     0.0,
			}
			fmt.Printf("new credit Account %s\n", transaction.CreditAccount)
			accounts[transaction.CreditAccount] = creditAccount
		}
		do_transaction(debitAccount, creditAccount, transaction.Amount)

	}
	accounts_array := make([]Account, 0, len(accounts))
	for _, value := range accounts {
		fmt.Printf("%+v\n", value)
		accounts_array = append(accounts_array, *value)
	}
	sort.Sort(ByAccount(accounts_array))
	c.JSON(http.StatusOK, accounts_array)
}
