package main

import (
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
  "github.com/plutov/packagemain/16-wails-desktop-app/cpustats/pkg/sys"
)
/*
func basic() string {
  return "Hello World!"
}*/



func main() {

  js := mewn.String("./frontend/dist/app.js")
  css := mewn.String("./frontend/dist/app.css")

  stats := &sys.Stats{}

  app := wails.CreateApp(&wails.AppConfig{
    Width:  824,
    Height: 568,
    Title:  "my-first-wails-app",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(stats)
  app.Run()
}
