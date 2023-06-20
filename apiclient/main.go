package main

import (
	"client"
	"fmt"
)

func main() {
	apiClient := client.APIClient{
		BaseURL:       "https://cpcli.cf.eu10.hana.ondemand.com", // Replace with your API base URL
		Username:      "P2006255217",
		Password:      "Carlos123*",
		GlobalAccount: "27eadf16trial-ga",
	}
	subaccountInfo := client.SubaccountInfo{
		SubacccountID: "75407dce-c7e7-4d84-a0ef-b012865dd4b0",
		Subdomain:     "27eadf16trial6",
		DisplayName:   "trial6",
		Region:        "us10",
	}
	err := apiClient.Login(&subaccountInfo)
	if err != nil {
		panic(err)
	}
	respList1, err := apiClient.CreateSubaccountCommand(&subaccountInfo)
	if err != nil {
		fmt.Println(string(err.Error()))
		panic(err)
	}
	if respList1 != nil {
		subaccountInfo = *respList1
	}
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
