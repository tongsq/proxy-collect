package bootstrap

import "fmt"

type StringList []string

func (l *StringList) String() string {
	return fmt.Sprintf("%s", *l)
}
func (l *StringList) Set(value string) error {
	*l = append(*l, value)
	return nil
}
