package service

type Image struct {
	Pool string `uri:"pool"`
	Name string `uri:"name"`
}

func (*Image) getImageInfo() {
	
}
