package service

import (
    "fmt"
    "gonum.org/v1/gonum/mat"
    "gonum.org/v1/gonum/stat"
    "prayerProcessor/tfidf" // Adjust the import path to where your tfidf package is located.
)

// Assuming Document implementation is provided.
// Here's a simple implementation of the Document interface based on the provided tfidf package.
type SimpleDocument struct {
    IDs []int
}

func (d SimpleDocument) IDs() []int {
    return d.IDs
}

// Convert text to a numerical vector representation using TF-IDF scores.
// This function needs to convert text into a format that your tfidf package can work with.
func TextToVector(docs []Document, tfidfCalculator *tfidf.TFIDF) [][]float64 {
    vectors := make([][]float64, len(docs))
    for i, doc := range docs {
        tfidfCalculator.Add(doc)
    }
    tfidfCalculator.CalculateIDF()

    for i, doc := range docs {
        vectors[i] = tfidfCalculator.Score(doc)
    }
    return vectors
}

// ComputePCA performs PCA on the given documents using the TF-IDF vector.
// Returns the principal components as a slice of floats.
func ComputePCA(docs []Document) ([]float64, error) {
    // Initialize the TFIDF calculator
    tfidfCalculator := tfidf.New()

    // Convert documents to vectors
    vectors := TextToVector(docs, tfidfCalculator)

    // Assuming vectors is a non-empty slice of TF-IDF vectors
    rows, cols := len(vectors), len(vectors[0])
    flatData := make([]float64, 0, rows*cols)
    for _, vector := range vectors {
        flatData = append(flatData, vector...)
    }
    data := mat.NewDense(rows, cols, flatData)

    var pc stat.PC
    ok := pc.PrincipalComponents(data, nil)
    if !ok {
        return nil, fmt.Errorf("PCA computation failed")
    }

    // Correctly use VectorsTo to extract principal components
    var vec mat.Dense
    pc.VectorsTo(&vec) // This line previously had incorrect usage
    
    // Extract the first principal component
    pc1 := vec.ColView(0)

    // Convert to []float64
    pc1Vec := make([]float64, pc1.Len())
    for i := 0; i < pc1.Len(); i++ {
        pc1Vec[i] = pc1.AtVec(i)
    }

    return pc1Vec, nil
}
