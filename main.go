package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	for true {
		var domain string
		fmt.Print("Enter domain to check: ")
		_, err := fmt.Scan(&domain)
		if err != nil {
			log.Fatalf("Error: could not read from input: %v\n", err)
		}
		//scanner := bufio.NewScanner(os.Stdin)

		//for scanner.Scan() {
		checkDomain(domain)
		//}

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
	fmt.Println()
}
