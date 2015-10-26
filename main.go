package main

import(
	//comment out if not used
    //"./lib/account"
	//"./lib/atm"
    //"./lib/branch"
    //"./lib/merchant"
)

func main() {

    //Demo Code for Requests (uncomment to run)

    //======================Account======================= (DONE)
    //account.GetAllAccounts()
    //account.GetAccountWithId("56241a13de4bf40b1711287b")
	//account.GetAccountsOfCustomer("56241a12de4bf40b17111f9d")
    //account.CreateAccount("56241a12de4bf40b17111f9d", "Checking", "Account to Delete", 44444, 44444, "")
    //account.CreateAccount("56241a12de4bf40b17111f9d", "Checking", "Account to Delete", 7777, 77777, "8888444455551111")
    //account.UpdateAccount("562cf1190afebb140066cd81", "Iron Man's Account", "")
    //account.UpdateAccount("562cf1190afebb140066cd81", "Mario's Account", "1234567812345678")
    //account.DeleteAccount("562b073f0afebb140066cd58")

    //=======================ATM=========================
    //atm.GetAllBranches(38.9283, -77.1753, 1)
    //atm.GetATMInfo("56241a12de4bf40b17111c65")

    //======================Branch======================= (DONE)
    //branch.GetAllBranches()
    //branch.GetBranchWithId("56241a12de4bf40b17111eb2")

    //=====================Merchant====================== (DONE)
    //merchant.GetAllMerchants(38.9283, -77.1753, 1)
    //merchant.CreateMerchant("Dunkin Donuts", "Food", "11006", "Capital One Dr.", "McLean", "VA", "20931", 38, -77)
    //merchant.CreateMerchant("Dunkin Donuts3", "", "11006", "Capital One Dr.", "McLean", "VA", "20931", 38, -77)
    //merchant.CreateMerchant("Dunkin Donuts", "", "", "", "", "", "", 38.223, -77.111)
    //merchant.CreateMerchant("Dunkin Donuts5", "Food", "", "", "", "", "", -999, -999)
    //merchant.GetMerchantInfo("562d1cf40afebb140066cd8a")
    //merchant.UpdateMerchant("562db1380afebb140066cda1", "Starbucks","","","","","","", 67.333,-1.43)
    //merchant.UpdateMerchant("562db1380afebb140066cda1", "Starbucks","","","","","","", -999,-999)
    //merchant.UpdateMerchant("562db1380afebb140066cda1", "Plan 9 Records", "", "11006", "Capital One Dr.", "McLean", "VA", "20931", 38, -77)
    //merchant.UpdateMerchant("562db1380afebb140066cda1", "Wootton High School", "Education", "", "", "", "", "", -999, -999)
}
