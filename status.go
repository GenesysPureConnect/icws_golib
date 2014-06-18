package icws_golib

import (
	"encoding/json"
)

//gets the status information for a user
func (i *Icws) GetStatus(userId string) (defaults ConfigRecord, err error) {

	body, err := i.httpGet("/status/user-statuses/" + userId)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &defaults)
	return

}

//Sets a status for a given user id
func (i *Icws) SetStatus(userId, statusKey string) (err error) {

	var statusData = map[string]string{
		"statusId": statusKey,
	}
	_, err = i.httpPut("/status/user-statuses/"+userId, statusData)

	return

}
