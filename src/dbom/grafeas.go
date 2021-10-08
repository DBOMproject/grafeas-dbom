/*
 * Copyright 2021 Unisys Corporation
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
package dbom

import (
	"fmt"

	"dbom-grafeas/dbom/config"

	grafeasConfig "github.com/grafeas/grafeas/go/config"
	"github.com/grafeas/grafeas/go/v1beta1/storage"
	"github.com/rs/zerolog"
)

type newDbomStorageFunc func(*config.DBoMConfig) (*DbomStorage, error)

type registerStorageTypeProviderFunc func(string, *grafeasConfig.StorageConfiguration) (*storage.Storage, error)

// DbomTypeProviderCreator returns a function for grafeas to register the dbom storage
func DbomTypeProviderCreator(newDBoM newDbomStorageFunc, logger *zerolog.Logger) registerStorageTypeProviderFunc {
	return func(storageType string, storageConfiguration *grafeasConfig.StorageConfiguration) (*storage.Storage, error) {
		var config *config.DBoMConfig

		logger.Info().Msg("registering dbom storage provider")

		if storageType != "dbom" {
			return nil, fmt.Errorf("unknown storage type %s, must be 'dbom'", storageType)
		}

		err := grafeasConfig.ConvertGenericConfigToSpecificType(storageConfiguration, &config)
		if err != nil {
			return nil, fmt.Errorf("unable to convert config for dbom: %s", err)
		}

		dbom, err := newDBoM(config)
		if err != nil {
			return nil, err
		}

		return &storage.Storage{
			Ps: dbom,
			Gs: dbom,
		}, nil
	}
}
