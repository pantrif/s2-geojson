[![Build Status](https://travis-ci.com/pantrif/s2-geojson.svg?branch=develop)](https://travis-ci.com/pantrif/s2-geojson)
[![GoDoc](https://godoc.org/github.com/pantrif/s2-geojson?status.png)](http://godoc.org/github.com/pantrif/s2-geojson)
[![Go Report Card](https://goreportcard.com/badge/github.com/pantrif/s2-geojson)](https://goreportcard.com/report/github.com/pantrif/s2-geojson)
[![codecov](https://codecov.io/gh/pantrif/s2-geojson/branch/develop/graph/badge.svg)](https://codecov.io/gh/pantrif/s2-geojson)


## Preview
<img width="1645" alt="preview" src="https://user-images.githubusercontent.com/9129280/77917392-31dca180-72a3-11ea-9e93-27c331731196.png">


## Summary

Draw a polygon on the map or paste a geoJSON and explore how the s2.RegionCoverer covers it with S2 cells depending on the min and max levels.
Drop a marker and inspect the corresponding S2 cell.

- Display s2 cells on leaflet map using the region coverer. 
- Convert geojson features to cell unions depending on the min and max levels (supported only Polygons and Points).
- Draw points and polygons.
- Check point & circle intersection with the geoJSON features.


## Quick start
```
 go run cmd/s2-geojson/main.go
```

## Docker 
```
docker run -p 8080:8080 --rm lmaroulis/s2-geojson
```



## License

This project is licensed under the MIT License - see [the LICENSE file](LICENSE) for details.