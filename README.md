# cardanocli

a Golang wrapper for executing cardano-cli commands

## Disclaimer

`cardano-cli` command line arguments and output formats are subjected to change, always use fixed version of `cardano-cli`.

USE THIS LIBRARY AT YOUR OWN RISK.

## Usage

```golang
package main

import (
	"go.guoyk.net/cardanocli"
)

func main() {
    cli := cardanocli.New()
    cli.SocketPath = "/path/to/cardano-node/node.socket"

    // example: get policy id from policy script
    {
    	var policyID string
        // cardano-cli transaction policyid --script-file /path/to/policy.script
        err := cli.Cmd().Transaction().
        	Policyid().
        	OptScriptFile("/path/to/policy.script").
        	Run(cardanocli.CollectStdout(&policyID))
        if err != nil {
            panic(err)
        }
    }
}
```

## Credits

Guo Y.K.， MIT License
