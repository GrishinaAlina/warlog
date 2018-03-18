// Copyright 2018 Shestakov Denis. All rights reserved.
// Use of this source code is governed by a Lil
// license that can be found in the LICENSE file.

// Пакет вводит собственную реализацию интерфейса ошибки
// для внутреннего кодирования места возникнования, причин,
// текстов сообщений, а также для единого стиля обработки.
package err

import (
	"fmt"
	"runtime/debug"
)

// Внешняя структура для сериализации и экспорта
type ErrV1 struct {
	Code    int    `json:"code"` // Код для отладки
	Status  int    `json:"-"`    // Код для HTTP-статуса
	Message string `json:"text"` // Cообщение об ошибке
	Stack   []byte `json:"-"`    // Стек для дебага
}

// Реализация интерфейса error
func (this ErrV1) Error() string {
	return fmt.Sprintf("[ %3d | %4d ] %s", this.Status, this.Code, this.Message)
}

// Хук для обработки при форматировании - печать стека только когда нужно
func (this ErrV1) Format(f fmt.State, c rune) {
	f.Write([]byte(this.Error()))
	if f.Flag('+') {
		fmt.Fprintf(f, "\n%s\n", this.Stack)
	}
}

// Основная функция. Формирование нового объекта ошибки (конструктор)
func V1(code int, args ...interface{}) error {
	var exist bool
	var result ErrV1

	// Граничные условия на код
	if code < UNKNOWN || code > MAX_ERR {
		code = UNKNOWN
	}

	// Создаем копию объекта ошибки из найденного шаблона
	if result, exist = tpl_v1[code]; !exist {
		code = UNKNOWN
		result = tpl_v1[UNKNOWN]
	}

	// В копии можем менять, что нужно
	result.Code = code
	result.Stack = debug.Stack()

	// Если передали аргументы для форматирования - используем
	// Наличие в шаблоне достаточного места для вставки не проверяется
	if len(args) > 0 {
		result.Message = fmt.Sprintf(result.Message, args...)
	}
	return result
}
