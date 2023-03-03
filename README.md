# env-util
A simple golang env util. 

## Installation
`go get github.com/natekinne97/env-util`

## Usage
`
import (
  "github.com/natekinne97/env-util"
  "os"
)

func somefunc(){
   envVar := envUtil.GetEnv("", os)
}

`
