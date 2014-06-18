package icws_golib

import (
	"log"
	"testing"
)

var testserver = "morbo.dev2000.com"
var testuser = "kevin.glinski"
var testpassword = "1234"

func TestLogin(t *testing.T) {
	log.Printf("LoginTest")

	icws := NewIcws()
	err := icws.Login("unitTest", testserver, testuser, testpassword)

	if err != nil {
		t.Error(err)
	}

}


func TestVersion(t *testing.T) {
    icws := NewIcws();

    err := icws.Login("unitTest", testserver, testuser, testpassword)

    if err != nil {
        t.Error(err)
    }

    version, err := icws.GetVersion();

    if err != nil {
        t.Error(err)
    }

    if(version.MajorVersion == ""){
        t.Error("Error getting version")
    }

    //log.Printf("%+v", version);
}


func TestFeatures(t *testing.T) {
    icws := NewIcws();

    err := icws.Login("unitTest", testserver, testuser, testpassword)

    if err != nil {
        t.Error(err)
    }

    features, err := icws.GetFeatures();

    if err != nil {
        t.Error(err)
    }

    if(features == nil || len(features) == 0){
        t.Error("Error getting features")
    }

    //log.Printf("%+v", features);
}
