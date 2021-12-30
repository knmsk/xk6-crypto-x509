# xk6-crypto-x509
A k6 extension to encrypt data with a PEM Public Key 

This is a [k6](https://go.k6.io/k6) extension using the
[xk6](https://github.com/grafana/xk6) system.

| :exclamation: This is a proof of concept, isn't supported by the k6 team, and may break in the future. USE AT YOUR OWN RISK! |
| ---------------------------------------------------------------------------------------------------------------------------- |
## Hash Supported:
* sha256
* sha512

## Usage

Import an entire module's contents:

```js
import * as cryptoX509 from "k6/x/crypto_x509";
```

## API

Functions:
* [encryptRsa](./docs/FUNCTION.md#encryptRsa)

## Example

```javascript
// example.js
import crypto from "k6/x/crypto-x509";

export default function () {
  console.log(
    crypto.encryptRsa(
      `-----BEGIN RSA PUBLIC KEY-----
      MIIBigKCAYEA0lavluO97KNSzvWMmVKjov0tFNLSEXZB5icjMY9zR1XXWS6F2mG0
      zmCSrsS0hGWIpq7O49UkjBY5IMLpOKFyKcIggftzjI6dq+Eql/ToOUPGoTLvu4Bx
      dL4afRPXpFX9QUnPN2gaNg31SUwf/ruKoDBJWKR/bymoreC34BLvk7J6wAbTaNxz
      Y59P4Doc4MV7bO9aDotw5LWHCFgfuWNypjsa4gVDQGbt7jCLDCIvksowYEcfpE7s
      nTuz8ojvJnLWg49jQbUk+4YntlQ/X8iSC9n5CsMnLTJu950PRpaq609TNuGv7pgr
      +VRFr69bdhfAppxvMok+EzFBl4KxedeauN5qe10Vbsr7ioMb5A3+gNy+Cdo1TUvf
      nWJOqZbos02DbmHL+UwSANBLgx1NX2pMYfQ9+h8FxedCOsuIvv5HEaa3vP/gE0iY
      eHJ6JSJNp6nQtW6llrt2ZYvGl/K/KuyvCzN0DatMpuXHn8/XEHqlEk2bzT6uTq6j
      VDUwqgzepMFtAgMBAAE=
      -----END RSA PUBLIC KEY-----`,
      "STRING TO BE ENCRYPTED"
    )
  );
}

```


## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:
  ```shell
  $ go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```shell
  $ xk6 build --with github.com/dgzlopes/xk6-exec@latest