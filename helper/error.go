package helper

var DataMessage = make(map[string]interface{})

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}
