package main

import (
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/bxcodec/faker/v3"
)

//go:embed dist
var dist embed.FS

type Person struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	IP    string `json:"ip"`
	Email string `json:"email"`
}

func main() {

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Connection", "Keep-Alive")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)

		ticker := time.NewTicker(time.Millisecond * 500)
		go func() {
			for {
				<-ticker.C
				p := Person{
					Name:  faker.Name(),
					Email: faker.Email(),
					IP:    faker.IPv4(),
					Phone: faker.Phonenumber(),
				}
				w.Write(ToJSON(p))
				w.(http.Flusher).Flush()
			}
		}()
		time.Sleep(time.Second * 100)
		ticker.Stop()
		fmt.Println("Finished: should return Content-Length: 0 here")
		w.Header().Set("Content-Length", "0")
	})
	d, _ := fs.Sub(dist, "dist")
	http.Handle("/", http.FileServer(http.FS(d)))
	http.ListenAndServe(":8080", nil)
}

func ToJSON(obj interface{}) []byte {
	buf, _ := json.Marshal(obj)
	return buf
}
