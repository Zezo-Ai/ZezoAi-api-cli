# Hostinger API Command Line Interface

A powerful command-line interface for managing Hostinger services through the Hostinger API.

For more information, please visit https://developers.hostinger.com.

# Installation

## Downloading a Release from GitHub
Download the latest binary from the [releases page](https://github.com/hostinger/api-cli/releases).

### Linux

Download, extract and move the binary:
```bash
cd ~
wget https://github.com/hostinger/api-cli/releases/download/<version>/hapi-<version>-linux-<arch>.tar.gz
tar xf ~/hapi-<version>-linux-<arch>.tar.gz
mv hapi /usr/local/bin
```

### macOS
WIP

### Windows
WIP

# Configuration

## File
Full example and possible values of the configuration file [can be found here](https://github.com/hostinger/api-cli/blob/main/hapi.yaml)

By default, configuration file will be read from `$HOME/.hapi.yaml`.

You can always define which configuration file to use by using `--config` argument, eg.: 
```bash
hapi --config path/to/your/config.yaml ...
```

## Environment variables
Instead of providing configuration file, you can provide configuration using environment variables.
Each config file property must have `HAPI_` prefix and key name in ALL CAPS, for example to provide
`api_token` parameter, you would export variable like this:

```bash
export HAPI_API_TOKEN=<your token> 
hapi vps vm list 
```

# Usage

See the [full reference documentation](blob/main/docs/hapi.md) for information about each available command.

# Enabling shell auto-completion (optional)

`hapi` has auto-complete support. This makes it easier to use the CLI and improves user experience by completing command
names by clicking TAB key. For example if you type `hapi vps dat<TAB>` with auto-completion enabled, shell will automatically append
rest of the command: `hapi vps data-centers`.

Auto-completion can be generated for multiple shells. The currently supported shells are:
- Bash
- Zsh
- fish
- PowerShell

After adding shell auto-completion, remember to refresh your shell profile by logging out from the shell and log back in.

Read more on [how to enable auto-completion](blob/main/AUTOCOMPLETE.md).