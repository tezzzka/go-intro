package homework5

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func FileRead(filename string) {
	//Добавили защиту от исключительной ситуации: отсутствия файла.
	file, err := os.OpenFile("testx.ini", os.O_APPEND|os.O_CREATE, 755)
	if err != nil {

		panic(err)
		return
	}
	defer file.Close()

	// getting size of file
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// reading file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	fmt.Println(string(bs))

}

type AddressBook struct {
	address []string
	desc    string
}

func CsvToStruct(Path string) (*AddressBook, error) {
	//var address []string
	var p = &AddressBook{}
	resp, e := http.Get(Path)

	if e != nil {
		return nil, e
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, e
	}
	txt := string(content)
	r := csv.NewReader(strings.NewReader(txt))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil, e
	}
	for key, val := range records {
		//address = append(address, val[key])
		//fmt.Println(val[0])
		p.address = append(p.address, val[key])
	}
	p.desc = time.Now().Format("2006-01-02T15:04:05Z07:00")
	return p, nil

}
