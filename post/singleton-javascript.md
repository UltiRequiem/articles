+++
title = "How to do a Singleton in JavaScript"
description = "How to Generate a Singleton in JavaScript"
date = "2021-10-30"
summary= "Learn how to generate a Singleton in JavaScript / Typescript"
tags = ["java", "typescript"]
draft = false
author = "UltiRequiem"
+++

You can generate a Singleton in JavaScript very easily using Modules, E.g:

```javascript
class DataBase {
  constructor() {
    DBDriver.connect();
  }

  async getLink(id) {
    // Do stuff
  }

  async deleteByID(id) {
    // Do other stuff
  }
}

export default new DataBase();
```

Yes, an ESM/CommonJS module is guaranteed to return the same instance,
every time you import it.

While this might be sound bad, because it is a mutable global state, it is very handy
in this case.

Also if you class don't use `this`, doesn't have a constructor, or all their methods are `static`,
consider using a literal object. Eg:

```javascript
const DataBaseTwo = {
  getLink: (id) => {
    // Do stuff
  },

  deleteByID: (id) => {
    // Do other stuff
  },
};

export default DataBaseTwo;
```

There is an [eslint rule for this if you want](https://github.com/sindresorhus/eslint-plugin-unicorn/blob/main/docs/rules/no-static-only-class.md).
