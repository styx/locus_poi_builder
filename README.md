# Download map
http://download.geofabrik.de/europe.html

# Choose required region
http://bboxfinder.com/#46.9,12.2,48.0,13.6
https://www.openstreetmap.org/export#map=9/47.5/12.9 - Export

# Drop not required data and convert everything to nodes
./osmconvert ./austria-latest.osm.pbf --all-to-nodes --drop-author --drop-version --max-objects=100000000 -b=12.2,46.9,13.6,48 -o=zalsb-zell-latest.osm.pbf

# Edit country code in DB