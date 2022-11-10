package MarsEngine

// Конфиг переменные

var (
	winWidth  = 600
	winHeight = 600
	winName   = "Window"
)

// Объекты движка

var (
	allObjects     = []gameObject{}
	activeQueue    = []gameObject{}
	collisionGroup = []string{}
)
