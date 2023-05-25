package controller

func CoordinateController() *Controller {
	cnt := Controller{}
	cnt.Path = "/coordinates"
	return &cnt
}
