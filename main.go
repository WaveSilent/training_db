package main

import (
	"fmt"
	"study/http_server"
)

func main() {
	fmt.Println("Запуск HTTP сервера!")

	if err := http_server.StartHTTPServer(); err != nil {
		fmt.Println("Во время сервера произошла ошибка:", err)
	} else {
		fmt.Println("Сервер завершился успешно!")
	}
}
