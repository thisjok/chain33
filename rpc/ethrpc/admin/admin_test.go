package admin

import (
	clientMocks "github.com/33cn/chain33/client/mocks"
	"github.com/33cn/chain33/queue"
	rpctypes "github.com/33cn/chain33/rpc/types"
	"github.com/33cn/chain33/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	admin *AdminApi
	qapi  *clientMocks.QueueProtocolAPI
	q     = queue.New("test")
)

func init() {
	qapi = &clientMocks.QueueProtocolAPI{}
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	q.SetConfig(cfg)
	admin = &AdminApi{}
	admin.cfg = cfg
}

func TestAdminApi_Peers(t *testing.T) {

	var peerlist types.PeerList

	var peer = &types.Peer{
		Addr:        "192.168.0.19",
		Port:        13803,
		Name:        "16Uiu2HAmBdwm5i6Ao6hBedNXHSM44ZUhM4243s5yJGAKPyHRjESw",
		MempoolSize: 0,
		Self:        false,
		Header: &types.Header{
			Version:    0,
			ParentHash: common.Hex2Bytes("0x56a00b884045e7087a6df93b58da8ba06d8127e3f9bf822b3c57271bf493aa62"),
			TxHash:     common.Hex2Bytes("0x91b60784d75497bc7c2ed681e51a985d933ec63ee6adf032ba6024a6adcb3dc9"),
			StateHash:  common.Hex2Bytes("0x593cfe7c38f21d87f2377983a911601140b679763e7e0439b830eab45dc26187"),
			Height:     2,
			BlockTime:  1647411454,
			TxCount:    1,
			Hash:       common.Hex2Bytes("0x560f19205f79a41bd89ccfc288dc89fe9902c33afafce9a43d00ed5c6c8e082e"),
			Difficulty: 0,
		},
		Version:     "6.0.0-3877-g1405d606a-1405d606@1.0.0",
		RunningTime: "0.044 minutes",
	}

	peerlist.Peers = append(peerlist.Peers, peer)

	testAdminApi_Peers(t, &peerlist)

}

func testAdminApi_Peers(t *testing.T, plist *types.PeerList) {
	qapi.On("PeerInfo", mock.Anything).Return(plist, nil)

	peers, err := admin.Peers()
	assert.Nil(t, err)
	t.Log("peers", peers)
}

func TestAdminApi_NodeInfo(t *testing.T) {
	var peerlist *types.PeerList = new(types.PeerList)

	var getPeer = types.P2PGetPeerReq{}

	var peer = &types.Peer{
		Addr:        "192.168.0.19",
		Port:        13803,
		Name:        "16Uiu2HAmBdwm5i6Ao6hBedNXHSM44ZUhM4243s5yJGAKPyHRjESw",
		MempoolSize: 0,
		Self:        true,
		Header: &types.Header{
			Version:    0,
			ParentHash: common.Hex2Bytes("0x56a00b884045e7087a6df93b58da8ba06d8127e3f9bf822b3c57271bf493aa62"),
			TxHash:     common.Hex2Bytes("0x91b60784d75497bc7c2ed681e51a985d933ec63ee6adf032ba6024a6adcb3dc9"),
			StateHash:  common.Hex2Bytes("0x593cfe7c38f21d87f2377983a911601140b679763e7e0439b830eab45dc26187"),
			Height:     2,
			BlockTime:  1647411454,
			TxCount:    1,
			Hash:       common.Hex2Bytes("0x560f19205f79a41bd89ccfc288dc89fe9902c33afafce9a43d00ed5c6c8e082e"),
			Difficulty: 0,
		},
		Version:     "6.0.0-3877-g1405d606a-1405d606@1.0.0",
		RunningTime: "0.044 minutes",
	}

	var peer2 = &types.Peer{
		Addr:        "192.168.0.20",
		Port:        13803,
		Name:        "12ABC2HAmBdwm5i6Ao6hBedNXHSM44ZUhM4243s5yJGAKPyHRjcdE",
		MempoolSize: 0,
		Self:        false,
		Header: &types.Header{
			Version:    0,
			ParentHash: common.Hex2Bytes("0x56a00b884045e7087a6df93b58da8ba06d8127e3f9bf822b3c57271bf493aa62"),
			TxHash:     common.Hex2Bytes("0x91b60784d75497bc7c2ed681e51a985d933ec63ee6adf032ba6024a6adcb3dc9"),
			StateHash:  common.Hex2Bytes("0x593cfe7c38f21d87f2377983a911601140b679763e7e0439b830eab45dc26187"),
			Height:     3,
			BlockTime:  1647411458,
			TxCount:    1,
			Hash:       common.Hex2Bytes("0x765f19205f79a41bd89ccfc288dc89fe9902c33afafce9a43d00ed5c6c8e283a"),
			Difficulty: 0,
		},
		Version:     "6.0.0-3877-g1405d606a-1405d606@1.0.0",
		RunningTime: "0.345 minutes",
	}

	peerlist.Peers = append(peerlist.Peers, peer, peer2)

	qapi.On("PeerInfo", mock.Anything).Return(&getPeer, peerlist)
	admin.cli.Init(q.Client(), qapi)

	var pr rpctypes.Peer
	for _, peer := range peerlist.Peers {

		if peer.Self != true {
			continue
		}

		pr.Addr = peer.GetAddr()
		pr.MempoolSize = peer.GetMempoolSize()
		pr.Name = peer.GetName()
		pr.Port = peer.GetPort()
		pr.Self = peer.GetSelf()
		pr.Header = &rpctypes.Header{
			BlockTime:  peer.Header.GetBlockTime(),
			Height:     peer.Header.GetHeight(),
			ParentHash: common.Bytes2Hex(peer.GetHeader().GetParentHash()),
			StateHash:  common.Bytes2Hex(peer.GetHeader().GetStateHash()),
			TxHash:     common.Bytes2Hex(peer.GetHeader().GetTxHash()),
			Version:    peer.GetHeader().GetVersion(),
			Hash:       common.Bytes2Hex(peer.GetHeader().GetHash()),
			TxCount:    peer.GetHeader().GetTxCount(),
		}

		pr.Version = peer.GetVersion()
		pr.LocalDBVersion = peer.GetLocalDBVersion()
		pr.StoreDBVersion = peer.GetStoreDBVersion()
		pr.RunningTime = peer.GetRunningTime()
		pr.FullNode = peer.GetFullNode()
		pr.Blocked = peer.GetBlocked()
	}

	testAdminApi_NodeInfo(t, &pr)
}

func testAdminApi_NodeInfo(t *testing.T, peer *rpctypes.Peer) {

	qapi.On("PeerInfo", mock.Anything).Return(peer, nil)

	node, err := admin.NodeInfo()
	assert.Nil(t, err)
	t.Log("node", node)
}
