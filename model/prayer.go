// /prayerProcessor/model/prayer.go

package model

// Prayer represents the incoming JSON payload.
type Prayer struct {
    Title       string  `json:"title"`
    Geolocation string  `json:"geolocation"` // You might want to split this into more precise struct fields depending on your needs.
    Text        string  `json:"text"`
}

// PCAResult represents the dimensional reduction result of a prayer's text.
type PCAResult struct {
    ReducedDimensions []float64 `json:"reducedDimensions"`
}

