# dictionary-cli

## Overview

This application is for playing with cobra and promptui.
You can add some notes like a dictionary by answering to questions in command line.

- item name
- definition of the item
- category of the item

## What is cobra?

Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.

https://github.com/spf13/cobra

## What is promptui?

Interactive prompt for command-line applications.

We built Promptui because we wanted to make it easy and fun to explore cloud services with manifold cli.

https://github.com/manifoldco/promptui

## How to play

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

## Examples

You'll be asked "What word would you like to make a note of?"

<img width="417" alt="Screenshot 2023-04-07 at 18 59 00" src="https://user-images.githubusercontent.com/41677855/230698526-9b710b9b-0fc9-4071-a011-588037723c1d.png">

Input a word.

<img width="454" alt="Screenshot 2023-04-07 at 18 59 08" src="https://user-images.githubusercontent.com/41677855/230698527-6bf76c8f-ff18-4c02-a224-f3742984cfcd.png">

You'll be asked "What is the definition of {word}?"

<img width="479" alt="Screenshot 2023-04-07 at 18 59 20" src="https://user-images.githubusercontent.com/41677855/230698528-c3d56415-f32b-491c-806c-8005cb88d6d2.png">

Input a description.

<img width="796" alt="Screenshot 2023-04-07 at 18 59 49" src="https://user-images.githubusercontent.com/41677855/230698529-66392537-50b1-435b-bf66-6f300d809b9c.png">

Choose on category.

<img width="783" alt="Screenshot 2023-04-07 at 19 00 04" src="https://user-images.githubusercontent.com/41677855/230698530-0ec5c35b-fa90-4285-aca6-c659f7f23897.png">

Complete!

<img width="796" alt="Screenshot 2023-04-07 at 18 58 46" src="https://user-images.githubusercontent.com/41677855/230698537-c7a99db0-4e2f-4052-bcc9-591ddca35799.png">

You'll see the word has added in your dictionary.

<img width="1168" alt="Screenshot 2023-04-07 at 19 00 33" src="https://user-images.githubusercontent.com/41677855/230698531-e9cddabb-9991-49d9-9427-34dce5f3ecc7.png">

Oh. Can't find a proper category?

<img width="1104" alt="Screenshot 2023-04-07 at 19 03 44" src="https://user-images.githubusercontent.com/41677855/230698532-84669b24-ad7d-4e48-ad6f-6027523e8c5d.png">

Don't worry. You can add new category.

<img width="1151" alt="Screenshot 2023-04-07 at 19 04 00" src="https://user-images.githubusercontent.com/41677855/230698534-c8a11d43-7785-4c3c-adb2-b318c555d9d8.png">

There you go.

<img width="982" alt="Screenshot 2023-04-07 at 19 13 42" src="https://user-images.githubusercontent.com/41677855/230698860-2c1ca557-1d57-462c-a0cc-6d792576c639.png">
