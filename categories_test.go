package main

import (
	"io/ioutil"
	"testing"

	"github.com/keltia/ssllabs"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestDisplayCategories(t *testing.T) {
	cntrs := map[string]int{
		"A": 666,
		"B": 0,
		"G": 1,
	}
	str := displayCategories(cntrs)
	assert.NotEmpty(t, str)
}

func TestHTTPCountsNil(t *testing.T) {
	cntrs := httpCounts(nil)
	assert.Empty(t, cntrs)
}

func TestHTTPCountsEmpty(t *testing.T) {
	cntrs := httpCounts(&TLSReport{})
	assert.NotEmpty(t, cntrs)
	assert.EqualValues(t, map[string]int{"Total": 0, "Broken": 0}, cntrs)
}

func TestHTTPCountsReport(t *testing.T) {
	ji, err := ioutil.ReadFile("testdata/site.json")
	require.NoError(t, err)

	// Simulate
	fIgnoreMozilla = true
	fIgnoreImirhil = true

	all, err := ssllabs.ParseResults(ji)
	require.NoError(t, err)

	sites, err := NewTLSReport(all)
	require.NoError(t, err)
	require.NotEmpty(t, sites)

	// Fake it
	sites.Sites[0].Mozilla = "A+"

	cntrs := httpCounts(sites)
	assert.NotEmpty(t, cntrs)
	assert.EqualValues(t, map[string]int{"A+": 1, "Total": 1, "Broken": 1}, cntrs)
}

func TestHTTPCountsReport_1(t *testing.T) {
	ji, err := ioutil.ReadFile("testdata/site.json")
	require.NoError(t, err)

	// Simulate
	fIgnoreMozilla = true
	fIgnoreImirhil = true

	all, err := ssllabs.ParseResults(ji)
	require.NoError(t, err)

	sites, err := NewTLSReport(all)
	require.NoError(t, err)
	require.NotEmpty(t, sites)

	// Fake it
	sites.Sites[0].Mozilla = "H"

	cntrs := httpCounts(sites)
	assert.NotEmpty(t, cntrs)
	assert.EqualValues(t, map[string]int{"H": 1, "Total": 0, "Broken": 1}, cntrs)
}

func TestCategoryCountsNil(t *testing.T) {
	cntrs := categoryCounts(nil)
	assert.Empty(t, cntrs)
}

func TestCategoryCountsEmpty(t *testing.T) {
	cntrs := categoryCounts([]ssllabs.Host{})
	assert.NotEmpty(t, cntrs)
	assert.EqualValues(t, map[string]int{"Total": 0, "X": 0, "Z": 0}, cntrs)
}

func TestCategoryCountsReport(t *testing.T) {
	ji, err := ioutil.ReadFile("testdata/site.json")
	require.NoError(t, err)

	// Simulate
	fIgnoreMozilla = true
	fIgnoreImirhil = true

	all, err := ssllabs.ParseResults(ji)
	require.NoError(t, err)

	good := map[string]int{
		"OCSP": 1, "Total": 1, "X": 0, "": 1, "Issues": 1,
		"HSTS": 1, "Z": 1, "A+": 1, "PFS": 1,
	}

	cntrs := categoryCounts(all)
	assert.NotEmpty(t, cntrs)
	assert.EqualValues(t, good, cntrs)
}

func TestCategoryCountsReportDES(t *testing.T) {
	ji, err := ioutil.ReadFile("testdata/reallybad.json")
	require.NoError(t, err)

	// Simulate
	fIgnoreMozilla = true
	fIgnoreImirhil = true

	all, err := ssllabs.ParseResults(ji)
	require.NoError(t, err)

	good := map[string]int{
		"OCSP": 1, "Total": 1, "X": 0, "Issues": 1,
		"HSTS": 1, "Z": 0, "A+": 1, "PFS": 1, "Sweet32": 1,
	}

	cntrs := categoryCounts(all)
	assert.NotEmpty(t, cntrs)
	assert.EqualValues(t, good, cntrs)
}

func TestCategoryCountsReportNull(t *testing.T) {
	ji, err := ioutil.ReadFile("testdata/null.json")
	require.NoError(t, err)

	// Simulate
	fIgnoreMozilla = true
	fIgnoreImirhil = true

	all, err := ssllabs.ParseResults(ji)
	require.NoError(t, err)

	good := map[string]int{
		"X": 1, "Z": 0, "Total": 0,
	}

	cntrs := categoryCounts(all)
	assert.NotEmpty(t, cntrs)
	assert.EqualValues(t, good, cntrs)
}
