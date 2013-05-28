package main

import (
	"fmt"

	"format"
)

var str = `	Lorem ipsum dolor sit amet, no vim duis mediocrem. Sit ex labore appetere hendrerit, sea ei partem quaestio, dicit audiam vivendum te per. Ei quo animal aliquam reprimique, ius suas adversarium no. In mel natum dicit interesset, id omnium aliquid antiopam usu. Has ipsum falli legimus ut, pro augue suscipit expetendis ut.

	Eos id malis luptatum intellegebat, enim facer in sit. Eos cu aperiam viderer consetetur, ancillae consetetur mediocritatem id eum, per cu liber ornatus imperdiet. Porro lobortis ius an, menandri antiopam per id. Aeterno eruditi voluptua vel ad, at duo aeque nullam assueverit, et postea labores torquatos eos. Ut sed sint vulputate.

	An affert tritani est. Iuvaret aliquando pro ne, fabulas menandri an eum. Ut ullum postea fastidii cum, homero partem iriure no pro, errem sadipscing id sit. Eam id detraxit sententiae signiferumque. Eum tamquam expetenda imperdiet ea, no mel altera equidem scripserit, idque tincidunt eu per. Nam iudico offendit philosophia an.

	Possit propriae an has, ne agam solum mei. Ea augue vivendum qui, aeterno consetetur cu pri, pro eu ferri possim eloquentiam. Has saepe eirmod abhorreant no, vel diceret vivendo legendos ad. Stet unum doming vix ea, nobis apeirian vim id. Populo scaevola disputando eos ne. Cu nec diam admodum liberavisse, elit vocent petentium ne mel, homero vivendo sit id.

	Affert sanctus definitionem eu duo, qui in purto populo theophrastus, quem euismod electram eam no. Habeo adolescens duo ne, ad usu corpora dignissim torquatos. Eos an dictas virtute, diam eripuit oporteat ea nec. Qui in graece feugiat. Nibh laudem id nam, ea prompta omnesque recteque qui. Stet utamur qualisque te quo.
`

func main() {
	for _, l := range format.Split(40, 4, str) {
		fmt.Println(l)
	}
}

