package main

func checkBalances() bool {
	// clientOptions := plaid.ClientOptions{
	// 	os.Getenv("PLAID_CLIENT_ID"),
	// 	os.Getenv("PLAID_SECRET_KEY"),
	// 	plaid.Development, // Available environments are Sandbox, Development, and Production
	// 	&http.Client{},    // This parameter is optional
	// }
	//
	// client, err := plaid.NewClient(clientOptions)
	// handleError(err)
	//
	// balanceResp, err := client.GetBalances(os.Getenv("ACCOUNT_TOKEN"))
	// handleError(err)
	//
	// if balanceResp.Accounts[2].Balances.Available-1350.00 >= 250.00 {
	// 	return true
	// }
	return false
}
