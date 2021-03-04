package main

// func main() {
// 	r := mux.NewRouter()

// 	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
// 	r.HandleFunc("/foo", fooHandler).Methods(http.MethodPost, http.MethodPatch, http.MethodOptions)
// 	r.Use(mux.CORSMethodMiddleware(r))

// 	http.ListenAndServe(":8080", r)
// }

// func fooHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	if r.Method == http.MethodOptions {
// 		return
// 	}

// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	fmt.Println(string(reqBody))

// 	w.Write([]byte("foo"))
// }
