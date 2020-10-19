package main

import (
	"fmt"
	"sort"

	"github.com/tezzzka/go-intro/homework4"
	"github.com/tezzzka/go-intro/homework4/abstract"
	"github.com/tezzzka/go-intro/homework4/calculator"
)

type userService struct {
	storage homework4.AccountStorage
}

func (s *userService) GetByID(id int64) (*abstract.User, error) {
	return s.storage.GetByID(id)
}

func (s *userService) NewUser(ID int64, Login string) error {
	return s.storage.NewUser(&abstract.User{
		ID:    ID,
		Login: Login,
	})
}

func (s *userService) SetPasswd(id int64, passwd string) error {
	s.storage.SetPasswd(id, passwd)
	return nil
}

func (s *userService) ResetPasswd(id int64, newpasswd string, root string) error {
	s.storage.ResetPasswd(id, newpasswd, root)
	return nil
}

func main() {
	mapStorage := abstract.Init()

	service := &userService{
		storage: mapStorage,
	}

	service.NewUser(1, "tzk")

	if user, e := service.GetByID(1); e != nil {
		panic("Cannot Get User By ID. Sorry.")
	} else {
		fmt.Println(user)
	}

	if pw := service.SetPasswd(1, "hello"); pw != nil {
		panic("Cannot set any password. Sorry.")
	}

	service.ResetPasswd(1, "PA$$W0RD", "Hello")
	//fmt.Println(service.GetByID(1))

	/*#2*/
	// X := abstract.SimpleInitContacts()
	// fmt.Print(&X.[])
	// sort.Slice(X, func(i, j int) bool {
	// 	return X[i].Cell < X[j].Cell
	// })
	// fmt.Println(abstract.SimpleInitContacts())

	group := []struct {
		name string
		age  int
	}{
		{"Gold", 14}, {"dave", 18}, {"kane", 12}, {"rain", 34},
	}

	sort.Slice(group, func(i, j int) bool {
		return group[i].age < group[j].age
	})
	// fmt.Println(group)

	// #3
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}
		if input == "help" {
			fmt.Println("This is manual page.")
		} else {
			if res, err := calculator.Calculate(input); err == nil {
				fmt.Printf("Результат: %v\n", res)
			} else {
				fmt.Println("Не удалось произвести вычисление")
			}
		}

	}

}
