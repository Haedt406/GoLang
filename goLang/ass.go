package main
import("fmt"
  	"io/ioutil"
//	"bufio"
//	"io"
//	"os"
//	"errors"
//	"unicode"
	"strconv"
)

func check(e error) {
	if e!= nil{
	panic(e)
	}
}

func main(){

	var fileToOpen, fileToWrite string // fileContents, parsedContents string
	fmt.Print("Enter the name of the file to open: ")
	fmt.Scanln(&fileToOpen)
	fmt.Print("Enter the name of the file to output to: ")
	fmt.Scanln(&fileToWrite)
	fmt.Println("file to open is: " + fileToOpen + ". File to Write is: " + fileToWrite)
	fileContents := openFile(fileToOpen)
	fmt.Println(fileContents)
	fmt.Println()
	parsedContents := parseNum(fileContents)
	fmt.Println(parsedContents)
}

func openFile(fileToOpen string) string{
	contents, err := ioutil.ReadFile(fileToOpen)
	check(err)
	return string(contents)
}

func parseNum(fileContents string) string{
	var parsedContents string
	fileContents := strconv.QuoteRuneToASCII(fileContents)
	for i := 0; i <len(fileContents); i++{
		fmt.Println(fileContents[i])
	//	t := strconv.QuoteRuneToASCII(fileContents[i])
		if i.IsNumber(){
		fmt.Println(" taken out!")
		}
	//	else{ 
	//	parsedContents.WriteString(fileContents[i])
	//	}
	}
	return string(parsedContents)

}
