package identicon

import "testing"

func TestSimpleIdenticonCreate(t *testing.T) {
	id := SimpleIdenticon{}
	svgElement := id.Create("test input", 8)
	if svgElement.ToString() != expectedSVG {
		t.Errorf("created svg should be: %v , instead got: %v", expectedSVG, svgElement.ToString())
	}
}

var expectedSVG string = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" shape-rendering="crispEdges"><rect fill="#00000000" x="0" y="0" width="1" height="1" /><rect fill="#00000000" x="15" y="0" width="1" height="1" /><rect fill="#00000000" x="0" y="15" width="1" height="1" /><rect fill="#00000000" x="15" y="15" width="1" height="1" /><rect fill="#ccff00" x="1" y="0" width="1" height="1" /><rect fill="#ccff00" x="14" y="0" width="1" height="1" /><rect fill="#ccff00" x="1" y="15" width="1" height="1" /><rect fill="#ccff00" x="14" y="15" width="1" height="1" /><rect fill="#00ff66" x="2" y="0" width="1" height="1" /><rect fill="#00ff66" x="13" y="0" width="1" height="1" /><rect fill="#00ff66" x="2" y="15" width="1" height="1" /><rect fill="#00ff66" x="13" y="15" width="1" height="1" /><rect fill="#00000000" x="3" y="0" width="1" height="1" /><rect fill="#00000000" x="12" y="0" width="1" height="1" /><rect fill="#00000000" x="3" y="15" width="1" height="1" /><rect fill="#00000000" x="12" y="15" width="1" height="1" /><rect fill="#00ff66" x="4" y="0" width="1" height="1" /><rect fill="#00ff66" x="11" y="0" width="1" height="1" /><rect fill="#00ff66" x="4" y="15" width="1" height="1" /><rect fill="#00ff66" x="11" y="15" width="1" height="1" /><rect fill="#00000000" x="5" y="0" width="1" height="1" /><rect fill="#00000000" x="10" y="0" width="1" height="1" /><rect fill="#00000000" x="5" y="15" width="1" height="1" /><rect fill="#00000000" x="10" y="15" width="1" height="1" /><rect fill="#00ff66" x="6" y="0" width="1" height="1" /><rect fill="#00ff66" x="9" y="0" width="1" height="1" /><rect fill="#00ff66" x="6" y="15" width="1" height="1" /><rect fill="#00ff66" x="9" y="15" width="1" height="1" /><rect fill="#00ff66" x="7" y="0" width="1" height="1" /><rect fill="#00ff66" x="8" y="0" width="1" height="1" /><rect fill="#00ff66" x="7" y="15" width="1" height="1" /><rect fill="#00ff66" x="8" y="15" width="1" height="1" /><rect fill="#00ff66" x="0" y="1" width="1" height="1" /><rect fill="#00ff66" x="15" y="1" width="1" height="1" /><rect fill="#00ff66" x="0" y="14" width="1" height="1" /><rect fill="#00ff66" x="15" y="14" width="1" height="1" /><rect fill="#ccff00" x="1" y="1" width="1" height="1" /><rect fill="#ccff00" x="14" y="1" width="1" height="1" /><rect fill="#ccff00" x="1" y="14" width="1" height="1" /><rect fill="#ccff00" x="14" y="14" width="1" height="1" /><rect fill="#ccff00" x="2" y="1" width="1" height="1" /><rect fill="#ccff00" x="13" y="1" width="1" height="1" /><rect fill="#ccff00" x="2" y="14" width="1" height="1" /><rect fill="#ccff00" x="13" y="14" width="1" height="1" /><rect fill="#00000000" x="3" y="1" width="1" height="1" /><rect fill="#00000000" x="12" y="1" width="1" height="1" /><rect fill="#00000000" x="3" y="14" width="1" height="1" /><rect fill="#00000000" x="12" y="14" width="1" height="1" /><rect fill="#00000000" x="4" y="1" width="1" height="1" /><rect fill="#00000000" x="11" y="1" width="1" height="1" /><rect fill="#00000000" x="4" y="14" width="1" height="1" /><rect fill="#00000000" x="11" y="14" width="1" height="1" /><rect fill="#00000000" x="5" y="1" width="1" height="1" /><rect fill="#00000000" x="10" y="1" width="1" height="1" /><rect fill="#00000000" x="5" y="14" width="1" height="1" /><rect fill="#00000000" x="10" y="14" width="1" height="1" /><rect fill="#00ff66" x="6" y="1" width="1" height="1" /><rect fill="#00ff66" x="9" y="1" width="1" height="1" /><rect fill="#00ff66" x="6" y="14" width="1" height="1" /><rect fill="#00ff66" x="9" y="14" width="1" height="1" /><rect fill="#00ff66" x="7" y="1" width="1" height="1" /><rect fill="#00ff66" x="8" y="1" width="1" height="1" /><rect fill="#00ff66" x="7" y="14" width="1" height="1" /><rect fill="#00ff66" x="8" y="14" width="1" height="1" /><rect fill="#ccff00" x="0" y="2" width="1" height="1" /><rect fill="#ccff00" x="15" y="2" width="1" height="1" /><rect fill="#ccff00" x="0" y="13" width="1" height="1" /><rect fill="#ccff00" x="15" y="13" width="1" height="1" /><rect fill="#00ff66" x="1" y="2" width="1" height="1" /><rect fill="#00ff66" x="14" y="2" width="1" height="1" /><rect fill="#00ff66" x="1" y="13" width="1" height="1" /><rect fill="#00ff66" x="14" y="13" width="1" height="1" /><rect fill="#ccff00" x="2" y="2" width="1" height="1" /><rect fill="#ccff00" x="13" y="2" width="1" height="1" /><rect fill="#ccff00" x="2" y="13" width="1" height="1" /><rect fill="#ccff00" x="13" y="13" width="1" height="1" /><rect fill="#00ff66" x="3" y="2" width="1" height="1" /><rect fill="#00ff66" x="12" y="2" width="1" height="1" /><rect fill="#00ff66" x="3" y="13" width="1" height="1" /><rect fill="#00ff66" x="12" y="13" width="1" height="1" /><rect fill="#ccff00" x="4" y="2" width="1" height="1" /><rect fill="#ccff00" x="11" y="2" width="1" height="1" /><rect fill="#ccff00" x="4" y="13" width="1" height="1" /><rect fill="#ccff00" x="11" y="13" width="1" height="1" /><rect fill="#00000000" x="5" y="2" width="1" height="1" /><rect fill="#00000000" x="10" y="2" width="1" height="1" /><rect fill="#00000000" x="5" y="13" width="1" height="1" /><rect fill="#00000000" x="10" y="13" width="1" height="1" /><rect fill="#00000000" x="6" y="2" width="1" height="1" /><rect fill="#00000000" x="9" y="2" width="1" height="1" /><rect fill="#00000000" x="6" y="13" width="1" height="1" /><rect fill="#00000000" x="9" y="13" width="1" height="1" /><rect fill="#00000000" x="7" y="2" width="1" height="1" /><rect fill="#00000000" x="8" y="2" width="1" height="1" /><rect fill="#00000000" x="7" y="13" width="1" height="1" /><rect fill="#00000000" x="8" y="13" width="1" height="1" /><rect fill="#00ff66" x="0" y="3" width="1" height="1" /><rect fill="#00ff66" x="15" y="3" width="1" height="1" /><rect fill="#00ff66" x="0" y="12" width="1" height="1" /><rect fill="#00ff66" x="15" y="12" width="1" height="1" /><rect fill="#00000000" x="1" y="3" width="1" height="1" /><rect fill="#00000000" x="14" y="3" width="1" height="1" /><rect fill="#00000000" x="1" y="12" width="1" height="1" /><rect fill="#00000000" x="14" y="12" width="1" height="1" /><rect fill="#00ff66" x="2" y="3" width="1" height="1" /><rect fill="#00ff66" x="13" y="3" width="1" height="1" /><rect fill="#00ff66" x="2" y="12" width="1" height="1" /><rect fill="#00ff66" x="13" y="12" width="1" height="1" /><rect fill="#00000000" x="3" y="3" width="1" height="1" /><rect fill="#00000000" x="12" y="3" width="1" height="1" /><rect fill="#00000000" x="3" y="12" width="1" height="1" /><rect fill="#00000000" x="12" y="12" width="1" height="1" /><rect fill="#00000000" x="4" y="3" width="1" height="1" /><rect fill="#00000000" x="11" y="3" width="1" height="1" /><rect fill="#00000000" x="4" y="12" width="1" height="1" /><rect fill="#00000000" x="11" y="12" width="1" height="1" /><rect fill="#00000000" x="5" y="3" width="1" height="1" /><rect fill="#00000000" x="10" y="3" width="1" height="1" /><rect fill="#00000000" x="5" y="12" width="1" height="1" /><rect fill="#00000000" x="10" y="12" width="1" height="1" /><rect fill="#ccff00" x="6" y="3" width="1" height="1" /><rect fill="#ccff00" x="9" y="3" width="1" height="1" /><rect fill="#ccff00" x="6" y="12" width="1" height="1" /><rect fill="#ccff00" x="9" y="12" width="1" height="1" /><rect fill="#00000000" x="7" y="3" width="1" height="1" /><rect fill="#00000000" x="8" y="3" width="1" height="1" /><rect fill="#00000000" x="7" y="12" width="1" height="1" /><rect fill="#00000000" x="8" y="12" width="1" height="1" /><rect fill="#00ff66" x="0" y="4" width="1" height="1" /><rect fill="#00ff66" x="15" y="4" width="1" height="1" /><rect fill="#00ff66" x="0" y="11" width="1" height="1" /><rect fill="#00ff66" x="15" y="11" width="1" height="1" /><rect fill="#00ff66" x="1" y="4" width="1" height="1" /><rect fill="#00ff66" x="14" y="4" width="1" height="1" /><rect fill="#00ff66" x="1" y="11" width="1" height="1" /><rect fill="#00ff66" x="14" y="11" width="1" height="1" /><rect fill="#00ff66" x="2" y="4" width="1" height="1" /><rect fill="#00ff66" x="13" y="4" width="1" height="1" /><rect fill="#00ff66" x="2" y="11" width="1" height="1" /><rect fill="#00ff66" x="13" y="11" width="1" height="1" /><rect fill="#00ff66" x="3" y="4" width="1" height="1" /><rect fill="#00ff66" x="12" y="4" width="1" height="1" /><rect fill="#00ff66" x="3" y="11" width="1" height="1" /><rect fill="#00ff66" x="12" y="11" width="1" height="1" /><rect fill="#00000000" x="4" y="4" width="1" height="1" /><rect fill="#00000000" x="11" y="4" width="1" height="1" /><rect fill="#00000000" x="4" y="11" width="1" height="1" /><rect fill="#00000000" x="11" y="11" width="1" height="1" /><rect fill="#00000000" x="5" y="4" width="1" height="1" /><rect fill="#00000000" x="10" y="4" width="1" height="1" /><rect fill="#00000000" x="5" y="11" width="1" height="1" /><rect fill="#00000000" x="10" y="11" width="1" height="1" /><rect fill="#00000000" x="6" y="4" width="1" height="1" /><rect fill="#00000000" x="9" y="4" width="1" height="1" /><rect fill="#00000000" x="6" y="11" width="1" height="1" /><rect fill="#00000000" x="9" y="11" width="1" height="1" /><rect fill="#00000000" x="7" y="4" width="1" height="1" /><rect fill="#00000000" x="8" y="4" width="1" height="1" /><rect fill="#00000000" x="7" y="11" width="1" height="1" /><rect fill="#00000000" x="8" y="11" width="1" height="1" /><rect fill="#ccff00" x="0" y="5" width="1" height="1" /><rect fill="#ccff00" x="15" y="5" width="1" height="1" /><rect fill="#ccff00" x="0" y="10" width="1" height="1" /><rect fill="#ccff00" x="15" y="10" width="1" height="1" /><rect fill="#00000000" x="1" y="5" width="1" height="1" /><rect fill="#00000000" x="14" y="5" width="1" height="1" /><rect fill="#00000000" x="1" y="10" width="1" height="1" /><rect fill="#00000000" x="14" y="10" width="1" height="1" /><rect fill="#00000000" x="2" y="5" width="1" height="1" /><rect fill="#00000000" x="13" y="5" width="1" height="1" /><rect fill="#00000000" x="2" y="10" width="1" height="1" /><rect fill="#00000000" x="13" y="10" width="1" height="1" /><rect fill="#00000000" x="3" y="5" width="1" height="1" /><rect fill="#00000000" x="12" y="5" width="1" height="1" /><rect fill="#00000000" x="3" y="10" width="1" height="1" /><rect fill="#00000000" x="12" y="10" width="1" height="1" /><rect fill="#00000000" x="4" y="5" width="1" height="1" /><rect fill="#00000000" x="11" y="5" width="1" height="1" /><rect fill="#00000000" x="4" y="10" width="1" height="1" /><rect fill="#00000000" x="11" y="10" width="1" height="1" /><rect fill="#ccff00" x="5" y="5" width="1" height="1" /><rect fill="#ccff00" x="10" y="5" width="1" height="1" /><rect fill="#ccff00" x="5" y="10" width="1" height="1" /><rect fill="#ccff00" x="10" y="10" width="1" height="1" /><rect fill="#ccff00" x="6" y="5" width="1" height="1" /><rect fill="#ccff00" x="9" y="5" width="1" height="1" /><rect fill="#ccff00" x="6" y="10" width="1" height="1" /><rect fill="#ccff00" x="9" y="10" width="1" height="1" /><rect fill="#00000000" x="7" y="5" width="1" height="1" /><rect fill="#00000000" x="8" y="5" width="1" height="1" /><rect fill="#00000000" x="7" y="10" width="1" height="1" /><rect fill="#00000000" x="8" y="10" width="1" height="1" /><rect fill="#00000000" x="0" y="6" width="1" height="1" /><rect fill="#00000000" x="15" y="6" width="1" height="1" /><rect fill="#00000000" x="0" y="9" width="1" height="1" /><rect fill="#00000000" x="15" y="9" width="1" height="1" /><rect fill="#00ff66" x="1" y="6" width="1" height="1" /><rect fill="#00ff66" x="14" y="6" width="1" height="1" /><rect fill="#00ff66" x="1" y="9" width="1" height="1" /><rect fill="#00ff66" x="14" y="9" width="1" height="1" /><rect fill="#00000000" x="2" y="6" width="1" height="1" /><rect fill="#00000000" x="13" y="6" width="1" height="1" /><rect fill="#00000000" x="2" y="9" width="1" height="1" /><rect fill="#00000000" x="13" y="9" width="1" height="1" /><rect fill="#00000000" x="3" y="6" width="1" height="1" /><rect fill="#00000000" x="12" y="6" width="1" height="1" /><rect fill="#00000000" x="3" y="9" width="1" height="1" /><rect fill="#00000000" x="12" y="9" width="1" height="1" /><rect fill="#00ff66" x="4" y="6" width="1" height="1" /><rect fill="#00ff66" x="11" y="6" width="1" height="1" /><rect fill="#00ff66" x="4" y="9" width="1" height="1" /><rect fill="#00ff66" x="11" y="9" width="1" height="1" /><rect fill="#ccff00" x="5" y="6" width="1" height="1" /><rect fill="#ccff00" x="10" y="6" width="1" height="1" /><rect fill="#ccff00" x="5" y="9" width="1" height="1" /><rect fill="#ccff00" x="10" y="9" width="1" height="1" /><rect fill="#00000000" x="6" y="6" width="1" height="1" /><rect fill="#00000000" x="9" y="6" width="1" height="1" /><rect fill="#00000000" x="6" y="9" width="1" height="1" /><rect fill="#00000000" x="9" y="9" width="1" height="1" /><rect fill="#ccff00" x="7" y="6" width="1" height="1" /><rect fill="#ccff00" x="8" y="6" width="1" height="1" /><rect fill="#ccff00" x="7" y="9" width="1" height="1" /><rect fill="#ccff00" x="8" y="9" width="1" height="1" /><rect fill="#ccff00" x="0" y="7" width="1" height="1" /><rect fill="#ccff00" x="15" y="7" width="1" height="1" /><rect fill="#ccff00" x="0" y="8" width="1" height="1" /><rect fill="#ccff00" x="15" y="8" width="1" height="1" /><rect fill="#00000000" x="1" y="7" width="1" height="1" /><rect fill="#00000000" x="14" y="7" width="1" height="1" /><rect fill="#00000000" x="1" y="8" width="1" height="1" /><rect fill="#00000000" x="14" y="8" width="1" height="1" /><rect fill="#00000000" x="2" y="7" width="1" height="1" /><rect fill="#00000000" x="13" y="7" width="1" height="1" /><rect fill="#00000000" x="2" y="8" width="1" height="1" /><rect fill="#00000000" x="13" y="8" width="1" height="1" /><rect fill="#00000000" x="3" y="7" width="1" height="1" /><rect fill="#00000000" x="12" y="7" width="1" height="1" /><rect fill="#00000000" x="3" y="8" width="1" height="1" /><rect fill="#00000000" x="12" y="8" width="1" height="1" /><rect fill="#ccff00" x="4" y="7" width="1" height="1" /><rect fill="#ccff00" x="11" y="7" width="1" height="1" /><rect fill="#ccff00" x="4" y="8" width="1" height="1" /><rect fill="#ccff00" x="11" y="8" width="1" height="1" /><rect fill="#ccff00" x="5" y="7" width="1" height="1" /><rect fill="#ccff00" x="10" y="7" width="1" height="1" /><rect fill="#ccff00" x="5" y="8" width="1" height="1" /><rect fill="#ccff00" x="10" y="8" width="1" height="1" /><rect fill="#ccff00" x="6" y="7" width="1" height="1" /><rect fill="#ccff00" x="9" y="7" width="1" height="1" /><rect fill="#ccff00" x="6" y="8" width="1" height="1" /><rect fill="#ccff00" x="9" y="8" width="1" height="1" /><rect fill="#00ff66" x="7" y="7" width="1" height="1" /><rect fill="#00ff66" x="8" y="7" width="1" height="1" /><rect fill="#00ff66" x="7" y="8" width="1" height="1" /><rect fill="#00ff66" x="8" y="8" width="1" height="1" /></svg>`
