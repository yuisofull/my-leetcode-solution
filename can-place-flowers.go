func canPlaceFlowers(flowerbed []int, n int) bool {
    for i,v:=range flowerbed {
        if i==0 {
            continue
        }
        if (v==0) && (flowerbed[i-1]==int(0)) && (flowerbed[i+1]==int(0)){
            n--
        }
        if i+1 == len(flowerbed){
            break
        }
    }
    return n==0
}