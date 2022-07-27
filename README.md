# 🥃 Snob

Snob is a simple CLI that fetches movies and TV shows info, written in Go.

<p align="center">
  <a href="https://github.com/nunogois/snob/blob/main/.github/screenshot.png?raw=true"><img src="https://github.com/nunogois/snob/blob/main/.github/screenshot.png?raw=true" /></a>
<p>

## Install

Homebrew:
  
  ```bash
  brew install nunogois/brews/snob
  ```

Or you can download the [latest release](https://github.com/nunogois/snob/releases/latest) instead.

Alternatively, you can clone this repository and then:

  ```bash
  go install
  ```

## Setup

Before using `snob`, you should get an OMDb API key by following [these steps](https://www.omdbapi.com/apikey.aspx).

After that, you can set your key like this:

  ```bash
  snob -k <your-api-key>
  ```

Once you have set your key, you can make a simple search:

  ```bash
  snob lost ark
  ```

Use `snob -h` to see all the available commands.

## Using

 - [OMDb API](http://www.omdbapi.com/)
 - [urfave/cli](https://github.com/urfave/cli)
 - [fatih/color](https://github.com/fatih/color)
 - [Viper](https://github.com/spf13/viper)
