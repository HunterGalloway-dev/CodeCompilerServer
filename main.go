package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/gorilla/mux"
	//"encoding/json"
)

// File strict that contains data related to a received file
type File struct {
	FileName    string
	FileContent string
}

// Packet that contains data received from student
type Packet struct {
	StudentID   string
	Language    string
	ExecuteFile int
	Files       []File
}

// RunConfig stores information on how to run code
type RunConfig struct {
	Lan        string
	RunFile    string
	Files      []File
	Test       bool
	TestOption string
	TestIn     string
	TestOut    string
}

// Output structure used to send back to client
type Output struct {
	Output string
}

func runCode(packet Packet) string {
	ret := ""

	files := make([]string, len(packet.Files))

	dname, err := ioutil.TempDir("./", "student"+packet.StudentID)
	defer os.RemoveAll(dname)
	check(err)
	for i, file := range packet.Files {
		fname := filepath.Join(dname, file.FileName)
		files[i] = fname
		err = ioutil.WriteFile(fname, []byte(file.FileContent), 0666)
		check(err)
	}

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", "java -cp components.jar "+files[packet.ExecuteFile])
		out, _ := cmd.CombinedOutput()
		//check(err)
		ret = string(out)
	} else {
		fmt.Println("We ain't in kansas pal")
	}

	return ret
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	r := mux.NewRouter()

	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	r.HandleFunc("/foo", fooHandler).Methods(http.MethodPost, http.MethodPatch, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))

	http.ListenAndServe(":8080", r)
	fmt.Println("Starting server...")

	x := []File{
		{
			FileName:    "Main.java",
			FileContent: "public class Main {public static void main(String[] args) {System.out.println(\"Test\")}}",
		},
	}

	runConfig := &RunConfig{
		Lan:        "java",
		RunFile:    "Main.java",
		Files:      x,
		Test:       false,
		TestOption: "",
		TestIn:     "",
		TestOut:    "",
	}

	fmt.Println(*runConfig)

	dname, err := ioutil.TempDir("./", "student")
	defer os.RemoveAll(dname)
	check(err)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	var receivedPacket Packet
	json.Unmarshal(reqBody, &receivedPacket)
	codeOutput := runCode(receivedPacket)

	fmt.Println(reqBody)

	var runConfig RunConfig
	json.Unmarshal(reqBody, &runConfig)
	fmt.Println("Test: ", runConfig)

	sendPacket := Output{
		Output: codeOutput,
	}

	fmt.Println(sendPacket)

	marshalledOutput, err := json.Marshal(sendPacket)
	check(err)

	fmt.Println(string(marshalledOutput))

	w.Write(marshalledOutput)
}
