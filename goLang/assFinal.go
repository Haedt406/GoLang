package main
import("fmt"
  	"io/ioutil"
	"io"
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
	fmt.Println("Success!!")
	return file.Sync()
}
func parseNum(fileContents string) string{
	reg, err := regexp.Compile("[0-9]")
	if err != nil{
		check(err)
	}
	parsedContents := reg.ReplaceAllString(fileContents,"")
	return string(parsedContents)
}
