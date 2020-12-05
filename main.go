package main

import "github.com/kindai-csg/d-blog-engine/infrastructure"

func main() {
	infrastructure.Router.Run(":2640")
}
