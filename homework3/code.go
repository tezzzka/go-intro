package homework3

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type AutoGeneral struct {
	Vendor        string
	Model         string
	Year          int16
	EngineVolume  float32
	EngineType    string
	IsEngineStart bool
	IsWindowsOpen bool
	Volume        float32
}

type driver struct {
	id   int16
	name string
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}

func (dv *driver) GetDriverID() int16 {
	return dv.id
}
func (dv *driver) GetDriverName() string {
	return dv.name
}
func (dv *driver) setDriverID(v int16) bool {
	dv.id = v
	return true
}
func (dv *driver) SetDriverName(val string, password string) bool {
	if len(val) != 0 && len(password) != 0 {
		/* Создаем структуру, в которую запишем данные из ф
		базы данных с паролем пользователя
		*/
		type UserDataJSON struct {
			DriverName string `json:"DriverName"`
			Password   string `json:"Password"`
		}

		/* Доступ к файлу json на чтение/запись ограничиваем
		   на уровне ОС. Позволяем получать доступ только
			процессу-демону нашей программе на Go. В этом файле
			схематично храним хэш пароля на операцию изменения
			именя водителя.
		*/
		db, err := ioutil.ReadFile("DB.json")

		//os.Create("db2.json") Мучился с путями windows :)
		if err != nil {
			fmt.Println(err)
		}

		UserData := &UserDataJSON{}

		if err := json.Unmarshal([]byte(db), UserData); err != nil {
			fmt.Println(err)
		}
		// Вытаскиваем Хэш пароля пользователя
		Psw := UserData.Password
		hash := GetMD5Hash(password)

		if result := hash == Psw; result == false {
			panic("Access denied to make any changes. Sorry.")
		}

		//fmt.Println(hash)
		dv.name = val
		dv.setDriverID(5)

		return true
	}
	return false
}

type Car struct {
	AutoGeneral
	driver
	Comment string
}

type Lorry struct {
	AutoGeneral
	driver
	route string //Маршрут следования служебного транспортного средства.
	acl   string //Разрешение на перевозку груза.
}
