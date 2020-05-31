# archivasa

A simple static blog generator written in GO, straight to the point, minimal configuration

## Features

* Post, pages in Markdown
* Works in the current working directory
* Can use themes
* Support for markdown files
* Output folder
* Custom frontmatter

### Frontmatter

Starts at the first file and ends with '---'

    <tag>:<value>|<values>

#### Example

    date:2019-03-7
    category:things
    tags:misc,whatever
    ---

    # Hello world

    World world

## TODO

[x] Add something to read config files
[ ] Add a builder
[x] Add license
[ ] Rearrange the source code
