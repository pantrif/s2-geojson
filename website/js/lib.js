
let map, geoJsonEditor, marker, circleLayer, drawnItems;
let polygon_cells = [];
let circle_cells = [];
let geoJsonLayers = [];

let coverUrl = '/cover';
let checkPointUrl = '/check_intersection';
let self, app;


let s2_geojson = {
    Init: function() {
        self = this;
        this.initEditor();
        this.initMap();
        this.initMapControls();
        this.bindEvents();
        return this;
    },
    initEditor: function() {
        geoJsonEditor = CodeMirror.fromTextArea(document.getElementById('geoJsonInput'), {
            mode: "javascript",
            theme: "default",
            lineNumbers: true,
        });
    },
    initMap : function() {
        map = L.map('map').setView([51.505, -0.09], 13);

        L.tileLayer('https://api.mapbox.com/styles/v1/{id}/tiles/{z}/{x}/{y}?access_token=pk.eyJ1IjoibWFwYm94IiwiYSI6ImNpejY4NXVycTA2emYycXBndHRqcmZ3N3gifQ.rJcFIG214AriISLbB6B5aw', {
            maxZoom: 18,
            attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors, ' +
                '<a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, ' +
                'Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
            id: 'mapbox/streets-v11',
            tileSize: 512,
            zoomOffset: -1
        }).addTo(map);
    },
    initMapControls : function() {
        drawnItems = new L.FeatureGroup();
        map.addLayer(drawnItems);
        map.addControl(
            new L.Control.Draw({
            draw: {
                circle: false,
                polyline: false,
                rectangle: false,
                circlemarker: false,
            },
            edit: {
                featureGroup: drawnItems
            }
        })
        );
    },
    bindEvents : function() {
        let max_level_circle_slider = document.getElementById("max_level_circle");
        let max_level_circle_value = document.getElementById("max_level_circle_value");
        max_level_circle_value.innerHTML = max_level_circle_slider.value;
        max_level_circle_slider.oninput = function() {
            max_level_circle_value.innerHTML = this.value;
        };
        let maxLevelGeoJsonInput = document.getElementById("max_level_geojson");
        let max_level_geojson_value = document.getElementById("max_level_geojson_value");
        max_level_geojson_value.innerHTML = maxLevelGeoJsonInput.value;
        maxLevelGeoJsonInput.oninput = function() {
            self.GeoJsonToMap();
            max_level_geojson_value.innerHTML = this.value;
        };

        let minLevelGeoJsonInput = document.getElementById("min_level_geojson");
        let min_level_geojson_value = document.getElementById("min_level_geojson_value");
        min_level_geojson_value.innerHTML = minLevelGeoJsonInput.value;
        minLevelGeoJsonInput.oninput = function() {
            self.GeoJsonToMap();
            min_level_geojson_value.innerHTML = this.value;
        };

        document.getElementById("lat").oninput = function() {
            self.checkPointIntersection();
        };
        document.getElementById("lng").oninput = function() {
            self.checkPointIntersection();
        };

        map.on('click', this.onMapClick);
        geoJsonEditor.on("change", self.onGeoJsonChange);

        map.on('draw:drawstart', function () {
            map.off('click', self.onMapClick);
            geoJsonEditor.off("change", self.onGeoJsonChange);
        });
        map.on('draw:drawstop', function () {
            map.on('click', self.onMapClick);
            geoJsonEditor.on("change", self.onGeoJsonChange);
        });
        map.on('draw:editstart', function () {
            map.off('click', self.onMapClick);
            geoJsonEditor.off("change", self.onGeoJsonChange);
        });
        map.on('draw:editstop', function () {
            map.on('click', self.onMapClick);
            geoJsonEditor.on("change", self.onGeoJsonChange);
        });
        map.on('draw:deletestart', function () {
            map.off('click', self.onMapClick);
            geoJsonEditor.off("change", self.onGeoJsonChange);
        });
        map.on('draw:deletestop', function () {
            map.on('click', self.onMapClick);
            geoJsonEditor.on("change", self.onGeoJsonChange);
        });
        map.on(L.Draw.Event.CREATED, self.onDrawCreated);
        map.on('draw:edited', self.onDrawEdited);
        map.on('draw:deleted', function () {
            self.clear();
        });
    },
    onMapClick : function(e) {

        let lat = e.latlng.lat;
        let lng = e.latlng.lng;

        document.getElementById("lat").value = lat;
        document.getElementById("lng").value = lng;

        self.checkPointIntersection();
    },
    checkPointIntersection: function() {
        let lat = document.getElementById("lat").value;
        let lng = document.getElementById("lng").value;

        if (marker) {
            map.removeLayer(marker)
        }
        if (circleLayer) {
            map.removeLayer(circleLayer)
        }
        marker = L.marker([lat, lng]).addTo(map);

        let max_level_circle = document.getElementById("max_level_circle").value;

        let radius = document.getElementById("radius").value;

        if (radius > 0) {
            circleLayer = L.circle([lat, lng], {radius: radius}).addTo(map);
        }

        self.removeCircleCells();

        let tokens = document.getElementById("cell_tokens").value;
        if (tokens !== '') {
            let params = "lat=" + lat + "&lng=" + lng  + "&max_level_circle=" + max_level_circle + "&radius=" + radius + "&tokens=" + tokens;
            self.postRequest(params, checkPointUrl, function (response) {
                let res = JSON.parse(response);
                let intersectsPointElem = document.getElementById("intersects_with_point");
                intersectsPointElem.innerHTML = "Features intersects with point: " + res.intersects_with_point;
                intersectsPointElem.className = "";
                intersectsPointElem.classList.add(res.intersects_with_point ? "success" : "error");

                let intersectsCircleElem = document.getElementById("intersects_with_circle");
                intersectsCircleElem.innerHTML = "Features intersects with circle: " + res.intersects_with_circle;
                intersectsCircleElem.className = "";
                intersectsCircleElem.classList.add(res.intersects_with_circle ? "success" : "error");

                if (radius > 0) {
                    let s2_cells = res.cells;
                    for (let i = 0; i < s2_cells.length; i++) {
                        circle_cells.push(L.polygon(s2_cells[i], {color: 'black'}).addTo(map));
                    }
                }
            });
        }
    },
    onDrawCreated : function(e) {
        self.clear();

        let type = e.layerType;
        if (type === 'polygon' || type === 'marker') {
            let layer = e.layer;
            geoJsonLayers.push(layer);
            let json = {
                "type": "FeatureCollection",
                "features": [layer.toGeoJSON(14)]
            };
            geoJsonEditor.setValue(JSON.stringify(json,null, 2));
            self.regionCover();
            drawnItems.addLayer(layer);
            if (type === 'polygon') {
                map.fitBounds(layer.getBounds());
            }
        }
    },
    onDrawEdited : function(e){
        self.clear();
        let layers = e.layers;
        geoJsonEditor.setValue(JSON.stringify(layers.toGeoJSON(14),null, 2));
        self.regionCover();
        layers.eachLayer(function(layer) {
            geoJsonLayers.push(layer);
            drawnItems.addLayer(layer);
        });
    },
    onGeoJsonChange : function() {
        document.getElementById("cell_tokens").value = '';
        self.GeoJsonToMap();
    },
    GeoJsonToMap : function() {
        let v = geoJsonEditor.getValue();

        self.removeGeoJsonLayers();

        if (v !== '') {
            self.regionCover();
            let layer = L.geoJSON(JSON.parse(v), {}).eachLayer(function (layer) {
                geoJsonLayers.push(layer);
                drawnItems.addLayer(layer);
            });
            map.fitBounds(layer.getBounds());
        }
    },
    regionCover : function() {
        self.removePolygonCells();

        let max_level_geojson = document.getElementById("max_level_geojson").value;
        let min_level_geojson = document.getElementById("min_level_geojson").value;

        let params = "max_level_geojson=" + max_level_geojson + "&min_level_geojson=" + min_level_geojson  + "&geojson=" + geoJsonEditor.getValue().trim();
        self.postRequest(params, coverUrl, function (response) {
            let res = JSON.parse(response);
            document.getElementById("cell_tokens").value = res.cell_tokens;
            let s2cells = res.cells;
            for (let i = 0; i < s2cells.length; i++) {
                polygon_cells.push(L.polygon(s2cells[i], {color: 'red'}).addTo(map));
            }
        });
    },
    postRequest : function(params, url, callback) {
        let xmlHttp = new XMLHttpRequest();
        xmlHttp.onreadystatechange = function () {
            if (xmlHttp.readyState === 4 && xmlHttp.status === 200) {
                callback(xmlHttp.responseText);
            }
        };
        xmlHttp.open(
            "POST",
            url,
            true);
        xmlHttp.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
        xmlHttp.send(params);
    },
    clear : function() {

        self.removeGeoJsonLayers();
        document.getElementById("geoJsonInput").value = '';
        geoJsonEditor.setValue("");

        self.removePolygonCells();
    },
    removePolygonCells : function () {
        for (let i = 0; i < polygon_cells.length; i++) {
            map.removeLayer(polygon_cells[i]);
        }
        polygon_cells = [];
    },
    removeCircleCells : function () {
        for (var i = 0; i < circle_cells.length; i++) {
            map.removeLayer(circle_cells[i]);
        }
        circle_cells = [];
    },
    removeGeoJsonLayers : function () {
        try {
            let layer;
            for (let i = 0; i < geoJsonLayers.length; i++) {
                layer = geoJsonLayers[i];
                drawnItems.removeLayer(geoJsonLayers[i]);
            }
        } catch (e) {
            console.log(e.message);
        }
        geoJsonLayers = [];
    },
};

document.addEventListener("DOMContentLoaded", function() {
    app = s2_geojson.Init();
});


