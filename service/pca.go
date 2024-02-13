package service

import (
    "gonum.org/v1/gonum/mat"
    "gonum.org/v1/gonum/stat"
    "github.com/go-nlp/tfidf"
    "strings"
)

// TextToVector converts prayer text to a numerical vector representation using TF-IDF.
func TextToVector(text string, corpus []string) []float64 {
    // Initialize a new TFIDF calculator
    calculator := tfidf.New()

    // Add documents to the corpus
    for _, doc := range corpus {
        calculator.AddDocs(strings.Fields(doc))
    }

    // Calculate TF-IDF for the given text
    // Splitting the text into words for vectorization
    words := strings.Fields(text)
    vector := make([]float64, len(words))
    for i, word := range words {
        vector[i] = calculator.TFIDF(word, words)
    }

    return vector
}

// ComputePCA performs PCA on the given text.
// Adjust the function signature as needed to pass the corpus.
func ComputePCA(text string, corpus []string) ([]float64, error) {
    // Convert the prayer text to a numerical vector using TF-IDF
    vector := TextToVector(text, corpus)

    // Create a matrix from the vector (in a real scenario, you would have multiple vectors)
    rows, cols := 1, len(vector) // Assuming one vector for simplicity; adjust as needed
    data := mat.NewDense(rows, cols, vector)

    // Compute the PCA
    var pc stat.PC
    ok := pc.PrincipalComponents(data, nil)
    if !ok {
        return nil, fmt.Errorf("PCA computation failed")
    }

    // Extract the first principal component
    var vec mat.VecDense
    pc.VectorsTo(nil).ColView(0, &vec)

    return vec.RawVector().Data, nil
}


