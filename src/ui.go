package main


func(app *Config) makeUI {
	// get the current weather of Kazan



//get toolbar
toolBar := app.getToolBar(app.MainWindow)
app.ToolBar = toolBar


go func() {
	for range time.Tick(time.Second * 30) {

		app.refreshTemperatureContent()
	}
}()