// Package projecterrors содержит типовые ошибки проекта
package projecterrors

import "errors"

// ErrEmailRequired указывает, что поле email равно nil.
var ErrEmailRequired = errors.New("field 'email' is required")

// ErrPasswordRequired указывает, что поле password равно nil.
var ErrPasswordRequired = errors.New("field 'password' is required")

// ErrReqBodyNilTask проверяет, что тело запроса равно nil.
var ErrReqBodyNilTask = errors.New("request body cannot be nil")

// ErrNilUserID указывает, что поле UserID равно nil.
var ErrReqBodyNilUserID = errors.New("user ID cannot be nil")

// ErrNotFoundTask указывает, что задача не найдена.
var ErrNotFoundTask = errors.New("cannot find the Task")

// ErrReqBodyNilUser проверяет, что тело запроса равно nil.
var ErrReqBodyNilUser = errors.New("request body cannot be nil")

// ErrNotFoundUser указывает, что задача не найдена.
var ErrNotFoundUser = errors.New("cannot find the Task")

// ErrNoTaskTable указывает отсутствие таблицы 'tasks' в БД (psql код: 42P01).
var ErrNoTaskTable = errors.New("таблица 'tasks' не найдена (код: 42P01)")

// ErrNoUserTable указывает отсутствие таблицы 'user' в БД (psql код: 42P01).
var ErrNoUserTable = errors.New("таблица 'user' не найдена (код: 42P01)")
