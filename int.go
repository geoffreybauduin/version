package version

import (
	"strconv"
	"strings"
)

type versionInt struct {
	nbr []int
}

func (v *versionInt) parse(versionString string) error {
	v.nbr = make([]int, 0)
	splitted := strings.Split(versionString, ".")
	for _, spl := range splitted {
		intVal, err := strconv.Atoi(spl)
		if err != nil {
			return err
		}
		v.nbr = append(v.nbr, intVal)
	}
	return nil
}

type cmpOperator struct {
	bothExhausted cmpDecision
	cmp           func(int, int) cmpDecision
}

type cmpDecision uint8

const (
	cmpDecision_STOP_FALSE cmpDecision = iota
	cmpDecision_STOP_TRUE
	cmpDecision_CONTINUE
)

var (
	operators = map[string]cmpOperator{
		"=": {
			bothExhausted: cmpDecision_STOP_TRUE,
			cmp: func(v1 int, v2 int) cmpDecision {
				if v1 == v2 {
					return cmpDecision_CONTINUE
				}
				return cmpDecision_STOP_FALSE
			},
		},
		">=": {
			bothExhausted: cmpDecision_STOP_TRUE,
			cmp: func(v1 int, v2 int) cmpDecision {
				if v1 > v2 {
					return cmpDecision_STOP_TRUE
				} else if v1 == v2 {
					return cmpDecision_CONTINUE
				}
				return cmpDecision_STOP_FALSE
			},
		},
		"<=": {
			bothExhausted: cmpDecision_STOP_TRUE,
			cmp: func(v1 int, v2 int) cmpDecision {
				if v1 < v2 {
					return cmpDecision_STOP_TRUE
				} else if v1 == v2 {
					return cmpDecision_CONTINUE
				}
				return cmpDecision_STOP_FALSE
			},
		},
		">": {
			bothExhausted: cmpDecision_STOP_FALSE,
			cmp: func(v1 int, v2 int) cmpDecision {
				if v1 > v2 {
					return cmpDecision_STOP_TRUE
				} else if v1 == v2 {
					return cmpDecision_CONTINUE
				}
				return cmpDecision_STOP_FALSE
			},
		},

		"<": {
			bothExhausted: cmpDecision_STOP_FALSE,
			cmp: func(v1 int, v2 int) cmpDecision {
				if v1 < v2 {
					return cmpDecision_STOP_TRUE
				} else if v1 == v2 {
					return cmpDecision_CONTINUE
				}
				return cmpDecision_STOP_FALSE
			},
		},
	}
)

func (v *versionInt) Number() interface{} {
	return v.nbr
}

func (v *versionInt) Is(condition string) bool {
	var operator string
	for op := range operators {
		if strings.HasPrefix(condition, op) && len(operator) < len(op) {
			operator = op
		}
	}
	condition = condition[len(operator):]
	if operator == "" {
		operator = "="
	}
	oth, err := New(condition)
	if err != nil { // invalid condition -> bye
		return false
	}
	i := 0
	v1 := v.nbr
	v2 := oth.(*versionInt).nbr
	for {
		dec := getDecision(operator, v1, v2, i)
		switch dec {
		case cmpDecision_STOP_FALSE:
			return false
		case cmpDecision_STOP_TRUE:
			return true
		}
		i++
	}
}

func getDecision(operator string, v1, v2 []int, i int) cmpDecision {
	v1exhausted := len(v1) <= i
	v2exhausted := len(v2) <= i
	if v1exhausted || v2exhausted {
		return operators[operator].bothExhausted
	}

	return operators[operator].cmp(v1[i], v2[i])
}
