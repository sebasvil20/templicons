# Templicons: Tabler icons made templ components
This is a collection of Tabler icons made as templ components for use in your projects.

> [!NOTE]
> I'm not associated in any form with Tabler icons, this is just a collection of their icons made as templ components.

## Installation
```bash
go get github.com/sebasvil20/templicons
```

## Usage
```go
package yourview

import (
    "github.com/sebasvil20/templicons/i"
    "github.com/sebasvil20/templicons/tabler"
)

templ YourView() {
    // with default props
    @icons.Menu2()
    
    // with custom props
    @icons.Menu2(i.Props{
        Size: icons.XS,
        Class: "text-red-500",
        StrokeWidth: "2",
        ID: "menu-icon",
        AriaLabel: "menu-icon",
    }
    )
}
```

### Props
Default props:
```go
Size:        MD, // XXS = "8", XS = "16", SM = "20", MD = "24", LG = "32", XL = "48", XXL = "64"
StrokeWidth: Regular, // Thin = "1", Regular = "1.5", Medium = "2", Bold = "2.5", Heavy = "3"
Class:       "text-current",
AriaLabel:   "icon",
```

The constants are just strings, so you can use any string you want. The default values are just for convenience.

## Credits
I have not created any of the icons, I just made them as templ components. All credits go to the original creators of the icons.
- [Tabler Icons](https://tabler.io/icons)