package help

import (
	"fmt"

	"github.com/charmbracelet/glamour"
	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/mark"
)

var Cmd = &bonzai.Cmd{
	Name:  `help`,
	Alias: `h|-h|--help|--h|/?`,
	Vers:  `v0.8.0`,
	Short: `display command help`,
	Long: `
		The {{code .Name}} command displays the help information for the
		immediate previous command unless it is passed arguments, in which
		case it resolves the arguments as if they were passed to the
		previous command and the help for the leaf command is displayed
		instead.`,

	Do: func(x *bonzai.Cmd, args ...string) (err error) {

		if len(args) > 0 {
			x, args, err = x.Caller().SeekInit(args...)
		} else {
			x = x.Caller()
		}

		md, err := mark.Bonzai(x)
		if err != nil {
			return err
		}

		renderer, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			//glamour.WithWordWrap(0),
			glamour.WithPreservedNewLines(),

			glamour.WithStylesFromJSONBytes([]byte(`{
				"document":{
					"margin":0,
					"block_prefix":"",
					"block_suffix":""
				},
				"heading": {
					"block_prefix":"",
					"block_suffix":""
				},
				"h1":{
					"background_color": "",
					"color": "11",
					"prefix": "",
					"suffix": "",
					"block_prefix": "",
					"block_suffix": ""
				},
				"h2":{
					"background_color": "",
					"color": "5"
				},
				"paragraph": {
					"margin": 6,
					"block_prefix":""
				},
				"list": {
					"margin": 6
				},
				"code_block": {
					"margin": 6
				},
				"blockquote": {
					"margin": 6
				},
				"code": {
					"color": "11",
					"background_color": "",
					"prefix":"",
					"suffix":""
				}

			}`)),
		)
		if err != nil {
			return fmt.Errorf("developer-error: %v", err)
		}

		rendered, err := renderer.Render(md)
		if err != nil {
			return fmt.Errorf("developer-error: %v", err)
		}

		fmt.Println("\u001b[2J\u001b[H" + rendered)

		return nil
	},
}
