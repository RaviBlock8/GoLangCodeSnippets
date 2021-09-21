package main

type bank struct{
	accountList map[string]account
	totalNoOfAccounts int
	bankName string
}

func createNewBank(bankName string) bank{
	newBank:=bank{bankName: "HDFC",totalNoOfAccounts: 0}
	return newBank
}

func (currentBank *bank)openNewAccountWithBank(holder string,address string,uniqueId int,passwords int) account{
	newAccount:=createAccount(holder,address,uniqueId,passwords)
	currentBank.accountList[newAccount.accountNumber]=newAccount
	currentBank.totalNoOfAccounts=len(currentBank.accountList)
	return newAccount
}

func (currentBank *bank) transferMoney(senderNumber string,receiverNumber string,password int,amount int) bool{
	senderAccount:=currentBank.accountList[senderNumber]
	receiverAccount:=currentBank.accountList[receiverNumber]
	if senderAccount.matchPassword(password)!=true{
		return false
	}
	senderAccount.deductBalance(amount);
	receiverAccount.addBalance(amount);
	return true;
}

func (currentBank *bank) addBeneficiariesToAccount(senderNumber string,password int,name string,relationWithHolder string,dob string) bool {
	senderAccount:=currentBank.accountList[senderNumber]
	if senderAccount.matchPassword(password)!=true{
		return false
	}
	senderAccount.addBeneficiaries(name,relationWithHolder,dob)
	return true
}