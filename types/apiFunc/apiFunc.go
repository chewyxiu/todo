package apiFunc

type APIFunc func(validatedRequest interface{}) (statusCode int, output interface{})