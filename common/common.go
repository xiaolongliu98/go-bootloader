package common

// load mode
const (
	ModeTest = "0"
	ModeProd = "1"
	ModeDev  = "2" // default
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
