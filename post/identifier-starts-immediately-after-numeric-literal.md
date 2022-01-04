+++
title = "Chain method to number literal"
description = "How to chain a method to a number literal"
date = "2022-01-04"
summary= "Learn how to chain a method to a number literal in JavaScript / Typescript"
tags = ["javascript", "typescript"]
draft = false 
author = "UltiRequiem"
+++

In JavaScript you can
[chain a method to an string literal](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String#instance_methods).

Example using
[`.toLowerCase`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/toLowerCase)
method:

```javascript
const sentence = "The quick brown fox jumps over the lazy dog.".toLowerCase();

console.log(sentence);
```

But this will not work with a number literal, example using [`.toString`]
(https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/toString):

```javascript
console.log(234.toString());
```

Depending of the runtime you
[can get different error messages](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Errors/Identifier_after_number#message),
for Example in [Deno](http://deno.land) you get:

```
parse error: Identifier cannot follow number
```

## How to "fix it"?

You can easily overcome this:

```javascript
console.log((234).toString());
```
