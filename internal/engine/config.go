package engine

//режим записи данных на диск, по умолчанию 0, при асинке есть риск потерять данные за секунду до отключения сервака с бд
const (
	asyncWalMode = 0
	syncWalMode  = 1
)

type DatabaseConfig struct {
	Name    string
	Port    uint32
	WalMode uint8
}
