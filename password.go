package icws_golib

import "fmt"

//changes the password for the given user
func (i *Icws) SetPassword(userId, password string, force bool) error {
	var passwordData = map[string]string{
		"password": password,
		"force":    fmt.Sprintf("%t", force),
	}
	_, err := i.httpPut("/configuration/users/"+userId+"/password", passwordData)

	return err

}
