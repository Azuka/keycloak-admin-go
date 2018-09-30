//go:generate gomodifytags -file $GOFILE -struct MultivaluedHashMap -add-options json=omitempty -add-tags json -w -transform camelcase
//go:generate easyjson -all $GOFILE

package keycloak

// MultivaluedHashMap multivalued map
type MultivaluedHashMap struct {
	Empty      bool    `json:"empty,ommitempty,omitempty"`
	LoadFactor float64 `json:"loadFactor,ommitempty,omitempty"`
	Threshold  int32   `json:"threshold,ommitempty,omitempty"`
}
