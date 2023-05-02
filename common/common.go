package common

// load mode
const (
	ModeTest = "test"
	ModeProd = "prod"
	ModeDev  = "dev" // default
)

// sort tag
const (
	Finally = "FINALLY"
	Firstly = "FIRSTLY"
)

// bootloader config
const (
	ConfigDisablePrint = "BOOTLOADER-DISABLE_PRINT"
	ConfigDisableSort  = "BOOTLOADER-DISABLE_SORT"
	ConfigMode         = "BOOTLOADER-MODE"
	ConfigProjectName  = "BOOTLOADER-PROJECT_NAME"
	ConfigProjectRoot  = "BOOTLOADER-PROJECT_ROOT"
)
