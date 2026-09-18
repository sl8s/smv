package smvgo

import (
	"fmt"
)

type IModels[T IModel] interface {
	Models() []T
	SetModels([]T)
	Clone() IModels[T]
	ToMaps() []map[string]any
}

func AddFromIModelsAndNewModel[T IModel](iModels IModels[T], newModel T) error {
	idByNewModel := newModel.Id()
	models := iModels.Models()
	for _, model := range models {
		if model.Id() == idByNewModel {
			return NewLocalError(
				"IModels",
				DeveloperByEnumGuilty(),
				fmt.Sprintf("Duplicate found in the model array based on the new model's ID: %#v -- %s", models, idByNewModel),
			)
		}
	}
	iModels.SetModels(append(models, newModel))
	return nil
}

func UpdateFromIModelsAndNewModelById[T IModel](iModels IModels[T], newModel T) error {
	models := iModels.Models()
	if len(models) <= 0 {
		return NewLocalError(
			"IModels",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("You cannot update an element in an array if the array is empty: %#v", models),
		)
	}
	hasUpdate := false
	idByNewModel := newModel.Id()
	newModels := make([]T, 0, len(models))
	for _, model := range models {
		if model.Id() != idByNewModel {
			newModels = append(newModels, model)
			continue
		}
		newModels = append(newModels, newModel)
		if hasUpdate {
			continue
		}
		hasUpdate = true
	}
	if !hasUpdate {
		return NewLocalError(
			"IModels",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("There are no updated elements in the array: %#v -- %#v", newModels, newModel),
		)
	}
	iModels.SetModels(newModels)
	return nil
}

func DeleteFromIModelsAndIdById[T IModel](iModels IModels[T], id string) error {
	models := iModels.Models()
	if len(models) <= 0 {
		return NewLocalError(
			"IModels",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("You cannot delete an element in an array if the array is empty: %#v", models),
		)
	}
	newModels := make([]T, 0)
	for _, model := range models {
		if model.Id() == id {
			continue
		}
		newModels = append(newModels, model)
	}
	if len(newModels) >= len(models) {
		return NewLocalError(
			"IModels",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("There are no deleted elements in the array.: %#v -- %s", newModels, id),
		)
	}
	iModels.SetModels(newModels)
	return nil
}

func AddFromIModelsAndNewModels[T IModel](iModels IModels[T], newModels []T) error {
	if len(newModels) <= 0 {
		return NewLocalError(
			"IModels",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("You cannot add an empty array: %#v", newModels),
		)
	}
	newModelsMap := make(map[string]T, len(newModels))
	for _, newModel := range newModels {
		idByNewModel := newModel.Id()
		if _, existsNewModel := newModelsMap[idByNewModel]; existsNewModel {
			return NewLocalError(
				"IModels",
				DeveloperByEnumGuilty(),
				fmt.Sprintf("There are duplicates in the new array of models: %#v -- %s", newModels, idByNewModel),
			)
		}
		newModelsMap[idByNewModel] = newModel
	}
	models := iModels.Models()
	for _, model := range models {
		idByModel := model.Id()
		if _, existsNewModel := newModelsMap[idByModel]; existsNewModel {
			return NewLocalError(
				"IModels",
				DeveloperByEnumGuilty(),
				fmt.Sprintf("Duplicate found in the model array based on the new model's ID: %#v -- %s", models, idByModel),
			)
		}
	}
	iModels.SetModels(append(models, newModels...))
	return nil
}

func UpdateFromIModelsAndNewModelsById[T IModel](iModels IModels[T], newModels []T) error {
	models := iModels.Models()
	if len(models) <= 0 || len(newModels) <= 0 {
		return NewLocalError(
			"IModels",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("You cannot update an empty array or add an empty array for an update: %#v -- %#v", models, newModels),
		)
	}
	newModelsMap := make(map[string]T, len(newModels))
	for _, newModel := range newModels {
		idByNewModel := newModel.Id()
		if _, existsNewModel := newModelsMap[idByNewModel]; existsNewModel {
			return NewLocalError(
				"IModels",
				DeveloperByEnumGuilty(),
				fmt.Sprintf("There are duplicates in the new array of models: %#v -- %s", newModels, idByNewModel),
			)
		}
		newModelsMap[idByNewModel] = newModel
	}
	hasUpdate := false
	secondNewModels := make([]T, 0, len(models))
	for _, model := range models {
		newModel, existsNewModel := newModelsMap[model.Id()]
		if !existsNewModel {
			secondNewModels = append(secondNewModels, model)
			continue
		}
		secondNewModels = append(secondNewModels, newModel)
		if hasUpdate {
			continue
		}
		hasUpdate = true
	}
	if !hasUpdate {
		return NewLocalError(
			"IModels",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("There are no updated elements in the array: %#v -- %#v", secondNewModels, newModels),
		)
	}
	iModels.SetModels(secondNewModels)
	return nil
}

func DeleteFromIModelsAndIdsById[T IModel](iModels IModels[T], ids []string) error {
	models := iModels.Models()
	if len(models) <= 0 || len(ids) <= 0 {
		return NewLocalError(
			"IModels",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("You cannot delete an empty array or add an empty array for an delete: %#v -- %#v", models, ids),
		)
	}
	structsMap := make(map[string]struct{}, len(ids))
	for _, id := range ids {
		if _, existsStruct := structsMap[id]; existsStruct {
			return NewLocalError(
				"IModels",
				DeveloperByEnumGuilty(),
				fmt.Sprintf("There are duplicates in the array of ids: %#v -- %s", ids, id),
			)
		}
		structsMap[id] = struct{}{}
	}
	newModels := make([]T, 0)
	for _, model := range models {
		if _, existsStruct := structsMap[model.Id()]; existsStruct {
			continue
		}
		newModels = append(newModels, model)
	}
	if len(newModels) >= len(models) {
		return NewLocalError(
			"IModels",
			DeveloperByEnumGuilty(),
			fmt.Sprintf("There are no deleted elements in the array: %#v -- %#v", newModels, ids),
		)
	}
	iModels.SetModels(newModels)
	return nil
}
