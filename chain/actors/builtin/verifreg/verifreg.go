// Code generated by: `make actors-gen`. DO NOT EDIT.
package verifreg

import (
	"fmt"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/go-state-types/cbor"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	builtin5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"
	builtin6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"
	builtin7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"

	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/lily/chain/actors/adt"
	"github.com/filecoin-project/lily/chain/actors/builtin"
)

func init() {

	builtin.RegisterActorState(builtin0.VerifiedRegistryActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load0(store, root)
	})

	builtin.RegisterActorState(builtin2.VerifiedRegistryActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load2(store, root)
	})

	builtin.RegisterActorState(builtin3.VerifiedRegistryActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load3(store, root)
	})

	builtin.RegisterActorState(builtin4.VerifiedRegistryActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load4(store, root)
	})

	builtin.RegisterActorState(builtin5.VerifiedRegistryActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load5(store, root)
	})

	builtin.RegisterActorState(builtin6.VerifiedRegistryActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load6(store, root)
	})

	builtin.RegisterActorState(builtin7.VerifiedRegistryActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load7(store, root)
	})

}

var (
	Address = builtin7.VerifiedRegistryActorAddr
	Methods = builtin7.MethodsVerifiedRegistry
)

func AllCodes() []cid.Cid {
	return []cid.Cid{
		builtin0.VerifiedRegistryActorCodeID,
		builtin2.VerifiedRegistryActorCodeID,
		builtin3.VerifiedRegistryActorCodeID,
		builtin4.VerifiedRegistryActorCodeID,
		builtin5.VerifiedRegistryActorCodeID,
		builtin6.VerifiedRegistryActorCodeID,
		builtin7.VerifiedRegistryActorCodeID,
	}
}

func Load(store adt.Store, act *types.Actor) (State, error) {
	switch act.Code {

	case builtin0.VerifiedRegistryActorCodeID:
		return load0(store, act.Head)

	case builtin2.VerifiedRegistryActorCodeID:
		return load2(store, act.Head)

	case builtin3.VerifiedRegistryActorCodeID:
		return load3(store, act.Head)

	case builtin4.VerifiedRegistryActorCodeID:
		return load4(store, act.Head)

	case builtin5.VerifiedRegistryActorCodeID:
		return load5(store, act.Head)

	case builtin6.VerifiedRegistryActorCodeID:
		return load6(store, act.Head)

	case builtin7.VerifiedRegistryActorCodeID:
		return load7(store, act.Head)

	}
	return nil, fmt.Errorf("unknown actor code %s", act.Code)
}

type State interface {
	cbor.Marshaler

	Code() cid.Cid

	RootKey() (address.Address, error)
	VerifiedClientDataCap(address.Address) (bool, abi.StoragePower, error)
	VerifierDataCap(address.Address) (bool, abi.StoragePower, error)
	ForEachVerifier(func(addr address.Address, dcap abi.StoragePower) error) error
	ForEachClient(func(addr address.Address, dcap abi.StoragePower) error) error

	verifiers() (adt.Map, error)
	verifiedClients() (adt.Map, error)
}

type VerifierInfo struct {
	Address address.Address
	DataCap abi.StoragePower
}

type VerifierChange struct {
	Before VerifierInfo
	After  VerifierInfo
}

type VerifierChanges struct {
	Added    []VerifierInfo
	Modified []VerifierChange
	Removed  []VerifierInfo
}
