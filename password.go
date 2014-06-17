package icws_golib

import "fmt"

//changes the password for the given user
func (i *Icws) SetPassword(userId, password string, force bool)(error){
    var passwordData = map[string]string{
        "password": password,
        "target": fmt.Sprintf("%b", force),
    }
    _, err, _ := i.httpPost("/configuration/users/" + userId + "/password", passwordData)

    return err

}
