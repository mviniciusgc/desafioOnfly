package enum

// Definição dos valores enum usando iota
type LogType int

// Definindo valores para LogType usando iota
const (
	Info LogType = iota
	Warning
	Error
	Debug
)

// String retorna uma representação legível de LogType
func (l LogType) String() string {
	return [...]string{"INFO", "WARNING", "ERROR", "DEBUG"}[l]
}
