// https://leetcode.com/problems/can-place-flowers

func canPlaceFlowers(flowerbed []int, n int) bool {
    for i,v:=range flowerbed {
        if n==0 {
            return true
        }
        if i == len(flowerbed)-1{
            if (v==0) && (flowerbed[i-1]==int(0)){
                flowerbed[i]=1
                n--
            }
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