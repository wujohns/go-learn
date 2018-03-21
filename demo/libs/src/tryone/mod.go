/**
 * 模块拆分简单案例
 * 需要注意的是暴露到外供调用的函数必须以大写字母开头
 *
 * @author wujohns
 * @date 18/3/21
 */
package tryone

import (
    "fmt"
)

func bb () {
    fmt.Println("Just BB");
}

func Sqrt (x float64) float64 {
    bb();
    z := 0.0;
    for i := 0; i < 1000; i++ {
        z -= (z*z - x)/(2*x)
    }
    return z
}