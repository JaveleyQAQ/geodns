package config

// 颜色常量
const (
	ColorReset  = "\033[0m"
	ColorPurple = "\033[35m"
	ColorGreen  = "\033[32m"
	ColorBlue   = "\033[34m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"
	ColorRed    = "\033[31m"
	ColorWhite  = "\033[37m"
	ColorOrange = "\033[33;31m"

	// 新增颜色常量
	ColorBrightRed    = "\033[91m"
	ColorBrightGreen  = "\033[92m"
	ColorBrightYellow = "\033[93m"
	ColorBrightBlue   = "\033[94m"
	ColorBrightPurple = "\033[95m"
	ColorBrightCyan   = "\033[96m"
	ColorBrightWhite  = "\033[97m"
	ColorMagenta      = "\033[35m"
	ColorLightBlue    = "\033[94m"
	ColorLightGreen   = "\033[92m"
	ColorLightRed     = "\033[91m"
	ColorLightYellow  = "\033[93m"
	ColorLightCyan    = "\033[96m"
	ColorLightPurple  = "\033[95m"
)

// DNS模式常量
const (
	ModeVercel     = "vercel"
	ModeCloudflare = "cloudflare"
)
