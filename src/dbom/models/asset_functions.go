package models

import "time"

var assetType string = "Software"
var assetManufacturer string = "Test"
var standardVersion float64 = 1.0
var assetModelNumber string = "N/A"

func NewAsset(assetMetadata interface{}, description *string, name *string, subType *string) *AssetDefinition {
	asset := AssetDefinition{}
	asset.AssetMetadata = assetMetadata
	asset.AssetDescription = description
	asset.DocumentName = name
	asset.AssetType = &assetType
	asset.AssetSubType = subType
	asset.AssetManufacturer = &assetManufacturer
	asset.DocumentCreator = &assetManufacturer
	asset.StandardVersion = &standardVersion
	manufactureSignature := ""
	asset.ManufactureSignature = &manufactureSignature
	asset.AssetModelNumber = &assetModelNumber
	currentTime := time.Now().String()
	asset.DocumentCreatedDate = &currentTime
	return &asset
}

func NewAttach(repoId, channelId, assetId, role, subRole *string) *AttachSubAssetRequest {
	attach := AttachSubAssetRequest{}
	attach.AssetID = assetId
	attach.ChannelID = channelId
	attach.RepoID = repoId
	attach.Role = role
	attach.SubRole = subRole
	return &attach
}
