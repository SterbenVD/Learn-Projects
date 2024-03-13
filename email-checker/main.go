package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("domain ...")

	for sc.Scan() {
		checkDomain(sc.Text())
		err := sc.Err()
		errchecker(err, "taking input")
	}

}

func errchecker(err error, where string) {
	if err != nil {
		fmt.Println("Error in %s: %v\n", where, err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMRC bool
	var SPF, DMRC string

	mxRecs, err := net.LookupMX(domain)

	errchecker(err, "looking up MX records")

	if len(mxRecs) > 0 {
		hasMX = true
	}

	txtRecs, err := net.LookupTXT(domain)

	errchecker(err, "looking up SPF in TXT records")

	for _, record := range txtRecs {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			SPF = record
			break
		}
	}

	dmarcRecs, err := net.LookupTXT("_dmarc." + domain)

	errchecker(err, "looking up DMRAC in TXT records")

	for _, record := range dmarcRecs {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMRC = true
			DMRC = record
			break
		}
	}

	fmt.Println("Domain: ", domain)
	fmt.Println("hasMX: ", hasMX)
	fmt.Println("hasSPF: ", hasSPF)
	if hasSPF == true {
		fmt.Println("SPF: ", SPF)
	}
	fmt.Println("hasDMRC: ", hasDMRC)
	if hasSPF == true {
		fmt.Println("DMRC: ", DMRC)
	}

}
