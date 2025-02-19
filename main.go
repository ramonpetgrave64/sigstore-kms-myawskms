package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sigstore/sigstore/pkg/signature/kms"

	"github.com/ramonpetgrave64/sigstore-kms-myawskms/aws"
	"github.com/sigstore/sigstore/pkg/signature/kms/cliplugin/common"
	"github.com/sigstore/sigstore/pkg/signature/kms/cliplugin/handler"
)

const expectedProtocolVersion = common.ProtocolVersion

func newSignerVerifier(initOptions *common.InitOptions) (kms.SignerVerifier, error) {
	ctx := context.TODO()
	// cliplugin will strip the part up to [plugin name]://[key ref],
	// but the existing code expects a specific regex, so we reconstruct.
	fullKeyResourceID := aws.ReferenceScheme + initOptions.KeyResourceID
	return aws.LoadSignerVerifier(ctx, fullKeyResourceID)
}

func main() {
	// we log to stderr, not stdout. stdout is reserved for the plugin return value.
	if protocolVersion := os.Args[1]; protocolVersion != expectedProtocolVersion {
		err := fmt.Errorf("expected protocol version: %s, got %s", expectedProtocolVersion, protocolVersion)
		handler.WriteErrorResponse(os.Stdout, err)
		panic(err)
	}

	pluginArgs, err := handler.GetPluginArgs(os.Args)
	if err != nil {
		handler.WriteErrorResponse(os.Stdout, err)
		panic(err)
	}

	signerVerifier, err := newSignerVerifier(pluginArgs.InitOptions)
	if err != nil {
		handler.WriteErrorResponse(os.Stdout, err)
		panic(err)
	}

	if _, err := handler.Dispatch(os.Stdout, os.Stdin, pluginArgs, signerVerifier); err != nil {
		// Dispatch() will have already called WriteResponse() with the error.
		panic(err)
	}
}
