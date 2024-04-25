package components

import "gitflic.ru/project/physicist2018/aerosol-decomposition/utlis"

type OpticalCoefs struct {
	AerosolMode `json:"aerosol_mode,omitempty"`
	Rh          int    `json:"rh"`
	Wvl         Vector `json:"wvl,omitempty"`
	Ext         Vector `json:"ext,omitempty"`
	Bck         Vector `json:"bck,omitempty"`
	B22         Vector `json:"b_22,omitempty"`
	MRe         Vector `json:"m_re,omitempty"`
	MIm         Vector `json:"m_im,omitempty"`
}

type OpticalDB []OpticalCoefs

// AerosolModeMixItem - элемент смеси частиц
type AerosolModeMixItem struct {
	OpticalCoefs `json:"optical_coefs,omitempty"`
	N            float64 `json:"n,omitempty"`
}

func (am AerosolModeMixItem) Value(r []float64) []float64 {
	ret := am.AerosolMode.Value(r)
	ret = utlis.Scale(am.N, 0, ret)
	return ret
}

func (am AerosolModeMixItem) MeanRadius() float64 {
	return am.AerosolMode.MeanRadius()
}

func (am AerosolModeMixItem) Area() float64 {
	return am.N * am.AerosolMode.Area()
}

func (am AerosolModeMixItem) Volume() float64 {
	return am.N * am.AerosolMode.Volume()
}

func (am AerosolModeMixItem) EffectiveRadius() float64 {
	return am.AerosolMode.EffectiveRadius()
}

func (am AerosolModeMixItem) Extinction() Vector {
	ret := make(Vector, DIM_SIZE)
	for i, ext := range am.Ext {
		ret[i] = am.N * ext
	}
	return ret
}

func (am AerosolModeMixItem) Backscatter() Vector {
	ret := make(Vector, DIM_SIZE)
	for i, bck := range am.Bck {
		ret[i] = am.N * bck
	}
	return ret
}

func (am AerosolModeMixItem) Backscatter22() Vector {
	ret := make(Vector, DIM_SIZE)
	for i, b22 := range am.B22 {
		ret[i] = am.N * b22
	}
	return ret
}

func (am AerosolModeMixItem) RefrReIdx() Vector {
	ret := make(Vector, DIM_SIZE)
	vol := am.Volume()
	for i, m := range am.MRe {
		ret[i] = vol * m
	}
	return ret
}

func (am AerosolModeMixItem) RefrImIdx() Vector {
	ret := make(Vector, DIM_SIZE)
	vol := am.Volume()
	for i, m := range am.MRe {
		ret[i] = vol * m
	}
	return ret
}
