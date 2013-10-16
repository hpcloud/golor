# golor

golor is a terminal color package for Go, supporting both 16-colors mode and
256 colors mode.

16-bit color printing:

```Go
// See examples/colors.go
fmt.Println(golor.Colorize("Green foreground", golor.G, -1))
fmt.Println(golor.Colorize("Cyan background", -1, golor.CYAN))
```

256-color printing:

```Go
// See examples/colors.go
index := 42  // index can vary from 0 to 255
fmt.Println(golor.Colorize("Foreground", index, -1))
```

Consistently assigning unique colors to strings:

```Go
// Assign an unique color for this string
name := "srid"
fmt.Println(golor.Colorize(process, golor.AssignColor(name), -1))
``

