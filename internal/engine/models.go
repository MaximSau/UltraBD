package engine

import "sync"

// константные типы данных для декода из ячейки
const (
	ByteContentType   = 0
	StringContentType = 1
	PhotoContentType  = 2
	JsonContentType   = 3
)

//тип метрик для поиска сходства между векторами
const (
	MethricTypeCosine = 0
	MethricTypeL2     = 1
)

//структура вектора, срез чисел с плавающей точкой,как правило от -1 до 1, бывают искл.
type vector []float32

//ВСЕ СТРУКТУРЫ ОТСОРТИРОВАННЫ В УБЫАЮЩЕМ ПОРЯДКЕ,ТК КОМПИЛЯТОРУ ЛЕГЧЕ РАБОТАТЬ С ТАКИМ ФОРМАТОМ!!!

// структура ячейки коллекции
type Record struct {
	Vector      vector         //вектор n-мерности, обычно соответствует данным, позволяет быстро искать совпадения, будет нормализован
	Content     []byte         //сами данные
	Meta        map[string]any // (опционально) метаданные для фильтрации
	Id          string         // id ячейки, НЕ ПУТАТЬ С ИНДЕКСОМ ЯЧЕЙКИ! поступает во время запроса на создание ячейки
	ContentType uint8          //тип данных (см выше)
}

// структура коллекции
type Collection struct {
	Mu          sync.RWMutex // мьютекс для избежания конфликтов записи и чтения
	Records     []Record     // слайс записей
	Name        string
	VectorSize  uint16 // мерность векторов для коллекции, в коллекцию можно добавлять данные только той мерности, которая указана в этом параметре
	MethricType uint8  //тип метрики совпадения векторов
}

type Database struct {
	Collections []Collection
	Config      *DatabaseConfig //  конфиг для бд с данными
}
