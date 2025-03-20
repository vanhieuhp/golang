package main

//
//import (
//	"database/sql"
//	"errors"
//)
//
//type DatabaseAccessor struct {
//
//	// data connection
//}
//
//type CustomError struct {
//	wrappedError err
//}
//
//func (e *CustomError) Error() string {
//	return "This is my custom error"
//}
//
//func (e *CustomError) Unwrap() error {
//	return &e.wrappedError
//}
//
//// Dependency Injection
//// google/wire
//func newDatabaseAccessor() (accessor *DatabaseAccessor, cleanUpFunc func(), err error) {
//	//return &DatabaseAccessor{}, func() {
//	//	// clean up task
//	//}, nil
//	if errors.Is(err, sql.ErrNoRows) {
//
//	}
//	return nil, nil, &CustomError{}
//}
//
//func main() {
//	accessor, cleanUp, err := newDatabaseAccessor()
//	if err != nil {
//		if cleanUp != nil {
//			cleanUp()
//		}
//		panic(err)
//	}
//
//	cleanUp()
//}
