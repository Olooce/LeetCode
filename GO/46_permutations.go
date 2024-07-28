package main

import "fmt"

// permute generates all permutations of a slice of integers
func permute(nums []int) [][]int {
    var results [][]int
    var path []int
    used := make([]bool, len(nums))
    backtrack(nums, used, path, &results)
    return results
}

// backtrack is a helper function that performs the DFS
func backtrack(nums []int, used []bool, path []int, results *[][]int) {
    if len(path) == len(nums) {
        // Make a copy of the current permutation and add it to the results
        perm := make([]int, len(path))
        copy(perm, path)
        *results = append(*results, perm)
        return
    }

    for i := 0; i < len(nums); i++ {
        if used[i] {
            continue
        }
        used[i] = true
        path = append(path, nums[i])
        backtrack(nums, used, path, results)
        // Backtrack
        path = path[:len(path)-1]
        used[i] = false
    }
}

func main() {
    nums := []int{1, 2, 3}
    result := permute(nums)
    fmt.Println(result)
}
