package inverter

type Inverter interface {
	IsExportEnabled() (bool, error)
	ExportEnable() error
	ExportDisable() error
}
