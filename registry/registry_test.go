package registry

import (
	"testing"

	"github.com/linki/0x-go/types"

	"github.com/stretchr/testify/suite"
)

type RegistrySuite struct {
	suite.Suite
}

func (suite *RegistrySuite) TestAssetRegistry() {
	for _, tt := range []struct {
		address string
		symbol  string
		digits  int
	}{
		{"0x2956356cd2a2bf3202f771f50d3d14a367b48070", "WETH", 18},
		{"0x744d70fdbe2ba4cf95131626614a1763df805b9e", "SNT", 18},
		{"0xaec2e87e0a235266d9c5adc9deb4b2e29b54d009", "SNGLS", 0},
		{"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", "WETH2", 18},
		{"0x2e071d2966aa7d8decb1005885ba1977d6038a65", "ROL", 16},
		{"0xe0b7927c4af23765cb51314a0e0521a9645f0e2a", "DGD", 9},
		{"0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0", "EOS", 18},
		{"0x888666ca69e0f178ded6d75b5726cee99a87d698", "ICN", 18},
		{"0xb63b606ac810a52cca15e44bb630fd42d8d1d83d", "MCO", 8},
		{"0xe7775a6e9bcf904eb39da2b68c5efb4f9360e08c", "TAAS", 6},
		{"0xcb94be6f13a1182e4a4b6140cb7bf2025d28e41b", "TRST", 6},
		{"0x701c244b988a513c945973defa05de933b23fe1d", "OAX", 18},
		{"0xb7cb1c96db6b22b0d3d9536e0108d062bd488f74", "WTC", 18},
		{"0xe41d2489571d322189246dafa5ebde1f4699f498", "ZRX", 18},
		{"0xe94327d07fc17907b4db788e5adf2ed424addff6", "REP", 18},
		{"0x419d0d8bdd9af5e606ae2232ed285aff190e711b", "FUN", 8},
		{"0xaf30d2a7e90d7dc361c8c4585e9bb7d2f6f15bc7", "1ST", 18},
		{"0xf433089366899d83a9f26a773d59ec7ecf30355e", "MTL", 8},
		{"0x7c5a0ce9267ed19b22f8cae653f198e3e8daf098", "SAN", 18},
		{"0xd0d6d6c5fe4a677d343cc433536bb717bae167dd", "ADT", 9},
		{"0x8f8221afbb33998d8584a2b05749ba73c37a938a", "REQ", 18},
		{"0xc66ea802717bfb9833400264dd12c2bceaa34a6d", "MKR", 18},
		{"0xb64ef51c888972c908cfacf59b47c1afbc0ab8ac", "STORJ", 8},
		{"0x08711d3b02c8758f2fb3ab4e80228418a7f8e39c", "EDG", 0},
		{"0x607f4c5bb672230e8672085532f7e901544a7375", "RLC", 9},
		{"0x667088b212ce3d06a1b553a7221e1fd19000d9af", "WINGS", 18},
		{"0x12480e24eb5bec1a9d4369cab6a80cad3c0a377a", "SUB", 2},
		{"0xaaaf91d9b90df800df4f55c205fd6989c977e73a", "TKN", 8},
		{"0x01afc37f4f85babc47c0e2d0eababc7fb49793c8", "WGNT", 18},
		{"0x4156d3342d5c385a87d264f90653733592000581", "SALT", 8},
		{"0x1776e1f26f98b1a5df9cd347953a26dd3cb46671", "NMR", 18},
		{"0x27dce1ec4d3f72c3e457cc50354f1f975ddef488", "AIR", 8},
		{"0xd26114cd6ee289accf82350c8d8487fedb8a0c07", "OMG", 18},
		{"0xb97048628db6b661d4c2aa833e95dbe1a905b280", "PAY", 18},
		{"0x1f573d6fb3f13d689ff844b4ce37794d79a7ff1c", "BNT", 18},
		{"0x5af2be193a6abca9c8817001f45744777db30756", "BQX", 8},
		{"0x12fef5e57bf45873cd9b62e9dbd7bfb99e32d73e", "CFI", 18},
		{"0x818fc6c2ec5986bc6e2cbf00939d90556ab12ce5", "KIN", 18},
		{"0xfa05a73ffe78ef8f1a739473e462c54bae6567d9", "LUN", 18},
		{"0x9a642d6b3368ddc662ca244badf32cda716005bc", "QTUM", 18},
		{"0x0abdace70d3790235af448c88547603b945604ea", "DNT", 18},
		{"0xab16e0d25c06cb376259cc18c1de4aca57605589", "FUCK", 4},
		{"0x514910771af9ca656af840dff83e8264ecf986ca", "LINK", 18},
		{"0xdd974d5c2e2928dea5f71b9825b8b646686bd200", "KNC", 18},
		{"0x27054b13b1b798b345b591a4d22e6562d47ea75a", "AST", 4},
		{"0xf0ee6b27b759c9893ce4f094b49ad28fd15a23e4", "ENG", 8},
		{"0xbeb9ef514a379b997e0798fdcc901ee474b6d9a1", "MLN", 18},
		{"0x6810e776880c02933d47db1b9fc05908e5386b96", "GNO", 18},
		{"0x0d8775f648430679a709e98d2b0cb6250d2887ef", "BAT", 18},
		{"0x960b236a07cf122663c4303350609a66a7b288c0", "ANT", 18},
		{"0x41e5560054824ea6b0732e656e3ad64e20e94e45", "CVC", 8},
		{"0xd4fa1460f537bb9085d22c7bccb5dd450ef28e3a", "PPT", 8},
		{"0x56ba2ee7890461f463f7be02aac3099f6d5811a8", "CAT", 18},
	} {
		suite.Require().Contains(Registry, tt.address, "registry doesn't contain '%s'", tt.address)
		suite.Equal(tt.symbol, Registry[tt.address], "address '%s' has unexpected asset symbol '%s', expected '%s'", tt.address, Registry[tt.address], tt.symbol)
		suite.Equal(tt.digits, Digits[tt.address], "address '%s' (asset '%s') has unexpected digits '%d', expected '%d'", tt.address, Registry[tt.address], Digits[tt.address], tt.digits)

		asset, ok := AssetRegistry[tt.address]

		suite.Require().True(ok, "registry doesn't contain '%s'", tt.address)
		suite.Equal(tt.address, asset.AssetData, "referencing address '%s' doesn't match asset's address '%s'", tt.address, asset.AssetData)
		suite.Equal(tt.symbol, asset.Symbol, "address '%s' has unexpected asset symbol '%s', expected '%s'", tt.address, asset.Symbol, tt.symbol)
		suite.Equal(tt.digits, asset.Digits, "address '%s' (asset '%s') has unexpected digits '%d', expected '%d'", tt.address, asset.Symbol, asset.Digits, tt.digits)
	}
}

func (suite *RegistrySuite) TestLookup() {
	for _, tt := range []struct {
		address string
		asset   types.Asset
	}{
		{"0x2956356cd2a2bf3202f771f50d3d14a367b48070", AssetRegistry["0x2956356cd2a2bf3202f771f50d3d14a367b48070"]}, // lower normal
		{"0x2956356cD2a2bf3202F771F50D3D14A367b48070", AssetRegistry["0x2956356cd2a2bf3202f771f50d3d14a367b48070"]}, // mixed case
		{"unknown address", types.UnknownAsset}, // unknown asset
	} {
		suite.Equal(tt.asset, Lookup(tt.address))
	}
}

func TestRegistrySuite(t *testing.T) {
	suite.Run(t, new(RegistrySuite))
}
