# Javascript PipeScript
[![npm version](https://badge.fury.io/js/pipescript.svg)](https://badge.fury.io/js/pipescript)

## Installing

You can either directly download `pipescript.js` from [the github](https://github.com/connectordb/pipescript/releases), or you can add it to your project:

```bash
npm install pipescript
```

you can then
```javascript
import pipescript from 'pipescript';
```
or if you directly downloaded `pipescript.js`:

```html
<script type="text/javascript" src="pipescript.js"></script>
```

## Usage

Code:
```javascript
pipescript.Script('$ > 5').Transform([{t: 123.34, d: 4},{t: 123.35, d: 6}])
```
Output:
```javascript
[{t:123.34,d:false},{t:123.35,d:true}]
```


## Building

This package requires golang 1.7. It uses gopherjs to create a javascript version of PipeScript.

To build:

```bash
git clone https://github.com/connectordb/pipescript
cd pipescript/main/javascript

go get ./...

npm install
npm run build
npm run test
```

This will create the `pipescript.js` file.

You can then try it in the browser by opening `tryme.html`
