# The Magic OSS Project

This project is magic.

## Installation

You can install Magic OSS Project via a Debian .deb package.

1. Download the DEB package.

```bash
$ curl
https://github.com/jamtur01/magicossproject/releases/magicossproject.deb
```

2. Install the package.

```bash
$ dpkg -i magicossproject.deb
```

3. Check that it works.

```bash
$ magicoss --version
```

## Usage

You can configure Magic OSS by creating a YAML configuration file.

1. Create a file.

```bash
$ touch magicoss.yml
```

2. Populate it with a basic configuration.

```yaml
---
Version: v1
interface: "localhost"
port: 8080
```

Here we can see the three configuration options.

* `Version` - Specifies the configuration version.
* `interface` - The interface `magicoss` runs on.
* `port` - The port `magicoss` runs on.

3. Then run the `magicoss` binary:

```bash
$ magicoss --config.file=magicoss.yml
```

Then it does magic.

## License

See the `LICENSE` file.

## Copyright

&copy; Copyright 2018 James Turnbull <james@lovedthanlost.net>
