package main

import (
    "bufio"
    "os"
    "log"
    "fmt"
    "strings"
    "path/filepath"
)

func writeToJSON(path string, svgContent string) {
	fileHandle, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fileHandle, _ = os.OpenFile(path, os.O_CREATE, 0666)
	}
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()
	fmt.Fprintln(writer, "{")
	fmt.Fprintln(writer, "src:`"+svgContent+"`")
	fmt.Fprintln(writer, "}")
	writer.Flush()
}

func parse_args() (file string) {
	if len(os.Args) < 2 {
		log.Fatal("usage: svg-json-converter <root_js_src_main_dir>\nEx. [PROJECT_ROOT/src]")
	}
	file = os.Args[1]
	if strings.HasSuffix(file, "/") {
		file = file[:len(file)-1]
	}
	return
}

func walkGrepAllSVG(path1 string) {    
    filepath.Walk(path1, func(path string, f os.FileInfo, err error) error {
        if !f.IsDir() && (strings.HasSuffix(path, "svg")) {
            fmt.Println("found:",path)
            f, err1 := os.Open(path)
            if err1 != nil {
                log.Fatal(err1)
            }
            defer f.Close()
            scanner := bufio.NewScanner(f)
            svgContent:= ""
            for scanner.Scan() {
                if(len(scanner.Text()) > 0 ){
                    svgContent+="\n";
                }
                svgContent= svgContent+scanner.Text()
            }
            if len(svgContent) > 0 {
                filename:= path[:len(path)-4]+".json"
                fmt.Println("replaced with: ", filename)
                writeToJSON(filename, svgContent)
                //remove old svg file
                os.Remove(path)
            }
        }
        return nil
    })
}

func main(){
    projectJSSrcMainDir := parse_args()
    walkGrepAllSVG(projectJSSrcMainDir)
}