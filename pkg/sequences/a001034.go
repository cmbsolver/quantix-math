package sequences

import (
	"fmt"
	"math/big"
)

// A001034: Orders of noncyclic simple groups (without repetitions).
// URL: https://oeis.org/A001034
// Data source: https://oeis.org/A001034/b001034.txt

// GetA001034Sequence returns the A001034 sequence based on either positional or max-value semantics.
func GetA001034Sequence(maxNumber *big.Int, isPositional bool) (*NumericSequence, error) {
	if isPositional {
		return GetA001034AtPosition(maxNumber)
	}
	return GenerateA001034Sequence(maxNumber)
}

// GenerateA001034Sequence generates all A001034 terms less than or equal to maxNumber.
func GenerateA001034Sequence(maxNumber *big.Int) (*NumericSequence, error) {
	if maxNumber.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("max number cannot be negative")
	}

	sequence := make([]*big.Int, 0)
	for _, term := range a001034Terms {
		if term.Cmp(maxNumber) > 0 {
			break
		}
		sequence = append(sequence, new(big.Int).Set(term))
	}

	if len(sequence) == 0 {
		return &NumericSequence{
			Name:     "Orders of noncyclic simple groups (A001034)",
			Number:   maxNumber,
			Sequence: sequence,
			Result:   big.NewInt(0),
		}, nil
	}

	return &NumericSequence{
		Name:     "Orders of noncyclic simple groups (A001034)",
		Number:   maxNumber,
		Sequence: sequence,
		Result:   new(big.Int).Set(sequence[len(sequence)-1]),
	}, nil
}

// GetA001034AtPosition returns the n-th A001034 term, with OEIS offset 1.
func GetA001034AtPosition(n *big.Int) (*NumericSequence, error) {
	if n.Cmp(big.NewInt(1)) < 0 {
		return nil, fmt.Errorf("position must be >= 1")
	}

	if !n.IsInt64() {
		return nil, fmt.Errorf("position is too large")
	}

	index := int(n.Int64()) - 1
	if index < 0 || index >= len(a001034Terms) {
		return nil, fmt.Errorf("position %s exceeds precomputed range (%d terms)", n.String(), len(a001034Terms))
	}

	result := new(big.Int).Set(a001034Terms[index])
	return &NumericSequence{
		Name:     "Orders of noncyclic simple groups (A001034)",
		Number:   n,
		Sequence: []*big.Int{result},
		Result:   result,
	}, nil
}

var a001034Terms = []*big.Int{
	big.NewInt(60), big.NewInt(168), big.NewInt(360), big.NewInt(504), big.NewInt(660), big.NewInt(1092), big.NewInt(2448), big.NewInt(2520), big.NewInt(3420), big.NewInt(4080),
	big.NewInt(5616), big.NewInt(6048), big.NewInt(6072), big.NewInt(7800), big.NewInt(7920), big.NewInt(9828), big.NewInt(12180), big.NewInt(14880), big.NewInt(20160), big.NewInt(25308),
	big.NewInt(25920), big.NewInt(29120), big.NewInt(32736), big.NewInt(34440), big.NewInt(39732), big.NewInt(51888), big.NewInt(58800), big.NewInt(62400), big.NewInt(74412), big.NewInt(95040),
	big.NewInt(102660), big.NewInt(113460), big.NewInt(126000), big.NewInt(150348), big.NewInt(175560), big.NewInt(178920), big.NewInt(181440), big.NewInt(194472), big.NewInt(246480), big.NewInt(262080),
	big.NewInt(265680), big.NewInt(285852), big.NewInt(352440), big.NewInt(372000), big.NewInt(443520), big.NewInt(456288), big.NewInt(515100), big.NewInt(546312), big.NewInt(604800), big.NewInt(612468),
	big.NewInt(647460), big.NewInt(721392), big.NewInt(885720), big.NewInt(976500), big.NewInt(979200), big.NewInt(1024128), big.NewInt(1123980), big.NewInt(1285608), big.NewInt(1342740), big.NewInt(1451520),
	big.NewInt(1653900), big.NewInt(1721400), big.NewInt(1814400), big.NewInt(1876896), big.NewInt(1934868), big.NewInt(2097024), big.NewInt(2165292), big.NewInt(2328648), big.NewInt(2413320), big.NewInt(2588772),
	big.NewInt(2867580), big.NewInt(2964780), big.NewInt(3265920), big.NewInt(3483840), big.NewInt(3594432), big.NewInt(3822588), big.NewInt(3940200), big.NewInt(4245696), big.NewInt(4680000), big.NewInt(4696860),
	big.NewInt(5515776), big.NewInt(5544672), big.NewInt(5663616), big.NewInt(5848428), big.NewInt(6004380), big.NewInt(6065280), big.NewInt(6324552), big.NewInt(6825840), big.NewInt(6998640), big.NewInt(7174332),
	big.NewInt(7906500), big.NewInt(8487168), big.NewInt(9095592), big.NewInt(9732420), big.NewInt(9951120), big.NewInt(9999360), big.NewInt(10200960), big.NewInt(10626828), big.NewInt(11093880), big.NewInt(11332452),
	big.NewInt(12068640), big.NewInt(12576732), big.NewInt(13685760), big.NewInt(14467068), big.NewInt(15039960), big.NewInt(15331992), big.NewInt(15927348), big.NewInt(16482816), big.NewInt(16776960), big.NewInt(17971200),
	big.NewInt(18132180), big.NewInt(19136208), big.NewInt(19958400), big.NewInt(20176632), big.NewInt(20890788), big.NewInt(21254100), big.NewInt(21993312), big.NewInt(23133960), big.NewInt(23522760), big.NewInt(24715248),
	big.NewInt(25947372), big.NewInt(27219780), big.NewInt(28090752), big.NewInt(29431740), big.NewInt(31285188), big.NewInt(32240400), big.NewInt(32537600), big.NewInt(34208760), big.NewInt(36779820), big.NewInt(37309020),
	big.NewInt(40031280), big.NewInt(40591152), big.NewInt(42302040), big.NewInt(42456960), big.NewInt(42573600), big.NewInt(43468932), big.NewInt(44352000), big.NewInt(45259200), big.NewInt(47721768), big.NewInt(48985860),
	big.NewInt(49626192), big.NewInt(50232960), big.NewInt(50923548), big.NewInt(54950880), big.NewInt(57750408), big.NewInt(59185140), big.NewInt(62125500), big.NewInt(63631512), big.NewInt(65935860), big.NewInt(70710120),
	big.NewInt(70915680), big.NewInt(71527572), big.NewInt(74017680), big.NewInt(79169940), big.NewInt(81833388), big.NewInt(86404068), big.NewInt(89226492), big.NewInt(92109720), big.NewInt(93084420), big.NewInt(96049728),
	big.NewInt(101130708), big.NewInt(104263632), big.NewInt(107460600), big.NewInt(108540600), big.NewInt(111823968), big.NewInt(115172892), big.NewInt(117442248), big.NewInt(118588020), big.NewInt(122070000), big.NewInt(125619480),
	big.NewInt(131687040), big.NewInt(132923532), big.NewInt(134217216), big.NewInt(135419688), big.NewInt(138297600), big.NewInt(139222212), big.NewInt(143095260), big.NewInt(144402060), big.NewInt(152410272), big.NewInt(155144028),
	big.NewInt(159305652), big.NewInt(164969340), big.NewInt(172235700), big.NewInt(174182400), big.NewInt(178200060), big.NewInt(185847120), big.NewInt(192119928), big.NewInt(193709880), big.NewInt(196916052), big.NewInt(197406720),
	big.NewInt(201791340), big.NewInt(205085832), big.NewInt(211341312), big.NewInt(211782000), big.NewInt(212427600), big.NewInt(216898668), big.NewInt(220355160), big.NewInt(227377920), big.NewInt(230944572), big.NewInt(239500800),
	big.NewInt(243721308), big.NewInt(244823040), big.NewInt(251596800), big.NewInt(253130388), big.NewInt(264737160), big.NewInt(266705460), big.NewInt(270178272), big.NewInt(276693420), big.NewInt(278720472), big.NewInt(282804228),
	big.NewInt(284860980), big.NewInt(295294440), big.NewInt(297411240), big.NewInt(310324812), big.NewInt(314710968), big.NewInt(316919460), big.NewInt(321367392), big.NewInt(337262628), big.NewInt(341898480), big.NewInt(344232252),
}
