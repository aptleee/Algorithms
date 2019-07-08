#include <iostream>
#include <random>
template <typename T>
void swap(T a[], int lo, int hi)
{
	int t = a[lo];
	a[lo] = a[hi];
	a[hi] = t;
}

template <typename T>
void reverse(T a[], int lo, int hi)
{
	while (lo < hi)
		swap(a, lo++, hi--);
}

// [lo..mid, mid+1..hi] -> [mid+1..hi, lo..mid]
template <typename T>
void rotate(T a[], int lo, int mid, int hi)
{
	reverse(a, lo, mid);
	reverse(a, mid + 1, hi);
	reverse(a, lo, hi);
}

template <typename T>
void mergeSort(T a[], int lo, int hi)
{
	if (lo < hi)
	{
		int mid = lo + (hi - lo) / 2;
		int i = lo, j = mid + 1;
		mergeSort(a, lo, mid);
		mergeSort(a, mid + 1, hi);
		while (i < j && j < hi + 1)
		{
			while (i < j && a[i] <= a[j])
				i++;
			if (i == j)
				break;
			int old_j = j;
			while (j < hi + 1 && a[i] > a[j])
				j++;
			rotate(a, i, old_j - 1, j - 1);
			i = i + j - old_j + 1;
		}
	}
}

int main()
{
	std::random_device rd;
	std::uniform_int_distribution<int> dist(0, 100);
	int a[50] = {};
	for (int k = 0; k < 50; k++)
		a[k] = dist(rd);
	mergeSort(a, 0, sizeof(a) / sizeof(a[0]));
	for (int i : a)
		std::cout << i << " ";
	return 0;
}
