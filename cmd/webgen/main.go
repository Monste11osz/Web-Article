package main

import (
	"WebPrac/pkg/models/mysql"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Для хранения зависимостей
type ProcessingLog struct {
	errLog  *log.Logger
	infoLog *log.Logger
	article *mysql.WebPracMod
}

func main() {

	addr := flag.String("addr", ":8000", "web address")
	flag.Parse()

	dsn := flag.String("mysql", "web:@/WebPrac?parseTime=true", "База данных")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errLog.Fatal(err)
	}

	defer db.Close()

	pl := &ProcessingLog{
		errLog:  errLog,
		infoLog: infoLog,
		article: &mysql.WebPracMod{DB: db},
	}

	//инициализация полей + свой логгер
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  pl.routers(),
	}

	infoLog.Printf("Запуск сервера %s", *addr)
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
