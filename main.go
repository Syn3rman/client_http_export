package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"


	"github.com/Syn3rman/httpExporter"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var logger = log.New(os.Stderr, "http-example", log.Ldate|log.Ltime|log.Llongfile)

func initTracer(url string) func() {
exporter, err := httpExporter.New(
		url,
		httpExporter.WithLogger(logger),
	)
	if err != nil {
		log.Fatal(err)
	}

	batcher := sdktrace.NewBatchSpanProcessor(exporter)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(batcher),
	)
	otel.SetTracerProvider(tp)

	return func() {
		_ = tp.Shutdown(context.Background())
	}
}
func main() {
	url := flag.String("http", "http://localhost:4000", "http url")
	flag.Parse()

	shutdown := initTracer(*url)
	defer shutdown()

	ctx := context.Background()

	tr := otel.GetTracerProvider().Tracer("component-main")
	ctx, span := tr.Start(ctx, "foo", trace.WithSpanKind(trace.SpanKindServer))
	<-time.After(6 * time.Millisecond)
	bar(ctx)
	<-time.After(6 * time.Millisecond)
	span.End()
}

func bar(ctx context.Context) {
	tr := otel.GetTracerProvider().Tracer("component-bar")
	_, span := tr.Start(ctx, "bar")
	<-time.After(6 * time.Millisecond)
	span.End()
}
