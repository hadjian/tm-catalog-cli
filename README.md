# Thing Model Catalog CLI

[![Go Report Card](https://goreportcard.com/badge/github.com/web-of-things-open-source/tm-catalog-cli)](https://goreportcard.com/report/github.com/web-of-thing-open-source/tm-catalog-cli) [![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/hadjian/tm-catalog-cli)](https://github.com/web-of-things-open-source/tm-catalog-cli/releases)[![](https://img.shields.io/github/actions/workflow/status/web-of-things-open-source/tm-catalog-cli/test.yml?branch=main&longCache=true&label=Test&logo=github%20actions&logoColor=fff)](https://github.com/web-of-things-open-source/tm-catalog-cli/actions?query=workflow%3ATest)

![Thing Model Catalog Logo](https://github.com/hadjian/tm-catalog-cli/raw/main/docs/media/tm-catalog-logo.svg)


---
Find, use and contribute device descriptions for industrial IoT devices!

⚠ This software is **experimental** and may not be fit for any purposes. 

The Thing Model Catalog Command Line Client, or ```tm-catalog-cli``` for short, is a tool for browsing, consuming, contributing and serving Thing Models.

Thing Models are simple device descriptions specified in the [W3C Thing Description][1] standard. Thing Models are to Thing Descriptions what classes are to instances in programming languages.

Thing Models let you describe industrial devices using a simple standardized JSON-based format, which is independent of the communication protocol. This enables a uniform access layer to the fragmented industrial protocol landscape we encounter today.

Thing Descriptions are to Modbus, BACnet, MQTT, DNP3 ... what HTML is to HTTP.

---

## Installation

1. Download the latest [release][2] for your operating system and architecture
2. Optionally rename to ```tm-catalog-cli``` to remove os/arch postfixes
3. Give it execution rights and move to a folder that is in your ```PATH```

## Quick Start

The ```tm-catalog-cli``` helps you to interact with a Thing Model catalog, which may be hosted on any git forge like github or create your own catalog in a git repository of your choosing. 

To enable a culture of sharing, we provide a canonical repository at [], but feel free to create your own open or private catalog as well.

To integrate publicly available and your own private Thing Models into your product, the ```tm-catalog-cli``` can be run as a server, exposing a REST API that can be protected with JWT tokens.

### Configure Autocompletion

1. Read the help of the ```completion``` command to find out which shells are supported
```bash
tm-catalog-cli completion -h
```

2. Follow the instructions of the shell specific help text
```bash
tm-catalog-cli completion <shell> -h
```

### Browse the canoncial Catalog

1. Configure the canonical repository
```bash
tm-catalog-cli remote add --type http thingmodels 'https://raw.githubusercontent.com/wot-oss/thingmodels'
```
2. List the contents of the canonical catalog
```bash
tm-catalog-cli list
```

The listed names are formatted as follows

```
<author>/<manufacturer>/<model>
```

You can specify a part of that path after the ```list``` command to filter the list for only parts of the list tree (use tab to auto-complete path parts):

```
tm-catalog-cli list nexus-x/siemens
```

### List Versions

Every model entry in the list may contain multiple versions, reflecting the evolution of the Thing Model (bugfixes, additions, changes in the device itself ...). List the available versions with the ```versions``` command:

```bash
tm-catalog-cli versions <name>
```

### Fetch a Thing Model

Like what you see? Fetch and store locally using the ```fetch``` command. It will print the Thing Model to stdout to enable unix-like piping:

```bash
tm-catalog-cli fetch <NAME>
```

If you just specify the name, the cli will fetch the latest version automatically. If you want to fetch a specific version, append the version string to the name, separated by a colon:

```bash
tm-catalog-cli fetch <NAME>:<SEMVER>
```

To store the Thing Model locally instead of printing to stdout, specify the ```-o``` flag and point it to a directory:

```bash
tm-catalog-cli fetch <NAME> -o .
```


[1]: https://www.w3.org/TR/wot-thing-description11/
[2]: https://github.com/web-of-things-open-source/tm-catalog-cli/releases
