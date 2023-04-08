# dictionary-cli

## Overview

---

This application is for playing with cobra and promptui.

You can add some notes like a dictionary by answering to questions in command line.

- item name
- definition of the item
- category of the item

## What is cobra?

---

Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.

https://github.com/spf13/cobra

## What is promptui?

---

Interactive prompt for command-line applications.

We built Promptui because we wanted to make it easy and fun to explore cloud services with manifold cli.

https://github.com/manifoldco/promptui

## How to play

---

There already exists both binary file and database thus you can skip step1 and step2.

### 1.Build this app first.

```
$ make build
```

### 2.Create a database(initialize).

```
$ make init
```

### 3.Add a note.

```
$ make new
```

### 4.Show all the lists you registered.

```
$ make list
```
