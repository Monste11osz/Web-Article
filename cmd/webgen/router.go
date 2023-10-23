package main

import "net/http"

func (pl *ProcessingLog) routers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", pl.home)
	mux.HandleFunc("/articles", pl.ShowArticles)
	mux.HandleFunc("/creat/article", pl.CreatArticle)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./iu/state/")})
	mux.Handle("/state/", http.StripPrefix("/state", fileServer))
	//mux.Handle("/state", http.NotFoundHandler())
	return mux
}
