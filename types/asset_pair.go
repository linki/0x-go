package types

type AssetPairs struct {
	Records []AssetPair
}

type AssetPair struct {
	AssetDataA Asset
	AssetDataB Asset
}
