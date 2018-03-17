// Copyright 2018 Shestakov Denis. All rights reserved.
// Use of this source code is governed by a Lil
// license that can be found in the LICENSE file.

package err

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ErrV1(t *testing.T) {
	// Для неизвестного кода ошибки - код по-умолчанию
	err := V1(100500)
	assert.EqualError(t, err, "[ 500 | 1010000 ] Неизвестная ошибка")

	err = V1(9019999)
	assert.EqualError(t, err, "[ 500 | 1010000 ] Неизвестная ошибка")

	// Подстановка в шаблон из параметров
	err = V1(TST_ERR, "печалька")
	assert.EqualError(t, err, "[ 400 | 9010000 ] Тестовое сообщение: 'печалька'")

	// Проверка форматирования и сериализации JSON
	assert.Contains(t, fmt.Sprintf("%+v", err), "debug/stack.go")
	data, e := json.Marshal(err)
	assert.NoError(t, e)
	assert.Equal(t, `{"code":9010000,"text":"Тестовое сообщение: 'печалька'"}`, string(data))
}
