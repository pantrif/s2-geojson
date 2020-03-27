# s2-geojson
- Display s2 cells on leaflet map using the region coverer. 
- Convert geojson features to cell unions depending on the min and max levels (supported only Polygons and Points).
- Draw points and polygons.
- Check intersection with the geojson features.


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