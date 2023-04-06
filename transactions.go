package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/gin-gonic/gin"
)


type Transaction struct {
	debitAccount string
	creditAccount string
	amount float64
}

type Account struct {
	account string
	debitCount int32
	creditCount int32
	balance float64
}

func (a *Account) withdraw(amount float64) {
	a.balance -= amount
	a.debitCount += 1
}

func (a *Account) deposit(amount float64) {
	a.balance += amount
	a.creditCount += 1
}

func do_transaction(debitAccount, creditAccout *Account, amount float64) {
	debitAccount.withdraw(amount)
	creditAccout.deposit(amount)
}

func Report(c *gin.Context){
	var transations []Transaction
	var accounts = make(map[string]Account)
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
	    c.AbortWithError(http.StatusBadRequest, err)
	    return
	}

	err = json.Unmarshal(body, &transations)
	if err != nil {
	    c.AbortWithError(http.StatusBadRequest, err)
	    return
	}
	for _, transaction := range transations {
		fmt.Printf("trans %+v", transaction)
		debitAccount, isPresent := accounts[transaction.debitAccount]
		if !isPresent {
			debitAccount = Account{
				account: transaction.debitAccount,
				debitCount: 0,
				creditCount: 0,
				balance: 0.0,
			}
			fmt.Printf("new debit account %s", transaction.debitAccount)
			accounts[transaction.debitAccount] = debitAccount
		}
		creditAccount, isPresent := accounts[transaction.creditAccount]
		if !isPresent {
			creditAccount = Account{
			account: transaction.creditAccount,
			debitCount: 0,
			creditCount: 0,
			balance: 0.0,
			}
			fmt.Printf("new credit Account %s", transaction.creditAccount)
			accounts[transaction.creditAccount] = creditAccount
		}
		do_transaction(&debitAccount, &creditAccount, transaction.amount)
		
	}
	accounts_array := make([]Account, 0, len(accounts))
	for _, value := range accounts {
		fmt.Printf("%+v\n", value)
		accounts_array = append(accounts_array, value)
	}
	c.JSON(http.StatusOK, accounts_array)
}
