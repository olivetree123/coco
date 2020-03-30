# Coco
[![Build Status](https://travis-ci.org/olivetree123/Coco.svg?branch=master)](https://travis-ci.org/olivetree123/Coco)  [![codecov](https://codecov.io/gh/olivetree123/Coco/branch/master/graph/badge.svg)](https://codecov.io/gh/olivetree123/Coco)  [![Codacy Badge](https://api.codacy.com/project/badge/Grade/14854a8ba1124a238ffecdcce7465b78)](https://www.codacy.com/manual/olivetree123/Coco?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=olivetree123/Coco&amp;utm_campaign=Badge_Grade)  ![GitHub](https://img.shields.io/github/license/olivetree123/coco)  
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

# Params
``` javascript
// GET
name := c.Params.ByName("name")

// POST JSON
var param DBCreateParam
decoder := json.NewDecoder(c.Request.Body)
err := decoder.Decode(&param)
```