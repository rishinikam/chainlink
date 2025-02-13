package orm

import (
	"math/big"
	"net/url"
	"time"

	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/store/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/contrib/sessions"
)

// ConfigReader represents just the read side of the config
type ConfigReader interface {
	AllowOrigins() string
	BlockBackfillDepth() uint64
	BridgeResponseURL() *url.URL
	CertFile() string
	ChainID() *big.Int
	ClientNodeURL() string
	CreateProductionLogger() *logger.Logger
	DatabaseMaximumTxDuration() time.Duration
	DatabaseTimeout() models.Duration
	DatabaseURL() url.URL
	DefaultHTTPAllowUnrestrictedNetworkAccess() bool
	DefaultHTTPLimit() int64
	DefaultHTTPTimeout() models.Duration
	DefaultMaxHTTPAttempts() uint
	Dev() bool
	EnableExperimentalAdapters() bool
	EthBalanceMonitorBlockDelay() uint16
	EthFinalityDepth() uint
	EthGasBumpPercent() uint16
	EthGasBumpThreshold() uint64
	EthGasBumpTxDepth() uint16
	EthGasBumpWei() *big.Int
	EthGasLimitDefault() uint64
	EthGasPriceDefault() *big.Int
	EthHeadTrackerHistoryDepth() uint
	EthHeadTrackerMaxBufferSize() uint
	EthLogBackfillBatchSize() uint32
	EthMaxGasPriceWei() *big.Int
	EthNonceAutoSync() bool
	EthRPCDefaultBatchSize() uint32
	EthTxResendAfterThreshold() time.Duration
	EthereumSecondaryURLs() []url.URL
	EthereumURL() string
	ExplorerAccessKey() string
	ExplorerSecret() string
	ExplorerURL() *url.URL
	FeatureExternalInitiators() bool
	FeatureFluxMonitor() bool
	FeatureOffchainReporting() bool
	GasUpdaterBlockDelay() uint16
	GasUpdaterBlockHistorySize() uint16
	GasUpdaterTransactionPercentile() uint16
	JSONConsole() bool
	KeyFile() string
	KeysDir() string
	LinkContractAddress() string
	LogLevel() LogLevel
	LogSQLStatements() bool
	LogToDisk() bool
	MaximumServiceDuration() models.Duration
	MigrateDatabase() bool
	MinIncomingConfirmations() uint32
	MinRequiredOutgoingConfirmations() uint64
	MinimumContractPayment() *assets.Link
	MinimumRequestExpiration() uint64
	MinimumServiceDuration() models.Duration
	OCRTraceLogging() bool
	OperatorContractAddress() common.Address
	OptimismGasFees() bool
	Port() uint16
	ReaperExpiration() models.Duration
	RootDir() string
	SecureCookies() bool
	SessionOptions() sessions.Options
	SessionSecret() ([]byte, error)
	SessionTimeout() models.Duration
	SetEthGasPriceDefault(value *big.Int) error
	TLSCertPath() string
	TLSHost() string
	TLSKeyPath() string
	TLSPort() uint16
	TLSRedirect() bool
	TriggerFallbackDBPollInterval() time.Duration
	tlsDir() string
}
