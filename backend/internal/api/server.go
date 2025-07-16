package api

func RunServer() {
	app := NewApp()
	app.Run(":8080")
}
