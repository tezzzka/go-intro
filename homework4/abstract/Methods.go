package abstract

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

type Acc struct {
	accounts map[int64]*User
}
func Init() *Acc {
	return &Acc{
		accounts: make(map[int64]*User),
	}
}
func (a *Acc) NewUser(u *User) error {
	a.accounts[u.ID] = u

	return nil
}
func (a *Acc) GetByID(id int64) (*User, error) {
	user, ok := a.accounts[id]
	if !ok {
		return nil, errors.New("User is not found.")
	}
	return user, nil
}
func (a *Acc) SetPasswd(id int64, passwd string) error {
	(a.accounts[id]).Passwd = passwd

	return nil
}
func (a *Acc) ResetPasswd(id int64, newpasswd string, root string) error {

	type pw struct {
		RootPassword string `json:"RootPassword"`
	}

	su, e := ioutil.ReadFile("secret.json")

	if e != nil {
		panic("No data available to receive. Please, check this point to engange in it. Thank you.")
	}
	Pw := &pw{}
	if err := json.Unmarshal([]byte(su), Pw); err != nil {
		panic("Cannot read a byte. Sorry.")
	}
	PswRoot := Pw.RootPassword
	PswUser := strings.ToUpper(GetMD5Hash(root))

	if PswRoot != PswUser {
		panic("Access denied. Password is incorrect.")
	}

	(a.accounts[id]).Passwd = newpasswd

	return nil
}

/* #2 */

type cn struct {
	contacts *Contacts
}



func CNInit() *cn {
	return &cn.contacts[0] {
		Name: "tezzzka",
		Cell: 9320909555,

	}
}

// func (a *Contacts) Add(Name string, Cell int16) {
// 	anew := Contact{
// 		Name: Name,
// 		Cell: Cell,
// 	}
// 	*a = append(*a, anew)
// }
