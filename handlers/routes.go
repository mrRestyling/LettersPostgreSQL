package handlers

func (h Handlers) SetRoutes() {

	h.MuxECHO.GET("/", Hello)
	h.MuxECHO.POST("/name", h.Name)

	h.MuxECHO.GET("/message", h.MessageReturn)
	h.MuxECHO.POST("/message", h.Message)

	h.MuxECHO.GET("/letter", h.LetterReturn)
	h.MuxECHO.POST("/letter", h.Letter)
}
