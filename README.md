<p align="center">
	<img src="https://file.anotherhadi.com/wtui-components/banner.png"
</p>

# What's this UI? - Components

<p>
    <a href="https://github.com/anotherhadi/wtui-components/releases"><img src="https://img.shields.io/github/release/anotherhadi/wtui-components.svg" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/anotherhadi/wtui-components?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://goreportcard.com/report/github.com/anotherhadi/wtui-components"><img src="https://goreportcard.com/badge/github.com/anotherhadi/wtui-components" alt="GoReportCard"></a>
</p>

The `wtui-components` package is a Go library that offers a set of components for developing Text-based User Interfaces (TUI) with [wtui](https://github.com/anotherhadi/wtui) or in standalone mode.
It includes features like input fields, selections, and other essential components for building interactive text-based applications.
Easy to set up, use, customize and modify. Some components are inspired by [charmbracelet/bubbles](https://github.com/charmbracelet/bubbles).

## Contents

- [What's this UI? - Components](#whats-this-ui---components)
  - [Contents](#contents)
  - [Installation](#installation)
  - [Components](#components)
    - [ANSI](#ansi)
    - [Asciimoji](#asciimoji)
    - [Asciitext](#asciitext)
    - [Breadcrumb](#breadcrumb)
    - [Checkbox](#checkbox)
    - [Combobox](#combobox)
    - [Confirm](#confirm)
    - [Floating Window](#floating-window)
    - [Get Char](#get-char)
    - [Get Size](#get-size)
    - [Image](#image)
    - [Input](#input)
    - [List](#list)
    - [Notification](#notification)
    - [Number Picker](#number-picker)
    - [Paragraph](#paragraph)
    - [RGBA Picker](#rgba-picker)
    - [RGB Picker](#rgb-picker)
    - [Selection](#selection)
    - [Shortcuts](#shortcuts)
    - [Spacer](#spacer)
    - [Table](#table)
  - [Options](#options)
  - [License](#license)

## Installation

```bash
go get https://github.com/anotherhadi/wtui-components@latest
```

## Components

Almost all components accept the wtopts.Opts struct (from `wtui-components/wtops`) which allows you to modify various parameters: Colors, Default values, Maximum and Minimum, Max rows & columns, etc.
You can refer to the examples folder for each component to learn more about the component and/or opts.

### ANSI

![ANSI Example](https://file.anotherhadi.com/wtui-components/ansi.gif)

The `ansi` package provides a simple and user-friendly way to add color and formatting to terminal output in Go applications.
Enhance your command-line interfaces with vibrant text, background colors, and text styles, making your output more readable and visually appealing.

You can found an example [here](https://github.com/anotherhadi/wtui-components/blob/main/ansi/example/main.go)

### Asciimoji

![Asciimoji Example](https://file.anotherhadi.com/wtui-components/asciimoji.gif)

The `asciimoji` package is a collection of emojis made from ASCII characters. From [asciimoji.com](https://asciimoji.com)

You can found an example [here](https://github.com/anotherhadi/wtui-components/blob/main/asciimoji/example/main.go)

### Asciitext

![Asciitext Example](https://file.anotherhadi.com/wtui-components/asciitext.gif)

The `asciitext` component is used to print large ASCII art text.

You can found an example with and without options [here](https://github.com/anotherhadi/wtui-components/blob/main/asciitext/example/main.go)

### Breadcrumb

![Breadcrumb Example](https://file.anotherhadi.com/wtui-components/breadcrumb.gif)

The `breadcrumb` component allows you to display breadcrumbs through a list of strings.

You can found an example [here](https://github.com/anotherhadi/wtui-components/blob/main/breadcrumb/example/main.go)

### Checkbox

![Checkbox Example](https://file.anotherhadi.com/wtui-components/checkbox.gif)

The `checkbox` component allows you to quicky prompt the user to choose multiple options.
Move with the arrow keys or JK, select an option with SPACE.

You can found an example with and without options [here](https://github.com/anotherhadi/wtui-components/blob/main/checkbox/example/main.go)

### Combobox

![Combobox Example](https://file.anotherhadi.com/wtui-components/combobox.gif)

The `combobox` component allows you to quickly prompt the user to choose an option. Move with the arrow keys or JK, select an option with CR, and change the filter with other keys.

You can found an example with and without options [here](https://github.com/anotherhadi/wtui-components/blob/main/combobox/example/main.go)

### Confirm

![Confirm Example](https://file.anotherhadi.com/wtui-components/confirm.gif)

The `confirm` component allows you to quickly prompt the user to choose between Yes/No, True/False, etc..
Move with the arrow keys or HJKL/YN, confirm selection with CR

You can found an example with and without options [here](https://github.com/anotherhadi/wtui-components/blob/main/confirm/example/main.go)

### Floating Window

![Floating Window Example](https://file.anotherhadi.com/wtui-components/floatingwindow.gif)

The `floatingwindow` component allow you to print boxes with different styles.

You can found an example with and without options [here](https://github.com/anotherhadi/wtui-components/blob/main/floatingwindow/example/main.go)

### Get Char

![Get Char Example](https://file.anotherhadi.com/wtui-components/getchar.gif)

The `getchar` package enables you to capture a single character or key input without the need for the user to press enter.

You can found an example [here](https://github.com/anotherhadi/wtui-components/blob/main/getchar/example/main.go)

### Get Size

![Get Size Example](https://file.anotherhadi.com/wtui-components/getsize.gif)

The `getsize` package allows you to obtain the size in columns and rows of the terminal.

You can found an example [here](https://github.com/anotherhadi/wtui-components/blob/main/getsize/example/main.go)

### Image

![Image Example](https://file.anotherhadi.com/wtui-components/image.gif)

The `image` component allows you to print some low quality image in the terminal

You can found an example [here](https://github.com/anotherhadi/wtui-components/blob/main/image/example/main.go)

### Input

![Input Example](https://file.anotherhadi.com/wtui-components/input.gif)

The `input` component allows you to prompt the user to type a string.

You can found an example with and without options [here](https://github.com/anotherhadi/wtui-components/blob/main/input/example/main.go)

### List

![List Example](https://file.anotherhadi.com/wtui-components/list.gif)

The `list` component allows you to quickly prompt the user to choose an option (With Title/Description). Move with the arrow keys or HL, select an option with CR.

You can found an example with and without options [here](https://github.com/anotherhadi/wtui-components/blob/main/list/example/main.go)

### Notification

![Notification Example](https://file.anotherhadi.com/wtui-components/notification.gif)

The `notification` component allows you to print some nice notifications.
You can change the notification's status through the "status" parameter.

You can found an example with and without options [here](https://github.com/anotherhadi/wtui-components/blob/main/notification/example/main.go)

### Number Picker

![Number Picker Example](https://file.anotherhadi.com/wtui-components/numberpicker.gif)

The `numberpicker` component allows you to quickly prompt the user to choose a number.
Increment/Decrement with the arrow keys or HJ/KL, type a number to change the input, and validate with CR.
You can change 'Maximum, Minimum, Increment, Decimals'  through Opts.

You can found an example with and without options [here](https://github.com/anotherhadi/wtui-components/blob/main/numberpicker/example/main.go)

### Paragraph

![Paragraph Example](https://file.anotherhadi.com/wtui-components/paragraph.gif)

The `paragraph` component is used to print strings with the same look and feel as other components.
You can use the markdown syntax for bold and italic.

You can found an example with and without options [here](https://github.com/anotherhadi/wtui-components/blob/main/paragraph/example/main.go)

### RGBA Picker

![RGBA Picker Example](https://file.anotherhadi.com/wtui-components/rgbapicker.gif)

The `rgbapicker` component allows you to quickly prompt the user to choose a rgba color (Red,Green,Blue,Opacity).
Increment/Decrement/Move with the arrow keys or HJKL, type a number to change a field, and validate with CR.

You can found an example with and without options [here](https://github.com/anotherhadi/wtui-components/blob/main/rgbapicker/example/main.go)

### RGB Picker

![RGB Picker Example](https://file.anotherhadi.com/wtui-components/rgbpicker.gif)

The `rgbpicker` component allows you to quickly prompt the user to choose a rgb color (Red,Green,Blue).
Increment/Decrement/Move with the arrow keys or HJ/KL, type a number to change a field, and validate with CR.

You can found an example with and without options [here](https://github.com/anotherhadi/wtui-components/blob/main/rgbpicker/example/main.go)

### Selection

![Selection Example](https://file.anotherhadi.com/wtui-components/selection.gif)

The `selection` component allows you to quickly prompt the user to choose an option.
Move with the arrow keys or JK, select an option with CR.

You can found an example with and without options [here](https://github.com/anotherhadi/wtui-components/blob/main/selection/example/main.go)

### Shortcuts

![Shortcuts Example](https://file.anotherhadi.com/wtui-components/shortcuts.gif)

The `shortcuts` component allows you to print a list of shortcuts with the same look and feel as other components.

You can found an example [here](https://github.com/anotherhadi/wtui-components/blob/main/shortcuts/example/main.go)

### Spacer

![Spacer Example](https://file.anotherhadi.com/wtui-components/spacer.gif)

The `spacer` component allows you to print space between components

You can found an example [here](https://github.com/anotherhadi/wtui-components/blob/main/spacer/example/main.go)

### Table

![Table Example](https://file.anotherhadi.com/wtui-components/table.gif)

The `table` component is ideal for creating tables with columns and other features for organized data representation.

You can found an example with and without options [here](https://github.com/anotherhadi/wtui-components/blob/main/table/example/main.go)

## Options

You can refer to the examples folder for each component to learn more about options (`wtui-components/wtops`).

## License

This project is licensed under the [MIT License](LICENSE).
