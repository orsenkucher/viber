package viber

import (
	"encoding/json"
	"fmt"
)

// User struct as part of UserDetails
type User struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	Country         string `json:"country"`
	Language        string `json:"language"`
	PrimaryDeviceOs string `json:"primary_device_os"`
	APIVersion      int    `json:"api_version"`
	ViberVersion    string `json:"viber_version"`
	Mcc             int    `json:"mcc"`
	Mnc             int    `json:"mnc"`
	DeviceType      string `json:"device_type"`
}

// UserDetails for Viber user
type UserDetails struct {
	Status        int    `json:"status"`
	StatusMessage string `json:"status_message"`
	MessageToken  int64  `json:"message_token"`
	User          `json:"user"`
}

// Online status struct
type online struct {
	Status        int          `json:"status"`
	StatusMessage string       `json:"status_message"`
	Users         []UserOnline `json:"users"`
}

// UserOnline response struct
type UserOnline struct {
	ID                  string `json:"id"`
	OnlineStatus        int    `json:"online_status"`
	OnlineStatusMessage string `json:"online_status_message"`
	LastOnline          int64  `json:"last_online,omitempty"`
}

// UserDetails of user id
func (v *Viber) UserDetails(id string) (UserDetails, error) {
	var u UserDetails
	s := struct {
		ID string `json:"id"`
	}{
		ID: id,
	}

	b, err := v.PostData(fmt.Sprintf("%s/get_user_details", ViberAPI), s)
	if err != nil {
		return u, err
	}

	if err := json.Unmarshal(b, &u); err != nil {
		return u, err
	}

	// viber error returned
	if u.Status != 0 {
		return u, Error{Status: u.Status, StatusMessage: u.StatusMessage}
	}

	return u, err

}

// UserOnline status
func (v *Viber) UserOnline(ids []string) ([]UserOnline, error) {
	var uo online
	req := struct {
		IDs []string `json:"ids"`
	}{
		IDs: ids,
	}
	b, err := v.PostData(fmt.Sprintf("%s/get_online", ViberAPI), req)
	if err != nil {
		return []UserOnline{}, err
	}

	if err := json.Unmarshal(b, &uo); err != nil {
		return []UserOnline{}, err
	}

	// viber error
	if uo.Status != 0 {
		return []UserOnline{}, Error{Status: uo.Status, StatusMessage: uo.StatusMessage}
	}

	return uo.Users, nil
}
