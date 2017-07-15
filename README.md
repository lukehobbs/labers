# Labers

A  Github Enterprise friendly CLI tool for managing Github issue labels.

## Installation

Get binaries for OS X / Linux from the latest [release].

Or use `go get`:

```
go get -u github.com/lukehobbs/labers
```

[release]: https://github.com/lukehobbs/labers/releases

## Usage

First, create a [GitHub token][tokens] with the appropriate scope. Then run `labers configure` to set up the environment file. You will be prompted for a domain name if you wish to interact with a Github Enterprise instance. After running `labers configure`, an environment file `labers.env` will be created in your current working directory.

Where `labers.env` looks like: 
```
GITHUB_TOKEN: 5555535553bbb632ba32e87d349418ca30a9d5e1f
GITHUB_URL: github.h0bbs.com
```

> - The token for public repos needs the `public_repo` scope.
> - The token for private repos needs the `repo` scope.

[tokens]: https://github.com/settings/tokens

### Saving labels

To save existing labels from a repository and into a file:
```
labers cp github://owner/name destination
```

### Copying labels

To copy labels from a repository to another repository:
```
labers cp github://source_owner/source_name github://destination_owner/destination_name
```

For example,
```
labers cp github://lukehobbs/labers github://lukehobbs/newrepo
```

To copy labels from a local file to a repository:
```
labers cp path/to/file.yaml github://destination_owner/destination_name
```
Where `file.yaml` is like:
```yml
Bug: c2929
Help Wanted: 000000
Fix: cccccc
Notes: fbca04
Debugging: bfd4f2
```

## Usage options

```
$ labers
Labers is a CLI that allows you to save, copy, and transfer
Github issue labels between repositories using the Github API and YAML.

Usage:
  labers [command]

Available Commands:
  cp          Copy labels from [source] to [destination]
  help        Help about any command
  init        Configure environment file

Flags:
  -h, --help   help for labers

Use "labers [command] --help" for more information about a command.
```
