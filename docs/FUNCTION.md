# Functions

## encryptRsa

▸ encryptRsa(`publicKey`: string, `data`: string, `algorithm`: string): string

| Name        | Type     | Description                                                |
| ----------- | -------- | ---------------------------------------------------------- |
| `publicKey` | `string` | Public key in PEM format to encrypt                        |
| `data`      | `string` | Data to be encrypted                                       |
| `algorithm` | `string` | Algorithm to be used. Accepts: sha256 and sha512 (default) |

**Returns**: string

The data encrypted using the chose algorithm