package converters

import "testing"

func TestConvertId(t *testing.T) {
	id := "@fareedkhandev/pandas-ai-the-future-of-data-analysis-8f0be9b5ab6f"
	convertedId := idFromPath(id)
	expected := "8f0be9b5ab6f"
	if convertedId != expected {
		t.Errorf("ConvertId(%s) = %s; want %s", id, convertedId, expected)
	}
}