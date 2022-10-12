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
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter domain to check: ")

		for scanner.Scan() {
			checkDomain(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("Error: could not read from input: %v\n", err)
		}
	}

}

func checkDomain(domain string) {
	var (
		hasMx, hasSPF, hasDMARC bool
		spfRecord, dmarcRecord  string
	)

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Println("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord")
	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
