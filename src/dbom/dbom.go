/*
 * Copyright 2022 Unisys Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package dbom implements dbom storage and the grafeas functions
package dbom

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	prpb "github.com/grafeas/grafeas/proto/v1beta1/project_go_proto"
	"github.com/rs/zerolog"

	"dbom-grafeas/dbom/client"
	"dbom-grafeas/dbom/client/asset"
	"dbom-grafeas/dbom/config"
	"dbom-grafeas/dbom/models"
)

const (
	projectDocumentKind     = "projects"
	occurrencesDocumentKind = "occurrences"
	notesDocumentKind       = "notes"
	sortField               = "createTime"
)

var projectPath string = "projects"
var notePath string = "notes"
var occurrencePath string = "occurrences"

// var dbomStorage.config.Channel string = "test_channel"
// var dbomStorage.config.Repo string = "DB1"
var eq string = "$eq"
var in string = "$in"
var and string = "$and"
var underscore = "_"
var slash = "/"
var assetSubType string = "assetSubType"
var assetMetadata string = "assetMetadata"
var occurrenceNote string = "assetMetadata.note_name"
var attachedChildren string = "attachedChildren"
var assetID string = "assetID"
var id string = "_id"
var parentAssetID string = "parentAsset.assetID"
var notImplemented string = "Not Implemented"

type DbomStorage struct {
	logger  *zerolog.Logger
	config  *config.DBoMConfig
	gateway *client.ChainsourceGateway
}

func NewDbomStorage(
	logger *zerolog.Logger,
	config *config.DBoMConfig) *DbomStorage {
	transport := client.TransportConfig{Host: config.Gateway, BasePath: client.DefaultBasePath, Schemes: client.DefaultSchemes}
	gateway := client.NewHTTPClientWithConfig(nil, &transport)
	return &DbomStorage{logger, config, gateway}
}

// CreateProject creates writes a project asset to the configured dbom channel
func (dbomStorage *DbomStorage) CreateProject(ctx context.Context, projectID string, project *prpb.Project) (*prpb.Project, error) {
	params := asset.NewAddAssetParams()
	projectName := createProjectID(projectID)
	params.SetAssetID(projectName)
	params.SetChannelID(dbomStorage.config.Channel)
	params.SetRepoID(dbomStorage.config.Repo)
	assetBody := models.NewAsset(project, &project.Name, &project.Name, &projectPath)
	params.SetBody(assetBody)
	_, err := dbomStorage.gateway.Asset.AddAsset(params)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	return project, nil
}

// GetProject returns the project asset from the configured dbom channel
func (dbomStorage *DbomStorage) GetProject(ctx context.Context, projectID string) (*prpb.Project, error) {
	params := asset.NewRetrieveAssetParams()
	params.SetChannelID(dbomStorage.config.Channel)
	params.SetRepoID(dbomStorage.config.Repo)
	params.SetAssetID(createProjectID(projectID))
	res, err := dbomStorage.gateway.Asset.RetrieveAsset(params)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	payload := res.Payload
	project := prpb.Project{}
	meta, _ := json.Marshal(payload.AssetMetadata)
	json.Unmarshal(meta, &project)
	return &project, nil
}

// ListProjects returns the project assets from the configured dbom channel
func (dbomStorage *DbomStorage) ListProjects(ctx context.Context, filter string, pageSize int, pageToken string) ([]*prpb.Project, string, error) {
	params := asset.NewQueryAssetsPostParams()
	params.SetChannelID(dbomStorage.config.Channel)
	params.SetRepoID(dbomStorage.config.Repo)
	query := models.QueryAssetDefinition{}
	query.Query = createEqualQuery(assetSubType, projectPath)
	params.SetBody(&query)
	res, err := dbomStorage.gateway.Asset.QueryAssetsPost(params)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, "", err
	}
	payload := res.Payload.(map[string]interface{})
	projects := []*prpb.Project{}
	for _, element := range payload {
		project := prpb.Project{}
		meta, _ := json.Marshal(element.(map[string]interface{})[assetMetadata])
		json.Unmarshal(meta, &project)
		projects = append(projects, &project)
	}
	return projects, "", nil
}

// DeleteProject is not supported by dbom
func (dbomStorage *DbomStorage) DeleteProject(ctx context.Context, projectID string) error {
	return errors.New(notImplemented)
}

// GetOccurrence returns the occurrence asset from the configured dbom channel
func (dbomStorage *DbomStorage) GetOccurrence(ctx context.Context, projectID, occurrenceID string) (*pb.Occurrence, error) {
	payload, err := dbomStorage.getAsset(ctx, createOccurrenceID(projectID, occurrenceID))
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	occurrence := pb.Occurrence{}
	meta, _ := json.Marshal(payload.AssetMetadata)
	json.Unmarshal(meta, &occurrence)
	return &occurrence, nil
}

// ListOccurrences returns the occurrence assets from the configured dbom channel
func (dbomStorage *DbomStorage) ListOccurrences(ctx context.Context, projectID, filter, pageToken string, pageSize int32) ([]*pb.Occurrence, string, error) {
	params := asset.NewQueryAssetsPostParams()
	params.SetChannelID(dbomStorage.config.Channel)
	params.SetRepoID(dbomStorage.config.Repo)
	query := models.QueryAssetDefinition{}
	projectQuery := createEqualQuery(parentAssetID, createProjectID(projectID))
	occurrenceQuery := createEqualQuery(assetSubType, occurrencePath)
	query.Query = createAndQuery(projectQuery, occurrenceQuery)
	params.SetBody(&query)
	res, err := dbomStorage.gateway.Asset.QueryAssetsPost(params)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, "", err
	}
	payload := res.Payload.(map[string]interface{})
	occurrences := []*pb.Occurrence{}
	for _, element := range payload {
		occurrence := pb.Occurrence{}
		meta, _ := json.Marshal(element.(map[string]interface{})[assetMetadata])
		json.Unmarshal(meta, &occurrence)
		occurrences = append(occurrences, &occurrence)
	}
	return occurrences, "", nil
}

// CreateOccurrence writes a occurrence asset to the configured dbom channel and attaches it to the specified note
func (dbomStorage *DbomStorage) CreateOccurrence(ctx context.Context, projectID, userID string, occurrence *pb.Occurrence) (*pb.Occurrence, error) {
	_, err := dbomStorage.GetProject(ctx, projectID)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	if occurrence.CreateTime == nil {
		occurrence.CreateTime = ptypes.TimestampNow()
	}
	if occurrence.UpdateTime == nil {
		occurrence.UpdateTime = occurrence.CreateTime
	}
	params := asset.NewAddAssetParams()
	projectName := createProjectID(projectID)
	occurrenceName := convertID(occurrence.Name)
	params.SetAssetID(occurrenceName)
	params.SetChannelID(dbomStorage.config.Channel)
	params.SetRepoID(dbomStorage.config.Repo)
	assetBody := models.NewAsset(occurrence, &occurrenceName, &occurrenceName, &occurrencePath)
	params.SetBody(assetBody)
	_, err = dbomStorage.gateway.Asset.AddAsset(params)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	fullNoteName := occurrence.NoteName
	if fullNoteName != "" {
		noteName := convertID(fullNoteName)
		attachParams := asset.NewAttachSubAssetParams()
		attachParams.SetAssetID(noteName)
		attachParams.SetChannelID(dbomStorage.config.Channel)
		attachParams.SetRepoID(dbomStorage.config.Repo)
		attachBody := models.NewAttach(&dbomStorage.config.Repo, &dbomStorage.config.Channel, &occurrenceName, &notePath, &occurrencePath)
		attachParams.SetBody(attachBody)
		_, err = dbomStorage.gateway.Asset.AttachSubAsset(attachParams)
		if err != nil {
			dbomStorage.logger.Err(err).Msg(err.Error())
			return nil, err
		}
	}
	attachParams := asset.NewAttachSubAssetParams()
	attachParams.SetAssetID(projectName)
	attachParams.SetChannelID(dbomStorage.config.Channel)
	attachParams.SetRepoID(dbomStorage.config.Repo)
	attachBody := models.NewAttach(&dbomStorage.config.Repo, &dbomStorage.config.Channel, &occurrenceName, &projectPath, &occurrencePath)
	attachParams.SetBody(attachBody)
	_, err = dbomStorage.gateway.Asset.AttachSubAsset(attachParams)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}

	return occurrence, nil
}

// BatchCreateOccurrences writes the occurrence assets to the configured dbom channel and attaches them to the specified notes
func (dbomStorage *DbomStorage) BatchCreateOccurrences(ctx context.Context, projectID string, userID string, occurrences []*pb.Occurrence) ([]*pb.Occurrence, []error) {
	errs := []error{}
	created := []*pb.Occurrence{}
	for _, o := range occurrences {
		occ, err := dbomStorage.CreateOccurrence(ctx, projectID, userID, o)
		if err != nil {
			errs = append(errs, err)
			continue
		} else {
			created = append(created, occ)
		}
	}
	return created, errs
}

// UpdateOccurrence updates the occurrence asset on the configured dbom channel
func (dbomStorage *DbomStorage) UpdateOccurrence(ctx context.Context, projectID, occurrenceID string, o *pb.Occurrence, mask *fieldmaskpb.FieldMask) (*pb.Occurrence, error) {
	occurrenceName := createOccurrenceID(projectID, occurrenceID)
	occurrenceAsset, err := dbomStorage.getAsset(ctx, occurrenceName)
	occurrence := pb.Occurrence{}
	meta, _ := json.Marshal(occurrenceAsset.AssetMetadata)
	json.Unmarshal(meta, &occurrence)
	if o.CreateTime == nil {
		o.CreateTime = occurrence.CreateTime
	}
	o.UpdateTime = ptypes.TimestampNow()
	occurrenceAsset.AssetMetadata = o
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	params := asset.NewUpdateAssetParams()
	params.SetAssetID(occurrenceName)
	params.SetChannelID(dbomStorage.config.Channel)
	params.SetRepoID(dbomStorage.config.Repo)
	assetBody := occurrenceAsset
	params.SetBody(assetBody)
	_, err = dbomStorage.gateway.Asset.UpdateAsset(params)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	return o, nil
}

// DeleteOccurrence is not supported by dbom
func (dbomStorage *DbomStorage) DeleteOccurrence(ctx context.Context, projectID, occurrenceID string) error {
	return errors.New(notImplemented)
}

// GetNote returns the note asset from the configured dbom channel
func (dbomStorage *DbomStorage) GetNote(ctx context.Context, projectID, noteID string) (*pb.Note, error) {
	payload, err := dbomStorage.getAsset(ctx, createNoteID(projectID, noteID))
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	note := pb.Note{}
	meta, _ := json.Marshal(payload.AssetMetadata)
	json.Unmarshal(meta, &note)
	return &note, nil
}

// ListNotes returns the note assets from the configured dbom channel
func (dbomStorage *DbomStorage) ListNotes(ctx context.Context, projectID, filter, pageToken string, pageSize int32) ([]*pb.Note, string, error) {
	params := asset.NewQueryAssetsPostParams()
	params.SetChannelID(dbomStorage.config.Channel)
	params.SetRepoID(dbomStorage.config.Repo)
	query := models.QueryAssetDefinition{}
	projectQuery := createEqualQuery(parentAssetID, createProjectID(projectID))
	noteQuery := createEqualQuery(assetSubType, notePath)
	query.Query = createAndQuery(projectQuery, noteQuery)
	params.SetBody(&query)
	res, err := dbomStorage.gateway.Asset.QueryAssetsPost(params)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, "", err
	}
	payload := res.Payload.(map[string]interface{})
	notes := []*pb.Note{}
	for _, element := range payload {
		note := pb.Note{}
		meta, _ := json.Marshal(element.(map[string]interface{})[assetMetadata])
		json.Unmarshal(meta, &note)
		notes = append(notes, &note)
	}
	return notes, "", nil
}

// CreateNote writes the note asset to the configured dbom channel
func (dbomStorage *DbomStorage) CreateNote(ctx context.Context, projectID, noteID, userID string, note *pb.Note) (*pb.Note, error) {
	_, err := dbomStorage.GetProject(ctx, projectID)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	if note.CreateTime == nil {
		note.CreateTime = ptypes.TimestampNow()
	}
	if note.UpdateTime == nil {
		note.UpdateTime = note.CreateTime
	}
	params := asset.NewAddAssetParams()
	projectName := createProjectID(projectID)
	noteName := createNoteID(projectID, noteID)
	params.SetAssetID(noteName)
	params.SetChannelID(dbomStorage.config.Channel)
	params.SetRepoID(dbomStorage.config.Repo)
	assetBody := models.NewAsset(note, &note.ShortDescription, &noteName, &notePath)
	params.SetBody(assetBody)
	_, err = dbomStorage.gateway.Asset.AddAsset(params)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	attachParams := asset.NewAttachSubAssetParams()
	attachParams.SetAssetID(projectName)
	attachParams.SetChannelID(dbomStorage.config.Channel)
	attachParams.SetRepoID(dbomStorage.config.Repo)
	attachBody := models.NewAttach(&dbomStorage.config.Repo, &dbomStorage.config.Channel, &noteName, &notePath, &notePath)
	attachParams.SetBody(attachBody)
	_, err = dbomStorage.gateway.Asset.AttachSubAsset(attachParams)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}

	return note, nil
}

// BatchCreateNotes writes the note assets to the configured dbom channel
func (dbomStorage *DbomStorage) BatchCreateNotes(ctx context.Context, projectID, userID string, notesWithNoteIDs map[string]*pb.Note) ([]*pb.Note, []error) {
	errs := []error{}
	created := []*pb.Note{}
	for id, o := range notesWithNoteIDs {
		note, err := dbomStorage.CreateNote(ctx, projectID, id, userID, o)
		if err != nil {
			errs = append(errs, err)
			continue
		} else {
			created = append(created, note)
		}
	}
	return created, errs
}

// UpdateNote updates the note asset on the configured dbom channel
func (dbomStorage *DbomStorage) UpdateNote(ctx context.Context, projectID, noteID string, note *pb.Note, mask *fieldmaskpb.FieldMask) (*pb.Note, error) {
	noteName := createNoteID(projectID, noteID)
	noteAsset, err := dbomStorage.getAsset(ctx, noteName)
	noteUpdate := pb.Note{}
	meta, _ := json.Marshal(noteAsset.AssetMetadata)
	json.Unmarshal(meta, &noteUpdate)
	if note.CreateTime == nil {
		note.CreateTime = noteUpdate.CreateTime
	}
	note.UpdateTime = ptypes.TimestampNow()
	noteAsset.AssetMetadata = note
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	params := asset.NewUpdateAssetParams()
	params.SetAssetID(noteName)
	params.SetChannelID(dbomStorage.config.Channel)
	params.SetRepoID(dbomStorage.config.Repo)
	assetBody := noteAsset
	params.SetBody(assetBody)
	_, err = dbomStorage.gateway.Asset.UpdateAsset(params)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	return note, nil
}

// DeleteNote is not supported by dbom
func (dbomStorage *DbomStorage) DeleteNote(ctx context.Context, projectID, noteID string) error {
	return errors.New(notImplemented)
}

// GetOccurrenceNote returns the note asset from the configured dbom channel for the given occurrence
func (dbomStorage *DbomStorage) GetOccurrenceNote(ctx context.Context, projectID, occurrenceID string) (*pb.Note, error) {
	occurrence, err := dbomStorage.GetOccurrence(ctx, projectID, occurrenceID)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	payload, err := dbomStorage.getAsset(ctx, convertID(occurrence.GetNoteName()))
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	note := pb.Note{}
	meta, _ := json.Marshal(payload.AssetMetadata)
	json.Unmarshal(meta, &note)
	return &note, nil
}

// ListNoteOccurrences returns the occurrence assets from the configured dbom channel for the given note
func (dbomStorage *DbomStorage) ListNoteOccurrences(ctx context.Context, projectID, noteID, filter, pageToken string, pageSize int32) ([]*pb.Occurrence, string, error) {
	/*noteParams := asset.NewQueryAssetsPostParams()
	noteParams.SetChannelID(dbomStorage.config.Channel)
	noteParams.SetRepoID(dbomStorage.config.Repo)
	noteQuery := models.QueryAssetDefinition{}
	noteID := createNoteID(projectID, noteID)
	noteQuery.Query = createEqualQuery(id, noteID)
	noteParams.SetBody(&noteQuery)
	note, err := dbomStorage.gateway.Asset.QueryAssetsPost(noteParams)
	notePayload := note.Payload.(map[string]interface{})[noteID]
	children := notePayload.(map[string]interface{})[attachedChildren]
	var noteIDs []string
	for _, element := range children.([]interface{}) {
		elementMap := element.(map[string]interface{})
		aID := elementMap[assetID].(string)
		noteIDs = append(noteIDs, aID)
	}
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, "", err
	}
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, "", err
	}
	params := asset.NewQueryAssetsPostParams()
	params.SetChannelID(dbomStorage.config.Channel)
	params.SetRepoID(dbomStorage.config.Repo)
	query := models.QueryAssetDefinition{}
	query.Query = createInQuery(id, noteIDs)
	params.SetBody(&query)
	res, err := dbomStorage.gateway.Asset.QueryAssetsPost(params)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, "", err
	}
	payload := res.Payload.(map[string]interface{})
	occurrences := []*pb.Occurrence{}
	for _, element := range payload {
		occurrence := pb.Occurrence{}
		meta, _ := json.Marshal(element.(map[string]interface{})[assetMetadata])
		json.Unmarshal(meta, &occurrence)
		occurrences = append(occurrences, &occurrence)
	}
	return occurrences, "", nil*/
	params := asset.NewQueryAssetsPostParams()
	params.SetChannelID(dbomStorage.config.Channel)
	params.SetRepoID(dbomStorage.config.Repo)
	query := models.QueryAssetDefinition{}
	query.Query = createEqualQuery(occurrenceNote, createFullNoteID(projectID, noteID))
	params.SetBody(&query)
	res, err := dbomStorage.gateway.Asset.QueryAssetsPost(params)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, "", err
	}
	payload := res.Payload.(map[string]interface{})
	occurrences := []*pb.Occurrence{}
	for _, element := range payload {
		occurrence := pb.Occurrence{}
		meta, _ := json.Marshal(element.(map[string]interface{})[assetMetadata])
		json.Unmarshal(meta, &occurrence)
		occurrences = append(occurrences, &occurrence)
	}
	return occurrences, "", nil
}

// GetVulnerabilityOccurrencesSummary is not implemented
func (dbomStorage *DbomStorage) GetVulnerabilityOccurrencesSummary(ctx context.Context, projectID, filter string) (*pb.VulnerabilityOccurrencesSummary, error) {
	return nil, errors.New(notImplemented)
}

func (dbomStorage *DbomStorage) getAsset(ctx context.Context, assetID string) (*models.AssetDefinition, error) {
	params := asset.NewRetrieveAssetParams()
	params.SetChannelID(dbomStorage.config.Channel)
	params.SetRepoID(dbomStorage.config.Repo)
	params.SetAssetID(assetID)
	res, err := dbomStorage.gateway.Asset.RetrieveAsset(params)
	if err != nil {
		dbomStorage.logger.Err(err).Msg(err.Error())
		return nil, err
	}
	return res.Payload, nil
}

// convertID removes slashes from an id
func convertID(id string) string {
	return strings.Replace(id, slash, underscore, -1)
}

// createProjectID creates an asset id for a project
func createProjectID(projectID string) string {
	return projectPath + underscore + projectID
}

// createFullNoteID creates the full project id for a note in a project
func createFullNoteID(projectID, noteID string) string {
	return projectPath + slash + projectID + slash + notePath + slash + noteID
}

// createNoteID creates an asset id for a note in a project
func createNoteID(projectID, noteID string) string {
	return projectPath + underscore + projectID + underscore + notePath + underscore + noteID
}

// createOccurrenceID creates an asset id for a occurrence in a project
func createOccurrenceID(projectID, occurrenceID string) string {
	return projectPath + underscore + projectID + underscore + occurrencePath + underscore + occurrenceID
}

// createEqualQuery creates a dbom equal query for a given field and value
func createEqualQuery(field, value string) map[string]interface{} {
	return map[string]interface{}{field: map[string]interface{}{
		eq: value,
	}}
}

// createInQuery creates a dbom in query for a given field and value
func createInQuery(field string, value []string) map[string]interface{} {
	return map[string]interface{}{field: map[string]interface{}{
		in: value,
	}}
}

// createAndQuery creates a dbom and query for a list of queries
func createAndQuery(queries ...map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		and: queries,
	}
}
