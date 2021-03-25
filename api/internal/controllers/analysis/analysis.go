// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package analysis

import (
	"github.com/google/uuid"

	"github.com/ZupIT/horusec-devkit/pkg/entities/analysis"
	"github.com/ZupIT/horusec-devkit/pkg/entities/cli"
	brokerService "github.com/ZupIT/horusec-devkit/pkg/services/broker"
	brokerConfig "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database"
)

type IController interface {
	SaveAnalysis(analysisData *cli.AnalysisData) (uuid.UUID, error)
	GetAnalysis(analysisID uuid.UUID) (*analysis.Analysis, error)
}

type Controller struct {
	broker        brokerService.IBroker
	databaseRead  database.IDatabaseRead
	databaseWrite database.IDatabaseWrite
	brokerConfig  brokerConfig.IConfig
}

func NewAnalysisController(broker brokerService.IBroker, brokerConfiguration brokerConfig.IConfig,
	databaseConnection *database.Connection) IController {
	return &Controller{
		broker:        broker,
		brokerConfig:  brokerConfiguration,
		databaseRead:  databaseConnection.Read,
		databaseWrite: databaseConnection.Write,
	}
}

func (c *Controller) SaveAnalysis(analysisData *cli.AnalysisData) (uuid.UUID, error) {
	panic("implement me")
}

func (c *Controller) GetAnalysis(analysisID uuid.UUID) (*analysis.Analysis, error) {
	panic("implement me")
}