package util

import (
	"encoding/csv"
	"log/slog"
	"os"

	"github.com/pkg/errors"
)

// ReadCSV reads CSV data from a file or stdout.
func ReadCSV(file string) ([][]string, error) {
	var err error
	in := os.Stdin
	if file != "" {
		in, err = os.Open(file)
		if err != nil {
			slog.Error("couldn't open input file", "error", err)
			return nil, errors.WithStack(err)
		}
		defer func() { _ = in.Close() }()
	}

	r := csv.NewReader(in)
	rows, err := r.ReadAll()
	if err != nil {
		slog.Error("couldn't read from input file", "error", err)
		return nil, errors.WithStack(err)
	}

	return rows, nil
}

// RemoveExpectedCSVHeader checks that the CSV's header is a non-empty prefix of the expected header,
// and removes it from the rows.
func RemoveExpectedCSVHeader(expectedHeader []string, rows [][]string) ([][]string, error) {
	if len(rows) == 0 {
		return nil, errors.WithStack(errors.New("expected CSV header but file is empty"))
	}

	header := rows[0]
	rows = rows[1:]
	if len(header) == 0 {
		return nil, errors.WithStack(errors.New("expected CSV header but first row has no fields"))
	}

	for i, field := range header[:min(len(header), len(expectedHeader))] {
		if field != expectedHeader[i] {
			return nil, errors.WithStack(errors.Errorf("unexpected column in CSV header: %v", field))
		}
	}

	return rows, nil
}

// WriteCSV writes CSV data to a file or stdout.
func WriteCSV(file string, rows [][]string) error {
	var err error
	out := os.Stdout
	if file != "" {
		out, err = os.Create(file)
		if err != nil {
			slog.Error("couldn't create output file", "error", err)
			return errors.WithStack(err)
		}
		defer func() { _ = out.Close() }()
	}

	w := csv.NewWriter(out)
	err = w.WriteAll(rows)
	if err != nil {
		slog.Error("couldn't write to output file", "error", err)
		return errors.WithStack(err)
	}

	return nil
}
