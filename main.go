package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	handleDeposits()
	handleWithdrawals()
}

func handleWithdrawals() {
	files, err := ioutil.ReadDir("withdrawals")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("Starting WITHDRAWAL:%s\n", file.Name())
		//Take it takes a while.-..
		time.Sleep(20 * time.Second)
		fmt.Printf("Done WITHDRAWAL:%s\n", file.Name())

		now := time.Now()
		fileName := fmt.Sprintf("%d-%02d-%02d-%02d-%02d-%02d.txt", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

		os.Rename("withdrawals/"+file.Name(),
			"processed/Withdrawal-"+fileName)
	}

}

func handleDeposits() {
	files, err := ioutil.ReadDir("deposits")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("Starting DEPOSIT:%s\n", file.Name())
		//Take it takes a while.-..
		time.Sleep(20 * time.Second)
		fmt.Printf("Done DEPOSIT:%s\n", file.Name())

		now := time.Now()
		fileName := fmt.Sprintf("%d-%02d-%02d-%02d-%02d-%02d.txt", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

		os.Rename("deposits/"+file.Name(),
			"processed/DEPOSIT-"+fileName)
	}
}
