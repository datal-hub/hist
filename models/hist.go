package models

type AsciiHist [128]int

func (h *AsciiHist) Add(secondHist AsciiHist) {
	for i := 0; i < len(h); i++ {
		h[i] += secondHist[i]
	}
}

func (h *AsciiHist) FromByteSlice(data []byte) {
	for i := 0; i < len(data); i++ {
		h[int(data[i])]++
	}
	return
}
