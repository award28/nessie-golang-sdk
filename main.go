package main

import(
	//comment out if not used
    //"./lib/account"
	//"./lib/atm"
    //"./lib/bill"
    //"./lib/branch"
    //"./lib/customer"
    "./lib/deposit"
    //"./lib/merchant"
    //"./lib/purchase"
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

    //=======================BILL======================== (DONE)
    //bill.GetBillsOfAccount("56241a13de4bf40b1711287b")
    //bill.GetBillWithId("56241a14de4bf40b1711318c")
    //bill.GetBillsOfCustomer("56241a12de4bf40b17111f9c")
    //bill.CreateBill("56241a13de4bf40b1711287a", "completed", "Comcast", "Cable", "2001-05-22", 22, 6.55)
    //bill.CreateBill("56241a13de4bf40b1711287a", "completed", "Nothing", "", "2001-05-22", -999, 116.55)
    //bill.CreateBill("56241a13de4bf40b1711287a", "completed", "Nothing", "Something", "", -999, 116.55)
    //bill.UpdateBill("562ef8440afebb140066cda9", "completed", "Verizon", "Cable", "2015-10-26", 1, 222.55)
    //bill.DeleteBill("562ef9db0afebb140066cdab")

    //======================Branch======================= (DONE)
    //branch.GetAllBranches()
    //branch.GetBranchWithId("56241a12de4bf40b17111eb2")

    //=====================Customer====================== (DONE)
    //customer.GetCustomerOfAccount("56241a13de4bf40b1711287a")
    //customer.GetAllCustomers()
    //customer.GetCustomerWithId("56241a12de4bf40b17111f9d")
    //customer.CreateCustomer("Robert", "Frost", "1111", "Infinity Loop", "Richmond", "VA", "22211")
    //customer.UpdateCustomer("56241a12de4bf40b17111f9d", "1112", "Infinity Loop", "Richmond", "VA", "22211")

    //=====================Deposit======================= 
    //deposit.GetDepositOfAccount("56241a13de4bf40b1711287c")
    //deposit.GetDepositById("562ff3980afebb140066cdae")
    //deposit.CreateDeposit("56241a13de4bf40b1711287b", "balance", "10/26/2015", "completed", 65.23, "paycheck")
    //deposit.CreateDeposit("56241a13de4bf40b1711287b", "balance", "10/26/2015", "completed", 65.24, "paycheck")
    //deposit.CreateDeposit("56241a13de4bf40b1711287b", "balance", "", "completed", 65.25, "paycheck")
    //deposit.CreateDeposit("56241a13de4bf40b1711287b", "balance", "10/26/2015", "", 65.26, "paycheck")
    //deposit.CreateDeposit("56241a13de4bf40b1711287b", "balance", "", "", 65.27, "")
    //deposit.UpdateDeposit("5636ec210afebb140066ce52", "rewards", 10000, "deposit doh")
    //deposit.UpdateDeposit("5636ec210afebb140066ce52", "rewards", -999, "deposit doh")
    //deposit.UpdateDeposit("5636ec210afebb140066ce52", "", 10000, "deposit doh")
    //deposit.UpdateDeposit("5636ec210afebb140066ce52", "rewards", 10000, "")
    //deposit.DeleteDeposit("5636ecd10afebb140066ce53")

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

    //=====================Purchase======================= 
    //purchase.GetPurchasesByAccount("56241a13de4bf40b1711287b")
    //purchase.GetPurchaseById("56304a7b0afebb140066cdb9")
    //purchase.CreatePurchase("56241a13de4bf40b1711287b", "562db1380afebb140066cda1", "balance", "10/16/2015", 112.31, "pending", "Lot of Coffee")
    //purchase.CreatePurchase("56241a13de4bf40b1711287b", "562db1380afebb140066cda1", "balance", "", 112.31, "pending", "Lot of Coffee")
    //purchase.CreatePurchase("56241a13de4bf40b1711287b", "562db1380afebb140066cda1", "balance", "10/16/2015", 112.31, "", "Lot of Coffee")
    //purchase.CreatePurchase("56241a13de4bf40b1711287b", "562db1380afebb140066cda1", "balance", "10/16/2015", 112.31, "pending", "")
    //purchase.UpdatePurchase("5636e3b80afebb140066ce4b", "56241a13de4bf40b1711287b", "balance", 100, "ajskdfjk")
    //purchase.UpdatePurchase("5636e3b80afebb140066ce4b", "56241a13de4bf40b1711287b", "balance", 100, "")
    //purchase.UpdatePurchase("5636e3b80afebb140066ce4b", "56241a13de4bf40b1711287b", "", 100, "")
    //purchase.DeletePurchase("5636e49e0afebb140066ce4d")
}


