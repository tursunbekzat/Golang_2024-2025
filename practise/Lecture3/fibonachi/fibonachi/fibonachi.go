package fibonachi


func Generate(n int) []int {
    fibList := make([]int, n)
    fibList[0], fibList[1] = 0, 1

    for i := 2; i < n; i++ {
        fibList[i] = fibList[i-1] + fibList[i-2]
    }

    return fibList
}