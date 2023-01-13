package main

import (
	"int_test/hash"
	randomStr "int_test/random"
	"int_test/store"
	"int_test/sum"
)

func main() {
	hashStore := store.NewHashStore()
	hashChannel := make(chan store.HashData, 10)
	go hashStore.ReceiveUpdate(hashChannel)
	for {
		newRandStr := randomStr.NewRandomString()
		hashedStr := hash.HashString(newRandStr)
		if hashedStr[len(hashedStr)-3:] == "000" {
			sumOfHashed := sum.SumIntegers(hashedStr)
			hashData := store.HashData{
				Hash: hashedStr,
				Sum:  sumOfHashed,
			}

			hashChannel <- hashData
		}
	}
}
