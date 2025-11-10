package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/JDGarner/go-template/internal/store/sqlc"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func TestIntegration(t *testing.T) {
	ctx := context.Background()

	testResources, err := setupTestResources(ctx, t)
	if err != nil {
		fmt.Printf("setup tests failed: %s", err)
		testResources.Cleanup(ctx, t)
		os.Exit(1)
	}

	t.Cleanup(func() {
		testResources.Cleanup(ctx, t)
	})

	t.Run("returns dummy item", func(t *testing.T) {
		id := uuid.New()
		expectedName := "hello"

		_, err := testResources.DB.ExecContext(
			ctx,
			`INSERT INTO dummy VALUES ($1, $2)`,
			id,
			expectedName,
		)
		if err != nil {
			t.Fatalf("failed to insert test data: %v", err)
		}

		url := fmt.Sprintf("%s/item/%s", testResources.AppURL, id.String())

		resp, err := http.Get(url)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected status 200, got %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}

		var result sqlc.Dummy
		err = json.Unmarshal(body, &result)
		if err != nil {
			t.Fatalf("failed to parse JSON response: %v. Body: %s", err, string(body))
		}

		if result.Name != expectedName {
			t.Errorf("expected name '%s', got '%s'", expectedName, result.Name)
		}

		if result.ID.String() != id.String() {
			t.Errorf("expected id '%s', got '%s'", id.String(), result.ID)
		}
	})

	t.Run("returns not found error", func(t *testing.T) {
		nonExistantID := uuid.New()

		url := fmt.Sprintf("%s/item/%s", testResources.AppURL, nonExistantID.String())

		resp, err := http.Get(url)
		if err != nil {
			t.Fatalf("failed to make request: %v", err)
		}

		if resp.StatusCode != http.StatusNotFound {
			t.Fatalf("expected status 404, got %d", resp.StatusCode)
		}
	})
}
