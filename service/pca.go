package service

import (
    "fmt"
    "gonum.org/v1/gonum/mat"
    "gonum.org/v1/gonum/stat"
    "github.com/go-nlp/tfidf" // Correct import path for go-nlp/tfidf
)

// Assuming we have a corpus and a way to fit the model before using it for a single text.
// For simplicity, we're directly using it here but in practice, you should fit this model
// with a corpus of documents to get meaningful results.

// Initialize a global TF-IDF vectorizer
var vectorizer = tfidf.New()

// TextToVector converts prayer text to a numerical vector representation using TF-IDF.
func TextToVector(text string) []float64 {
    // Convert the single text into a "document" for the TF-IDF vectorizer
    // Normally, you would first fit the vectorizer on a larger corpus to calculate IDF values.
    vectorizer.AddDocs(text)
    vector := vectorizer.Transform([]string{text})

    // The Transform method might return a sparse matrix or a dense slice depending on the library's implementation.
    // You need to convert this output to a dense slice if it's not already in that format.
    // Below is a pseudo-code as the actual implementation depends on how the TF-IDF library represents the transformed vector.

    // Assuming vector is a slice of tfidf.Document, where each Document represents a vector for the text.
    var tfidfVector []float64
    for _, wordVector := range vector[0].WordWeights {
        tfidfVector = append(tfidfVector, wordVector)
    }
    return tfidfVector
}

// ComputePCA performs PCA on the given text using the TF-IDF vector.
// Returns the principal components as a slice of floats.
func ComputePCA(text string) ([]float64, error) {
    vector := TextToVector(text)
    rows, cols := 1, len(vector) // Adjust as needed for multiple vectors
    data := mat.NewDense(rows, cols, vector)

    var pc stat.PC
    ok := pc.PrincipalComponents(data, nil)
    if !ok {
        return nil, fmt.Errorf("PCA computation failed")
    }

    // Extract the first principal component correctly
    var vec mat.Dense
    pc.VectorsTo(&vec)

    // Assuming we want the first principal component
    var pc1 mat.VecDense
    pc1.ColViewOf(&vec, 0)

    return pc1.RawVector().Data, nil
}
