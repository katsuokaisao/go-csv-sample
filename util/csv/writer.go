package csv

import (
	"encoding/csv"
	"io"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

type Row struct {
	Cells []string
}

type CSVWriter interface {
	WriteRows(w io.Writer, rows []Row) error
	WriteRow(w io.Writer, r *Row) error
}

type csvWriter struct {
	header  []string
	useCRLF bool
	encoder *encoding.Encoder
}

func NewCSVWriter(header []string, useCRLF bool, encoder *encoding.Encoder) CSVWriter {
	return &csvWriter{
		header:  header,
		useCRLF: useCRLF,
		encoder: encoder,
	}
}

func (w *csvWriter) WriteRows(writer io.Writer, rows []Row) error {
	csvWriter := csv.NewWriter(transform.NewWriter(writer, w.encoder))
	csvWriter.UseCRLF = w.useCRLF
	defer csvWriter.Flush()

	if len(w.header) > 0 {
		if err := csvWriter.Write(w.header); err != nil {
			return err
		}
	}

	for _, r := range rows {
		if err := csvWriter.Write(r.Cells); err != nil {
			return err
		}
	}

	return nil
}

func (w *csvWriter) WriteRow(writer io.Writer, r *Row) error {
	csvWriter := csv.NewWriter(transform.NewWriter(writer, w.encoder))
	csvWriter.UseCRLF = w.useCRLF
	defer csvWriter.Flush()

	return csvWriter.Write(r.Cells)
}
