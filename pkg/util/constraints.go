package util

type Hashable interface {
	~bool |
		~string |
		~int |
		~uint |
		~byte |
		~rune |
		float32 |
		float64 |
		complex64 |
		complex128
}
