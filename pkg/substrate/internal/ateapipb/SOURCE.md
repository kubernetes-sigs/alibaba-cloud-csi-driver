# Substrate API projection

The `.proto`, generated Go messages and generated gRPC client/server in this
directory are a minimal projection of `pkg/proto/ateapipb` from the deployed
Substrate integration, extracted at source revision
`abe771016ae50264d6018885a910fdbbd7a4f0de`. They are not a verbatim copy of
that package.

The bridge performs two read-only lookups: `Control.GetActor` and
`Control.GetActorTemplate`. The projection keeps those RPCs and the message
closure they need - `ResourceMetadata` (atespace, name, uid, annotations),
`Actor` (metadata, actor_template), `ObjectRef`, `ActorTemplate` (metadata),
`GetActorRequest` and `GetActorTemplateRequest`. Message, RPC, field names and
field numbers are verbatim from the source revision; omitted fields are marked
`reserved` at their original numbers so a refresh cannot silently reuse them.
The `go_package` option points at this repository because the copy is generated
here.

## Why not depend on the public Substrate module

The public `github.com/agent-substrate/substrate` API currently has no
`ResourceMetadata.annotations`, and that annotation is the carrier for this
bridge's publish workflow. Depending on the public module would silently drop
the one deployed extension the bridge relies on; its monolithic module also
introduces unrelated dependency upgrades. The annotation is the only reason
this projection exists: everything else the bridge uses matches the public API.
Publishing an annotation-capable, lightweight Substrate API module remains the
prerequisite for replacing this projection with a normal public module
dependency. No claim is made that the bridge works with an unextended public
Substrate server.

## Refreshing

Do not edit the `.proto`, message fields or RPC definitions by hand. Refresh by
re-extracting the same RPCs and message closure from a new Substrate source
revision, then regenerate with protoc 25.3, protoc-gen-go v1.36.11-devel and
protoc-gen-go-grpc v1.6.2:

```
protoc -I . --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  pkg/substrate/internal/ateapipb/ateapi.proto
```
