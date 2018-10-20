package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "Go Server")
		fmt.Fprint(w, `<html>
      <body>
        <h1>Hello Gopher</h1>

				<h3>Articles</h3>
				<ul>
					<li><a href="/articles">/articles/</a></li>
					<li><a href="/articles/">/articles/</a></li>
					<li><a href="/articles/latest">/articles/latest</a></li>
				</ul>

					<h3>Articles</h3>
					<ul>
						<li><a href="/users">/users</a></li>
						<li><a href="/users/">/users/</a></li>
					</ul>

					<li><a href="/missed-page">Missed page</a></li>

      </body>
    </html>`)
	})

	mux.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /articles/")
	})

	mux.HandleFunc("/articles/latest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /articles/latest")
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from /users")
	})

	mux.HandleFunc("/missed-page", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://google.com", http.StatusMovedPermanently)

		// http.Error(w, "I can't find it :(", http.StatusNotFound)
		// http.NotFound(w, r)
	})
	http.ListenAndServe(":3000", mux)
}
