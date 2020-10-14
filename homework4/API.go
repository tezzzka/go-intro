package homework4

import "github.com/tezzzka/go-intro/homework4/abstract"

type AccountStorage interface {
	GetByID(id int64) (*abstract.User, error)
	NewUser(*abstract.User) error
	SetPasswd(id int64, passwd string) error
	ResetPasswd(id int64, newpasswd string, root string) error
	// RemoveUser(id int64, root string) error
	// SetACL(rules string, root string) error
}
