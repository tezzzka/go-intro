package abstract

type User struct {
	ID    int64
	Login string
	userSec
	userGeneral
}

type userGeneral struct {
	Name        string
	Description string
	Address     string
	phone       int16
}

type userSec struct {
	Passwd   string
	checksum float64
	secret   string
}

type AddressBook struct {
	User
	cell map[int16]*int64
}
