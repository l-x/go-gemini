package gogemini

type Handler interface {
	ServeGemini(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

type funcHandler struct {
	hf HandlerFunc
}

func (h funcHandler) ServeGemini(w ResponseWriter, r *Request) {
	h.hf(w, r)
}
