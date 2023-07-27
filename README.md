# Sign Arbitrary Messages

This application allows you to sign an arbitrary message using any Cosmos SDK based application. This app has been tested with Archway's `archwayd` and Terra's `terrad`.

## Requirements

- go >= 1.20 (https://go.dev/dl)

## Configuration

Make a copy of the **.env.example** file and rename it to **.env**. The following information will explain the values contained within:

- **WALLET_KEY**: This is the name of the key created using the `.. key add ..` command. For example, with Archway this would be `archwayd keys add <key-name>`.
- **APP_NAME**: This is the name of the blockchain application that will be used to sign the message. For **Archway**, this is `archwayd`, and for **Terra**, this is `terrad`.
- **APP_ROOT_DIR**: This is where the application stores its data and configuration files. For **Archway** on a mac and linux, this usually `~/.archway` and for **Terra** this is usually `~/.terra`.
- **BACKEND**: This is the keyring backend used for the key that will be used to sign the message. Some examples include `os`, `file` and  `test`. **Note**: Based on my tests, only the `test` backend works with `terrad`.
- **MSG**: This is the message that the user will sign.

## Execute

In order to sign the message, first clone the repository. Then, navigate to the directory containing the code and execute the following commands:
- `go build`
- `go run main.go`

You should see the **Signature** and the **Public Key** printed out. These should be sent to whoever needs to verify the signature.
