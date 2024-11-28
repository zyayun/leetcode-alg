def BoubbleSort(arr):

    for i in range(len(arr)):
        for j in range(len(arr) -i -1):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]




if __name__ == "__main__":
    arr = [6,1,2,3,4,5]
    BoubbleSort(arr)
    print(arr)
