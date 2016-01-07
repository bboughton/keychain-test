# KeyChain Test

## Setup

To setup your keychain run the following in your shell

```
security add-generic-password -a AWS_SECRET_KEY -s foo.bar -w foo
security add-generic-password -a AWS_ACCESS_KEY -s foo.bar -w bar
```

## Teardown

To remove the keychain items run
```
security delete-generic-password -a AWS_SECRET_KEY -s foo.bar
security delete-generic-password -a AWS_ACCESS_KEY -s foo.bar
```
