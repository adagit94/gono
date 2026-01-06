package goter

import (
	"slices"
	"strings"
	"github.com/adagit94/gono/gotils"
)

func genSegConfs(segs []string) []segmentConf {
	segsConfs := gotils.MapSlice(segs, func(seg string, i int) segmentConf {
		isDyn := strings.HasPrefix(seg, ":")

		if isDyn {
			extractedSeg, _ := strings.CutPrefix(seg, ":")
			seg = extractedSeg
		}

		return segmentConf{segment: seg, static: !isDyn}
	})

	return segsConfs
}

func sortRoutes(confs []routeConf) {
	slices.SortFunc(confs, func(a, b routeConf) int {
		aSegsLen, bSegsLen := len(a.segments), len(b.segments)
		minSegs := min(aSegsLen, bSegsLen)

		for i := range minSegs {
			aSeg, bSeg := a.segments[i], b.segments[i]

			if (aSeg.static && bSeg.static) || (!aSeg.static && !bSeg.static) {
				continue
			}

			if aSeg.static {
				return -1
			}

			if bSeg.static {
				return 1
			}
		}

		return aSegsLen - bSegsLen
	})
}
