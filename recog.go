package imgo

import (
	"bytes"
	"io"
)

import (
	"math"
)

func GetVector(body io.ReadCloser) (vector []uint8, cos float64, err error) {
	matrix, err2 := ResizeForMatrix2(body, 80, 60)
	if err2 != nil {
		err = err2
		return
	}

	vector = Matrix2Vector(matrix)
	cos = math.Sqrt(Dot(vector, vector))
	return
}

//calculate Cosine Similarity of two images, input two file path
func CosineSimilarity2(vector []uint8, cos float64, src2 string) (cossimi float64, err error) {
	matrix2, err2 := ResizeForMatrix(src2, 80, 60)
	if err2 != nil {
		err = err2
		return
	}

	myx := vector
	myy := Matrix2Vector(matrix2)
	cos1 := Dot(myx, myy)
	//	cos21 := math.Sqrt(Dot(myx, myx))
	cos22 := math.Sqrt(Dot(myy, myy))

	cossimi = cos1 / (cos * cos22)
	return
}

//calculate Cosine Similarity of two images, input two file path
func CosineSimilarity(src1 string, src2 string) (cossimi float64, err error) {
	matrix1, err1 := ResizeForMatrix(src1, 80, 60)
	if err1 != nil {
		err = err1
		return
	}

	matrix2, err2 := ResizeForMatrix(src2, 80, 60)
	if err2 != nil {
		err = err2
		return
	}

	myx := Matrix2Vector(matrix1)
	myy := Matrix2Vector(matrix2)
	cos1 := Dot(myx, myy)
	cos21 := math.Sqrt(Dot(myx, myx))
	cos22 := math.Sqrt(Dot(myy, myy))

	cossimi = cos1 / (cos21 * cos22)
	return
}

//binaryzation process of image matrix , threshold can use 127 to test
func Binaryzation(src [][][]uint8, threshold int) [][][]uint8 {
	imgMatrix := RGB2Gray(src)

	height := len(imgMatrix)
	width := len(imgMatrix[0])
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var rgb int = int(imgMatrix[i][j][0]) + int(imgMatrix[i][j][1]) + int(imgMatrix[i][j][2])
			if rgb > threshold {
				rgb = 255
			} else {
				rgb = 0
			}
			imgMatrix[i][j][0] = uint8(rgb)
			imgMatrix[i][j][1] = uint8(rgb)
			imgMatrix[i][j][2] = uint8(rgb)
		}
	}

	return imgMatrix
}

//GetFingerprint use Perceptual Hash Algorithm to get fingerprint from a pircture
func GetFingerprint(src string) (fp string, err error) {
	imgMatrix, err1 := ResizeForMatrix(src, 8, 8)
	if err1 != nil {
		return "", err1
	}

	//convert RGB to Gray
	h, w := len(imgMatrix), len(imgMatrix[0])
	gray := make([]byte, w*h)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			gray[x+y*8] = byte((imgMatrix[x][y][0]*30 + imgMatrix[x][y][1]*59 + imgMatrix[x][y][2]*11) / 100)
		}
	}

	//calculate average value of color of picture
	sum := 0
	for _, v := range gray {
		sum += int(v)
	}
	avg := byte(sum / len(gray))

	var buffer bytes.Buffer
	for _, v := range gray {
		if avg >= v {
			buffer.WriteByte('1')
		} else {
			buffer.WriteByte('0')
		}
	}
	fp = buffer.String()
	return
}
