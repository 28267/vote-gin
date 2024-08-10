package main

import router "vote-gin/router"

func main() {

	r := router.Router()
	r.Run(":7777")

}
