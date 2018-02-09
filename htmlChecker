package main 

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	//посылаем запрост на сервер методом get и в resp записываем ответ сервера
  //send request and writing response in resp
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("http.Get: %v", err)
	}
	//откладываем закрытие соединения, оно будет произведено после чтения 
  //defer connection closing
	defer resp.Body.Close()

	//средство io вытаскиивает из resp байты данных
  //get data from resp
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll: %v", err)
	}
	//преобразуем байты в строку и ищем там "word"
  //convert bytes into string and find "word"
	fmt.Println(strings.Contains(string(bytes), "word"), resp)
}
