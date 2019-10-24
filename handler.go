package coco

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"mime"
	"net/http"
	"strings"
)

//Handler 所有 handler 都会被转换成该类型
type Handler func(coco *Coco) Result

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	coco := NewCoco()
	coco.Request = r
	coco.ResponseWriter = w
	coco.Params = httprouter.ParamsFromContext(r.Context())
	resp := h(coco)
	var response []byte
	if resp.Type == "binary" {
		response = resp.BinaryData
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", resp.FileName))
		w.Header().Set("Content-Type", MimeTypeByName(resp.FileName))
	} else {
		response, err = json.Marshal(resp)
		if err != nil {
			panic(err)
		}
	}
	w.Write(response)
}

func MimeTypeByName(name string) string {
	names := strings.Split(name, ".")
	ext := "." + names[len(names)-1]
	return mime.TypeByExtension(ext)
}
