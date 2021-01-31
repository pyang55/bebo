package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/deet/govenmo"
)

var account = govenmo.Account{
	AccessToken: os.Getenv("VENMO_TOKEN"),
}
var amt float64 = 0.0

// gets user venmo information
func accountRefresh() string {
	err := account.Refresh()
	if err != nil {
		fmt.Printf("help")
	}
	info := fmt.Sprintf("Username:%s\nFriends Count:%d\nBalance:$%d\n", account.Username, uint64(account.FriendsCount), uint64(account.Balance))
	return info
}

func venmoAction(email string, description string, action string, amount string) error {
	target := govenmo.Target{}
	target.Email = email
	amt, _ = strconv.ParseFloat(amount, 64)
	if action == "charge" {
		_, err := account.PayOrCharge(target, -amt, description, "private")
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else {
		_, err := account.PayOrCharge(target, amt, description, "private")
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
