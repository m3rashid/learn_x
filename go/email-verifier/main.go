package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")
	fmt.Printf("Enter domain: ")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("ERROR: Could nor read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDmarc bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc" + domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDmarc = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("domain: %v\nhasMX: %v\nhasSPF: %v\nspfRecord: %v\nhasDMARC: %v\ndmarcRecord: %v\n", domain, hasMX, hasSPF, spfRecord, hasDmarc, dmarcRecord)
}
