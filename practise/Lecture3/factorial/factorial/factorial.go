package factorial


func Calculate(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    return n * Calculate(n-1)
}