package main

import "fmt"

type stringMap = map[string]string



func main() {
	bookmarks := make(stringMap, 3)

	for {
		action := scanAction()
		switch action {
		case 1:
			fmt.Println(bookmarks)
		case 2:
			name, address := scanBookmark()
			bookmarks[name] = address
		case 3:
			name := scanBookmarkNameForDel()
			delete(bookmarks, name)
		case 4:
			return
		default: 
			fmt.Println("Неправильное действие")
		}
	}
}

func scanAction() int{
	fmt.Println("\n1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	var ans int
	fmt.Scan(&ans)
	return ans
}

func scanBookmark() (string, string){
	var name, address string
	fmt.Print("Введите название: ")
	fmt.Scan(&name)
	fmt.Print("Введите адрес: ")
	fmt.Scan(&address)
	return name, address
}

func scanBookmarkNameForDel() string {
	var name string
	fmt.Print("Введите название для удаления: ")
	fmt.Scan(&name)
	return name
}