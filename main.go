package main

import (
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/google/uuid"
	"github.com/shinobi-mtr/be-wish/lib"
)

func handleUUIDPost(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "uuid")
	if err := uuid.Validate(filename); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cl, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := make([]byte, cl)

	n, err := r.Body.Read(data)
	if n != cl && err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	filename = path.Join("./public", filename)
	if err := lib.AppendDataToFile(filename, data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func handleUUIDGet(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "uuid")
	offset := 0

	if err := uuid.Validate(filename); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cok, err := r.Cookie(filename)
	if err == nil {
		offset, err = strconv.Atoi(cok.Value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	data, err := lib.GetDataFromFile(path.Join("./public", filename), int64(offset))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     filename,
		Value:    strconv.Itoa(offset + len(data)),
		HttpOnly: true,
	})

	w.Write(data)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(httprate.LimitByIP(1, 1*time.Minute))

	r.Post("/{uuid}", handleUUIDPost)
	r.Get("/{uuid}", handleUUIDGet)

	http.ListenAndServe(":3000", r)
}
