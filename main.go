package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/obstmi/gosea/status"
)

func main() {

	//log.Print("starting service")

	//logfile, err := os.Open("messages.log")
	logfile, err := os.Create("messages.log")

	if err != nil {
		log.Printf("error opening log file: %s", err.Error())
		os.Exit(1) // entspricht log.Fatal()
	}

	defer logfile.Close() // verschiebt die Ausführung an das Ende der Funktion

	// _ = ungenutzteVariable // -> zu Testzwecken, um ungenutzte Variablen zu übergeben

	logger := log.New(os.Stdout, "gosea", log.LstdFlags)
	//logger := log.New(logfile, "gosea", log.LstdFlags)

	logger.Print("starting service")

	sigChan := make(chan os.Signal) // channel muss initialisiert werden. Ohne Parameter default = 1 Datentyp ("Pipeline")
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", status.Health) // Wieso geht das? func status.Health hat doch 2 Parameter? -> Funktion wird übergeben

	srv := &http.Server{ // Belegung mit Werten mit Kommata, Initialisierung ohne Komma
		Addr:    ":8000",
		Handler: mux,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil /* && !errors.Is(err, http.ErrServerClosed) */ { // Fehler nicht ausgeben, wenn Server beendet wird
			log.Printf("error starting server: %s", err.Error())

		}
	}()

	//go srv.ListenAndServe() // Funktion wird in einem eigenen Kontext ausgeführt

	logger.Print("starting server")

	<-sigChan
	//sig := <-sigChan // aus Channel lesen

	srv.Close()
	close(sigChan // könnte ich auch nach dem make(chan..) mit defer close(sigChan) machen

	logger.Print("stopping server") // kommt an dieser Stelle nicht dahin, wenn der Dienst mit Strg+c "abgeschossen" wird, auch nicht mit defer()

	// channel und go-Keyword sind die Aspekte der Nebenläufigkeit
}
