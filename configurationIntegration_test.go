package icws_golib

import ("testing"
"fmt"
)

var server = "http://morbo.dev2000.com:8018"
var user = "devlab_user"
var password = "1234"

func TestGetConfigurationRecord(t *testing.T){

    var icws = NewIcws();
    icws.Login("Test", server, user, password);
    record, err := icws.GetConfigurationRecord("user", user, "extension" )

    if err != nil{
        t.Error(fmt.Sprintf("%s",err));
    }

    if(record== nil){
        t.Error("Record is NIL");
    }

}

func TestGetConfigurationRecord_WithInvalidId(t *testing.T){

    var icws = NewIcws();
    icws.Login("Test", server, user, password);
    record, err := icws.GetConfigurationRecord("user", "sdfkasjdfkalsjdhfaskfj", "extension" )

    if err == nil{
        t.Error("should have returned an error");
    }

    if(record != nil){
        t.Error("Record should have been NIL");
    }

}

func TestSelectConfigurationRecords(t *testing.T){
    var icws = NewIcws();
    icws.Login("Test", server, user, password);
    records, err := icws.SelectConfigurationRecords("user", "extension", "" )

    if err != nil{
        t.Error(fmt.Sprintf("%s",err));
    }

    if(records== nil){
        t.Error("Record is NIL");
    }

    if(len(records) == 0){
        t.Error("no results returned");
    }

}
