package main

// go packages documentation url https://pkg.go.dev/std
import (
	"fmt"
	"sort"
	"strings"
)

// a global variable can be decleared like this
var nickname = "jason"

// or this
var nicknamex string = "reacher"

var d int = 987990

func main() {
	// this is how to declear a string in go and in go you only use double quot for string
	// and it a variable is decleard like this name string you can only assign to it
	var name string = "stephen"
	var surname = "okpeku"
	var emptytest string
	// this can only be used inside a function in go
	middlename := "ighodaro"
	// fmt.Println this is used to console or echo variable in go
	fmt.Println(name, surname, middlename, nickname, nicknamex)
	// for number if you set the type to int you cant use decmical point on it
	var a int = 3747
	var b = 2738
	c := 37478
	fmt.Println(a, b, c, d)

	// you can mulpulate existing variable or an empty variable in go
	name = "Peter"
	emptytest = "testing testing"
	fmt.Println(name, emptytest)

	// you can set how long long you number and decmical point can be in go
	// check this url to learn more https://go.dev/ref/spec#Numeric_types

	// int8 the set of all signed  8-bit integers (-128 to 127)
	var numOne int8 = 127
	var numTwo int16 = 2457
	fmt.Println(numOne, numTwo)

	// for demical point you have to set the type

	var numThree float32 = 24.4
	var numFour float64 = 4785785058874.98
	var numSix float64 = 384784.8
	// this is the shorthand of float64
	numFive := 459858895.0585
	fmt.Println(numThree, numFour, numFive, numSix)
	// fmt.Print this does not add space or new line you have to do it manually
	// fmt.Println this automacatlly add a new line
	fmt.Print("hello")
	fmt.Print("world \n")
	fmt.Print("Stephen ")

	// string format and concatantion
	// %v is both works on string and number and if you use it twice or five time or more you have to declear two or five or more variable
	fmt.Printf("my name is %v and my surname is %v \n", name, surname)
	//  %q this only work for string if you want the number to show here you have to conver the number to a string
	fmt.Printf("my nickname is %q and my age is %q \n", nickname, b)
	// %T is is use to show the variable if it is a string, number or boolean
	fmt.Printf("this variable is what type %T \n", numFour)
	// %f this is used for demecial numbers and to make not show 24.400000 you can use this to approximate %0.1 that means one demecial place it will be 24.4
	fmt.Printf("%0.1f \n", numThree)

	// this save the string format and you and assign it to a variable
	var word = fmt.Sprintf("my name is %v and my surname is %v \n", name, surname)
	fmt.Println(word)

	// array
	//   in go array delceared like this, [3] it is used to set the lenght of the array and int is used to define the type variable
	// in the array  and you must also  add  this [3]int after = in go and {} is the array
	var allnumarr [3]int = [3]int{34, 35, 57}
	fmt.Println(allnumarr)

	// short hand of array
	// if you want to add or push to an array that the length is already set it will not work you can only mupulate what exist in the
	// array
	var arrnum = [4]int{4, 5, 7, 8}
	arr := [3]int{45, 67, 56}
	fmt.Println(arrnum, arr)

	// i change 7 in the array to 9 in arrnum
	arrnum[2] = 9
	fmt.Println(arrnum)

	// string in array

	var arrname [4]string = [4]string{"stephen", "david", "paul", "samuel"}
	// shorthand
	var arrnames = [4]string{"tobi", "ayo", "kenule", "lukeman"}
	arrnam := [3]string{"dog", "cat", "rat"}
	arrnam[1] = "horse"
	fmt.Println(arrname, "\n", arrnames, "\n", arrnam)

	// Slice this are array the have no limit and new items can be addded to the array
	var arrwords = []string{"school", "shoe", "gym"}
	//  you have to assign it back to the array to overwrite old and insert the new one
	arrwords = append(arrwords, "stephen")
	fmt.Println(arrwords)

	// this is used to check the lenght of an array in go
	fmt.Println(len(arrname))

	// Slice Range
	// this is used to get from one part of the array to another part of the array
	// arrnam[0:2] it will start from index 0 to 1
	rangeOne := arrnam[0:2]
	fmt.Println(rangeOne)

	// it will start from index one to last item in the array
	rangeTwo := arrwords[1:]
	fmt.Println(rangeTwo)
	// it will start from index 0 to index 3
	rangeThree := arrnames[:4]
	fmt.Println(rangeThree)

	wordx := "hello word nice to meet you"
	// strings.Contains() this is used to check if a string exist in a string
	fmt.Println(strings.Contains(wordx, "hello"))

	// this is use to check the index of any letter or letters in the string
	fmt.Println(strings.Index(wordx, "ll"))

	// this is used to replace a string in an existing string
	fmt.Println(strings.ReplaceAll(wordx, "hello", "hi"))

	// split this is used to convert the string into an array
	fmt.Println(strings.Split(wordx, " "))

	// sorting and searching of array of String and Int

	agenum := []int{60, 5, 70, 3, 20, 30, 15, 0, 300}

	// sort.Ints :- this is used to sort array that have number
	sort.Ints(agenum)
	fmt.Println(agenum)

	// sort.SearchInts this is used search the int array and you have to assign to another variable
	searchint := sort.SearchInts(agenum, 15)
	fmt.Println(searchint)

	wordsarr := []string{"goat", "cow", "tea", "logo", "suger", "zen"}
	//this is used sort a string array
	sort.Strings(wordsarr)
	fmt.Println(wordsarr)

	// sort.SearchStrings this is used search the string array and you have to assign to another variable

	searchstring := sort.SearchStrings(wordsarr, "logo")
	fmt.Println(searchstring)

	// looping
	// x := 0
	// for x < 5 {
	// 	// first way to for loop
	// 	fmt.Printf("the value is: %v \n", x)
	// 	x++
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("the value is: %v \n", i)
	// }

	// this is how to loop through a slice array in go
	wordsarrloop := []string{"goat", "cow", "tea", "logo", "suger", "zen"}
	// for i := 0; i < len(wordsarrloop); i++ {
	// 	fmt.Println("here", wordsarrloop[i])
	// }

	// foreach loop in go
	// for index, value := range wordsarrloop {
	// 	fmt.Printf("the index %v and  word %v \n", index, value)
	// }

	// if i do not want to display this index or value use _
	for _, value := range wordsarrloop {
		fmt.Printf("the   word %v \n", value)
	}
}