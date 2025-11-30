package service

import (
	"bufio"
	"context"
	"os"
	"sync"
)

type PromoLocation struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

type PromoService interface {
	FindPromo(ctx context.Context, code string, files []string, maxMatches int) ([]PromoLocation, error)
}

type promoService struct{}

func NewPromoService() PromoService {
	return &promoService{}
}

func (ps *promoService) FindPromo(
	ctx context.Context,
	code string,
	files []string,
	maxMatches int,
) ([]PromoLocation, error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	type result struct {
		loc PromoLocation
		err error
	}

	resultsCh := make(chan result)
	var wg sync.WaitGroup

	for _, path := range files {
		wg.Add(1)

		go func(p string) {
			defer wg.Done()

			file, err := os.Open(p)
			if err != nil {
				select {
				case resultsCh <- result{err: err}:
				case <-ctx.Done():
				}
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			lineNum := 0

			for scanner.Scan() {
				lineNum++

				// Stop scanning immediately when cancelled
				select {
				case <-ctx.Done():
					return
				default:
				}

				if scanner.Text() == code {
					// Non-blocking send to avoid blocking on channel
					select {
					case resultsCh <- result{
						loc: PromoLocation{
							File: p,
							Line: lineNum,
						},
					}:
					case <-ctx.Done():
						return
					}
				}
			}

			if err := scanner.Err(); err != nil {
				select {
				case resultsCh <- result{err: err}:
				case <-ctx.Done():
				}
			}

		}(path)
	}

	// Close resultsCh when all goroutines finish
	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	var matches []PromoLocation

	for r := range resultsCh {
		// ignore errors from individual files
		if r.err == nil {
			matches = append(matches, r.loc)
		}

		// As soon as we have enough matches:
		if len(matches) >= maxMatches {
			// Immediately cancel other goroutines
			cancel()
			return matches, nil
		}
	}

	// Fewer than required matches
	return matches, nil
}
