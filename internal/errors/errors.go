package errors

// Private Variables
var errorCodeToString = map[int]string{
    UnknownError: "Given error code is not recognised. Consider checking! Given code: '%d'",
    UnknownVerb: "Given verb for throwing an error is unknown. Consider checking. Given verb: '%s'",
    EmptyError: "%s",
}

// Public Constants
const (
    UnknownError int = iota
    UnknownVerb
    EmptyError
)
