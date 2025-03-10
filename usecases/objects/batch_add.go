//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package objects

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/semi-technologies/weaviate/entities/errorcompounder"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/usecases/objects/validation"
)

// AddObjects Class Instances in batch to the connected DB
func (b *BatchManager) AddObjects(ctx context.Context, principal *models.Principal,
	objects []*models.Object, fields []*string,
) (BatchObjects, error) {
	err := b.authorizer.Authorize(principal, "create", "batch/objects")
	if err != nil {
		return nil, err
	}

	unlock, err := b.locks.LockConnector()
	if err != nil {
		return nil, NewErrInternal("could not acquire lock: %v", err)
	}
	defer unlock()

	before := time.Now()
	b.metrics.BatchInc()
	defer b.metrics.BatchOp("total_uc_level", before.UnixNano())
	defer b.metrics.BatchDec()

	return b.addObjects(ctx, principal, objects, fields)
}

func (b *BatchManager) addObjects(ctx context.Context, principal *models.Principal,
	classes []*models.Object, fields []*string,
) (BatchObjects, error) {
	beforePreProcessing := time.Now()
	if err := b.validateObjectForm(classes); err != nil {
		return nil, NewErrInvalidUserInput("invalid param 'objects': %v", err)
	}

	batchObjects := b.validateObjectsConcurrently(ctx, principal, classes, fields)
	b.metrics.BatchOp("total_preprocessing", beforePreProcessing.UnixNano())

	var (
		res BatchObjects
		err error
	)

	beforePersistence := time.Now()
	defer b.metrics.BatchOp("total_persistence_level", beforePersistence.UnixNano())
	if res, err = b.vectorRepo.BatchPutObjects(ctx, batchObjects); err != nil {
		return nil, NewErrInternal("batch objects: %#v", err)
	}

	return res, nil
}

func (b *BatchManager) validateObjectForm(classes []*models.Object) error {
	if len(classes) == 0 {
		return fmt.Errorf("cannot be empty, need at least one object for batching")
	}

	return nil
}

func (b *BatchManager) validateObjectsConcurrently(ctx context.Context, principal *models.Principal,
	classes []*models.Object, fields []*string,
) BatchObjects {
	fieldsToKeep := determineResponseFields(fields)
	c := make(chan BatchObject, len(classes))

	wg := new(sync.WaitGroup)

	// Generate a goroutine for each separate request
	for i, object := range classes {
		wg.Add(1)
		go b.validateObject(ctx, principal, wg, object, i, &c, fieldsToKeep)
	}

	wg.Wait()
	close(c)
	return objectsChanToSlice(c)
}

func (b *BatchManager) validateObject(ctx context.Context, principal *models.Principal,
	wg *sync.WaitGroup, concept *models.Object, originalIndex int, resultsC *chan BatchObject,
	fieldsToKeep map[string]struct{},
) {
	defer wg.Done()

	var id strfmt.UUID

	ec := &errorcompounder.ErrorCompounder{}

	// Auto Schema
	err := b.autoSchemaManager.autoSchema(ctx, principal, concept)
	ec.Add(err)

	if concept.ID == "" {
		// Generate UUID for the new object
		uid, err := generateUUID()
		id = uid
		ec.Add(err)
	} else {
		if _, err := uuid.Parse(concept.ID.String()); err != nil {
			ec.Add(err)
		}
		id = concept.ID
	}

	// Create Action object
	object := &models.Object{}
	object.LastUpdateTimeUnix = 0
	object.ID = id
	object.Vector = concept.Vector

	if _, ok := fieldsToKeep["class"]; ok {
		object.Class = concept.Class
	}
	if _, ok := fieldsToKeep["properties"]; ok {
		object.Properties = concept.Properties
	}

	if object.Properties == nil {
		object.Properties = map[string]interface{}{}
	}
	now := unixNow()
	if _, ok := fieldsToKeep["creationTimeUnix"]; ok {
		object.CreationTimeUnix = now
	}
	if _, ok := fieldsToKeep["lastUpdateTimeUnix"]; ok {
		object.LastUpdateTimeUnix = now
	}
	class, err := b.schemaManager.GetClass(ctx, principal, object.Class)
	ec.Add(err)
	if class == nil {
		ec.Add(fmt.Errorf("class '%s' not present in schema", object.Class))
	} else {
		// not possible without the class being present
		err = validation.New(b.vectorRepo.Exists, b.config).Object(ctx, object, class)
		ec.Add(err)

		err = b.modulesProvider.UpdateVector(ctx, object, class, nil, b.findObject, b.logger)
		ec.Add(err)

	}

	*resultsC <- BatchObject{
		UUID:          id,
		Object:        object,
		Err:           ec.ToError(),
		OriginalIndex: originalIndex,
		Vector:        object.Vector,
	}
}

func objectsChanToSlice(c chan BatchObject) BatchObjects {
	result := make([]BatchObject, len(c))
	for object := range c {
		result[object.OriginalIndex] = object
	}

	return result
}

func unixNow() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
