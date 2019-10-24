# Coco
Coco is an http web framework for golang.

# Install
``` shell
go get -u github.com/olivetree123/coco
```

# Tutorial
``` javascript
import (
    "github.com/olivetree123/coco"
)

func HelloHandler(c *coco.Coco) coco.Result {
    return coco.APIResponse("Hello Coco !")
}

func main() {
    c := coco.NewCoco()
    c.AddRouter("GET", "/", HelloHandler)
    c.Run()
}
``` 