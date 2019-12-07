package adapters

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang/glog"
	"github.com/seadiaz/adoption/src/server/details/adapters/usecases/entities"
)

const persistenceTypeTool = "tools"

// ToolRepository ...
type ToolRepository struct {
	persistence Persistence
}

type persistedTool struct {
	ID     string
	Name   string
	Labels []persistedLabel
}

type persistedLabel struct {
	Kind  string
	Value string
}

func createPersistedToolFromTool(entity *entities.Tool) *persistedTool {
	return &persistedTool{
		ID:     entity.ID.String(),
		Name:   entity.Name,
		Labels: createPersistedLabelListFromLabelList(entity.Labels),
	}
}

func createToolFromPersistedTool(pEntity *persistedTool) *entities.Tool {
	return &entities.Tool{
		ID:     entities.RecoverID(pEntity.ID),
		Name:   pEntity.Name,
		Labels: createLabelListFromPersistedLabelList(pEntity.Labels),
	}
}

func createPersistedToolListFromToolList(list []*entities.Tool) []*persistedTool {
	output := make([]*persistedTool, 0, 0)
	for _, item := range list {
		entity := createPersistedToolFromTool(item)
		output = append(output, entity)
	}

	return output
}

func createToolListFromPersistedToolList(list []*persistedTool) []*entities.Tool {
	output := make([]*entities.Tool, 0, 0)
	for _, item := range list {
		entity := createToolFromPersistedTool(item)
		output = append(output, entity)
	}

	return output
}

func createPersistedLabelFromLabel(entity *entities.Label) persistedLabel {
	return persistedLabel{
		Kind:  entity.Kind,
		Value: entity.Value,
	}
}

func createPersistedLabelListFromLabelList(list []*entities.Label) []persistedLabel {
	output := make([]persistedLabel, 0, 0)
	for _, item := range list {
		entity := createPersistedLabelFromLabel(item)
		output = append(output, entity)
	}

	return output
}

func createLabelFromPersistedLabel(pEntity persistedLabel) *entities.Label {
	return &entities.Label{
		Kind:  pEntity.Kind,
		Value: pEntity.Value,
	}
}

func createLabelListFromPersistedLabelList(list []persistedLabel) []*entities.Label {
	output := make([]*entities.Label, 0, 0)
	for _, item := range list {
		entity := createLabelFromPersistedLabel(item)
		output = append(output, entity)
	}

	return output
}

func (t *persistedTool) MarshalBinary() (data []byte, err error) {
	return json.Marshal(t)
}

func (t *persistedTool) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &t); err != nil {
		glog.Error(err)
		return err
	}

	return nil
}

// Clone ...
func (t *persistedTool) Clone() PersistedData {
	return &persistedTool{}
}

// CreateToolRepository ...
func CreateToolRepository(persistence Persistence) *ToolRepository {
	return &ToolRepository{
		persistence: persistence,
	}
}

// GetAllTools ...
func (r *ToolRepository) GetAllTools() ([]*entities.Tool, error) {
	output := make([]*entities.Tool, 0, 0)
	proto := &persistedTool{}
	items, _ := r.persistence.GetAll(persistenceTypeTool, proto)
	for _, item := range items {
		entity := createToolFromPersistedTool(item.(*persistedTool))
		output = append(output, entity)
	}

	return output, nil
}

// SaveTool ...
func (r *ToolRepository) SaveTool(entity *entities.Tool) (*entities.Tool, error) {
	if entity.ID.String() == "" {
		return nil, fmt.Errorf("ID is missing for tool %s", entity.Name)
	}
	pTool := createPersistedToolFromTool(entity)
	if err := r.persistence.Create(persistenceTypeTool, entity.ID.String(), pTool); err != nil {
		glog.Error(err)
		return nil, err
	}

	return entity, nil
}

// FindToolByID ...
func (r *ToolRepository) FindToolByID(id string) (*entities.Tool, error) {
	proto := &persistedTool{}
	pTool, err := r.persistence.Find(persistenceTypeTool, id, proto)
	if err != nil {
		glog.Warning(err)
		return nil, errors.New("error getting tool")
	}
	if pTool == nil {
		return nil, errors.New("tool doesn't exists")
	}

	entity := createToolFromPersistedTool(pTool.(*persistedTool))
	return entity, nil
}
