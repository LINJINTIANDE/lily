// Code generated by: `make actors-gen`. DO NOT EDIT.

package market

import (
	"bytes"
	"fmt"

	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"
	actorstypes "github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/go-state-types/manifest"
	market2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	adt2 "github.com/filecoin-project/specs-actors/v2/actors/util/adt"

	lotusactors "github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/adt"
)

var _ State = (*state2)(nil)

func load2(store adt.Store, root cid.Cid) (State, error) {
	out := state2{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func make2(store adt.Store) (State, error) {
	out := state2{store: store}

	ea, err := adt2.MakeEmptyArray(store).Root()
	if err != nil {
		return nil, err
	}

	em, err := adt2.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}

	out.State = *market2.ConstructState(ea, em, em)

	return &out, nil
}

type state2 struct {
	market2.State
	store adt.Store
}

func (s *state2) StatesChanged(otherState State) (bool, error) {
	otherState2, ok := otherState.(*state2)
	if !ok {
		// there's no way to compare different versions of the state, so let's
		// just say that means the state of balances has changed
		return true, nil
	}
	return !s.State.States.Equals(otherState2.State.States), nil
}

func (s *state2) States() (DealStates, error) {
	stateArray, err := adt2.AsArray(s.store, s.State.States)
	if err != nil {
		return nil, err
	}
	return &dealStates2{stateArray}, nil
}

func (s *state2) ProposalsChanged(otherState State) (bool, error) {
	otherState2, ok := otherState.(*state2)
	if !ok {
		// there's no way to compare different versions of the state, so let's
		// just say that means the state of balances has changed
		return true, nil
	}
	return !s.State.Proposals.Equals(otherState2.State.Proposals), nil
}

func (s *state2) Proposals() (DealProposals, error) {
	proposalArray, err := adt2.AsArray(s.store, s.State.Proposals)
	if err != nil {
		return nil, err
	}
	return &dealProposals2{proposalArray}, nil
}

type dealStates2 struct {
	adt.Array
}

func (s *dealStates2) Get(dealID abi.DealID) (DealState, bool, error) {
	var deal2 market2.DealState
	found, err := s.Array.Get(uint64(dealID), &deal2)
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}
	deal := fromV2DealState(deal2)
	return deal, true, nil
}

func (s *dealStates2) ForEach(cb func(dealID abi.DealID, ds DealState) error) error {
	var ds2 market2.DealState
	return s.Array.ForEach(&ds2, func(idx int64) error {
		return cb(abi.DealID(idx), fromV2DealState(ds2))
	})
}

func (s *dealStates2) decode(val *cbg.Deferred) (DealState, error) {
	var ds2 market2.DealState
	if err := ds2.UnmarshalCBOR(bytes.NewReader(val.Raw)); err != nil {
		return nil, err
	}
	ds := fromV2DealState(ds2)
	return ds, nil
}

func (s *dealStates2) array() adt.Array {
	return s.Array
}

func fromV2DealState(v2 market2.DealState) DealState {
	return dealStateV2{v2}
}

type dealStateV2 struct {
	ds2 market2.DealState
}

func (d dealStateV2) SectorStartEpoch() abi.ChainEpoch {
	return d.ds2.SectorStartEpoch
}

func (d dealStateV2) LastUpdatedEpoch() abi.ChainEpoch {
	return d.ds2.LastUpdatedEpoch
}

func (d dealStateV2) SlashEpoch() abi.ChainEpoch {
	return d.ds2.SlashEpoch
}

func (d dealStateV2) Equals(other DealState) bool {
	if ov2, ok := other.(dealStateV2); ok {
		return d.ds2 == ov2.ds2
	}

	if d.SectorStartEpoch() != other.SectorStartEpoch() {
		return false
	}
	if d.LastUpdatedEpoch() != other.LastUpdatedEpoch() {
		return false
	}
	if d.SlashEpoch() != other.SlashEpoch() {
		return false
	}

	return true
}

var _ DealState = (*dealStateV2)(nil)

type dealProposals2 struct {
	adt.Array
}

func (s *dealProposals2) Get(dealID abi.DealID) (*DealProposal, bool, error) {
	var proposal2 market2.DealProposal
	found, err := s.Array.Get(uint64(dealID), &proposal2)
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}

	proposal, err := fromV2DealProposal(proposal2)
	if err != nil {
		return nil, true, xerrors.Errorf("decoding proposal: %w", err)
	}

	return &proposal, true, nil
}

func (s *dealProposals2) ForEach(cb func(dealID abi.DealID, dp DealProposal) error) error {
	var dp2 market2.DealProposal
	return s.Array.ForEach(&dp2, func(idx int64) error {
		dp, err := fromV2DealProposal(dp2)
		if err != nil {
			return xerrors.Errorf("decoding proposal: %w", err)
		}

		return cb(abi.DealID(idx), dp)
	})
}

func (s *dealProposals2) decode(val *cbg.Deferred) (*DealProposal, error) {
	var dp2 market2.DealProposal
	if err := dp2.UnmarshalCBOR(bytes.NewReader(val.Raw)); err != nil {
		return nil, err
	}

	dp, err := fromV2DealProposal(dp2)
	if err != nil {
		return nil, err
	}

	return &dp, nil
}

func (s *dealProposals2) array() adt.Array {
	return s.Array
}

func fromV2DealProposal(v2 market2.DealProposal) (DealProposal, error) {

	label, err := labelFromGoString(v2.Label)

	if err != nil {
		return DealProposal{}, xerrors.Errorf("error setting deal label: %w", err)
	}

	return DealProposal{
		PieceCID:     v2.PieceCID,
		PieceSize:    v2.PieceSize,
		VerifiedDeal: v2.VerifiedDeal,
		Client:       v2.Client,
		Provider:     v2.Provider,

		Label: label,

		StartEpoch:           v2.StartEpoch,
		EndEpoch:             v2.EndEpoch,
		StoragePricePerEpoch: v2.StoragePricePerEpoch,

		ProviderCollateral: v2.ProviderCollateral,
		ClientCollateral:   v2.ClientCollateral,
	}, nil
}

func (s *state2) DealProposalsAmtBitwidth() int {
	return 3
}

func (s *state2) DealStatesAmtBitwidth() int {
	return 3
}

func (s *state2) ActorKey() string {
	return manifest.MarketKey
}

func (s *state2) ActorVersion() actorstypes.Version {
	return actorstypes.Version2
}

func (s *state2) Code() cid.Cid {
	code, ok := lotusactors.GetActorCodeID(s.ActorVersion(), s.ActorKey())
	if !ok {
		panic(fmt.Errorf("didn't find actor %v code id for actor version %d", s.ActorKey(), s.ActorVersion()))
	}

	return code
}

func (s *state2) GetProviderSectors() (map[abi.SectorID][]abi.DealID, error) {

	return nil, nil

}

func (s *state2) GetProviderSectorsByDealID(dealIDMap map[abi.DealID]bool) (map[abi.DealID]abi.SectorID, error) {

	return nil, nil

}
