// https://leetcode.com/problems/can-place-flowers

func canPlaceFlowers(flowerbed []int, n int) bool {
    for i,v:=range flowerbed {
        if n==0 {
            return true
        }
        if i+1 == len(flowerbed){
            break
        }
        if i==0 {
            if (v==0) && (flowerbed[i+1]==int(0)){
                flowerbed[i]=1
                n--
            }
            continue
        }
        if (v==0) && (flowerbed[i-1]==int(0)) && (flowerbed[i+1]==int(0)){
            flowerbed[i]=1
            n--
        }
    }
    return n==0
}