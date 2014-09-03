package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "net/http"
)

func PanicIf(err error){
  if err != nil{
    panic(err)
  }
}

func main() {
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
    Layout: "layout",
    Extensions: []string{".tmpl", ".html"},
  }))
  m.Get("/", HomePage)
  m.Run()
}

func HomePage(ren render.Render, r *http.Request,){
  ren.HTML(200, "index", nil)
}