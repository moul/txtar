# `txtar`

Generate txtar from CLI.

Using https://golang.org/x/tools/txtar.

## Install

    go install moul.io/txtar

## Example

    $ txtar .
    -- README.md --
    # `txtar`

    Generate txtar from CLI.
    
    Using https://golang.org/x/tools/txtar.

    ## Install

        go install moul.io/txtar

    ## Example

        $ txtar .
        [redacted]
    -- go.mod --
    module moul.io/txtar

    go 1.20

    require golang.org/x/tools v0.12.0
    -- go.sum --
    golang.org/x/tools v0.12.0 h1:YW6HUoUmYBpwSgyaGaZq1fHjrBjX1rlpZ54T6mu2kss=
    golang.org/x/tools v0.12.0/go.mod h1:Sc0INKfu04TlqNoRA1hgpFZbhYXHPr4V5DzpSBTPqQM=
    -- main.go --
    package main

    import (
    	"flag"
    	"fmt"
    	"io/ioutil"
    	"os"
    	"path/filepath"
    	"strings"

    	"golang.org/x/tools/txtar"
    )

    func main() {
    	[redacted]
    }
