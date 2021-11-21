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

(Although this is technically not a type in JavaScript, an `[]` or `new Array()`
its of type `object`, I'm considering it a type for educational reasons.)

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
> this later, but basically we use them to execute JavaScript in a computer.