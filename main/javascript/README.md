# Javascript PipeScript
[![npm version](https://badge.fury.io/js/pipescript.svg)](https://badge.fury.io/js/pipescript)

This package requires golang 1.5. It uses gopherjs to create a javascript version.

To build:

```bash
git clone https://github.com/connectordb/pipescript
cd pipescript/main/javascript

go get ./...

npm install
npm run build
npm run test
```

This will create the pipescript.js file.

You can then try it in the browser by opening `tryme.html`
