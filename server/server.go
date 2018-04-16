package server

func Start() {

	r := NewRouter()

	r.Run(":6060") // listen and serve on 0.0.0.0:6060
}
