package main

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

var TRUE = pulumi.BoolPtr(true)
var FALSE = pulumi.BoolPtr(false)


type String string
func (s String) R() string{
	return string(s)
}
func (s String) S() pulumi.StringInput{
	return pulumi.String(s)
}
func (s String) SP() pulumi.StringPtrInput{
	return pulumi.StringPtr(string(s))
}

type Int int
func (i Int) R() int{
	return int(i)
}
func (i Int) I() pulumi.IntInput{
	return pulumi.Int(i)
}
func (i Int) SP() pulumi.IntPtrInput{
	return pulumi.IntPtr(int(i))
}

type Bool bool
func (b Bool) R() bool{
	return bool(b)
}
func (b Bool) BP() pulumi.BoolPtrInput{
	return pulumi.BoolPtr(bool(b))
}
func (b Bool) B() pulumi.BoolInput{
	return pulumi.Bool(b)
}