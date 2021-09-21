package main

import "crypto/sha1"

type person struct{
	name string
	relationWithHolder string
	dob string
}

type account struct{
	accountNumber string
	balance int
	holder string
	address string
	beneficiaries map[int]person
	password string
}

func createAccount(holder string,address string,uniqueId int,passwords int) *account{
	var newAccount account
	newAccount.holder=holder
	newAccount.address=address
	h:=sha1.New()
	h.Write([]byte{byte(uniqueId)})
	newAccount.accountNumber=string(h.Sum(nil))
	passwordHash:=sha1.New()
	// So here basically I am changing the internal state , depending on which we will get hash.
	passwordHash.Write([]byte{byte(passwords)})
	newAccount.password=string(passwordHash.Sum(nil))
	return &newAccount
}

func (userAccount *account) addBalance(value uint){
	userAccount.balance+=int(value)
}

func (userAccount *account) deductBalance(value uint){
	userAccount.balance-=int(value)
}

func (userAccount *account) matchPassword(password int) bool{
	passwordHash:=sha1.New()
	passwordHash.Write([]byte{byte(password)})
	if string(passwordHash.Sum(nil))!=userAccount.password{
		return false
	}
	return true
}

func (userAccount *account) addBeneficiaries(name string,relationWithHolder string,dob string){
	userAccount.beneficiaries[len(userAccount.beneficiaries)]=person{name,relationWithHolder,dob}
}

