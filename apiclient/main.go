package main

import "fmt"

func main() {
	apiClient := APIClient{
		BaseURL:       "https://cpcli.cf.eu10.hana.ondemand.com", // Replace with your API base URL
		Username:      "P2006255217",
		Password:      "--------",
		GlobalAccount: "27eadf16trial-ga",
	}
	subaccountInfo := SubaccountInfo{
		SubacccountID: "cb96c8bf-a5aa-46c4-9e7d-c488edaf7b26",
		Subdomain:     "27eadf16trial7",
		DisplayName:   "trial7",
		Region:        "us10",
	}
	err := apiClient.Login(&subaccountInfo)
	if err != nil {
		panic(err)
	}
	// respList1, err := apiClient.CreateSubaccountCommand(&subaccountInfo)
	// if err != nil {
	// 	fmt.Println(string(err.Error()))
	// 	panic(err)
	// }
	// if respList1 != nil {
	// 	subaccountInfo = *respList1
	// }
	respList, err := apiClient.GetSubaccountCommand(&subaccountInfo)
	if err != nil {
		fmt.Println(string(err.Error()))
		panic(err)
	}
	if respList != nil {
		subaccountInfo = *respList
	}
	subaccountInfo.DisplayName = "trail15"
	respList2, err := apiClient.UpdateSubaccountCommand(&subaccountInfo)
	if err != nil {
		fmt.Println(string(err.Error()))
		panic(err)
	}

	if respList2 != nil {
		subaccountInfo = *respList2
	}

	respList3, err := apiClient.DeleteSubaccountCommand(&subaccountInfo)
	if err != nil {
		fmt.Println(string(err.Error()))
		panic(err)
	}

	if respList3 != nil {
		subaccountInfo = *respList3
	}
}
