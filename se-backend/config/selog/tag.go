package selog

import "go.uber.org/zap"

const DefaultTagKey = "tag"

// ServerTag indica falhas relacionadas ao servidor.
var ServerTag = zap.String(DefaultTagKey, "SERVER")

// MigrationTag indica falhas na execução do versionamento do banco de dados.
var MigrationTag = zap.String(DefaultTagKey, "MIGRATION")
