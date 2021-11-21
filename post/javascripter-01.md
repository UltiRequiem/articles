+++
title = "From beginner to JavaScripter, Part One"
description = "Variables, Data Types and Functions (From beginner to JavaScripter, Part One)"
date = "2021-11-20"
summary= "Learn about variables, constants, data types and functions of JavaScript."
tags = ["javascript"]
draft = false
author = "UltiRequiem"
+++

# From beginner to JavaScripter

In this series of tutorials you will learn JavaScript in a practical
way by doing projects and challenges.

In this chapter you will learn the basics of JavaScript.

It is not necessary that you have anything installed in this part,
but in the following parts you will need a text editor,
like [NeoVim](https://neovim.io) or [VSCode](https://code.visualstudio.com),
installed and [Node.js](https://nodejs.org).

You can execute the code from this article in [your browser's console](https://stackoverflow.com/a/51145033/14720975).

## Variables

A variable is a place that you can assign to a certain value.

An a constant is the same but cannot be reassigned.

We can initialize and assign a value to a constant variable like this:

```javascript
const numberEight = 8;
```

Lets break this snippets in parts:

- `const`: Stands for constant, is a keyword to define a variable that cannot change

- `numberEight`: The name of the constant

- `= 8`: We use the `=` symbol to assign, and `8` as the value assigned

Resuming all the line you are basically initializing a constant
and assigning it the value of `8`, which is of type `number`.

If you want a variable that can change you would
use the keyword `let`, instead of `const`. Example:

```javascript
let myFavoriteWord = "Pizza";
```

Here we are initializing a variable that may change in the future,
and assigning it the value of `"Pizza"`, which is of type `string`.

Now that you know what is a variable and how to assign one,
lets talk about the most common types on JavaScript.

## Basic Types

A type its a classification or categorization of a value.

Every variable, no matter if its defined with `let` or `const`, have a type.

The basic types are:

- `number`

It can be any number minor to `9007199254740991`.

E.g:

```javascript
const decemberNumberMonth = 12;
```

- `string`

It can be any text.

E.g:

```javascript
const myName = "John Doe";
```

- `boolean`

It can be `true` or `false`.

E.g:

```javascript
const iLikePizza = true;
```

- `array`

> Although this is technically not a type in JavaScript, an `[]` or `new Array()`
its of type `object`, I'm considering it a type for educational reasons.

E.g:

```javascript
const coolNames = ["Dayah", "Pedro", "Dazai", "Atsushi"];
```

You can put any type inside an `array`, even other `array`:

```javascript
const arrayOfArrays = [["Hello", "World"], false, 89, [true, false]];
```

## Functions

A `function` is a self-contained block of code, that may receive elements, if specified,
and may return something.

There are some global scoped function, this means that you can use it anywhere,
that are already defined. E.g:

- [`parseInt`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/parseInt)

- [`encodeURI`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURI)

- [`eval`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/eval)

- etc

Depending where you are running JavaScript there are going to be
different global scoped functions, for example, the `prompt` function
doesn't exists on [Node.js](https://nodejs.org) while in [Deno](https://deno.land)
or the browser it exists.

> Node.js and Deno are runtime for JavaScript, I will explain more in depth
> this later, but basically we use them to execute JavaScript outside the browser.

If you want to print a value or a variable in the console, you need to use
the `console.log()` function.

E.g:

```javascript
const helloWorld = "Hello World";
console.log(helloWorld);
```

You can also define your own function, for example:

```javascript
function sum(numberOne, numberTwo) {
  const numberOnePlusNumberTwo = numberOne + numberTwo;
  return numberOnePlusNumberTwo;
}
```

Again lets break this snippet into parts:

1. You start with the `function` keyword to tell JavaScript
   that what follows is a `function`.

2. We open `{` at the beginning of the function and `}` at the to
   define the range of the block.

3. We store the value of `numberOne` plus `numberTwo` in
   the `numberOnePlusNumberTwo` constant.

4. We use the `return` keyword to return the `numberOnePlusNumberTwo` value
   as result of the `function`

Congratulations! You have written your first function in JavaScript, so far you
learned about variables, data types and functions.

Here is an example using all we have used until the moment:

```javascript
function sum(numberOne, numberTwo) {
  const numberOnePlusNumberTwo = numberOne + numberTwo;
  return numberOnePlusNumberTwo;
}

const myFavoriteNumber = sum(3, 4);

console.log(myFavoriteNumber);
```

This is the first part of a series of tutorial, don't forget to follow me
on [Twitter](https://twitter.com/UltiRequiem) and
[Reddit](https://www.reddit.com/u/UltiRequiem) as UltiRequiem to know
when the next part comes out.
