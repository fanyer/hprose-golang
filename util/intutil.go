/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * util/intutil.go                                        *
 *                                                        *
 * int to bytes util for Go.                              *
 *                                                        *
 * LastModified: Aug 17, 2016                             *
 * Author: Ma Bingyao <andot@hprose.com>                  *
 *                                                        *
\**********************************************************/

package util

import "math"

var digits = [...]byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

var digit3Table = ([]byte)("" +
	"000001002003004005006007008009010011012013014015016017018019" +
	"020021022023024025026027028029030031032033034035036037038039" +
	"040041042043044045046047048049050051052053054055056057058059" +
	"060061062063064065066067068069070071072073074075076077078079" +
	"080081082083084085086087088089090091092093094095096097098099" +
	"100101102103104105106107108109110111112113114115116117118119" +
	"120121122123124125126127128129130131132133134135136137138139" +
	"140141142143144145146147148149150151152153154155156157158159" +
	"160161162163164165166167168169170171172173174175176177178179" +
	"180181182183184185186187188189190191192193194195196197198199" +
	"200201202203204205206207208209210211212213214215216217218219" +
	"220221222223224225226227228229230231232233234235236237238239" +
	"240241242243244245246247248249250251252253254255256257258259" +
	"260261262263264265266267268269270271272273274275276277278279" +
	"280281282283284285286287288289290291292293294295296297298299" +
	"300301302303304305306307308309310311312313314315316317318319" +
	"320321322323324325326327328329330331332333334335336337338339" +
	"340341342343344345346347348349350351352353354355356357358359" +
	"360361362363364365366367368369370371372373374375376377378379" +
	"380381382383384385386387388389390391392393394395396397398399" +
	"400401402403404405406407408409410411412413414415416417418419" +
	"420421422423424425426427428429430431432433434435436437438439" +
	"440441442443444445446447448449450451452453454455456457458459" +
	"460461462463464465466467468469470471472473474475476477478479" +
	"480481482483484485486487488489490491492493494495496497498499" +
	"500501502503504505506507508509510511512513514515516517518519" +
	"520521522523524525526527528529530531532533534535536537538539" +
	"540541542543544545546547548549550551552553554555556557558559" +
	"560561562563564565566567568569570571572573574575576577578579" +
	"580581582583584585586587588589590591592593594595596597598599" +
	"600601602603604605606607608609610611612613614615616617618619" +
	"620621622623624625626627628629630631632633634635636637638639" +
	"640641642643644645646647648649650651652653654655656657658659" +
	"660661662663664665666667668669670671672673674675676677678679" +
	"680681682683684685686687688689690691692693694695696697698699" +
	"700701702703704705706707708709710711712713714715716717718719" +
	"720721722723724725726727728729730731732733734735736737738739" +
	"740741742743744745746747748749750751752753754755756757758759" +
	"760761762763764765766767768769770771772773774775776777778779" +
	"780781782783784785786787788789790791792793794795796797798799" +
	"800801802803804805806807808809810811812813814815816817818819" +
	"820821822823824825826827828829830831832833834835836837838839" +
	"840841842843844845846847848849850851852853854855856857858859" +
	"860861862863864865866867868869870871872873874875876877878879" +
	"880881882883884885886887888889890891892893894895896897898899" +
	"900901902903904905906907908909910911912913914915916917918919" +
	"920921922923924925926927928929930931932933934935936937938939" +
	"940941942943944945946947948949950951952953954955956957958959" +
	"960961962963964965966967968969970971972973974975976977978979" +
	"980981982983984985986987988989990991992993994995996997998999")

var digit2Table = ([]byte)("" +
	"0001020304050607080910111213141516171819" +
	"2021222324252627282930313233343536373839" +
	"4041424344454647484950515253545556575859" +
	"6061626364656667686970717273747576777879" +
	"8081828384858687888990919293949596979899")

var minInt32Buf = [...]byte{'-', '2', '1', '4', '7', '4', '8', '3', '6', '4', '8'}

var minInt64Buf = [...]byte{
	'-', '9', '2', '2', '3', '3', '7', '2', '0', '3',
	'6', '8', '5', '4', '7', '7', '5', '8', '0', '8'}

// GetIntBytes returns the []byte representation of i in base 10.
func GetIntBytes(i int64) []byte {
	if (i >= 0) && (i <= 9) {
		return digits[i : i+1]
	}
	if (i >= 10) && (i < 100) {
		p := 2 * i
		return digit2Table[p : p+2]
	}
	if (i >= 100) && (i < 1000) {
		p := 3 * i
		return digit3Table[p : p+3]
	}
	if i == math.MinInt64 {
		return minInt64Buf[:]
	}
	var buf [20]byte
	var sign byte
	if i < 0 {
		sign = '-'
		i = -i
	}
	off := 20
	var q, p int64
	digit3 := digit3Table
	digit2 := digit2Table
	for i >= 100 {
		q = i / 1000
		p = (i - (q * 1000)) * 3
		i = q
		off -= 3
		buf[off] = digit3[p]
		buf[off+1] = digit3[p+1]
		buf[off+2] = digit3[p+2]
	}
	for i >= 10 {
		q = i / 100
		p = (i - (q * 100)) * 2
		i = q
		off -= 2
		buf[off] = digit2[p]
		buf[off+1] = digit2[p+1]
	}
	if i > 0 {
		off--
		buf[off] = digits[i]
	}
	if sign == '-' {
		off--
		buf[off] = sign
	}
	return buf[off:]
}

// GetUintBytes returns the []byte representation of i in base 10.
func GetUintBytes(i uint64) []byte {
	if (i >= 0) && (i <= 9) {
		return digits[i : i+1]
	}
	if (i >= 10) && (i < 100) {
		p := 2 * i
		return digit2Table[p : p+2]
	}
	if (i >= 100) && (i < 1000) {
		p := 3 * i
		return digit3Table[p : p+3]
	}
	var buf [20]byte
	off := 20
	var q, p uint64
	digit3 := digit3Table
	digit2 := digit2Table
	for i >= 100 {
		q = i / 1000
		p = (i - (q * 1000)) * 3
		i = q
		off -= 3
		buf[off] = digit3[p]
		buf[off+1] = digit3[p+1]
		buf[off+2] = digit3[p+2]
	}
	for i >= 10 {
		q = i / 100
		p = (i - (q * 100)) * 2
		i = q
		off -= 2
		buf[off] = digit2[p]
		buf[off+1] = digit2[p+1]
	}
	if i > 0 {
		off--
		buf[off] = digits[i]
	}
	return buf[off:]
}
