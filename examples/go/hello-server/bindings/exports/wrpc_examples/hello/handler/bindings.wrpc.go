// Generated by `wit-bindgen-wrpc-go` 0.6.0. DO NOT EDIT!
package handler

import (
	bytes "bytes"
	context "context"
	binary "encoding/binary"
	fmt "fmt"
	io "io"
	slog "log/slog"
	math "math"
	wrpc "wrpc.io/go"
)

type Handler interface {
	Hello(ctx__ context.Context) (string, error)
}

func ServeInterface(s wrpc.Server, h Handler) (stop func() error, err error) {
	stops := make([]func() error, 0, 1)
	stop = func() error {
		for _, stop := range stops {
			if err := stop(); err != nil {
				return err
			}
		}
		return nil
	}

	stop0, err := s.Serve("wrpc-examples:hello/handler", "hello", func(ctx context.Context, w wrpc.IndexWriteCloser, r wrpc.IndexReadCloser) {
		defer func() {
			if err := w.Close(); err != nil {
				slog.DebugContext(ctx, "failed to close writer", "instance", "wrpc-examples:hello/handler", "name", "hello", "err", err)
			}
		}()
		slog.DebugContext(ctx, "calling `wrpc-examples:hello/handler.hello` handler")
		r0, err := h.Hello(ctx)
		if cErr := r.Close(); cErr != nil {
			slog.ErrorContext(ctx, "failed to close reader", "instance", "wrpc-examples:hello/handler", "name", "hello", "err", err)
		}
		if err != nil {
			slog.WarnContext(ctx, "failed to handle invocation", "instance", "wrpc-examples:hello/handler", "name", "hello", "err", err)
			return
		}

		var buf bytes.Buffer
		writes := make(map[uint32]func(wrpc.IndexWriter) error, 1)

		write0, err := (func(wrpc.IndexWriter) error)(nil), func(v string, w io.Writer) (err error) {
			n := len(v)
			if n > math.MaxUint32 {
				return fmt.Errorf("string byte length of %d overflows a 32-bit integer", n)
			}
			if err = func(v int, w io.Writer) error {
				b := make([]byte, binary.MaxVarintLen32)
				i := binary.PutUvarint(b, uint64(v))
				slog.Debug("writing string byte length", "len", n)
				_, err = w.Write(b[:i])
				return err
			}(n, w); err != nil {
				return fmt.Errorf("failed to write string byte length of %d: %w", n, err)
			}
			slog.Debug("writing string bytes")
			_, err = w.Write([]byte(v))
			if err != nil {
				return fmt.Errorf("failed to write string bytes: %w", err)
			}
			return nil
		}(r0, &buf)
		if err != nil {
			slog.WarnContext(ctx, "failed to write result value", "i", 0, "wrpc-examples:hello/handler", "name", "hello", "err", err)
			return
		}
		if write0 != nil {
			writes[0] = write0
		}
		slog.DebugContext(ctx, "transmitting `wrpc-examples:hello/handler.hello` result")
		_, err = w.Write(buf.Bytes())
		if err != nil {
			slog.WarnContext(ctx, "failed to write result", "wrpc-examples:hello/handler", "name", "hello", "err", err)
			return
		}
		if len(writes) > 0 {
			for index, write := range writes {
				w, err := w.Index(index)
				if err != nil {
					slog.ErrorContext(ctx, "failed to index writer", "index", index, "wrpc-examples:hello/handler", "name", "hello", "err", err)
					return
				}
				index := index
				write := write
				go func() {
					if err := write(w); err != nil {
						slog.WarnContext(ctx, "failed to write nested result value", "index", index, "wrpc-examples:hello/handler", "name", "hello", "err", err)
					}
				}()
			}
		}
	})
	if err != nil {
		return nil, fmt.Errorf("failed to serve `wrpc-examples:hello/handler.hello`: %w", err)
	}
	stops = append(stops, stop0)
	return stop, nil
}
