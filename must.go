package tryerr

func Must[T any](r T, err error) T {
	if err != nil {
		panic(err)
	}
	return r
}

func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

func Must2[T1, T2 any](r1 T1, r2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return r1, r2
}
