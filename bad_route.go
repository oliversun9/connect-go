package rerpc

import (
	"context"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// NewBadRouteHandler always returns gRPC and Twirp's equivalent of the
// standard library's http.StatusNotFound. To be fully compatible with the
// Twirp specification, mount this handler at the root of your API (so that it
// handles any requests for invalid protobuf methods).
func NewBadRouteHandler(opts ...HandlerOption) *Handler {
	return NewHandler(
		"", "", "", // protobuf method, service, package names
		func() proto.Message { return &emptypb.Empty{} }, // unused req msg
		func(ctx context.Context, _ proto.Message) (proto.Message, error) {
			path := "???"
			if md, ok := HandlerMeta(ctx); ok {
				path = md.Spec.Path
			}
			return nil, Wrap(CodeNotFound, newBadRouteError(path))
		},
		opts...,
	)
}
