package main

import(
	 "fmt"
	//  "os"
	// "strings"
	"transform"
	"io"
	"net/http"
	"path/filepath"
	"log"

)


func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		html:= `<html><body>
		<form action="/upload" method="post" enctype="multipart/form-data">
		<input type="file" name="image" >
		<button type="submit">Upload Image</button>
		</form>
		</body>
		</html>`
		fmt.Fprint(w,html)
	})
	mux.HandleFunc("/upload",func(w http.ResponseWriter, r *http.Request){
		file,header,err:=r.FormFile("image")
		if err != nil {
			http.Error(w,err.Error(),http.StatusBadRequest)
			return
		}
		defer file.Close()
		ext:=filepath.Ext(header.Filename)[1:] /* this excludes the . in the
		name extension say ".png" */
		_=ext
		out,err:=primitive.Transform(file,50)
		switch ext{
		case "jpg":
			fallthrough
		case "jpeg":
			w.Header().Set("Content-type","image/jpeg")
		case "png":
			w.Header().Set("Content-type","image/png")
		default:
			http.Error(w,fmt.Sprintf("invalid image type %s",ext),http.StatusInternalServerError)
			return
		}
		// w.Header().Set("Content-type","image/png")
		io.Copy(w,out)
		
	})
	log.Fatal(http.ListenAndServe(":3000",mux))
}