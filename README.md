# archivasa

A simple static blog generator written in GO. Straight to the point, minimal configuration.

## Features

* Post and pages in Markdown
* Works in the current working directory
* Can use themes
* Support for markdown files

## Installation

### First step

    go clone https://gitgub.com/dotoscat/archivasa
    cd archivasa

### The main cli

    go install gitgub.com/dotoscat/archivasa/cmd/archivasa

## Usage

This is a command line interface. In your current working directory you need the next tree structure:

    * theme
        * css
            * main.css
        * templates
            * basic.tmpl
            * document.tmpl
            * postspage.tmpl
    * content
        * pages
        * posts

plus a config file named *config.txt*

the contents of this file may have the next, for example

    title = My blog
    postsperpage = 7

then do this

    archivasa

this will generate a folder *output* in your CWD with the site ready to be uploaded anywhere

### Preview

It is recomended to use a local server to preview the generated site. For example, if you have Python installed you can make

    python -m http.server -d output

to serve the files

### Frontmatter

Starts at the first file and ends with '---'

    <tag>=<value>[,<value, ...]

#### Example

    date:2019-03-7
    tags:misc,whatever
    ---

    # Hello world

    World world
