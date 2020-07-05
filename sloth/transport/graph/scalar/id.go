package scalar
//
//import (
//	"fmt"
//	"io"
//	"strconv"
//)
//
//type ID int
//
//
//func (id *ID) UnmarshalGQL(v interface{}) error {
//	value, ok := v.(string)
//	if !ok {
//		return fmt.Errorf("id must be string")
//	}
//
//	result, err := strconv.Atoi(value)
//	if err != nil {
//		return err
//	}
//
//	*id = ID(result)
//
//	return nil
//}
//
//// MarshalGQL implements the graphql.Marshaler interface
//func (id ID) MarshalGQL(w io.Writer) {
//	w.Write([]byte(strconv.Itoa(int(id))))
//}