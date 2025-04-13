# Templicons: Tabler icons made templ components
This is a collection of Tabler icons made as templ components for use in your projects.

> [!NOTE]
> I'm not associated in any form with Tabler icons, this is just a collection of their icons made as templ components.

## Pre-requisites
- [Go](https://golang.org/) 1.24 or later
- [Templ](https://templ.guide/) 0.3.857 or later

## Installation
```bash
go get github.com/sebasvil20/templicons
```

## Features
- Lightweight: Each icon is an inline SVG component, so you can use them without any external dependencies. Don't worry about this pkg weight, Go compiler will remove all the unused icons from your binary.
- Customizable: You can easily customize the size, color, and other properties of the icons using `i.Props` structure.
- Easy to use: The icons are easy to use and integrate into your projects.
- Over 5880 icons: The library includes over 5880 icons from Tabler, so you can find the perfect icon for your project.

## Usage
```go
package yourview

import (
    "github.com/sebasvil20/templicons/i"
    "github.com/sebasvil20/templicons/tabler"
)

templ YourView() {
    // with default props
    @tabler.Menu2()
    
    // with custom props
    @tabler.Menu2(i.Props{
        Size: "128",
        Class: "text-red-500",
        StrokeWidth: "2",
        ID: "menu-icon",
        AriaLabel: "menu-icon",
    })
    
    // You can also use the color constants or the size constans
    @tabler.Menu2(i.Props{
        Size: i.SizeXS,
        Class: i.StrokeWarning,
        StrokeWidth: i.StrokeBold,
        ID: "menu-icon",
        AriaLabel: "menu-icon",
    })
}
```

### Props
Default props:
```go
Size:        "24",
StrokeWidth: "1.5",
Class:       "text-current",
AriaLabel:   "icon",
```
- `Size`: The size of the icon in pixels. You can use the size constants defined in the library or just use the size you want in pixels.
- `StrokeWidth`: The stroke width of the icon in pixels. You can use the stroke width constants defined in the library or just use the size you want in pixels.
- `Class`: The class/es of the icon. You can use the color constants defined in the library or just use a list of classes you want.
- `ID`: The id of the icon. This is useful if you want to use the icon in a form or if you want to use it with JavaScript.
- `AriaLabel`: The aria-label of the icon. This is useful for accessibility purposes.

### Color constants
Color constants are utilities to use the Tailwind CSS colors in the icons. You can use them in the `Class` prop.
In order to use it, you need to have defined the colors in your Tailwind CSS config file. If you use tailwind css component libraries like
[Daisy UI](https://daisyui.com/) you have it covered.

This is a list of the colors you can use, separated by filled and stroked:
```go
    StrokePrimary         = "text-primary"
    StrokePrimaryContent  = "text-primary-content"
    StrokeSecondary       = "text-secondary"
    StrokeSecondaryContent = "text-secondary-content"
    StrokeAccent          = "text-accent"
    StrokeAccentContent   = "text-accent-content"
    StrokeNeutral         = "text-neutral"
    StrokeNeutralContent  = "text-neutral-content"
    StrokeBase            = "text-base-content"
    StrokeInfo            = "text-info"
    StrokeInfoContent     = "text-info-content"
    StrokeSuccess         = "text-success"
    StrokeSuccessContent  = "text-success-content"
    StrokeWarning         = "text-warning"
    StrokeWarningContent  = "text-warning-content"
    StrokeError           = "text-error"
    StrokeErrorContent    = "text-error-content"
    StrokeCurrent         = "text-current"
    FillPrimary         = "fill-primary"
    FillPrimaryContent  = "fill-primary-content"
    FillSecondary       = "fill-secondary"
    FillSecondaryContent = "fill-secondary-content"
    FillAccent          = "fill-accent"
    FillAccentContent   = "fill-accent-content"
    FillNeutral         = "fill-neutral"
    FillNeutralContent  = "fill-neutral-content"
    FillBase            = "fill-base-content"
    FillInfo            = "fill-info"
    FillInfoContent     = "fill-info-content"
    FillSuccess         = "fill-success"
    FillSuccessContent  = "fill-success-content"
    FillWarning         = "fill-warning"
    FillWarningContent  = "fill-warning-content"
    FillError           = "fill-error"
    FillErrorContent    = "fill-error-content"
    FillCurrent         = "fill-current"
```

In order to use this constants, you need to have Tailwind CSS configured in your project and add this to the `input.css` (Tailwind V4) or equivalent in
your tailwind version:
```js
@source "${GOPATH}/pkg/mod/github.com/sebasvil20/templicons@*/i/*.{go,templ}";
```
Replace ` ${GOPATH}` with your Go path. This will include all the icons in the Tailwind CSS build process.
```
go env GOPATH
```

> [!NOTE]
> Constants are just for convenience, you can use any color you want in the `Class` prop. 

You can also use all Tailwind CSS classes you want in the `Class` prop.

### Size constants
Defined in pixels, constant utilities for you to keep consistency in your project.
```go
    SizeXXS = "8"
    SizeXS  = "16"
    SizeSM  = "20"
    SizeMD  = "24"
    SizeLG  = "32"
    SizeXL  = "48"
    SizeXXL = "64"
```
You can use them in the `Size` prop or just use the size you want in pixels.

### Stroke width constants
Defined in pixels, constant utilities for you to keep consistency in your project. This is only valid for outlined icons.
```go
    StrokeThin    = "1"
    StrokeRegular = "1.5"
    StrokeMedium  = "2"
    StrokeBold    = "2.5"
    StrokeHeavy   = "3"
```
You can use them in the `StrokeWidth` prop or just use the size you want in pixels.

## Credits
I have not created any of the icons, I just made them as templ components. All credits go to the original creators of the icons.
- [Tabler Icons](https://tabler.io/icons)