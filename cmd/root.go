package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/rinetd/transfer/codec"
	"github.com/rinetd/transfer/version"
	"github.com/urfave/cli"
)

func Run() error {
	app := cli.NewApp()
	app.Name = `Transfer`
	app.Email = "https://github.com/rinetd"
	app.Usage = `Translate YAML, JSON, TOML, HCL, XML, properties ...
	 transfer [-f] [-s input.yaml] [-o output.json] /path/to/input.yaml [/path/to/output.json]`
	app.UsageText = ` Default output format: json
	$ transfer -f data/input.yaml           (output "./data/input.json")
	$ transfer -f data/input.yaml out.json  (output "./out.json")
	$ transfer -f -s data/input.yaml -o hcl (output "./data/input.hcl")
	$ transfer -f -s data/input.yaml -o /root/out.toml (output "/root/out.toml")
	$ transfer -f -o yaml data/input.json   (output "data/input.yaml")`
	app.Version = version.Version
	app.Action = func(c *cli.Context) {
		Parse(c)
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "s, input",
			// Value: "yaml",
			Usage: "input Type: json, yaml, toml, hcl, xml",
		},
		cli.StringFlag{
			Name: "o, output",
			// Value: "json",
			Usage: "output file : json, yaml, toml, hcl, xml",
		},
		cli.BoolFlag{
			Name:  "f,force",
			Usage: "force cover output file",
		},
	}

	return app.Run(os.Args)
}

var conf codec.Transform

func Parse(c *cli.Context) error {
	var in, out string
	f := c.BoolT("f")
	// fmt.Println(f)
	// conf := Transform{}
	in = c.String("s")
	conf.InputType = codec.Typ(in)
	out = c.String("o")
	conf.OutputType = codec.Typ(out)
	// fmt.Println(in, conf.InputType, out, conf.OutputType)
	conf.Setin()
	conf.Setout()

	switch c.NArg() {
	case 0:
		if conf.InputType == codec.FileTypeUnknown {
			cli.ShowAppHelp(c)
			return nil
		}
		// 未指定输出，根据输入确定默认输出
		if conf.OutputType == codec.FileTypeUnknown {
			conf.OutputType = codec.FileTypeJSON
			filename := strings.TrimSuffix(in, path.Ext(in))
			out = filename + "." + conf.OutputType
		}

		// if conf.InputType == codec.FileTypeUnknown {
		// 	conf.InputType = codec.FileTypeJSON
		// 	conf.OutputType = codec.FileTypeYAML
		// 	for {
		// 		conf.PipeLine(os.Stdin, os.Stdout)
		// 	}
		// 	// return errors.New("Unknown Input Type")
		// }
		// // 未指定输出，根据输入确定默认输出
		// if conf.OutputType == codec.FileTypeUnknown {
		// 	conf.OutputType = codec.FileTypeYAML
		// 	filename := strings.TrimSuffix(in, conf.InputType)
		// 	out = filename + conf.OutputType
		// }

	case 1:
		if conf.InputType == codec.FileTypeUnknown {
			// 如果未指定输入 就是输入
			in = c.Args()[0]
			conf.InputType = codec.Typ(in)
			if conf.InputType == codec.FileTypeUnknown {
				return errors.New("Unknown Input Type")
			}

			// 未指定输出，根据输入确定默认输出
			if conf.OutputType == codec.FileTypeUnknown {
				conf.OutputType = codec.FileTypeJSON
				filename := strings.TrimSuffix(in, path.Ext(in))
				out = filename + "." + conf.OutputType
			}

		} else {
			// 那就是输出了
			out = c.Args()[0]
			conf.OutputType = codec.Typ(out)

			if conf.OutputType == codec.FileTypeUnknown {
				return errors.New("Unknown Output Type")
			}
			if conf.OutputType == conf.InputType {
				return errors.New("Unknown Output Type")
			}
		}

	case 2:
		if conf.InputType == codec.FileTypeUnknown {
			in = c.Args()[0]
			conf.InputType = codec.Typ(in)
			if conf.InputType == codec.FileTypeUnknown {
				return errors.New("Unknown Input Type")
			}

		}
		if conf.OutputType == codec.FileTypeUnknown {
			out = c.Args()[1]
			conf.OutputType = codec.Typ(out)
			if conf.OutputType == codec.FileTypeUnknown {
				return errors.New("Unknown Output Type")
			}

		}

	default:
	}
	if out == conf.OutputType {
		filename := strings.TrimSuffix(in, path.Ext(in))
		out = filename + "." + conf.OutputType
	}
	// fmt.Println(in, out)
	src, err := os.Open(in)
	if err != nil {
		return err
	}
	defer src.Close()
	conf.Reader = src

	if !f {
		out = strings.TrimSuffix(out, "."+conf.OutputType) + "-" + strconv.FormatInt(time.Now().Unix(), 10) + "." + conf.OutputType
	}
	// if _, err := os.Stat(filepath.Dir(out)); os.IsNotExist(err) {
	os.MkdirAll(filepath.Dir(out), 0777)
	// }
	dst, err := os.Create(out)
	if err != nil {
		return err
	}
	defer dst.Close()
	conf.Writer = dst

	conf.PipeLine()
	fmt.Println("  input :", in)
	fmt.Println("  output:", out)
	return nil

}

func Action(c *cli.Context) {
	fmt.Println(c.NArg())
	if c.NArg() > 0 {
		fmt.Println(c.Args()[c.NArg()-1])
		for index := 0; index < c.NArg(); index++ {
			fmt.Println(c.Args()[index])
		}

	}
	fmt.Println(conf)
	fmt.Println(c.Args()[c.NArg()-1])
	for index, value := range c.Args() {
		fmt.Println(index, value)
	}
}
