package service

import "log"

type Image struct {
	Pool string `uri:"pool" json:"pool"`
	Name string `uri:"name" json:"name"`
	Size int8 `uri:"size" json:"size"`
}

func (image *Image) Create() {
	log.Fatalf("创建pool:%s, name:%s, size:%s", image.Pool, image.Name, image.Size)
}