# The Magic OSS Project

This project is magc.

## Installation

You can install Magic OSS Project

1. Install the package.

```bash
$ dpkg -i magicossproject.deb
```

2. Check that it works.

```bash
$ magicoss --version
```

## Usuage

You can configuration Magic OSS by creating a YAML configuration file.

--- 
Version: v1
interface: 
  "localhost"
port:
  8080
      

And run the `magicoss` binary:

```bash
$ magicoss --config.file=magicoss.yml
```

Then it does magic.

## License

Apache 2.0.

## Copyright

&copy; Copyright 2018 James Turnbull <james@lovedthanlost.net>
