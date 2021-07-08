# JWT Tool

This tool generates JWT Tokens and copies the result to the clipboard.

## Installation

- Download the .zip from the latest release
- Extract the .exe from the .zip
- Put the .exe somewhere and add it to your PATH

## Configuration

The tool will look for a file in your $HOME directory called `.jwt-tool.yaml`. In my case, this is `C:\Users\rob.bailey\.jwt-tool.yaml`. 

In this configuration file, currently only a single entry is needed (`secretKey`).

Mine looks like:
```
secretKey: superSecretKey
```

## Usage
From a terminal/command line, enter `jwt-tool`. This should run the tool. 

```
Usage:
  jwt-tool [flags]

Flags:
      --config string      config file (default is $HOME/.jwt-tool.yaml)
  -h, --help               help for jwt-tool
      --secretKey string   The JWT Secret to sign the token with. If blank, a value will be found in $HOME/.jwt-tool.yaml (default "secret")
      --userId string      The user id of the user you want to login (A random string is used by default)

```