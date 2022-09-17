package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	hostname, _ := os.Hostname()

	fmt.Fprintf(os.Stdout, "Server has started on %s.\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")

		result := struct {
			hostname string
		}{
			hostname: hostname,
		}

		json := convertToJSON(result)
		writeResponse(w, string(json))
	})

	http.HandleFunc("/date", func(w http.ResponseWriter, r *http.Request) {
		date, _ := time.Now().MarshalJSON()

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")

		result := struct {
			hostname string
			date     string
		}{
			hostname: hostname,
			date:     string(date),
		}

		json := convertToJSON(result)
		writeResponse(w, string(json))
	})

	http.HandleFunc("/joke", func(w http.ResponseWriter, r *http.Request) {

		client := &http.Client{}
		req, _ := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
		req.Header.Add("Accept", "text/plain")

		response, err := client.Do(req)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			return
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			return
		}

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")

		result := struct {
			hostname string
			joke     string
		}{
			hostname: hostname,
			joke:     string(responseData),
		}

		json := convertToJSON(result)
		writeResponse(w, string(json))
	})

	listenError := http.ListenAndServe(":"+port, nil)
	log.Fatal(listenError)
}

func writeResponse(w io.Writer, response string) {
	fmt.Fprintln(w, response)
	fmt.Fprintln(os.Stdout, response)
}

func convertToJSON(data any) string {
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return ""
	}

	return string(json)
}
