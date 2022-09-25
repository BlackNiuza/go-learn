package main

import (
	"fmt"
	"strconv"
)

type INT uint64

// ================================================================
type Recurable interface {
	Fibonacci(i INT, callback func(i INT, res INT) bool) INT
}

//================================================================

type FbisException struct {
	error
	msg string
}

func (fe *FbisException) Error() string {
	return "FbisException: " + fe.msg
}

//================================================================

type Fbis struct {
	// 1. 继承
	Recurable
	vals [1000]INT
}

// 2. 重载
func (f *Fbis) Fibonacci(i INT, callback func(i INT, res INT) bool) INT {
	return f.fibonacci0(i, callback)
}

// 3. 封装，可见性
func (f *Fbis) fibonacci0(i INT, callback func(i INT, res INT) bool) INT {
	var res INT = 0
	if f.vals[i] != 0 {
		res = f.vals[i]
	} else {
		if i <= 2 {
			res = 1
		} else {
			res = f.Fibonacci(i-1, callback) + f.Fibonacci(i-2, callback)
		}
		f.vals[i] = res
		callback(i, res)
	}
	return res
}

// 5. Public异常
func (f *Fbis) HasError() (str string, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()
	str = f.hasError0()
	return
}

// 6. private异常
func (f *Fbis) hasError0() string {
	exception := new(FbisException)
	exception.msg = "haha"
	panic(exception)
}

//================================================================

func main() {
	print := func(i INT, res INT) bool {
		fmt.Println("i=" + strconv.FormatUint(uint64(i), 10) + ", res=" + strconv.FormatUint(uint64(res), 10))
		return true
	}
	var fbis *Fbis = new(Fbis)
	fmt.Println("==========================================")
	fmt.Printf(">> fbis=%d\n", fbis.Fibonacci(100, print))
	fmt.Println("==========================================")
	// 4. 转换父类
	var r Recurable = fbis
	fmt.Printf(">> r=%d\n", r.Fibonacci(100, print))
	fmt.Println("==========================================")
	// 7. 测试异常
	if _, err := fbis.HasError(); err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Println("No Error")
	}
}
