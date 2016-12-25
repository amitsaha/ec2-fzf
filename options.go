package ec2fzf

import (
	"flag"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
)

type Options struct {
	Region       string
	UsePrivateIp bool
	Template     string
}

func ParseOptions() (Options, error) {
	options := Options{
		Region:       "us-east-1",
		UsePrivateIp: false,
		Template:     `{{index .Tags "Name"}}`,
	}

	path, err := homedir.Expand("~/.config/ec2-fzf")
	if err != nil {
		return Options{}, err
	}
	toml.DecodeFile(path, &options)

	region := flag.String("region", options.Region, "The AWS region")
	usePrivateIp := flag.Bool("private", options.UsePrivateIp, "return the private IP address of the instance rather than the public dns")
	template := flag.String("template", options.Template, "Template")

	flag.Parse()

	return Options{
		Region:       *region,
		UsePrivateIp: *usePrivateIp,
		Template:     *template,
	}, nil
}