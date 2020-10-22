package main
import("fmt"
  	"io/ioutil"
//	"bufio"
	"io"
//	"os"
//	"errors"
//	"unicode/utf8"
//	"strconv"
	"regexp"
	"os"
	
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
	parsedContents := parseNum(fileContents)
	fmt.Println(parsedContents)
//	f, err := os.Create(fileToWrite)
//	f, err := filepath.Abs(filepath.Dir(os.Args[0]))
//	check(err)
//	fmt.Println(f)
	writeFile(parsedContents,fileToWrite)
}

func openFile(fileToOpen string) string{
	contents, err := ioutil.ReadFile(fileToOpen)
	check(err)
	return string(contents)
}
func  writeFile(parsedContents string, fileToWrite string) error{
	file, err := os.Create(fileToWrite)
	check(err)
	defer file.Close()
 	_,err = io.WriteString(file, parsedContents)
	return file.Sync()
//	err = f.Close()
}
func parseNum(fileContents string) string{
//	var parsedContents string
	reg, err := regexp.Compile("[0-9]")
	if err != nil{
		check(err)
	}
	parsedContents := reg.ReplaceAllString(fileContents,"")
//	fileContents := strconv.QuoteRuneToASCII(fileContents)
//	fileContents = utf8.DecodeRuneInString(fileContents)
//	for i := 0; i <len(fileContents); i++{
	//	fmt.Print(fileContents[i])
	//	if _, err := strconv.ParseInt(fileContents[i],10,32); err == nil{
	//	fmt.Print("number")
	//	}
	//	else {
	//	parsedContents += fileContents[i]
	//	}
	//	j := strconv.QuoteToASCII(fileContents[i])
	//	fmt.Print(j)
	//	if unicode.IsNumber(fileContents){
	//	fmt.Println(" taken out!")
	//	}
	//	else{ 
	//	parsedContents.WriteString(fileContents[i])
	//	}
//	}
	return string(parsedContents)
}
