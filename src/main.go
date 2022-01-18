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

// Package main contains the main entrypoint function
package main

import (
	"dbom-grafeas/dbom"
	"dbom-grafeas/dbom/config"
	"dbom-grafeas/helpers"

	"github.com/grafeas/grafeas/go/v1beta1/server"
	grafeasStorage "github.com/grafeas/grafeas/go/v1beta1/storage"
)

func main() {
	logger := helpers.GetLogger("Grafeas DBoM Storage")
	registerStorageTypeProvider := dbom.DbomTypeProviderCreator(func(c *config.DBoMConfig) (*dbom.DbomStorage, error) {
		return dbom.NewDbomStorage(&logger, c), nil
	}, &logger)

	err := grafeasStorage.RegisterStorageTypeProvider("dbom", registerStorageTypeProvider)
	if err != nil {
		logger.Err(err).Msg("Error when registering dbom storage")
	} else {
		logger.Info().Msg("Starting Grafeas Server with DBoM Storage")
		err = server.StartGrafeas()
		if err != nil {
			logger.Err(err).Msg("Failed to start Grafeas server...")
		}
	}
}
