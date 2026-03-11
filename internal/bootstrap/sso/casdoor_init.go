package sso

import (
	"apis/configs"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

func InitCasDoor(c *configs.CasdoorConfig) {
	endpoint := c.GetEndpoint()
	clientID := c.GetClientID()
	clientSecret := c.GetClientSecret()
	organizationName := c.GetOrganizationName()
	applicationName := c.GetApplicationName()
	certificate := `
	-----BEGIN CERTIFICATE-----
	MIIE2TCCAsGgAwIBAgIDAeJAMA0GCSqGSIb3DQEBCwUAMCYxDjAMBgNVBAoTBWFk
	bWluMRQwEgYDVQQDDAtjZXJ0XzR4bzNzZzAeFw0yNjAzMTExNDE4MTZaFw00NjAz
	MTExNDE4MTZaMCYxDjAMBgNVBAoTBWFkbWluMRQwEgYDVQQDDAtjZXJ0XzR4bzNz
	ZzCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAOgble9wfQg248/UJnC2
	H0gkiCZJOvbOiV6aFTIWBdiVquj8TLTRuj9pQ3gBqhPx9YaQPwPQFbLWqEsEZt3k
	M6w/xyR97tPZgFSidkvWwG+du1TYzhJU97RwtQytyTdfk45d0cMUNrxMOlACqPxk
	Xow6pON3NdUc1DwkPjP8MMeSWJVYFJNNYGXFEl2C3QFSRMSYpKE813/zT4m8ElDs
	scVvYwz4W1EDtUFa2wxdkshlo/WRZROvEQEtexpCO5UquVgtuN54aFMZT0Funr17
	NrczN0zu2yeE12Ow30JgGGLin60AZLgHzCE5NY4TwayMBe2OCObLlOMBu1zupXnJ
	l64zU92Shu7oAtVHNcFbJPxM7immr/scnvWxibnZwbysx0tVsDwj4YBtWu9xaelg
	oplZwWU54yzhIJswN49njcDXGAsyvMJcMHtmLRHzFNAHO9m2ojr04HwjqhiawgvT
	8r9C+XptCutf6OIhpVDgKJorAsea6X22R/sCpCoWtQtTT6VSPjwZaJU7SWcZjqiC
	k8sNxqn/RDn1udml1t+fTl+lk/6M3nKt5y8fT1owY2y3xBm611Xao9FKiGuXHz6O
	dvMthXFB41NzRJ2/aPYCp3PHgBsI/DKK0IUWQ52CaHzMrxUaKzohfw3D9u999+IS
	1QxKNUvfO5dUeGNWGPpuiLZHAgMBAAGjEDAOMAwGA1UdEwEB/wQCMAAwDQYJKoZI
	hvcNAQELBQADggIBAJzsdotKQouu92zfW1raMisiis7ZtX/Xk9z5W6vSq3dpxgPc
	C2YT6WwO36tUoAyHq3JaEUkoSWunpVupuFtJfabzkf4mFcbwbpIkgtABvghRGxcb
	BwYYeaL9vn8oCDhWKDOgvtUQ8p2AB701Cmn+zV3IreGW8T+wi6rlui+9L+VJKASy
	oezkxFONHDxBzQ9TlFIdiQuv+Z/auiR+M+nqrUWLubNvhbq28wRQfo+ALyvAUszV
	DWxu+1isCss1c2c/cNZeI2kL1lolJT8RVHPrqaZxLdQ2FgK99lsteY7Sq0DRm+gL
	iVASYyV//AXNhj9XV//6XCA7l+xG3W/Rp2TE4XzwJHwVZ7/nAjDV84IPPXig0qx9
	hL5M17b+8lN0EfY5R5AXNrngUthYSDeq/UM/FLeWGh8uEeWtoawFsHwnaH0eUfiy
	yZL65ZF2MdJG8E6fSSWM+FBjNzG1v2uUChdKDcqU1+aJWKv8mxzbQSm228lpHk+6
	0BCRhddOAZPGV+J4DkKu5DXtvl23o4k5I7VTfXLo7NWMJHZbo3Pg2/6EgHl5rHvZ
	uPrFwL2CbZGItHCBZIgN1GuHyRkC4R1uJ3y5DP6mtVMkajU0BSl0mXkGCIGlR6BO
	xibtg2UVnxLyj7EWZOEVAES1MbTraX+MKoFN5AHXr3DrvJ3n+WHaFpcxwjjl
	-----END CERTIFICATE-----
	`

	casdoorsdk.InitConfig(endpoint, clientID, clientSecret, certificate, organizationName, applicationName)
}
