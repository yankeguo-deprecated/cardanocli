# cardanocli

a Golang wrapper for executing cardano-cli commands

I know it's stupid to invoke `cardano-cli` other than communicate with `cardano-node` via `node.socket` directly, but
this is maybe the best solution for Golang right now.

## Usage

```golang
package main

import "go.guoyk.net/cardanocli"

func main() {
    cli := cardanocli.New()
    cli.SocketPath = "/path/to/cardano-node/node.socket"

    // example: get policy id from policy script
    (
        buf := &bytes.Buffer{}
        // cardano-cli transaction policyid --script-file /path/to/policy.script
        x := cli.Cmd().Transaction().Policyid().OptScriptFile("/path/to/policy.script").Exec()   
        x.Stdout = buf
        err := x.Run()
        if err != nil {
            panic(err)
        }
        policyID := strings.TrimSpace(buf.String())
    )
}
```

## Credits

Guo Y.K.， MIT License
