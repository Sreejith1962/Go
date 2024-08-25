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
	fmt.Println("domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error : Could not read from input ", err)
	}

	/*	var str string
		for {
			fmt.Scanln(&str)
			checkDomain(str)
		}
	*/
}

func checkDomain(domain string) {
	var hasMX, hasDMARC, hasSPF bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Println(err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}
	textRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Println(err)
	}

	for _, record := range textRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Println(err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}

	}

	fmt.Println(domain, ",", hasMX, ",", hasSPF, ",", spfRecord, ",", hasDMARC, ",", dmarcRecord)
}
