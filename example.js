import * as crypto from "k6/x/crypto-x509";

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
