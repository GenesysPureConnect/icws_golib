PACKAGE DOCUMENTATION

package icws_golib
    import "github.com/interactiveintelligence/icws_golib"

    the icws_golib package wraps CIC ICWS fuctionality in a GO library. To
    get started, call icws_golib.NewIcws() to instansiate a new Icws struct.
    Then call icws.Login(...) and continue on from there.


TYPES

type ConfigRecord map[string]interface{}



type Icws struct {
    CurrentToken, CurrentCookie, CurrentSession, CurrentServer string
}


func NewIcws() (icws *Icws)
    Creates a new ICWS struct


func (i *Icws) GetConfigurationRecord(configurationType, id, properties string) (record ConfigRecord, err error)
    Gets a record for the ID of a specific configuration type.

func (i *Icws) InteractionAction(action, interactionId, attribute string) (result ConfigRecord, err error)
    Performs an action on a given interaction. e.g. Pickup, Hold,
    Disconnect, etc

func (i *Icws) Login(applicationName, server, username, password string) (err error)
    Logs into a CIC server. Server should be a url e.g.
    https://MyServer:8019

func (i *Icws) LoginMarketPlaceApp(applicationName, server, username, password, marketplaceLicense, markeplaceAppKey string) (err error)
    Logs into a CIC server for a MarketPlace application using the app's
    custom license. Server should be a url e.g. https://MyServer:8019

func (i *Icws) MakeCall(target string) (result ConfigRecord, err error)
    Creates a new call

func (i *Icws) SelectConfigurationRecords(objectType, selectFields, where string) (records []ConfigRecord, err error)
    Returns a list of matching records for the given object type.

func (i *Icws) SetPassword(userId, password string, force bool) error
    changes the password for the given user



