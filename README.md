# Ransom

Fun command-line utility that converts a string into emoji codes.

> [!NOTE]
> You must have a custom emoji for each letter.

<picture>
  <source media="(prefers-color-scheme: light)" srcset="https://github.com/user-attachments/assets/0642fd90-3070-4e14-9da2-d23143a09edc">
  <img width="450" alt="Ransom Screenshot" src="https://github.com/user-attachments/assets/7dacd900-4837-4e20-9624-ff525e157b00">
</picture>

## Installation

```shell
go install github.com/gabe565/ransom@latest
```

## Usage

Run the command with any number of arguments. The arguments will be converted into emoji codes and copied to the clipboard.

```shell
$ ransom hello world
:h::e::l::l::o::blank::w::o::r::l::d:
INFO Copied to clipboard
```
