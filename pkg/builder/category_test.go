package builder

import (
    "testing"
)

const (
    cat1 = "category1"
    cat2 = "category2"
)

func TestCategory(t *testing.T) {
    document1 := Document{Name: "document1"}
    document2 := Document{Name: "document2"}
    document3 := Document{Name: "document3"}
    document4 := Document{Name: "document4"}
    categories := newCategory()
    categories.AddDocument(cat1, &document1)
    categories.AddDocument(cat2, &document2)
    categories.AddDocument(cat1, &document3)
    if categories.Contains(&document1) == false {
        t.Fatalf("document1 not in category %s", cat1)
    }
    if categories.Contains(&document2) == false {
        t.Fatalf("document2 not in category %s", cat2)
    }
    if categories.Contains(&document3) == false {
        t.Fatalf("document3 not in category %s", cat1)
    }
    if categories.Contains(&document4) == true {
        t.Fatalf("document4 IN category!")
    }
    t.Logf("categories: %p %v\n", categories, categories)
}
