package icws_golib

import (
    "encoding/json"
    "strings"


)

type ConfigRecord map[string]interface{}


//Gets a record for the ID of a specific configuration type.
func (i *Icws) GetConfigurationRecord(configurationType, id, properties string) (record ConfigRecord, err error) {

	if !strings.HasSuffix(configurationType, "s") {
		configurationType += "s"
	}

	body, err := i.httpGet("/configuration/" + configurationType + "/" + id + "?select=" + properties)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &record)
	return
}

//Returns a list of matching records for the given object type.
func (i *Icws) SelectConfigurationRecords(objectType, selectFields, where string) (records []ConfigRecord, err error) {

	if !strings.HasSuffix(objectType, "s") {
		objectType += "s"
	}

	var selectString string
	if selectFields == "*" {
		selectString = ""
	} else {
		selectString = "select=" + selectFields
	}

	var whereString string
	if len(where) == 0 {
		whereString = ""
	} else {
		whereString = "&where=" + where
	}

	body, err := i.httpGet("/configuration/" + objectType + "?" + selectString + whereString)
	if err != nil {
		return
	}

	var result map[string][]ConfigRecord
	err = json.Unmarshal(body, &result)

	records = result["items"]
	return

}
