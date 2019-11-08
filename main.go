package main

import (
	"flag"
	"fmt"
	"net/http"

	"golang.org/x/text/encoding/htmlindex"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	_archiveHttpPresenter "github.com/alunegov/k3archive/archive/presentation/http"
	_archiveUsecase "github.com/alunegov/k3archive/archive/usecase"
	"github.com/alunegov/k3archive/mappers"
	_chunksMappers "github.com/alunegov/k3archive/mappers/chunks"
	"github.com/alunegov/k3archive/middleware"
	"github.com/alunegov/k3archive/models"
	_chunksModels "github.com/alunegov/k3archive/models/chunks"
)

func main() {
	// default fs root is for Korsar3RPi
	fsRoot := flag.String("fs_root", "./data", "a FS root")
	port := flag.Int("port", 3100, "a server port")
	encName := flag.String("enc", "windows-1251", "a text encoding of file data")
	debug := flag.Bool("debug", false, "a debug flag")
	flag.Parse()

	// init mappers and archive types
	enc, err := htmlindex.Get(*encName)
	if err != nil {
		fmt.Println(err)
		return
	}

	textDecoder := enc.NewDecoder()

	chunkMappers := _chunksMappers.ChunkMappers{
		_chunksModels.ChunkComment:      _chunksMappers.NewCommentChunkMapper(textDecoder),
		_chunksModels.ChunkSignalHeader: _chunksMappers.NewSignalHeaderChunkMapper(),
	}

	rosChunkFileMapper := _chunksMappers.NewRosChunkFileMapper(chunkMappers)

	archiveTypes := []struct {
		path   string
		mapper models.FileDataMapper
	}{
		{"Center", nil},
		{"Rms", mappers.NewRmsMapper(rosChunkFileMapper)},
		{"Signal", mappers.NewSignalMapper(rosChunkFileMapper)},
	}

	mux := httprouter.New()
	mux.HandleOPTIONS = true

	for _, archiveType := range archiveTypes {
		archiveUseCase := _archiveUsecase.NewArchiveUsecase(*fsRoot, archiveType.path, archiveType.mapper)
		_archiveHttpPresenter.UseArchiveApiHandler(mux, archiveUseCase, archiveType.path)
		_archiveHttpPresenter.UseArchiveHandler(mux, *fsRoot, archiveType.path)
	}

	c := cors.New(cors.Options{
		AllowedMethods:     []string{http.MethodGet, http.MethodDelete},
		AllowedHeaders:     []string{"authorization"},
		AllowCredentials:   true,
		OptionsPassthrough: true,
		Debug:              *debug,
	})

	handler := c.Handler(mux)

	handler = middleware.Logger(*debug, handler)

	addr := fmt.Sprintf(":%d", *port)
	if err := http.ListenAndServe(addr, handler); err != nil {
		fmt.Println(err)
	}
}
