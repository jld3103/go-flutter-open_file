# open_file

This Go package implements the host-side of the Flutter [open_file](https://github.com/crazecoder/open_file) plugin.

## Usage

Import as:

```go
import open_file "github.com/jld3103/open_file"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&open_file.OpenFilePlugin{}),
```
