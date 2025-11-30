package service

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func createTempPromoFile(t *testing.T, dir, filename, content string) string {
	t.Helper()

	path := filepath.Join(dir, filename)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp promo file: %v", err)
	}
	return path
}

func TestPromoService_FindPromo_SingleMatch(t *testing.T) {
	ps := NewPromoService()

	dir := t.TempDir()
	file1 := createTempPromoFile(t, dir, "promo1.txt", "HELLO\nDISCOUNT50\nWORLD\n")
	file2 := createTempPromoFile(t, dir, "promo2.txt", "NOTHING\nHERE\n")
	file3 := createTempPromoFile(t, dir, "promo3.txt", "PROMO_X\nPROMO_Y\n")

	files := []string{file1, file2, file3}

	ctx := context.Background()
	matches, err := ps.FindPromo(ctx, "DISCOUNT50", files, 1)
	if err != nil {
		t.Fatalf("promo search error: %v", err)
	}

	if len(matches) != 1 {
		t.Fatalf("expected 1 match, got %d", len(matches))
	}

	if matches[0].File != file1 {
		t.Fatalf("expected match in %s, got %s", file1, matches[0].File)
	}
}

func TestPromoService_FindPromo_MultipleMatches_StopsEarly(t *testing.T) {
	ps := NewPromoService()

	dir := t.TempDir()
	file1 := createTempPromoFile(t, dir, "promo1.txt", "PROMO123\n")
	file2 := createTempPromoFile(t, dir, "promo2.txt", "abc\nPROMO123\n")
	file3 := createTempPromoFile(t, dir, "promo3.txt", "no match")

	files := []string{file1, file2, file3}

	ctx := context.Background()
	// maxMatches = 2 → should stop after two goroutines produce results
	matches, err := ps.FindPromo(ctx, "PROMO123", files, 2)
	if err != nil {
		t.Fatalf("promo search error: %v", err)
	}

	if len(matches) != 2 {
		t.Fatalf("expected 2 matches, got %d", len(matches))
	}
}

func TestPromoService_NoMatches(t *testing.T) {
	ps := NewPromoService()

	dir := t.TempDir()
	file1 := createTempPromoFile(t, dir, "promo1.txt", "nope\n")
	file2 := createTempPromoFile(t, dir, "promo2.txt", "still nope\n")

	files := []string{file1, file2}

	ctx := context.Background()
	matches, err := ps.FindPromo(ctx, "NOT_EXIST", files, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(matches) != 0 {
		t.Fatalf("expected 0 matches, got %d", len(matches))
	}
}
