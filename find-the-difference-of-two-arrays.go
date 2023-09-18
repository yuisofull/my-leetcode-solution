func findDifference(nums1 []int, nums2 []int) [][]int {
    var res [][]int
    m1 := map[int]int{}
    m2 := map[int]int{}
    res = append(res, []int{})
    res = append(res, []int{})

    for _, v1 := range nums1 {
        a := true
        for _, v2 := range nums2 {
            if v1 == v2 {
                a = false
            }
        }
        if a && m1[v1] == 0{
            res[0] = append(res[0], v1)
            m1[v1]++
        }
    }
    for _, v1 := range nums2 {
        a := true
        for _, v2 := range nums1 {
            if v1 == v2 {
                a = false
            }
        }
        if a && m2[v1] == 0 {
            res[1] = append(res[1], v1)
            m2[v1]++
        }
    }
    return res
}