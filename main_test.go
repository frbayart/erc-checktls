package main

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/gobuffalo/packr"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestGetResults(t *testing.T) {
	ji, err := ioutil.ReadFile("testdata/site.json")
	require.NoError(t, err)

	buf, err := getResults("testdata/site.json")
	require.NoError(t, err)

	assert.Equal(t, ji, buf)
}

func TestGetResultsNothing(t *testing.T) {
	buf, err := getResults("testdata/site.nowhere")
	require.Error(t, err)
	require.Empty(t, buf)
}

func TestReadContractFile(t *testing.T) {
	// We embed the file now
	box := packr.NewBox("./files")

	cntrs, err := readContractFile(box)
	assert.NoError(t, err)
	assert.NotEmpty(t, cntrs)
}

func TestLoadTemplates(t *testing.T) {
	// We embed the file now
	box := packr.NewBox("./files")

	str, err := loadTemplates(box)
	assert.NoError(t, err)
	assert.NotEmpty(t, str)
	assert.Equal(t, 2, len(str))
}

func TestCheckOutput(t *testing.T) {
	fh := checkOutput("")
	assert.NotEmpty(t, fh)
	assert.EqualValues(t, os.Stdout, fh)
}

func TestCheckOutput_1(t *testing.T) {
	temp, err := ioutil.TempDir("", "test")
	require.NoError(t, err)

	defer os.RemoveAll(temp)

	fn := path.Join(temp, "foo.out")
	fh := checkOutput(fn)
	assert.NotEmpty(t, fh)

	fi, err := os.Stat(fn)
	assert.NoError(t, err)
	assert.NotNil(t, fi)
	assert.NotEmpty(t, fi)
}

func TestWriteCSV(t *testing.T) {
	cntrs := map[string]int{
		"A": 666,
		"B": 42,
		"F": 1,
	}

	https := map[string]int{
		"A":  666,
		"B+": 37,
		"F":  42,
	}

	err := WriteCSV(os.Stderr, nil, cntrs, https)
	assert.Error(t, err)
}

func TestWriteCSV2(t *testing.T) {
	cntrs := map[string]int{
		"A": 666,
		"B": 42,
		"F": 1,
	}

	https := map[string]int{
		"A":  666,
		"B+": 37,
		"F":  42,
	}

	err := WriteCSV(os.Stderr, &TLSReport{}, cntrs, https)
	assert.Error(t, err)
}

func TestWriteHTML(t *testing.T) {
	cntrs := map[string]int{
		"A": 666,
		"B": 42,
		"F": 1,
	}

	https := map[string]int{
		"A":  666,
		"B+": 37,
		"F":  42,
	}

	err := WriteHTML(os.Stderr, nil, cntrs, https)
	assert.Error(t, err)
}

func TestWriteHTML2(t *testing.T) {
	cntrs := map[string]int{
		"A": 666,
		"B": 42,
		"F": 1,
	}

	https := map[string]int{
		"A":  666,
		"B+": 37,
		"F":  42,
	}

	err := WriteHTML(os.Stderr, &TLSReport{}, cntrs, https)
	assert.Error(t, err)
}
