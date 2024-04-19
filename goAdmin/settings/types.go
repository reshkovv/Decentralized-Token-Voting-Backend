package settings

type Config struct {
	MumbaiRPC       string `yaml:"mumbaiRPC"`
	EthereumRPC     string `yaml:"ethereumRPC"`
	PolygonRPC      string `yaml:"polygonRPC"`
	SabrechainRPC   string `yaml:"sabrechainRPC"`
	SepoliaRPC      string `yaml:"sepoliaRPC"`
	MumbaiChainID   int    `yaml:"mumbaiChainID"`
	SepoliaChainID  int    `yaml:"sepoliaChainID"`
	EthereumChainID int    `yaml:"ethereumChainID"`
	PolygonChainID  int    `yaml:"polygonChainID"`
	AdminPrivateKey string `yaml:"adminPrivateKey"`
	AdminAddress    string `yaml:"adminAddress"`
	GovernorAddress string `yaml:"governorAddress"`
	GovTokenAddress string `yaml:"govTokenAddress"`
	TimeLockAddress string `yaml:"timeLockAddress"`
	ProposalID      string `yaml:"proposalID"`
}
