package engine

import "math"

// длинна вектора, квадррат суммы элементов вектора
func Magnitude(vec vector) float32 {
	var sumSquares float64
	for _, elem := range vec {
		sumSquares += float64(elem * elem)
	}
	if math.Abs(sumSquares-1.0) <= 1e-5 {
		return 1.0
	}
	return float32(math.Sqrt(sumSquares))
}

//функция нормализации вектора
func VectorNormalize(vec vector) vector {
	magnitude := Magnitude(vec)
	resVec := make(vector, len(vec))
	if magnitude == 1.0 {
		return vec
	}
	for ind, elem := range vec {
		resVec[ind] = elem / magnitude
	}

	return resVec
}

//поиск совпадений засчет скалярного произведения, но тк векторы нормализованы, то знаминатель всегда равен 1
// длина векторов на входе строго равна между собой
func CalcCosineScore(vec1, vec2 vector) float32 {
	var score float32

	//хак для компилятора, чтобы он убедился что массивы имеют нужную длинну
	length := len(vec1) - 1
	_ = vec1[length]
	_ = vec2[length]

	for i := 0; i < len(vec1); i++ {
		score += vec1[i] * vec2[i]
	}
	return score
}

//иная метрика через евклидово расстояние, чем ближе к нулю, чем больше совпадение
func CalcL2Score(vec1, vec2 vector) float32 {
	return 2 - 2*CalcCosineScore(vec1, vec2)
}
