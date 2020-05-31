package status

import (
	"fmt"
	"net/http"
)

func Health(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("content-type", "text/plain")

	/*
		output := strings.NewReader("Ok")

		io.Copy(writer, output)
	*/

	/*
		writer.Write([]byte("Supi - alles klappt"))
	*/

	status := "ok"
	fmt.Fprintf(writer, "Health: %s", status)
}
