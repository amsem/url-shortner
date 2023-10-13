package utils

import (
    "math"
    "strings"
)

const (
    base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    b = 62
)

func ToEncode(n int) string {
    r := n % b
    res := string(base[r])
    div := n / b
    q := int(math.Floor(float64(div)))

    for q != 0{
        r = q % b
        temp := q / b
        q = int(math.Floor(float64(temp)))
        res = string(base[int(r)]) + res
    }
    return string(res)
}

func ToDeCode(u string) int {
    res := 0
    for _, r:= range u {
        res = (b * res) + strings.Index(base, string(r))
    }
    return res
}
