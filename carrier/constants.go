package carrier

type Currency string

var (
	CurrencyTRY Currency = "TRY"
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	CurrencyGBP Currency = "GBP"
)

type Incoterm string

var (
	IncotermDDP Incoterm = "DDP"
	IncotermDDU Incoterm = "DDU"
)

type SignatureConfirmation string

var (
	SignatureConfirmationStandard SignatureConfirmation = "Standard"
	SignatureConfirmationAdult    SignatureConfirmation = "Adult"
)

type PType string

var (
	PaymentTypeSender    PType = "Sender"
	PaymentTypeRecipient PType = "Recipient"
	PaymentTypeCODCredit PType = "CODCredit"
	PaymentTypeCODCash   PType = "CODCash"
)

type DimensionUnit string

var (
	PackageDimensionCM DimensionUnit = "cm"
	PackageDimensionIn DimensionUnit = "in"
)

type WeightUnit string

var (
	PackageWeightKg WeightUnit = "kg"
	PackageWeightLb WeightUnit = "lb"
)

type LabelFormat string

var LabelFormatZPL LabelFormat = "zpl"
var LabelFormatPNG LabelFormat = "png"
var LabelFormatPDF LabelFormat = "pdf"

type LabelSize string

var LabelSizeThermal4x4 LabelSize = "thermal4x4"
var LabelSizeThermal6x4 LabelSize = "thermal6x4"
var LabelSizeNormalA4 LabelSize = "normalA4"
var LabelSizeNormalA5 LabelSize = "normalA5"
var LabelSizeNormalA6 LabelSize = "normalA6"

const DefaultLabelSize = "normalA4"

type Identifier string

var IdentifierShipmentID Identifier = "shipment_id"
var IdentifierTrackingID Identifier = "tracking_id"
var IdentifierCarrierShipmentID Identifier = "carrier_shipment_id"
