#!/usr/bin/make -f

.PHONY: pb

pb:
	@ protoc --go_out=plugins=grpc:. chat/chat.proto
