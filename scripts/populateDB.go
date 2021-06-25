// file responsable to populate the created API

package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"yawoen.com/scripts/pkg"
)

//#region Package variables
var logError *log.Logger
var logInfo *log.Logger

const string_sep = ";"                // Field CSV sep
const short_timeout = 1 * time.Second // Maximum delay until closes the application

//#endregion

func main() {
	// check it out the parameters and loads the menu options
	opt := ReadArgs()

	// Menu handler
	switch opt.Option {
	case "-h":
		opt.Help()
		return
	case "-sql":
		opt.ToSQL()
		return
	case "-request":
		opt.ToHttpRequest()
		return
	default:
		opt.Help()
		return
	}

}

// Startup function (Setup)
func init() {
	// init the loggers
	logError = log.New(os.Stderr, "", 0) // print to stderr buffer
	logInfo = log.New(os.Stdout, "", 0)  // print to stdout buffer (default)
}

//#region: Functions

type Options struct {
	Filename      string   // the filename name
	Option        string   // if should request to an api or generate an sql file
	Method        string   // -request: PUT|PATCH|DELETE|GET|POST
	Url           string   // -request: http://www.google.com
	HeaderPayload []string // -request: "name;zip"
}

// Read the parameters arguments
func ReadArgs() Options {
	var opt Options
	opt.HeaderPayload = []string{}

	// check if exists elements in parameters
	if os.Args == nil || len(os.Args) < 2 {
		logError.Printf("You should pass arguments.\nCheck it out by using -h flag. Example:\n\n")
		return opt
	}

	// Parse the option
	opt.Option = os.Args[1]
	if opt.Option != "-h" && opt.Option != "-sql" && opt.Option != "-request" {
		logError.Fatal("The option parameter must be one of those:\n-sql\n-request\n-h Shows the help menu")
	}
	if opt.Option == "-h" {
		return opt
	}

	// Parse the filename and check if exists
	opt.Filename = os.Args[2]
	if _, err := os.Stat(opt.Filename); os.IsNotExist(err) {
		logError.Fatal("The file doesn't exists! :(")
	}
	// check if is a csv file
	if opt.Filename[len(opt.Filename)-3:] != "csv" {
		logError.Fatal("The file should be a CSV file")
	}

	if len(os.Args) >= 5 {
		// Parse the method options
		opt.Method = os.Args[3]
		if (opt.Method != "PUT") && (opt.Method != "PATCH") && (opt.Method != "DELETE") && (opt.Method != "POST") && (opt.Method != "GET") {
			logError.Fatal("The method should be one of PUT|PATCH|DELETE|GET|POST")
		}

		// Parse endpoint
		opt.Url = os.Args[4]
		if !govalidator.IsURL(opt.Url) {
			logError.Fatal("The ENDPOINT should be a valid URL string")
		}

		// Parse map
		if len(os.Args) == 6 {
			opt.HeaderPayload = strings.Split(os.Args[5], ";")
			if len(opt.HeaderPayload) == 0 {
				logError.Fatal("The CSV header use the semicollon ';' to separate the headers")
			}
		}
	}
	return opt
}

// Prints the helper function to this file
func (opt *Options) Help() {
	logInfo.Println("Open a CSV file concurrently and execute some action according to flag passed to")
	logInfo.Println("Usage:")
	logInfo.Println("populateDB [OPTION] [FILENAME] [...ARGS]")
	logInfo.Println("\nWhere:")
	logInfo.Println("[OPTION] is a flag that can be:")
	logInfo.Println("\t-sql\t\tGenerate a sql file inserting to the associated table")
	logInfo.Println("\t-request\tDoes an http request to the ARGS")
	logInfo.Println("\n[...ARGS] is only used when passed the -request option. It makes a JSON request to the endpoint specified")
	logInfo.Println("Usage: METHOD ENDPOINT CSV_HEADER")
	logInfo.Println("\tMETHOD \t\tis some of them PUT|GET|POST|DELETE|PATCH")
	logInfo.Println("\tENDPOINT \tis a URL string")
	logInfo.Println("\tCSV_HEADER \tis a mapping to csv positions switching the field send in JSON payload. It must have the same number of columns that the csv file.")
	logInfo.Println("\n\tExample: PUT \"http://localhost:3000/company\" \"name;zip\"")
}

// Generate a new SQL file with the same name inserting in the endpoint specified
// TODO: Implement this function
func (opt *Options) ToSQL() {}

// Make http requests to an specified endpoint
//
// TODO: implement
func (opt *Options) ToHttpRequest() {
	//#region Abre o arquivo
	logInfo.Println("Opening the file")
	file, err := os.Open(opt.Filename)
	if err != nil {
		logError.Panic("Couldn't open the file", err)
		return
	}
	defer file.Close() // Close the file when reaches the end of this escope
	//#endregion

	//#region Reading file
	logInfo.Println("Creating the buffer")
	fileStream := bufio.NewReader(file)

	logInfo.Println("Reading header")
	firstLine, err := fileStream.ReadBytes('\n') // read file and walk the pointer until finds \n char
	buf := make([]byte, 0)                       // creates a buffer
	if err != nil {                              // check if some error occurred
		logError.Panic(err)
	} else {
		buf = append(buf, firstLine...) // otherwise, append this new buffer to the old one
	}

	logInfo.Println("Spliting header...")
	header := strings.Split(strings.TrimSpace(string(buf)), string_sep)
	// Replacing the header if wasn't passed none to script
	logInfo.Println("Header Payload: ", opt.HeaderPayload)
	if len(opt.HeaderPayload) == 0 {
		logInfo.Println("Changing the payload header")
		opt.HeaderPayload = header
		logInfo.Printf("|%s|, |%s|, |%s|\n", header[0], header[1], header[2])
	}
	//#endregion

	//#region walking over file using chunks of data
	pkg.ProcessFile(*fileStream, func(s *string) {
		ctx, cancel := context.WithTimeout(context.Background(), short_timeout)
		defer cancel()

		finished := make(chan string, 1) // When finished the current work, emits a signal
		go func() {                      // The function itself
			// time.Sleep(2 * time.Second)                            // debug only
			word := strings.Split(strings.TrimSpace(*s), string_sep) // split the line
			if len(word) != len(opt.HeaderPayload) {                 // some error occurred with this line
				logError.Panic(word)
			}
			payload := make(map[string]string) // create the payload map
			for i, word := range word {        // otherwise, map the header to the word value
				payload[opt.HeaderPayload[i]] = word
			}

			jsonPayload, err := json.MarshalIndent(payload, "", " ") // convert to json payload
			if err != nil {
				logError.Panic("Occurred an error when reading the response")
			}
			request := bytes.NewBuffer(jsonPayload) // convert the json to bytes handler

			// Does the requisitions
			req, _ := http.NewRequest(opt.Method, opt.Url, request) // create a new basic request config
			req.Header.Set("Content-Type", "application/json")      // set the body encoding
			client := &http.Client{}                                // creates a new client
			response, err := client.Do(req)                         // gets the response from server
			if err != nil {
				logInfo.Println("The API responded with", err)
			}
			defer response.Body.Close() // finish the listener

			// Parse the body element
			body, err := ioutil.ReadAll(response.Body) // Read the response from server
			if err != nil {
				logError.Panic("Occurred an error when reading the response")
			}

			finished <- string(body) // return the current payload element
		}()

		// Configure the function to stop all the operations if the api took too long to
		// respond
		select {
		case payload := <-finished:
			// NOTE: debug only
			log.Println(payload)

		case <-ctx.Done():
			logError.Fatal("The api tooked too long to respond, deadline exceed")
		}

	})
	//#endregion
}

//#endregion
