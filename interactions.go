package icws_golib

import (
	"encoding/json"
)

//Creates a new call
func (i *Icws) MakeCall(target string) (result ConfigRecord, err error) {

	var callData = map[string]string{
		"__type": "urn:inin.com:interactions:createCallParameters",
		"target": target,
	}
	body, err := i.HttpPost("/interactions", callData)

	if err != nil {
		return
	}

	err = json.Unmarshal(body, &result)
	return

}

//Performs an action on a given interaction.  e.g. Pickup, Hold, Disconnect, etc
func (i *Icws) InteractionAction(action, interactionId, attribute string) (result ConfigRecord, err error) {

	var body []byte
	if action == "get" {
		body, err = i.httpGet("/interactions/" + interactionId + "?select=" + attribute)

		if err != nil {
			return
		}

		err = json.Unmarshal(body, &result)

	} else if action == "set" {
		/*body, err, _ := httpPost(server + "/icws/" + session + "/interactions/" + interactionId " , callData)

		  err = json.Unmarshal(body, &result)
		*/
	} else {
		var isOn = "false"
		if attribute == "on" || attribute == "yes" || attribute == "1" {
			isOn = "true"
		}
		var callData = map[string]string{
			"on": isOn,
		}

		if len(attribute) == 0 {
			callData = nil
		}

		_, err = i.HttpPost("/interactions/"+interactionId+"/"+action, callData)

	}

	return

}
