package main

import (
	"io"
	"log"

	geo "github.com/paulmach/go.geo"
	"github.com/qedus/osmpbf"
)

func hasTags(tags map[string]string) bool {
	n := len(tags)
	return n != 0
}

func process(decoder *osmpbf.Decoder) {
	nodes = make(map[int64]*osmpbf.Node)
	ways = make(map[int64]*osmpbf.Node)
	relations = make(map[int64]*osmpbf.Relation)

	var nc, wc, rc uint64
	for {
		if v, err := decoder.Decode(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			switch v := v.(type) {
			case *osmpbf.Node:
				onNode(v)
				nc++
			case *osmpbf.Way:
				onWay(v)
				wc++
			case *osmpbf.Relation:
				onRelation(v)
				rc++
			default:
				log.Fatalf("unknown type %T\n", v)
			}
		}
	}

	log.Printf("Nodes: %d, Ways: %d, Relations: %d\n", nc, wc, rc)
}

func onNode(node *osmpbf.Node) {
	nodes[node.ID] = node
}

func onWay(way *osmpbf.Way) {
	ways[way.ID] = wayToNode(nodes, way)
}

func onRelation(relation *osmpbf.Relation) {
	relations[relation.ID] = relation
}

func wayToNode(nodes map[int64]*osmpbf.Node, way *osmpbf.Way) *osmpbf.Node {
	points := geo.NewPointSet()

	node := osmpbf.Node{
		ID:   way.ID,
		Tags: way.Tags,
	}

	for _, each := range way.NodeIDs {
		point, pointPresent := nodes[each]

		if !pointPresent {
			log.Println("denormalize failed for way:", way.ID, "node not found:", each)
			return nil
		}

		points.Push(geo.NewPoint(point.Lon, point.Lat))
	}

	centroid := points.GeoCentroid()
	node.Lat = centroid.Lat()
	node.Lon = centroid.Lng()

	return &node
}
