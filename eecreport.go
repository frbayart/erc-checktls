// eecreport.go

/*
This file contains our EEC-specific data/func
*/
package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"strings"

	"erc-checktls/imirhil"
	"erc-checktls/ssllabs"
)

// EECReport is the data we want to extract
// Private functions

// getResults read the JSON array generated and gone through jq
func getResults(file string) (res []byte, err error) {
	fh, err := os.Open(file)
	if err != nil {
		return
	}
	defer fh.Close()

	res, err = ioutil.ReadAll(fh)
	return
}

// Public functions

// NewTLSReport generates everything we need for display/export
func NewTLSReport(reports *ssllabs.LabsReports) (e *TLSReport, err error) {
	e = TLSReport{Date:time.Now()}
	e.Sites = make(map[string]TLSSite, len(*reports))

	for _, site := range *reports {
		endp := site.Endpoints[0]
		det := endp.Details
		cert := endp.Details.Cert

		// make space
		siteData := make(EECLine, 17)

		// [0] = site
		siteData = append(siteData, site)

		// [1] = contract
		siteData = append(siteData, contracts[site])

		// [2] = grade
		siteData = append(siteData, fmt.Sprintf("%s %d bits",
			det.Key.Alg,
			det.Key.Size))

		// [3] = signature
		siteData = append(siteData, det.Cert.SigAlg)

		// [4] = issuer
		siteData = append(siteData, det.Cert.IssuerLabel)

		// [5] = validity
		siteData = append(siteData, time.Unix(cert.NotAfter, 0).String())

		// [6] = path
		siteData = append(siteData, fmt.Sprintf("%d", len(det.Chain.Certs)))

		// [7] = issues
		siteData = append(siteData, fmt.Sprintf("%d", det.Chain.Issues))

		// [8] = protocols
		protos := make([]string, 10)
		for _, p := range det.Protocols {
			protos = append(protos, fmt.Sprintf("%sv%s", p.Name, p.Version))
		}
		siteData = append(siteData, strings.Join(protos, ","))

		// [9] = RC4
		if det.SupportsRC4 {
			siteData = append(siteData, "YES")
		} else {
			siteData = append(siteData, "NO")
		}

		// [10] = PFS
		if det.SupportsRC4 {
			siteData = append(siteData, "YES")
		} else {
			siteData = append(siteData, "NO")
		}

		// [11] = OCSP Stapling
		if det.OcspStapling {
			siteData = append(siteData, "YES")
		} else {
			siteData = append(siteData, "NO")
		}

		// [12] = HSTS
		if det.OcspStapling {
			siteData = append(siteData, "YES")
		} else {
			siteData = append(siteData, "NO")
		}

		// [13] = HSTS
		if det.HstsPolicy.Status == "present" {
			siteData = append(siteData, "YES")
		} else {
			siteData = append(siteData, "NO")
		}

		// [14] = ALPN
		if det.SupportsAlpn {
			siteData = append(siteData, "YES")
		} else {
			siteData = append(siteData, "NO")
		}

		// [15] = Drown vuln
		if det.DrownVulnerable {
			siteData = append(siteData, "YES")
		} else {
			siteData = append(siteData, "NO")
		}

		// [16] = imirhil score unless ignored
		if !fIgnoreImirhil {
			siteData = append(siteData, imirhil.GetScore(site))
		} else {
			siteData = append(siteData, "")
		}
	}
	return
}

/* Display for one report
func (rep *ssllabs.LabsReport) String() {
	host := rep.Host
	grade := rep.Endpoints[0].Grade
	details := rep.Endpoints[0].Details
	log.Printf("Looking at %s/%s — grade %s", host, contracts[host], grade)
	if fVerbose {
		log.Printf("  Ciphers: %d", details.Suites.len())
	} else if fReallyVerbose {
		for _, cipher := range details.Suites.List {
			log.Printf("  %s: %d bits", cipher.Name, cipher.CipherStrength)
		}
	}
}*/

// toLine groups part of the data into a single array
func (r *TLSSite) toLine() {

}

// ToCSV generate a CSV file from a given report
func (r *TLSSite) ToCSV() {

}

// getContract retrieve the site's contract from the DB
func readContractFile(file string) (contracts map[string]string, err error) {
	var (
		fh *os.File
	)

	_, err = os.Stat(file)
	if err != nil {
		return
	}

	if fh, err = os.Open(file); err != nil {
		return
	}
	defer fh.Close()

	all := csv.NewReader(fh)
	allSites, err := all.ReadAll()

	contracts = make(map[string]string)
	for _, site := range allSites {
		contracts[site[0]] = site[1]
	}
	return
}
