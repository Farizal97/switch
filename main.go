package main

import "fmt"

//

func main() {

	students := []map[string]string{
		{"name":"Agung","score": "A"},
		{"name":"Lim","score":"B"},
		{"name":"Lin","score" : "C"},
	}

	for _,student := range students {
		fmt.Println(student["name"], "-",student["score"])
	}

	
	// myMap :=  map[string]string{
	// 	"ruby":"is beautiful",
	// 	"go":"is super fast",
	// }

	// value:= myMap["phyton"]
	// fmt.Println(value)
}

	// for key, value := range myMap{
	// 	fmt.Println("Key:",key,"value:",value)
	// }

	// 	fmt.Println("===========")

	// 	delete(myMap,"ruby")

		
	// for key, value := range myMap{
	// 	fmt.Println("Key:",key,"value:",value)
	// }
		//fmt.Println(MyMap)
	 //gamingConsoles := []string{"playstation4",""}

	//fmt.Println(gamingConsoles)

	// gamingConsoles = append(gamingConsoles,"Playstation4")





	// var languages [5] string
	// languages[0] = "go"
	// languages[1] = "php"
	// languages[2] = "js"
	// languages[3] = "java"
	// languages[4] = "xcode"

	// languages := [5]string{"Ruby","phyton","js","java","xcode"}

	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println(length)

	// title:= "golang the best"

//

	// for index,letter :=  range title {
	// 	letterString := string(letter)
		//if index%2 == 0 {
		//fmt.Println("index:",index,"letter :",string(letter))

		// switch letterString{
		// case "a","i","u","e","o":
		// 	fmt.Println("index:",index,"letter:",string(letter))
		// }

		// }
	



	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("saya sedang Belajar Go:",i)
	// }

	// i := 1 
	// for i <=100 {
	// 	fmt.Println("saya sedang Belajar Go:",i)
	// 	i++
	// }


	//number := 20

	//switch number {
	//case 1:
	//	fmt.Println("Satu")
	//case 2:
	//	fmt.Println("Dua")
	//case 3:
	//	fmt.Println("Tiga")
	//default:
	//	fmt.Println("DEFAULT")
	//}
