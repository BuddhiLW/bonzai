package mark_test

import (
	"fmt"
	"text/template"

	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/mark"
)

type Thing struct {
	Name  string
	Count int
}

func (a Thing) Summary() string {
	return fmt.Sprintf("%v %v", a.Name, a.Count)
}

func ExampleFill() {

	/* cannot declare type with method within function, but this is it

	type Thing struct {
		Name  string
		Count int
	}

	func (a Thing) Summary() string {
		return fmt.Sprintf("%v %v", a.Name, a.Count)
	}

	*/

	thing := Thing{`Thing`, 20}
	tmpl := `
	{{hello}}, my name is {{.Name}} with {{.Count}}. Summary: {{.Summary}}`
	funcs := template.FuncMap{}
	funcs[`hello`] = func() string { return `Hello` }

	out, err := mark.Fill(thing, funcs, tmpl)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(out)

	// Output:
	// Hello, my name is Thing with 20. Summary: Thing 20

}

func ExampleUsage_withHiddenCmds() {
	var subFooHiddenCmd = &bonzai.Cmd{
		Name:  `iamhidden`,
		Short: `i am hidden`,
	}

	var subFooCmd = &bonzai.Cmd{
		Name:  `subfoo`,
		Alias: `sf`,
		Short: `under the foo command`,
	}

	var fooCmd = &bonzai.Cmd{
		Name:  `foo`,
		Alias: `f`,
		Short: `foo this command`,
		Cmds:  []*bonzai.Cmd{subFooCmd, subFooHiddenCmd.AsHidden()},
		// Cmds:  []*bonzai.Cmd{subFooCmd, subFooHiddenCmd},
	}

	var barCmd = &bonzai.Cmd{
		Name:  `bar`,
		Alias: `b`,
		Short: `bar this command`,
	}

	var Cmd = &bonzai.Cmd{
		Name:  `mycmd`,
		Alias: `my|cmd`,
		Short: `my command short summary`,
		Cmds:  []*bonzai.Cmd{fooCmd, barCmd},
		Long: `
			Here is a long description.
			On multiple lines.`,
	}

	Cmd.Run()
	out, _ := mark.Bonzai(Cmd)
	fmt.Println(string(out))

	// Output:
	// # Name
	//
	// `mycmd` (`my`|`cmd`) — my command short summary
	//
	// # Usage
	//
	//     my|cmd|mycmd COMMAND
	//
	// # Commands
	//
	//     foo      ← foo this command
	//       subfoo ← under the foo command
	//     bar      ← bar this command
	//
	// # Description
	//
	// Here is a long description.
	// On multiple lines.

}

func ExampleUsage_missingShort() {
	var subFooHiddenCmd = &bonzai.Cmd{
		Name:  `iamhidden`,
		Short: `i am hidden`,
	}

	var subFooCmd = &bonzai.Cmd{
		Name:  `subfoo`,
		Alias: `sf`,
		Short: `under the foo command`,
	}

	var fooCmd = &bonzai.Cmd{
		Name:  `foo`,
		Alias: `f`,
		//Short: `foo this command`,
		Cmds: []*bonzai.Cmd{subFooCmd, subFooHiddenCmd.AsHidden()},
		// Cmds:  []*bonzai.Cmd{subFooCmd, subFooHiddenCmd},
	}

	var barCmd = &bonzai.Cmd{
		Name:  `bar`,
		Alias: `b`,
		Short: `bar this command`,
	}

	var Cmd = &bonzai.Cmd{
		Name:  `mycmd`,
		Alias: `my|cmd`,
		Short: `my command short summary`,
		Cmds:  []*bonzai.Cmd{fooCmd, barCmd},
		Long: `
			Here is a long description.
			On multiple lines.`,
	}

	Cmd.Run()
	out, _ := mark.Bonzai(Cmd)
	fmt.Println(out)

	// Output:
	// # Name
	//
	// `mycmd` (`my`|`cmd`) — my command short summary
	//
	// # Usage
	//
	//     my|cmd|mycmd COMMAND
	//
	// # Commands
	//
	//     foo
	//       subfoo ← under the foo command
	//     bar      ← bar this command
	//
	// # Description
	//
	// Here is a long description.
	// On multiple lines.

}

func ExampleUsage_middle() {
	var subFooHiddenCmd = &bonzai.Cmd{
		Name:  `iamhidden`,
		Short: `i am hidden`,
	}

	var subFooCmd = &bonzai.Cmd{
		Name:  `subfoo`,
		Alias: `sf`,
		Short: `under the foo command`,
	}

	var fooCmd = &bonzai.Cmd{
		Name:  `foo`,
		Alias: `f`,
		//Short: `foo this command`,
		Cmds: []*bonzai.Cmd{subFooCmd, subFooHiddenCmd.AsHidden()},
		// Cmds:  []*bonzai.Cmd{subFooCmd, subFooHiddenCmd},
	}

	var barCmd = &bonzai.Cmd{
		Name:  `bar`,
		Alias: `b`,
		Short: `bar this command`,
	}

	var Cmd = &bonzai.Cmd{
		Name:  `mycmd`,
		Alias: `my|cmd`,
		Short: `my command short summary`,
		Cmds:  []*bonzai.Cmd{fooCmd, barCmd},
		Long: `
			Here is a long description.
			On multiple lines.`,
	}

	_ = Cmd

	Cmd.Run()
	out, err := mark.Bonzai(fooCmd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)

	// Output:
	// # Name
	//
	// `foo` (`f`)
	//
	// # Usage
	//
	//     f|foo COMMAND
	//
	// # Commands
	//
	//     subfoo ← under the foo command

}

func ExampleUsage_longFirstName() {

	var fooCmd = &bonzai.Cmd{
		Name: `foo`,
		//Short: `a foo`,
		Do: func(*bonzai.Cmd, ...string) error {
			return nil
		},
	}

	var Cmd = &bonzai.Cmd{
		Name:  `help-test`,
		Alias: `h|ht`,
		Short: `just a help test`,
		Opts:  `some|-y|--yaml`,
		Cmds:  []*bonzai.Cmd{fooCmd, fooCmd.WithName(`foo2`)},
		Def:   fooCmd,
	}

	Cmd.Run(`foo`) // for default
	out, _ := mark.Bonzai(Cmd)
	fmt.Println(string(out))

	// Output:
	// # Name
	//
	// `help-test` (`h`|`ht`) — just a help test
	//
	// # Usage
	//
	//     h|ht|help-test COMMAND|some|-y|--yaml
	//
	// # Commands
	//
	//     foo ← (default)
	//     foo2
	//
}
