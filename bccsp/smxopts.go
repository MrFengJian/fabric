package bccsp

const (
	// SM4
	SM4 = "SM4"
	// SM3
	SM3 = "SM3"
	// SM2
	SM2 = "SM2"
)

// SM2KeyGenOpts contains options for SM2 key generation.
type SM2KeyGenOpts struct {
	Temporary bool
}

// Algorithm returns the key generation algorithm identifier (to be used).
func (opts *SM2KeyGenOpts) Algorithm() string {
	return SM2
}

// Ephemeral returns true if the key to generate has to be ephemeral,
// false otherwise.
func (opts *SM2KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

// SM4KeyGenOpts contains options for SM2 key generation.
type SM4KeyGenOpts struct {
	Temporary bool
}

// Algorithm returns the key generation algorithm identifier (to be used).
func (opts *SM4KeyGenOpts) Algorithm() string {
	return SM4
}

// Ephemeral returns true if the key to generate has to be ephemeral,
// false otherwise.
func (opts *SM4KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

//SM4ImportKeyOpts  实现  bccsp.KeyImportOpts 接口
type SM4ImportKeyOpts struct {
	Temporary bool
}

// Algorithm returns the key importation algorithm identifier (to be used).
func (opts *SM4ImportKeyOpts) Algorithm() string {
	return SM4
}

// Ephemeral returns true if the key generated has to be ephemeral,
// false otherwise.
func (opts *SM4ImportKeyOpts) Ephemeral() bool {
	return opts.Temporary
}

//SM2PrivateKeyImportOpts  实现  bccsp.KeyImportOpts 接口
type SM2PrivateKeyImportOpts struct {
	Temporary bool
}

// Algorithm returns the key importation algorithm identifier (to be used).
func (opts *SM2PrivateKeyImportOpts) Algorithm() string {
	return SM2
}

// Ephemeral returns true if the key generated has to be ephemeral,
// false otherwise.
func (opts *SM2PrivateKeyImportOpts) Ephemeral() bool {
	return opts.Temporary
}

//SM2PublicKeyImportOpts  实现  bccsp.KeyImportOpts 接口
type SM2PublicKeyImportOpts struct {
	Temporary bool
}

// Algorithm returns the key importation algorithm identifier (to be used).
func (opts *SM2PublicKeyImportOpts) Algorithm() string {
	return SM2
}

// Ephemeral returns true if the key generated has to be ephemeral,
// false otherwise.
func (opts *SM2PublicKeyImportOpts) Ephemeral() bool {
	return opts.Temporary
}