package main

import "sync"

type DB struct {
	mx sync.RWMutex
	m  map[string]string
}

func NewDB() *DB {
	return &DB{
		m: make(map[string]string),
	}
}

func (db *DB) Load(key string) (string, bool) {
	db.mx.RLock()
	defer db.mx.RUnlock()
	val, ok := db.m[key]
	return val, ok
}

func (db *DB) Store(key string, value string) {
	db.mx.Lock()
	defer db.mx.Unlock()
	db.m[key] = value
}
