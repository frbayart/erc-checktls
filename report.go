// report.go

/*
This file contains func for generating the report
*/
package main

import (
	"fmt"
	"io"
	"time"

	"github.com/atotto/encoding/csv"
	"github.com/keltia/ssllabs"
	"github.com/pkg/errors"
)

var (
	yesno = map[bool]string{
		true:  "YES",
		false: "NO",
	}
)

// Private functions

func getSSLablsVersion(site ssllabs.Host) string {
	debug("%#v", site)
	return fmt.Sprintf("%s/%s", site.EngineVersion, site.CriteriaVersion)
}

// NewTLSReport generates everything we need for display/export
func NewTLSReport(reports []ssllabs.Host) (e *TLSReport, err error) {
	if len(reports) == 0 {
		return nil, fmt.Errorf("empty list")
	}

	e = &TLSReport{
		Date:    time.Now(),
		SSLLabs: getSSLablsVersion(reports[0]),
	}

	verbose("%d sites found.\n", len(reports))

	// Now analyze each site
	for _, site := range reports {
		current := NewTLSSite(site)
		e.Sites = append(e.Sites, current)
	}
	return e, nil
}

// ToCSV output a CSV file from a report
func (r *TLSReport) ToCSV(w io.Writer) (err error) {
	wh := csv.NewWriter(w)
	debug("%v\n", r.Sites)
	if err = wh.WriteStructHeader(r.Sites[0]); err != nil {
		return errors.Wrap(err, "can not write csv header")
	}

	err = wh.WriteStructAll(r.Sites)
	return errors.Wrap(err, "can not write csv file")
}
