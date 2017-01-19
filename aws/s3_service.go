package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Server struct {
	S3Svc      *s3.S3
	Bucket     string
	ContentMap map[string]string
}

func (s *Server) checkAuth(username string, password string, ok bool) bool {
	/// @todo : impl me
	log.Printf("Username: '%s', Password: '%s', Ok: %t", username, password, ok)
	return ok
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !s.checkAuth(username, password, ok) {
		w.Header().Set("WWW-Authenticate", "Basic realm=\"Zzz\"")
		http.Error(w, "Authorization failed", 401)
		return
	}

	names := r.Header["Name-File"]
	name := ""
	if len(names) == 1 {
		name = names[0]
	}
	if "" == name {
		http.Error(w, "Not found header 'NAME-FILE'", 400)
		return
	}

	ext := path.Ext(name)
	ctype := s.ContentMap[ext]
	if ctype == "" {
		http.Error(w, "Invalid header content 'NAME-FILE'", 400)
		return
	}

	uploader := s3manager.NewUploaderWithClient(s.S3Svc)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:        r.Body,
		Bucket:      aws.String(s.Bucket),
		Key:         aws.String(name),
		ContentType: &ctype,
	})
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	str := fmt.Sprintf("File: '%s' successfully uploaded to: %s", name, result.Location)
	log.Println(str)
	w.Write([]byte(str))
}

func main() {
	s := Server{
		S3Svc:      s3.New(session.New(&aws.Config{Region: aws.String("eu-central-1")})),
		Bucket:     "tankzter-storage",
		ContentMap: make(map[string]string),
	}
	s.ContentMap[".png"] = "image/png"
	s.ContentMap[".jpg"] = "image/jpeg"
	s.ContentMap[".jpeg"] = "image/jpeg"
	s.ContentMap[".mp4"] = "video/mp4"
	s.ContentMap[".webm"] = "video/webm"

	http.HandleFunc("/upload", s.handler)
	http.ListenAndServe(":11111", nil)
}
