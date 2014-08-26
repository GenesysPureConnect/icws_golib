package icws_golib

import (
	"fmt"
	"testing"
	"os"
)

func TestGetConfigurationRecord(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("skipping test in short mode.")
	}
	var icws = NewIcws()
	icws.Login("Test", testserver, testuser, testpassword)
	record, err := icws.GetConfigurationRecord("user", testuser, "extension")

	if err != nil {
		t.Error(fmt.Sprintf("%s", err))
	}

	if record == nil {
		t.Error("Record is NIL")
	}

}

func TestGetConfigurationRecord_WithInvalidId(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("skipping test in short mode.")
	}

	var icws = NewIcws()
	icws.Login("Test", testserver, testuser, testpassword)
	record, err := icws.GetConfigurationRecord("user", "sdfkasjdfkalsjdhfaskfj", "extension")

	if err == nil {
		t.Error("should have returned an error")
	}

	if record != nil {
		t.Error("Record should have been NIL")
	}

}

func TestSelectConfigurationRecords(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("skipping test in short mode.")
	}
	var icws = NewIcws()
	icws.Login("Test", testserver, testuser, testpassword)
	records, err := icws.SelectConfigurationRecords("user", "extension", "")

	if err != nil {
		t.Error(fmt.Sprintf("%s", err))
	}

	if records == nil {
		t.Error("Record is NIL")
	}

	if len(records) == 0 {
		t.Error("no results returned")
	}

}
