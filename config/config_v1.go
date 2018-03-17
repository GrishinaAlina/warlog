// Copyright 2018 Shestakov Denis. All rights reserved.
// Use of this source code is governed by a Lil
// license that can be found in the LICENSE file.

// Package config provides app configuration reading and validating.
package config

import (
	"sync"
)

// Внутренняя глобальная переменная для хранения полученных значений
var (
	v1      = new(cfg_v1)
	v1_lock = new(sync.RWMutex)
)

// Точка входа - главный метод получения текущего экземпляра конфига
func V1() ConfigV1 {
	v1_lock.RLock()
	defer v1_lock.RUnlock()
	return v1
}

// Вызов перезагрузки из указанного файла
func ReloadV1(path string) error {
	return nil
}

// Публичный интерфейс, вся работа идет через него
type ConfigV1 interface {
	Validate() error
}

// Внутренняя структура для хранения конкретных значений
type cfg_v1 struct {
}

// Базовая точка входа для валидации
func (this *cfg_v1) Validate() error {
	return nil
}
