PACKAGE DOCUMENTATION

package icws_golib
    import "github.com/interactiveintelligence/icws_golib"

    the icws_golib package wraps CIC ICWS fuctionality in a GO library. To
    get started, call icws_golib.NewIcws() to instansiate a new Icws struct.
    Then call icws.Login(...) and continue on from there.


TYPES

type ConfigRecord map[string]interface{}



type Icws struct {
    CurrentToken, CurrentCookie, CurrentSession, CurrentServer, UserId string
}


func NewIcws() (icws *Icws)
    Creates a new ICWS struct


func (i *Icws) Defaults(configurationType string) (defaults ConfigRecord, err error)
    gets the default values for a configuration type

func (i *Icws) GetConfigurationRecord(configurationType, id, properties string) (record ConfigRecord, err error)
    Gets a record for the ID of a specific configuration type.

func (i *Icws) GetFeatures() (features []ServerFeature, err error)
    Gets a list of features and their version from the server

func (i *Icws) GetStatus(userId string) (defaults ConfigRecord, err error)
    gets the status information for a user

func (i *Icws) GetVersion() (version ServerVersion, err error)
    gets server version details. Does not need to be connected to the server

func (i *Icws) InteractionAction(action, interactionId, attribute string) (result ConfigRecord, err error)
    Performs an action on a given interaction. e.g. Pickup, Hold,
    Disconnect, etc

func (i *Icws) Login(applicationName, server, username, password string) (err error)
    Logs into a CIC server. Server should be a server name e.g.
    MyServer.domain.com

func (i *Icws) LoginMarketPlaceApp(applicationName, server, username, password, marketplaceLicense, markeplaceAppKey string) (err error)
    Logs into a CIC server for a MarketPlace application using the app's
    custom license. Server should be a url e.g. https://MyServer:8019

func (i *Icws) MakeCall(target string) (result ConfigRecord, err error)
    Creates a new call

func (i *Icws) SelectConfigurationRecords(objectType, selectFields, where string) (records []ConfigRecord, err error)
    Returns a list of matching records for the given object type.

func (i *Icws) SetPassword(userId, password string, force bool) error
    changes the password for the given user

func (i *Icws) SetStatus(userId, statusKey string) (err error)
    Sets a status for a given user id


type ServerFeature struct {
    //Id of the features
    FeatureId string
    //version of the feature
    Version int
}
    Definition of a server feature



type ServerVersion struct {
    //The product's two-digit release year. For the release "CIC 2015 R1" this value will be "15"
    MajorVersion string
    //The product's release number. For the release "CIC 2015 R1" this value will be "1".
    MinorVersion string
    //The patch number of the release. The value "0" indicates the release without any patches. For the release "CIC 2015 R1" this value will be "0", and for the release "CIC 2015 R1 Patch2" this value will be "2".
    Su string
    //The product line identifier.
    ProductId string
    //The codebase identifier.
    CodeBaseId string
    //The build number.
    Build string
    //The display string for the release. This does not include patch information.
    ProductReleaseDisplayString string
    //The display string for the release including patch information. This string is recommended for use on application "About" screens.
    ProductPatchDisplayString string
}
    Version information for the server




