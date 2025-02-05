An example AWS plugin for use with [sigstore](https://github.com/sigstore/sigstore/tree/main/pkg/signature/kms/cliplugin).

This repo's `aws` package was [copied](https://github.com/sigstore/sigstore/tree/4b62818325b78ea76c0149b940e4b7fea31142b3/pkg/signature/kms/aws) directly from sigstore, also removing it's `init()`.

The main addition is `main.go` to run as a separate program, and tests are in [.github/workflows/test.yml](.github/workflows/test.yml)

Pending official releases of sigstore and cosign, use in cosign requires a [special build](https://github.com/sigstore/cosign/compare/main...ramonpetgrave64:cosign:cliplugin-no-builtin-aws).