package helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var (
	adjective = []string{
		"Весёлый",
		"Тихий",
		"Странный",
		"Грозный",
		"Причудливый",
		"Смеющийся",
		"Полосатый",
		"Гремучий",
		"Кусачий",
	}

	noun = []string{
		"Барсук",
		"Боров",
		"Волк",
		"Дрозд",
		"Дикобраз",
		"Баран",
		"Кролик",
		"Койот",
		"Крот",
		"Медведь",
		"Орел",
	}
)

// GetRandomName - генерация случайного имени для пользователя
func GetRandomName() string {
	return fmt.Sprintf("%s %s", getAdjective(), getNoun())
}

// getAdjective - получение случайного прилагательного
func getAdjective() string {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(adjective))))
	if err != nil {
		num = big.NewInt(0)
	}

	return adjective[num.Int64()]
}

// getNoun - получение случайного существительного
func getNoun() string {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(len(noun))))
	if err != nil {
		num = big.NewInt(0)
	}

	return noun[num.Int64()]
}
