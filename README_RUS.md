# LRU Cache [![Go Reference](https://pkg.go.dev/badge/github.com/tarvarrs/lru-cache.svg)](https://pkg.go.dev/github.com/tarvarrs/lru-cache) [![Go Report Card](https://goreportcard.com/badge/github.com/tarvarrs/lru-cache)](https://goreportcard.com/report/github.com/tarvarrs/lru-cache) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Реализация потокобезопасного LRU (Least Recently Used) кеша с дженериками на Go 1.26+.  
Поддерживает операции `Get`, `Set`, и `Clear` с временной сложностью O(1).

## Преимущества

- **Дженерики** – работает с comparable ключами и any значениями.
- **Потокобезопасность** – использует мьютекс для конкурентного доступа.
- **O(1) операции** – основан на хеш-мапе и двусвязном списке.
- **Удаление при переполнении** – автоматически удаляет самые давно использованные пары при переполнении емкости.
- **Отсутствие зависимостей** – используется только стандартная библиотека.

## Установка

```bash
go get github.com/tarvarrs/lru-cache@latest
```
