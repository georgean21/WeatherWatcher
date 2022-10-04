package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
)

type Config struct {
	App        fyne.App
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	MainWindow fyne.Window
}

var myApp Config

func main() {
	// create a fyne application

	fyneApp := app.NewWithID("georgean.weatherwatcher.preferences")

	myApp.App = fyneApp

	/*Create our loggers.That logger writes to standard error
	and prints the date and time of each logged message*/
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// open a connection to the db

	// create a db repo

	//create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("WeatherWatcher")
	myApp.MainWindow.Resize(fyne.NewSize(150, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	//myApp.makeUI()

	//show & run the app
	myApp.MainWindow.ShowAndRun()
}
