package tryerr

import "fmt"

func CatchErrorAndHandle(handler func(msg string, err error)) {
	if err := recover(); err != nil {
		if se, ok := err.(specialError); ok {
			handler(se.msg, se.err)
		} else {
			panic(err)
		}
	}
}

func CatchError(pErr *error) {
	if err := recover(); err != nil {
		if se, ok := err.(specialError); ok {
			*pErr = fmt.Errorf("%s: %w", se.msg, se.err)
		} else {
			panic(err)
		}
	}
}

type specialError struct {
	err error
	msg string
}

type baseRetWrapper struct {
	err error
}

func (brw baseRetWrapper) handleErr(msg string, args ...any) {
	if brw.err != nil {
		panic(specialError{
			err: brw.err,
			msg: fmt.Sprintf(msg, args...),
		})
	}
}

type retWrapper[T any] struct {
	baseRetWrapper
	result T
}

func (rw retWrapper[T]) Err(msg string, args ...any) T {
	rw.handleErr(msg, args...)
	return rw.result
}

type retWrapper0 struct {
	baseRetWrapper
}

func (rw retWrapper0) Err(msg string, args ...any) {
	rw.handleErr(msg, args...)
}

type retWrapper2[T, V any] struct {
	baseRetWrapper
	resultT T
	resultV V
}

func (rw retWrapper2[T, V]) Err(msg string, args ...any) (T, V) {
	rw.handleErr(msg, args...)
	return rw.resultT, rw.resultV
}

type retWrapper3[T, V, U any] struct {
	baseRetWrapper
	resultT T
	resultV V
	resultU U
}

func (rw retWrapper3[T, V, U]) Err(msg string, args ...any) (T, V, U) {
	rw.handleErr(msg, args...)
	return rw.resultT, rw.resultV, rw.resultU
}

func Try[T any](result T, err error) retWrapper[T] {
	return retWrapper[T]{
		baseRetWrapper: baseRetWrapper{err: err},
		result:         result,
	}
}

func Try0(err error) retWrapper0 {
	return retWrapper0{baseRetWrapper{err: err}}
}

func Try2[T, V any](resultT T, resultV V, err error) retWrapper2[T, V] {
	return retWrapper2[T, V]{
		baseRetWrapper: baseRetWrapper{err: err},
		resultT:        resultT,
		resultV:        resultV,
	}
}

func Try3[T, V, U any](resultT T, resultV V, resultU U, err error) retWrapper3[T, V, U] {
	return retWrapper3[T, V, U]{
		baseRetWrapper: baseRetWrapper{err: err},
		resultT:        resultT,
		resultV:        resultV,
		resultU:        resultU,
	}
}
