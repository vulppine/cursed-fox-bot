package cursedfoxbot

import (
	"context"
)

type PubSubMessage struct {
	Data []byte
}

func GooglePubSubEntryPoint(ctx context.Context, m PubSubMessage) {
	MakeCursedFox()
}
