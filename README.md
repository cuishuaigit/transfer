# transfer [![GoDoc](https://godoc.org/github.com/rinetd/transfer?status.png)](https://godoc.org/github.com/rinetd/transfer)[![Build Status](https://travis-ci.org/rinetd/transfer.svg?branch=master)](https://travis-ci.org/rinetd/transfer)

[中文文档](README.zh.md)

Converts from one encoding to another. 

Supported formats HCL ⇄ JSON ⇄ YAML⇄TOML⇄XML⇄plist⇄pickle⇄properties ... 
### install

```
$ go get github.com/rinetd/transfer
```
### Download

```
https://github.com/rinetd/transfer/releases
```

### usage

```
usage:

	transfer [-f] [-s input.yaml] [-o output.json] /path/to/input.yaml [/path/to/output.json]

Converts from one encoding to another. Supported formats (and their file extensions):

	- JSON (.json)
	- TOML (.toml)
	- YAML (.yaml or .yml)
	- HCL (.hcl or .tf)
	- XML (.xml)
	- MSGPACK (.msgpack)
	- PLIST (.plist)
	- BSON (.bson)
	- PICKLE (.pickle)
	- PROPERTIES (.prop or .props or .properties)

```

### docker usage

```
# build the transfer image
docker build -o rientd/transfer .
```


### examples

Convert data/input.yml TO data/input.json
$ transfer -f data/input.yaml           (output "./data/input.json")

```
$ transfer -f data/input.yaml out.json  (output "./out.json")
$ transfer -f -s data/input.yaml -o /root/out.toml (output "/root/out.toml")
$ transfer -f -s data/input.yaml -o hcl (output "./data/input.hcl")
$ transfer -f -o yaml data/input.json   (output "data/input.yaml")
```
```yaml
Author:
  email: rinetd@163.com
  github: rinetd
menu:
  main:
  - Identifier: categories
    Name: categories
    Pre: <i class='fa fa-category'></i>
    URL: /categories/
    Weight: -102
  - Identifier: tags
    Name: tags
    Pre: <i class='fa fa-oags'></i>
    URL: /tags/
    Weight: -101
theme: hueman

```

```json
{
	"Author": {
		"email": "rinetd@163.com",
		"github": "rinetd"
	},
	"menu": {
		"main": [
			{
				"Identifier": "categories",
				"Name": "categories",
				"Pre": "<i class='fa fa-category'></i>",
				"URL": "/categories/",
				"Weight": -102
			},
			{
				"Identifier": "tags",
				"Name": "tags",
				"Pre": "<i class='fa fa-oags'></i>",
				"URL": "/tags/",
				"Weight": -101
			}
		]
	},
	"theme": "hueman"
}
```

```hcl
$ transfer main.json main.hcl
```
