package server

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var rawCmd = &cobra.Command{
	Use: "raw",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting raw server")
		srv := http.Server{
			Addr:        ":8081",
			Handler:     http.HandlerFunc(handler),
			IdleTimeout: 5 * time.Second,
		}
		log.Fatalln(srv.ListenAndServe())

	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	time.Sleep(time.Second * 15)
	_, _ = w.Write([]byte("ok"))
}
