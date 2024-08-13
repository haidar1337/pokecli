module github.com/haidar1337/pokecli

go 1.22.6

require internal/api v1.0.0

replace internal/api => ./internal/api

require internal/pokecache v1.0.0

replace internal/pokecache => ./internal/pokecache