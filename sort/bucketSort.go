package sort


func insertionSort(A []float64) {
	for i := 1; i < len(A); i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && A[j] > key {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
	}
}

func bucketSort(array []float64, bucketSize int) []float64 {
	var max, min float64
	for _, n := range array {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	nBuckets := int(max-min)/bucketSize + 1
	buckets := make([][]float64, nBuckets)

	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]float64, 0)
	}

	for _, n := range array {
		idx := int(n-min) / bucketSize
		buckets[idx] = append(buckets[idx], n)
	}

	sorted := make([]float64, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			insertionSort(bucket)
			sorted = append(sorted, bucket...)
		}
	}

	return sorted
}

func main() {

}

